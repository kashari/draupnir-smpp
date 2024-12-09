package pdu

import "github.com/kashari/draupnir/constants"

// ReplaceSmResp PDU.
type ReplaceSmResp struct {
	base
}

// NewReplaceSmResp returns ReplaceSmResp.
func NewReplaceSmResp() PDU {
	c := &ReplaceSmResp{
		base: newBase(),
	}
	c.CommandID = constants.REPLACE_SM_RESP
	return c
}

// NewReplaceSmRespFromReq returns ReplaceSmResp.
func NewReplaceSmRespFromReq(req *ReplaceSm) PDU {
	c := NewReplaceSmResp().(*ReplaceSmResp)
	if req != nil {
		c.SequenceNumber = req.SequenceNumber
	}
	return c
}

// CanResponse implements PDU interface.
func (c *ReplaceSmResp) CanResponse() bool {
	return false
}

// GetResponse implements PDU interface.
func (c *ReplaceSmResp) GetResponse() PDU {
	return nil
}

// Marshal implements PDU interface.
func (c *ReplaceSmResp) Marshal(b *ByteBuffer) {
	c.base.marshal(b, nil)
}

// Unmarshal implements PDU interface.
func (c *ReplaceSmResp) Unmarshal(b *ByteBuffer) error {
	return c.base.unmarshal(b, nil)
}
