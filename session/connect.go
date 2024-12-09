package draupnir

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"time"

	"github.com/kashari/draupnir/constants"
	"github.com/kashari/draupnir/pdu"
)

var (
	// NonTLSDialer is non-tls connection dialer.
	NonTLSDialer = func(addr string) (net.Conn, error) {
		return net.Dial("tcp", addr)
	}
)

// Dialer is connection dialer.
type Dialer func(addr string) (net.Conn, error)
type State byte

// Auth represents basic authentication to SMSC.
type Auth struct {
	// SMSC is SMSC address.
	SMSC       string
	SystemID   string
	Password   string
	SystemType string
}

type BindError struct {
	CommandStatus constants.CommandStatusType
}

func (err BindError) Error() string {
	return fmt.Sprintf("binding error (%s): %s", err.CommandStatus, err.CommandStatus.Desc())
}

func newBindRequest(s Auth, bindingType pdu.BindingType, addressRange pdu.AddressRange) (bindReq *pdu.BindRequest) {
	bindReq = pdu.NewBindRequest(bindingType)
	bindReq.SystemID = s.SystemID
	bindReq.Password = s.Password
	bindReq.SystemType = s.SystemType
	bindReq.AddressRange = addressRange
	return
}

type Connection struct {
	systemID string
	conn     net.Conn
	reader   *bufio.Reader
}

// Connector is connection factory interface.
type Connector interface {
	Connect() (conn *Connection, err error)
	GetBindType() pdu.BindingType
}

type connector struct {
	dialer       Dialer
	auth         Auth
	bindingType  pdu.BindingType
	addressRange pdu.AddressRange
}

func (c *connector) GetBindType() pdu.BindingType {
	return c.bindingType
}

func (c *connector) Connect() (conn *Connection, err error) {
	conn, err = connect(c.dialer, c.auth.SMSC, newBindRequest(c.auth, c.bindingType, c.addressRange))
	return
}

func connect(dialer Dialer, addr string, bindReq *pdu.BindRequest) (c *Connection, err error) {
	conn, err := dialer(addr)
	if err != nil {
		return
	}

	c = NewConnection(conn)

	_, err = c.WritePDU(bindReq)
	if err != nil {
		_ = conn.Close()
		return
	}

	var (
		p    pdu.PDU
		resp *pdu.BindResp
	)

	for {
		if p, err = pdu.Parse(c); err != nil {
			_ = conn.Close()
			return
		}

		if pd, ok := p.(*pdu.BindResp); ok {
			resp = pd
			break
		}
	}

	if resp.CommandStatus != constants.ESME_ROK {
		err = BindError{CommandStatus: resp.CommandStatus}
		_ = conn.Close()
	} else {
		c.systemID = resp.SystemID
	}

	return
}

// TXConnector returns a Transmitter (TX) connector.
func TXConnector(dialer Dialer, auth Auth) Connector {
	return &connector{
		dialer:      dialer,
		auth:        auth,
		bindingType: pdu.Transmitter,
	}
}

// RXConnector returns a Receiver (RX) connector.
func RXConnector(dialer Dialer, auth Auth, opts ...connectorOption) Connector {
	c := &connector{
		dialer:      dialer,
		auth:        auth,
		bindingType: pdu.Receiver,
	}
	for _, opt := range opts {
		opt(c)
	}
	return c
}

// TRXConnector returns a Transceiver (TRX) connector.
func TRXConnector(dialer Dialer, auth Auth, opts ...connectorOption) Connector {
	c := &connector{
		dialer:      dialer,
		auth:        auth,
		bindingType: pdu.Transceiver,
	}
	for _, opt := range opts {
		opt(c)
	}
	return c
}

type connectorOption func(c *connector)

func WithAddressRange(addressRange pdu.AddressRange) connectorOption {
	return func(c *connector) {
		c.addressRange = addressRange
	}
}


// NewConnection returns a Connection.
func NewConnection(conn net.Conn) (c *Connection) {
	c = &Connection{
		conn:   conn,
		reader: bufio.NewReaderSize(conn, 128<<10),
	}
	return
}

// Read reads data from the connection.
// Read can be made to time out and return an Error with Timeout() == true
// after a fixed time limit; see SetDeadline and SetReadDeadline.
func (c *Connection) Read(b []byte) (n int, err error) {
	n, err = c.reader.Read(b)
	return
}

