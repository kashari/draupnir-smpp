package pdu

import (
	"errors"
	"io"

	"github.com/kashari/draupnir/constants"
)

// SubmitSmResp PDU.
type SubmitSmResp struct {
	base
	MessageID string
}

// NewSubmitSmResp returns new SubmitSmResp.
func NewSubmitSmResp() PDU {
	c := &SubmitSmResp{
		base:      newBase(),
		MessageID: constants.DFLT_MSGID,
	}
	c.CommandID = constants.SUBMIT_SM_RESP
	return c
}

// NewSubmitSmRespFromReq returns new SubmitSmResp.
func NewSubmitSmRespFromReq(req *SubmitSm) PDU {
	c := NewSubmitSmResp().(*SubmitSmResp)
	if req != nil {
		c.SequenceNumber = req.SequenceNumber
	}
	return c
}

// CanResponse implements PDU interface.
func (c *SubmitSmResp) CanResponse() bool {
	return false
}

// GetResponse implements PDU interface.
func (c *SubmitSmResp) GetResponse() PDU {
	return nil
}

// Marshal implements PDU interface.
func (c *SubmitSmResp) Marshal(b *ByteBuffer) {
	c.base.marshal(b, func(b *ByteBuffer) {
		b.Grow(len(c.MessageID) + 1)

		_ = b.WriteCString(c.MessageID)
	})
}

// Unmarshal implements PDU interface.
func (c *SubmitSmResp) Unmarshal(b *ByteBuffer) error {
	return c.base.unmarshal(b, func(b *ByteBuffer) (err error) {
		c.MessageID, err = b.ReadCString()
		if errors.Is(err, io.EOF) {
			return nil
		}
		return
	})
}
