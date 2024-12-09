package draupnir

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"sync/atomic"
	"time"

	"github.com/kashari/draupnir/pdu"
	cmap "github.com/orcaman/concurrent-map/v2"
	"golang.org/x/exp/maps"
)

// RequestStore interface used for WindowedRequestTracking
type RequestStore interface {
	Set(ctx context.Context, request Request) error
	Get(ctx context.Context, sequenceNumber int32) (Request, bool)
	List(ctx context.Context) []Request
	Delete(ctx context.Context, sequenceNumber int32) error
	Clear(ctx context.Context) error
	Length(ctx context.Context) (int, error)
}

type DefaultStore struct {
	store cmap.ConcurrentMap[string, Request]
}

func NewDefaultStore() DefaultStore {
	return DefaultStore{
		store: cmap.New[Request](),
	}
}

func (s DefaultStore) Set(ctx context.Context, request Request) error {
	select {
	case <-ctx.Done():
		fmt.Println("Task cancelled")
		return ctx.Err()
	default:
		s.store.Set(strconv.Itoa(int(request.PDU.GetSequenceNumber())), request)
		return nil
	}
}

func (s DefaultStore) Get(ctx context.Context, sequenceNumber int32) (Request, bool) {
	select {
	case <-ctx.Done():
		fmt.Println("Task cancelled")
		return Request{}, false
	default:
		return s.store.Get(strconv.Itoa(int(sequenceNumber)))
	}
}

func (s DefaultStore) List(ctx context.Context) []Request {
	select {
	case <-ctx.Done():
		return []Request{}
	default:
		return maps.Values(s.store.Items())
	}
}

func (s DefaultStore) Delete(ctx context.Context, sequenceNumber int32) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
		s.store.Remove(strconv.Itoa(int(sequenceNumber)))
		return nil
	}
}

func (s DefaultStore) Clear(ctx context.Context) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
		s.store.Clear()
		return nil
	}
}

func (s DefaultStore) Length(ctx context.Context) (int, error) {
	select {
	case <-ctx.Done():
		return 0, ctx.Err()
	default:
		return s.store.Count(), nil
	}
}


var (
	ErrWindowSizeEqualZero                   = errors.New("request window size cannot be 0")
	ErrExpireCheckTimerNotSet                = errors.New("ExpireCheckTimer cannot be 0 if PduExpireTimeOut is set")
	ErrStoreAccessTimeOutEqualZero           = errors.New("StoreAccessTimeOut window size cannot be 0")
	ErrWindowSizeNotAvailableOnReceiverBinds = errors.New("window size not available on receiver binds")
)

// Session represents session for TX, RX, TRX.
type Session struct {
	c Connector

	originalOnClosed func(State)
	settings         Settings

	rebindingInterval time.Duration

	trx atomic.Value // transceivable

	state        int32
	rebinding    int32
	requestStore RequestStore
}

type SessionOption func(session *Session)

// NewSession creates new session for TX, RX, TRX.
//
// Session will `non-stop`, automatically rebind (create new and authenticate connection with SMSC) when
// unexpected error happened.
//
// `rebindingInterval` indicates duration that Session has to wait before rebinding again.
//
// Setting `rebindingInterval <= 0` will disable `auto-rebind` functionality.
func NewSession(c Connector, settings Settings, rebindingInterval time.Duration, opts ...SessionOption) (session *Session, err error) {
	// Loop through each option

	if settings.ReadTimeout <= 0 || settings.ReadTimeout <= settings.EnquireLink {
		return nil, fmt.Errorf("invalid settings: ReadTimeout must greater than max(0, EnquireLink)")
	}
	var requestStore RequestStore = nil
	if settings.WindowedRequestTracking != nil {
		requestStore = NewDefaultStore()
		if settings.MaxWindowSize == 0 {
			return nil, ErrWindowSizeEqualZero
		}
		if settings.StoreAccessTimeOut == 0 {
			return nil, ErrStoreAccessTimeOutEqualZero
		}
		if settings.PduExpireTimeOut > 0 && settings.ExpireCheckTimer == 0 {
			return nil, ErrExpireCheckTimerNotSet
		}
	}

	conn, err := c.Connect()
	if err == nil {
		session = &Session{
			c:                 c,
			rebindingInterval: rebindingInterval,
			originalOnClosed:  settings.OnClosed,
			requestStore:      requestStore,
		}

		for _, opt := range opts {
			opt(session)
		}

		if rebindingInterval > 0 {
			newSettings := settings
			newSettings.OnClosed = func(state State) {
				switch state {
				case ExplicitClosing:
					return

				default:
					if session.originalOnClosed != nil {
						session.originalOnClosed(state)
					}
					session.rebind()
				}
			}
			session.settings = newSettings
		} else {
			session.settings = settings
		}

		// bind to session
		trans := newTransceivable(conn, session.settings, session.requestStore)
		trans.start()
		session.trx.Store(trans)
	}
	return
}

func WithRequestStore(store RequestStore) SessionOption {
	return func(s *Session) {
		s.requestStore = store
	}
}

func (s *Session) bound() *transceivable {
	r, _ := s.trx.Load().(*transceivable)
	return r
}

// Transmitter returns bound Transmitter.
func (s *Session) Transmitter() Transmitter {
	return s.bound()
}

// Receiver returns bound Receiver.
func (s *Session) Receiver() Receiver {
	return s.bound()
}

// Transceiver returns bound Transceiver.
func (s *Session) Transceiver() Transceiver {
	return s.bound()
}

func (s *Session) GetWindowSize() (int, error) {
	if s.c.GetBindType() == pdu.Transmitter || s.c.GetBindType() == pdu.Transceiver {
		size, err := s.bound().GetWindowSize()
		if err != nil {
			return 0, err
		}
		return size, nil
	}
	return 0, ErrWindowSizeNotAvailableOnReceiverBinds
}

// Close session.
func (s *Session) Close() (err error) {
	if atomic.CompareAndSwapInt32(&s.state, Alive, Closed) {
		err = s.close()
	}
	return
}

func (s *Session) close() (err error) {
	if b := s.bound(); b != nil {
		err = b.Close()
	}
	return
}

func (s *Session) rebind() {
	if atomic.CompareAndSwapInt32(&s.rebinding, 0, 1) {
		_ = s.close()

		for atomic.LoadInt32(&s.state) == Alive {
			conn, err := s.c.Connect()
			if err != nil {
				if s.settings.OnRebindingError != nil {
					s.settings.OnRebindingError(err)
				}
				time.Sleep(s.rebindingInterval)
			} else {
				// bind to session
				trans := newTransceivable(conn, s.settings, s.requestStore)
				trans.start()
				s.trx.Store(trans)

				// reset rebinding state
				atomic.StoreInt32(&s.rebinding, 0)
				if s.settings.OnRebind != nil {
					s.settings.OnRebind()
				}

				return
			}
		}
	}
}
