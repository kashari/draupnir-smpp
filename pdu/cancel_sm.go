package pdu

import "github.com/kashari/draupnir/constants"

// CancelSm PDU is issued by the ESME to cancel one or more previously submitted short messages
// that are still pending delivery. The command may specify a particular message to cancel, or
// all messages for a particular source, destination and service_type are to be cancelled.
type CancelSm struct {
	base
	ServiceType string
	MessageID   string
	SourceAddr  Address
	DestAddr    Address
}

// NewCancelSm returns CancelSm PDU.
func NewCancelSm() PDU {
	c := &CancelSm{
		base:        newBase(),
		ServiceType: constants.DFLT_SRVTYPE,
		MessageID:   constants.DFLT_MSGID,
		SourceAddr:  NewAddress(),
		DestAddr:    NewAddress(),
	}
	c.CommandID = constants.CANCEL_SM
	return c
}

// CanResponse implements PDU interface.
func (c *CancelSm) CanResponse() bool {
	return true
}

// GetResponse implements PDU interface.
func (c *CancelSm) GetResponse() PDU {
	return NewCancelSmRespFromReq(c)
}

// Marshal implements PDU interface.
func (c *CancelSm) Marshal(b *ByteBuffer) {
	c.base.marshal(b, func(b *ByteBuffer) {
		b.Grow(len(c.ServiceType) + len(c.MessageID) + 2)

		_ = b.WriteCString(c.ServiceType)
		_ = b.WriteCString(c.MessageID)
		c.SourceAddr.Marshal(b)
		c.DestAddr.Marshal(b)
	})
}

// Unmarshal implements PDU interface.
func (c *CancelSm) Unmarshal(b *ByteBuffer) error {
	return c.base.unmarshal(b, func(b *ByteBuffer) (err error) {
		if c.ServiceType, err = b.ReadCString(); err == nil {
			if c.MessageID, err = b.ReadCString(); err == nil {
				if err = c.SourceAddr.Unmarshal(b); err == nil {
					err = c.DestAddr.Unmarshal(b)
				}
			}
		}
		return
	})
}
