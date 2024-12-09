package pdu

import "github.com/kashari/draupnir/constants"

// QuerySmResp PDU.
type QuerySmResp struct {
	base
	MessageID    string
	FinalDate    string
	MessageState byte
	ErrorCode    byte
}

// NewQuerySmResp returns new QuerySm PDU.
func NewQuerySmResp() PDU {
	c := &QuerySmResp{
		base:         newBase(),
		FinalDate:    constants.DFLT_DATE,
		MessageState: constants.DFLT_MSG_STATE,
		ErrorCode:    constants.DFLT_ERR,
	}
	c.CommandID = constants.QUERY_SM_RESP
	return c
}

// NewQuerySmRespFromReq returns new QuerySm PDU.
func NewQuerySmRespFromReq(req *QuerySm) PDU {
	c := NewQuerySmResp().(*QuerySmResp)
	if req != nil {
		c.SequenceNumber = req.SequenceNumber
	}
	return c
}

// CanResponse implements PDU interface.
func (c *QuerySmResp) CanResponse() bool {
	return false
}

// GetResponse implements PDU interface.
func (c *QuerySmResp) GetResponse() PDU {
	return nil
}

// Marshal implements PDU interface.
func (c *QuerySmResp) Marshal(b *ByteBuffer) {
	c.base.marshal(b, func(b *ByteBuffer) {
		b.Grow(len(c.MessageID) + len(c.FinalDate) + 4)

		_ = b.WriteCString(c.MessageID)
		_ = b.WriteCString(c.FinalDate)
		_ = b.WriteByte(c.MessageState)
		_ = b.WriteByte(c.ErrorCode)
	})
}

// Unmarshal implements PDU interface.
func (c *QuerySmResp) Unmarshal(b *ByteBuffer) error {
	return c.base.unmarshal(b, func(b *ByteBuffer) (err error) {
		if c.MessageID, err = b.ReadCString(); err == nil {
			if c.FinalDate, err = b.ReadCString(); err == nil {
				if c.MessageState, err = b.ReadByte(); err == nil {
					c.ErrorCode, err = b.ReadByte()
				}
			}
		}
		return
	})
}
