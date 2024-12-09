package pdu

import (
	"encoding/binary"
	"io"
	"sync/atomic"

	"github.com/kashari/draupnir/constants"
	"github.com/kashari/draupnir/errors"
)

type PDU interface {
	// Marshal PDU to buffer.
	Marshal(*ByteBuffer)

	// Unmarshal PDU from buffer.
	Unmarshal(*ByteBuffer) error

	// CanResponse indicates that PDU could response to SMSC.
	CanResponse() bool

	// GetResponse PDU.
	GetResponse() PDU

	// RegisterOptionalParam assigns an optional param.
	RegisterOptionalParam(Field)

	// GetHeader returns PDU header.
	GetHeader() Header

	// IsOk returns true if command status is OK.
	IsOk() bool

	// IsGNack returns true if PDU is GNack.
	IsGNack() bool

	// AssignSequenceNumber assigns sequence number auto-incrementally.
	AssignSequenceNumber()

	// ResetSequenceNumber resets sequence number.
	ResetSequenceNumber()

	// GetSequenceNumber returns assigned sequence number.
	GetSequenceNumber() int32

	// SetSequenceNumber manually sets sequence number.
	SetSequenceNumber(int32)
}

func nextSequenceNumber(s *int32) (v int32) {
	// & 0x7FFFFFFF: cater for integer overflow
	// Allowed range is 0x01 to 0x7FFFFFFF. This
	// will still result in a single invalid value
	// of 0x00 every ~2 billion PDUs (not too bad):
	if v = atomic.AddInt32(s, 1) & 0x7FFFFFFF; v <= 0 {
		v = 1
	}
	return
}

// Header represents PDU header.
type Header struct {
	CommandLength  int32
	CommandID      constants.CommandIDType
	CommandStatus  constants.CommandStatusType
	SequenceNumber int32
}

// ParseHeader parses PDU header.
func ParseHeader(v [16]byte) (h Header) {
	h.CommandLength = int32(binary.BigEndian.Uint32(v[:]))
	h.CommandID = constants.CommandIDType(binary.BigEndian.Uint32(v[4:]))
	h.CommandStatus = constants.CommandStatusType(binary.BigEndian.Uint32(v[8:]))
	h.SequenceNumber = int32(binary.BigEndian.Uint32(v[12:]))
	return
}

// Unmarshal from buffer.
func (c *Header) Unmarshal(b *ByteBuffer) (err error) {
	var id, status int32
	c.CommandLength, err = b.ReadInt()
	if err == nil {
		id, err = b.ReadInt()
		if err == nil {
			c.CommandID = constants.CommandIDType(id)
			if status, err = b.ReadInt(); err == nil {
				c.CommandStatus = constants.CommandStatusType(status)
				c.SequenceNumber, err = b.ReadInt()
			}
		}
	}
	return
}

var sequenceNumber int32

// AssignSequenceNumber assigns sequence number auto-incrementally.
func (c *Header) AssignSequenceNumber() {
	c.SetSequenceNumber(nextSequenceNumber(&sequenceNumber))
}

// ResetSequenceNumber resets sequence number.
func (c *Header) ResetSequenceNumber() {
	c.SequenceNumber = 1
}

// GetSequenceNumber returns assigned sequence number.
func (c *Header) GetSequenceNumber() int32 {
	return c.SequenceNumber
}

// SetSequenceNumber manually sets sequence number.
func (c *Header) SetSequenceNumber(v int32) {
	c.SequenceNumber = v
}

// Marshal to buffer.
func (c *Header) Marshal(b *ByteBuffer) {
	b.Grow(16)
	b.WriteInt(c.CommandLength)
	b.WriteInt(int32(c.CommandID))
	b.WriteInt(int32(c.CommandStatus))
	b.WriteInt(c.SequenceNumber)
}


type base struct {
	Header
	OptionalParameters map[Tag]Field
}

func newBase() (v base) {
	v.OptionalParameters = make(map[Tag]Field)
	v.AssignSequenceNumber()
	return
}

type pduGenerator func() PDU

