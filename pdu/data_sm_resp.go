package pdu

import "github.com/kashari/draupnir/constants"

// DataSmResp PDU.
type DataSmResp struct {
	base
	MessageID string
}

// NewDataSmResp returns DataSmResp.
func NewDataSmResp() PDU {
	c := &DataSmResp{
		base:      newBase(),
		MessageID: constants.DFLT_MSGID,
	}
	c.CommandID = constants.DATA_SM_RESP
	return c
}

// NewDataSmRespFromReq returns DataSmResp.
func NewDataSmRespFromReq(req *DataSm) PDU {
	c := NewDataSmResp().(*DataSmResp)
	if req != nil {
		c.SequenceNumber = req.SequenceNumber
	}
	return c
}

// CanResponse implements PDU interface.
func (c *DataSmResp) CanResponse() bool {
	return false
}

// GetResponse implements PDU interface.
func (c *DataSmResp) GetResponse() PDU {
	return nil
}

// Marshal implements PDU interface.
func (c *DataSmResp) Marshal(b *ByteBuffer) {
	c.base.marshal(b, func(b *ByteBuffer) {
		b.Grow(len(c.MessageID) + 1)

		_ = b.WriteCString(c.MessageID)
	})
}

// Unmarshal implements PDU interface.
func (c *DataSmResp) Unmarshal(b *ByteBuffer) error {
	return c.base.unmarshal(b, func(b *ByteBuffer) (err error) {
		c.MessageID, err = b.ReadCString()
		return
	})
}
