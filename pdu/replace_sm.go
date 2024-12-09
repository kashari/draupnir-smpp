package pdu

import "github.com/kashari/draupnir/constants"

// ReplaceSm PDU is issued by the ESME to replace a previously submitted short message
// that is still pending delivery. The matching mechanism is based on the message_id and
// source address of the original message. Where the original submit_sm ‘source address’
// was defaulted to NULL, then the source address in the replace_sm command should also be NULL.
type ReplaceSm struct {
	base
	MessageID            string
	SourceAddr           Address
	ScheduleDeliveryTime string
	ValidityPeriod       string
	RegisteredDelivery   byte
	Message              ShortMessage
}

// NewReplaceSm returns ReplaceSm PDU.
func NewReplaceSm() PDU {
	message, _ := NewShortMessage("")
	message.withoutDataCoding = true
	c := &ReplaceSm{
		base:                 newBase(),
		SourceAddr:           NewAddress(),
		ScheduleDeliveryTime: constants.DFLT_SCHEDULE,
		ValidityPeriod:       constants.DFLT_VALIDITY,
		RegisteredDelivery:   constants.DFLT_REG_DELIVERY,
		Message:              message,
	}
	c.CommandID = constants.REPLACE_SM
	return c
}

// CanResponse implements PDU interface.
func (c *ReplaceSm) CanResponse() bool {
	return true
}

// GetResponse implements PDU interface.
func (c *ReplaceSm) GetResponse() PDU {
	return NewReplaceSmRespFromReq(c)
}

// Marshal implements PDU interface.
func (c *ReplaceSm) Marshal(b *ByteBuffer) {
	c.base.marshal(b, func(b *ByteBuffer) {
		b.Grow(len(c.MessageID) + len(c.ScheduleDeliveryTime) + len(c.ValidityPeriod) + 4)

		_ = b.WriteCString(c.MessageID)
		c.SourceAddr.Marshal(b)
		_ = b.WriteCString(c.ScheduleDeliveryTime)
		_ = b.WriteCString(c.ValidityPeriod)
		_ = b.WriteByte(c.RegisteredDelivery)
		c.Message.Marshal(b)
	})
}

// Unmarshal implements PDU interface.
func (c *ReplaceSm) Unmarshal(b *ByteBuffer) error {
	return c.base.unmarshal(b, func(b *ByteBuffer) (err error) {
		if c.MessageID, err = b.ReadCString(); err == nil {
			if err = c.SourceAddr.Unmarshal(b); err == nil {
				if c.ScheduleDeliveryTime, err = b.ReadCString(); err == nil {
					if c.ValidityPeriod, err = b.ReadCString(); err == nil {
						if c.RegisteredDelivery, err = b.ReadByte(); err == nil {
							err = c.Message.Unmarshal(b, false)
						}
					}
				}
			}
		}
		return
	})
}