var pduMap = map[constants.CommandIDType]pduGenerator{
	constants.BIND_TRANSMITTER:      NewBindTransmitter,
	constants.BIND_TRANSMITTER_RESP: NewBindTransmitterResp,
	constants.BIND_TRANSCEIVER:      NewBindTransceiver,
	constants.BIND_TRANSCEIVER_RESP: NewBindTransceiverResp,
	constants.BIND_RECEIVER:         NewBindReceiver,
	constants.BIND_RECEIVER_RESP:    NewBindReceiverResp,
	constants.UNBIND:                NewUnbind,
	constants.UNBIND_RESP:           NewUnbindResp,
	constants.OUTBIND:               NewOutbind,
	constants.SUBMIT_SM:             NewSubmitSm,
	constants.SUBMIT_SM_RESP:        NewSubmitSmResp,
	constants.SUBMIT_MULTI:          NewSubmitMulti,
	constants.SUBMIT_MULTI_RESP:     NewSubmitMultiResp,
	constants.DELIVER_SM:            NewDeliverSm,
	constants.DELIVER_SM_RESP:       NewDeliverSmResp,
	constants.DATA_SM:               NewDataSm,
	constants.DATA_SM_RESP:          NewDataSmResp,
	constants.QUERY_SM:              NewQuerySm,
	constants.QUERY_SM_RESP:         NewQuerySmResp,
	constants.CANCEL_SM:             NewCancelSm,
	constants.CANCEL_SM_RESP:        NewCancelSmResp,
	constants.REPLACE_SM:            NewReplaceSm,
	constants.REPLACE_SM_RESP:       NewReplaceSmResp,
	constants.ENQUIRE_LINK:          NewEnquireLink,
	constants.ENQUIRE_LINK_RESP:     NewEnquireLinkResp,
	constants.ALERT_NOTIFICATION:    NewAlertNotification,
	constants.GENERIC_NACK:          NewGenericNack,
}

// CreatePDUFromCmdID creates PDU from cmd id.
func CreatePDUFromCmdID(cmdID constants.CommandIDType) (PDU, error) {
	if g, ok := pduMap[cmdID]; ok {
		return g(), nil
	}
	return nil, errors.ErrUnknownCommandID
}

// GetHeader returns pdu header.
func (c *base) GetHeader() Header {
	return c.Header
}

func (c *base) unmarshal(b *ByteBuffer, bodyReader func(*ByteBuffer) error) (err error) {
	fullLen := b.Len()

	if err = c.Header.Unmarshal(b); err == nil {

		// try to unmarshal body
		if bodyReader != nil {
			err = bodyReader(b)
		}

		if err == nil {
			// command length
			cmdLength := int(c.CommandLength)

			// got - total read byte(s)
			got := fullLen - b.Len()
			if got > cmdLength {
				err = errors.ErrInvalidPDU
				return
			}

			// body < command_length, still have optional parameters ?
			if got < cmdLength {
				var optParam []byte
				if optParam, err = b.ReadN(cmdLength - got); err == nil {
					err = c.unmarshalOptionalParam(optParam)
				}
				if err != nil {
					return
				}
			}

			// validate again
			if b.Len() != fullLen-cmdLength {
				err = errors.ErrInvalidPDU
			}
		}
	}

	return
}

func (c *base) unmarshalOptionalParam(optParam []byte) (err error) {
	buf := NewBuffer(optParam)
	for buf.Len() > 0 {
		var field Field
		if err = field.Unmarshal(buf); err == nil {
			c.OptionalParameters[field.Tag] = field
		} else {
			return
		}
	}
	return
}

// Marshal to buffer.
func (c *base) marshal(b *ByteBuffer, bodyWriter func(*ByteBuffer)) {
	bodyBuf := NewBuffer(nil)

	// body
	if bodyWriter != nil {
		bodyWriter(bodyBuf)
	}

	// optional body
	for _, v := range c.OptionalParameters {
		v.Marshal(bodyBuf)
	}

	// write header
	c.CommandLength = int32(constants.PDU_HEADER_SIZE + bodyBuf.Len())
	c.Header.Marshal(b)

	// write body and its optional params
	b.WriteBuffer(bodyBuf)
}

// RegisterOptionalParam register optional param.
func (c *base) RegisterOptionalParam(tlv Field) {
	c.OptionalParameters[tlv.Tag] = tlv
}

// IsOk is status ok.
func (c *base) IsOk() bool {
	return c.CommandStatus == constants.ESME_ROK
}

// IsGNack is generic n-ack.
func (c *base) IsGNack() bool {
	return c.CommandID == constants.GENERIC_NACK
}

// Parse PDU from reader.
func Parse(r io.Reader) (pdu PDU, err error) {
	var headerBytes [16]byte

	if _, err = io.ReadFull(r, headerBytes[:]); err != nil {
		return
	}

	header := ParseHeader(headerBytes)
	if header.CommandLength < 16 || header.CommandLength > constants.MAX_PDU_LEN {
		err = errors.ErrInvalidPDU
		return
	}

	// read pdu body
	bodyBytes := make([]byte, header.CommandLength-16)
	if len(bodyBytes) > 0 {
		if _, err = io.ReadFull(r, bodyBytes); err != nil {
			return
		}
	}

	// try to create pdu
	if pdu, err = CreatePDUFromCmdID(header.CommandID); err == nil {
		buf := NewBuffer(make([]byte, 0, header.CommandLength))
		_, _ = buf.Write(headerBytes[:])
		if len(bodyBytes) > 0 {
			_, _ = buf.Write(bodyBytes)
		}
		err = pdu.Unmarshal(buf)
	}

	return
}