// Write writes data to the connection.
// Write can be made to time out and return an Error with Timeout() == true
// after a fixed time limit; see SetDeadline and SetWriteDeadline.
func (c *Connection) Write(b []byte) (n int, err error) {
	n, err = c.conn.Write(b)
	return
}

// WritePDU data to the connection.
func (c *Connection) WritePDU(p pdu.PDU) (n int, err error) {
	buf := pdu.NewBuffer(make([]byte, 0, 64))
	p.Marshal(buf)
	n, err = c.conn.Write(buf.Bytes())
	return
}

// Close closes the connection.
// Any blocked Read or Write operations will be unblocked and return errors.
func (c *Connection) Close() error {
	return c.conn.Close()
}

// LocalAddr returns the local network address.
func (c *Connection) LocalAddr() net.Addr {
	return c.conn.LocalAddr()
}

// RemoteAddr returns the remote network address.
func (c *Connection) RemoteAddr() net.Addr {
	return c.conn.RemoteAddr()
}

// SetDeadline sets the read and write deadlines associated
// with the connection. It is equivalent to calling both
// SetReadDeadline and SetWriteDeadline.
//
// A deadline is an absolute time after which I/O operations
// fail with a timeout (see type Error) instead of
// blocking. The deadline applies to all future and pending
// I/O, not just the immediately following call to Read or
// Write. After a deadline has been exceeded, the connection
// can be refreshed by setting a deadline in the future.
//
// An idle timeout can be implemented by repeatedly extending
// the deadline after successful Read or Write calls.
//
// A zero value for t means I/O operations will not time out.
//
// Note that if a TCP connection has keep-alive turned on,
// which is the default unless overridden by Dialer.KeepAlive
// or ListenConfig.KeepAlive, then a keep-alive failure may
// also return a timeout error. On Unix systems a keep-alive
// failure on I/O can be detected using
// errors.Is(err, syscall.ETIMEDOUT).
func (c *Connection) SetDeadline(t time.Time) error {
	return c.conn.SetDeadline(t)
}

// SetReadDeadline sets the deadline for future Read calls
// and any currently-blocked Read call.
// A zero value for t means Read will not time out.
func (c *Connection) SetReadDeadline(t time.Time) error {
	return c.conn.SetReadDeadline(t)
}

// SetReadTimeout is equivalent to ReadDeadline(now + timeout)
func (c *Connection) SetReadTimeout(t time.Duration) error {
	return c.conn.SetReadDeadline(time.Now().Add(t))
}

// SetWriteDeadline sets the deadline for future Write calls
// and any currently-blocked Write call.
// Even if write times out, it may return n > 0, indicating that
// some of the data was successfully written.
// A zero value for t means Write will not time out.
func (c *Connection) SetWriteDeadline(t time.Time) error {
	return c.conn.SetWriteDeadline(t)
}

// SetWriteTimeout is equivalent to WriteDeadline(now + timeout)
func (c *Connection) SetWriteTimeout(t time.Duration) error {
	return c.conn.SetWriteDeadline(time.Now().Add(t))
}

// Transceiver interface.
type Transceiver interface {
	io.Closer
	Submit(pdu.PDU) error
	SystemID() string
}

// Transmitter interface.
type Transmitter interface {
	io.Closer
	Submit(pdu.PDU) error
	SystemID() string
}

// Receiver interface.
type Receiver interface {
	io.Closer
	SystemID() string
}

// Request represent a request tracked by the RequestStore
type Request struct {
	pdu.PDU
	TimeSent time.Time
}

// Response represents a response from a Request in the RequestStore
type Response struct {
	pdu.PDU
	OriginalRequest Request
}

// PDUCallback handles received PDU.
type PDUCallback func(pdu pdu.PDU, responded bool)

// AllPDUCallback handles all received PDU.
//
// This pdu is NOT responded to automatically, manual response/handling is needed
// and the bind can be closed by retuning true on closeBind.
type AllPDUCallback func(pdu pdu.PDU) (responsePdu pdu.PDU, closeBind bool)

// PDUErrorCallback notifies fail-to-submit PDU with along error.
type PDUErrorCallback func(pdu pdu.PDU, err error)

