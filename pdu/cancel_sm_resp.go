package pdu

import "github.com/kashari/draupnir/constants"

// CancelSmResp PDU.
type CancelSmResp struct {
	base
}

// NewCancelSmResp returns CancelSmResp.
func NewCancelSmResp() PDU {
	c := &CancelSmResp{
		base: newBase(),
	}
	c.CommandID = constants.CANCEL_SM_RESP
	return c
}

// NewCancelSmRespFromReq returns CancelSmResp.
func NewCancelSmRespFromReq(req *CancelSm) PDU {
	c := NewCancelSmResp().(*CancelSmResp)
	if req != nil {
		c.SequenceNumber = req.SequenceNumber
	}
	return c
}

// CanResponse implements PDU interface.
func (c *CancelSmResp) CanResponse() bool {
	return false
}

// GetResponse implements PDU interface.
func (c *CancelSmResp) GetResponse() PDU {
	return nil
}

// Marshal implements PDU interface.
func (c *CancelSmResp) Marshal(b *ByteBuffer) {
	c.base.marshal(b, nil)
}

// Unmarshal implements PDU interface.
func (c *CancelSmResp) Unmarshal(b *ByteBuffer) error {
	return c.base.unmarshal(b, nil)
}
