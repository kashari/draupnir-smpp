package pdu

import "github.com/kashari/draupnir/constants"

// SubmitMultiResp PDU.
type SubmitMultiResp struct {
	base
	MessageID     string
	UnsuccessSMEs UnsuccessSmes
}

// NewSubmitMultiResp returns new SubmitMultiResp.
func NewSubmitMultiResp() PDU {
	c := &SubmitMultiResp{
		base:          newBase(),
		MessageID:     constants.DFLT_MSGID,
		UnsuccessSMEs: NewUnsuccessSmes(),
	}
	c.CommandID = constants.SUBMIT_MULTI_RESP
	return c
}

// NewSubmitMultiRespFromReq returns new SubmitMultiResp.
func NewSubmitMultiRespFromReq(req *SubmitMulti) PDU {
	c := NewSubmitMultiResp().(*SubmitMultiResp)
	if req != nil {
		c.SequenceNumber = req.SequenceNumber
	}
	return c
}

// CanResponse implements PDU interface.
func (c *SubmitMultiResp) CanResponse() bool {
	return false
}

// GetResponse implements PDU interface.
func (c *SubmitMultiResp) GetResponse() PDU {
	return nil
}

// Marshal implements PDU interface.
func (c *SubmitMultiResp) Marshal(b *ByteBuffer) {
	c.base.marshal(b, func(b *ByteBuffer) {
		b.Grow(len(c.MessageID) + 1)

		_ = b.WriteCString(c.MessageID)
		c.UnsuccessSMEs.Marshal(b)
	})
}

// Unmarshal implements PDU interface.
func (c *SubmitMultiResp) Unmarshal(b *ByteBuffer) error {
	return c.base.unmarshal(b, func(b *ByteBuffer) (err error) {
		if c.MessageID, err = b.ReadCString(); err == nil {
			err = c.UnsuccessSMEs.Unmarshal(b)
		}
		return
	})
}
