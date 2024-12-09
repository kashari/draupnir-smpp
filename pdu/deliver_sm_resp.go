package pdu

import "github.com/kashari/draupnir/constants"

// DeliverSmResp PDU.
type DeliverSmResp struct {
	base
	MessageID string
}

// NewDeliverSmResp returns new DeliverSmResp.
func NewDeliverSmResp() PDU {
	c := &DeliverSmResp{
		base:      newBase(),
		MessageID: constants.DFLT_MSGID,
	}
	c.CommandID = constants.DELIVER_SM_RESP
	return c
}

// NewDeliverSmRespFromReq returns new DeliverSmResp.
func NewDeliverSmRespFromReq(req *DeliverSm) PDU {
	c := NewDeliverSmResp().(*DeliverSmResp)
	if req != nil {
		c.SequenceNumber = req.SequenceNumber
	}
	return c
}

// CanResponse implements PDU interface.
func (c *DeliverSmResp) CanResponse() bool {
	return false
}

// GetResponse implements PDU interface.
func (c *DeliverSmResp) GetResponse() PDU {
	return nil
}

// Marshal implements PDU interface.
func (c *DeliverSmResp) Marshal(b *ByteBuffer) {
	c.base.marshal(b, func(b *ByteBuffer) {
		b.Grow(len(c.MessageID) + 1)

		_ = b.WriteCString(c.MessageID)
	})
}

// Unmarshal implements PDU interface.
func (c *DeliverSmResp) Unmarshal(b *ByteBuffer) error {
	return c.base.unmarshal(b, func(b *ByteBuffer) (err error) {
		c.MessageID, err = b.ReadCString()
		return
	})
}
