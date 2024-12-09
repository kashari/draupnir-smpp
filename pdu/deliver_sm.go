package pdu

import "github.com/kashari/draupnir/constants"

// DeliverSm PDU is issued by the SMSC to send a message to an ESME.
// Using this command, the SMSC may route a short message to the ESME for delivery.
type DeliverSm struct {
	base
	ServiceType          string
	SourceAddr           Address
	DestAddr             Address
	EsmClass             byte
	ProtocolID           byte
	PriorityFlag         byte
	ScheduleDeliveryTime string // not used
	ValidityPeriod       string // not used
	RegisteredDelivery   byte
	ReplaceIfPresentFlag byte // not used
	Message              ShortMessage
}

// NewDeliverSm returns DeliverSm PDU.
func NewDeliverSm() PDU {
	message, _ := NewShortMessage("")
	c := &DeliverSm{
		base:                 newBase(),
		ServiceType:          constants.DFLT_SRVTYPE,
		SourceAddr:           NewAddress(),
		DestAddr:             NewAddress(),
		EsmClass:             constants.DFLT_ESM_CLASS,
		ProtocolID:           constants.DFLT_PROTOCOLID,
		PriorityFlag:         constants.DFLT_PRIORITY_FLAG,
		ScheduleDeliveryTime: constants.DFLT_SCHEDULE,
		ValidityPeriod:       constants.DFLT_VALIDITY,
		RegisteredDelivery:   constants.DFLT_REG_DELIVERY,
		ReplaceIfPresentFlag: constants.DFTL_REPLACE_IFP,
		Message:              message,
	}
	c.CommandID = constants.DELIVER_SM
	return c
}

// CanResponse implements PDU interface.
func (c *DeliverSm) CanResponse() bool {
	return true
}

// GetResponse implements PDU interface.
func (c *DeliverSm) GetResponse() PDU {
	return NewDeliverSmRespFromReq(c)
}

// Marshal implements PDU interface.
func (c *DeliverSm) Marshal(b *ByteBuffer) {
	c.base.marshal(b, func(b *ByteBuffer) {
		b.Grow(len(c.ServiceType) + len(c.ScheduleDeliveryTime) + len(c.ValidityPeriod) + 10)

		_ = b.WriteCString(c.ServiceType)
		c.SourceAddr.Marshal(b)
		c.DestAddr.Marshal(b)
		_ = b.WriteByte(c.EsmClass)
		_ = b.WriteByte(c.ProtocolID)
		_ = b.WriteByte(c.PriorityFlag)
		_ = b.WriteCString(c.ScheduleDeliveryTime)
		_ = b.WriteCString(c.ValidityPeriod)
		_ = b.WriteByte(c.RegisteredDelivery)
		_ = b.WriteByte(c.ReplaceIfPresentFlag)
		c.Message.Marshal(b)
	})
}

// Unmarshal implements PDU interface.
func (c *DeliverSm) Unmarshal(b *ByteBuffer) error {
	return c.base.unmarshal(b, func(b *ByteBuffer) (err error) {
		if c.ServiceType, err = b.ReadCString(); err == nil {
			if err = c.SourceAddr.Unmarshal(b); err == nil {
				if err = c.DestAddr.Unmarshal(b); err == nil {
					if c.EsmClass, err = b.ReadByte(); err == nil {
						if c.ProtocolID, err = b.ReadByte(); err == nil {
							if c.PriorityFlag, err = b.ReadByte(); err == nil {
								if c.ScheduleDeliveryTime, err = b.ReadCString(); err == nil {
									if c.ValidityPeriod, err = b.ReadCString(); err == nil {
										if c.RegisteredDelivery, err = b.ReadByte(); err == nil {
											if c.ReplaceIfPresentFlag, err = b.ReadByte(); err == nil {
												err = c.Message.Unmarshal(b, (c.EsmClass&constants.SM_UDH_GSM) > 0)
											}
										}
									}
								}
							}
						}
					}
				}
			}
		}
		return
	})
}
