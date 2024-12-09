package pdu

import "github.com/kashari/draupnir/constants"

// QuerySm PDU is issued by the ESME to query the status of a previously submitted short message.
// The matching mechanism is based on the SMSC assigned message_id and source address. Where the
// original submit_sm, data_sm or submit_multi ‘source address’ was defaulted to NULL, then the
// source address in the query_sm command should also be set to NULL.
type QuerySm struct {
	base
	MessageID  string
	SourceAddr Address
}

// NewQuerySm returns new QuerySm PDU.
func NewQuerySm() PDU {
	c := &QuerySm{
		SourceAddr: NewAddress(),
	}
	c.CommandID = constants.QUERY_SM
	return c
}

// CanResponse implements PDU interface.
func (c *QuerySm) CanResponse() bool {
	return true
}

// GetResponse implements PDU interface.
func (c *QuerySm) GetResponse() PDU {
	return NewQuerySmRespFromReq(c)
}

// Marshal implements PDU interface.
func (c *QuerySm) Marshal(b *ByteBuffer) {
	c.base.marshal(b, func(b *ByteBuffer) {
		b.Grow(len(c.MessageID) + 1)

		_ = b.WriteCString(c.MessageID)
		c.SourceAddr.Marshal(b)
	})
}

// Unmarshal implements PDU interface.
func (c *QuerySm) Unmarshal(b *ByteBuffer) error {
	return c.base.unmarshal(b, func(b *ByteBuffer) (err error) {
		if c.MessageID, err = b.ReadCString(); err == nil {
			err = c.SourceAddr.Unmarshal(b)
		}
		return
	})
}