// ErrorCallback notifies happened error while reading PDU.
type ErrorCallback func(error)

// ClosedCallback notifies closed event due to State.
type ClosedCallback func(State)

// RebindCallback notifies rebind event due to State.
type RebindCallback func()

// WindowedRequestTracking settings for TX (transmitter) and TRX (transceiver) request store.
type WindowedRequestTracking struct {

	// OnReceivedPduRequest handles received PDU request from SMSC.
	//
	// User can also decide to close bind by retuning true, default is false
	OnReceivedPduRequest AllPDUCallback

	// OnExpectedPduResponse handles expected PDU response from SMSC.
	// Only triggered when the original request is found in the window cache
	//
	// Handle is optional
	// If not set, response will be dropped
	OnExpectedPduResponse func(Response)

	// OnUnexpectedPduResponse handles unexpected PDU response from SMSC.
	// Only triggered if the original request is not found in the window cache
	//
	// Handle is optional
	// If not set, response will be dropped
	OnUnexpectedPduResponse func(pdu.PDU)

	// OnExpiredPduRequest handles expired PDU request with no response received
	//
	// Mandatory: the PduExpireTimeOut must be set
	// Handle is optional
	// If not set, expired PDU will be removed from cache
	// the bind can be closed by retuning true on closeBind.
	OnExpiredPduRequest func(pdu.PDU) (closeBind bool)

	// OnClosePduRequest will return all PDU request found in the store when the bind closes
	OnClosePduRequest func(pdu.PDU)

	// Set the number of second to expire a request sent to the SMSC
	//
	// Zero duration disables pdu expire check and the cache may fill up over time with expired PDU request
	// Recommended: eual or less to the value set in ReadTimeout + EnquireLink
	PduExpireTimeOut time.Duration

	// The time period between each check of the expired PDU in the cache
	//
	// Zero duration disables pdu expire check and the cache may fill up over time with expired PDU request
	// Recommended: Less or half the time set in for PduExpireTimeOut
	// Don't be too aggressive, there is a performance hit if the check is done often
	ExpireCheckTimer time.Duration

	// The maximum number of pending request sent to the SMSC
	//
	// Maximum value is 255
	MaxWindowSize uint8

	// if enabled, EnquireLink and Unbind request will be responded to automatically
	EnableAutoRespond bool

	// Set the number of millisecond to expire a request to store or retrieve data from request store
	//
	// Value must be greater than 0
	// 200 to 1000 is a good starting point
	StoreAccessTimeOut time.Duration
}

type Settings struct {
	// ReadTimeout is timeout for reading PDU from SMSC.
	// Underlying net.Conn will be stricted with ReadDeadline(now + timeout).
	// This setting is very important to detect connection failure.
	//
	// Must: ReadTimeout > max(0, EnquireLink)
	ReadTimeout time.Duration

	// WriteTimeout is timeout for submitting PDU.
	WriteTimeout time.Duration

	// EnquireLink periodically sends EnquireLink to SMSC.
	// The duration must not be smaller than 1 minute.
	//
	// Zero duration disables auto enquire link.
	EnquireLink time.Duration

	// OnPDU handles received PDU from SMSC.
	//
	// `Responded` flag indicates this pdu is responded automatically,
	// no manual respond needed.
	//
	// Will be ignored if OnAllPDU or WindowedRequestTracking is set
	OnPDU PDUCallback

	// OnAllPDU handles all received PDU from SMSC.
	//
	// This pdu is NOT responded to automatically,
	// manual response/handling is needed
	//
	// User can also decide to close bind by retuning true, default is false
	//
	// Will be ignored if WindowedRequestTracking is set
	OnAllPDU AllPDUCallback

	// OnReceivingError notifies happened error while reading PDU
	// from SMSC.
	OnReceivingError ErrorCallback

	// OnSubmitError notifies fail-to-submit PDU with along error.
	OnSubmitError PDUErrorCallback

	// OnRebindingError notifies error while rebinding.
	OnRebindingError ErrorCallback

	// OnClosed notifies `closed` event due to State.
	OnClosed ClosedCallback

	// OnRebind notifies `rebind` event due to State.
	OnRebind RebindCallback

	// SMPP Bind Window tracking feature config
	*WindowedRequestTracking

	response func(pdu.PDU)
}