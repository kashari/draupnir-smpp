package pdu

import "github.com/kashari/draupnir/constants"

// SubmitSM PDU is used by an ESME to submit a short message to the SMSC for onward
// transmission to a specified short message entity (SME). The submit_sm PDU does
// not support the transaction message mode.
type SubmitSm struct {
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

// NewSubmitSM returns SubmitSM PDU.
func NewSubmitSm() PDU {
	message, _ := NewShortMessage("")
	c := &SubmitSm{
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
	c.CommandID = constants.SUBMIT_SM
	return c
}

// ShouldSplit check if this the user data of submitSM PDU
func (c *SubmitSm) ShouldSplit() bool {
	// GSM standard mandates that User Data must be no longer than 140 octet
	return len(c.Message.messageData) > constants.SM_GSM_MSG_LEN
}

// CanResponse implements PDU interface.
func (c *SubmitSm) CanResponse() bool {
	return true
}

// GetResponse implements PDU interface.
func (c *SubmitSm) GetResponse() PDU {
	return NewSubmitSmRespFromReq(c)
}

// Split split a single long text message into multiple SubmitSM PDU,
// Each have the TPUD within the GSM's User Data limit of 140 octet
// If the message is short enough and doesn't need splitting,
// Split() returns an array of length 1
func (c *SubmitSm) Split() (multiSubSM []*SubmitSm, err error) {
	multiSubSM = []*SubmitSm{}

	multiMsg, err := c.Message.split()
	if err != nil {
		return
	}

	esmClass := c.EsmClass // no need to "or" with SM_UDH_GSM when a message has a single part
	if len(multiMsg) > 1 {
		esmClass = c.EsmClass | constants.SM_UDH_GSM // must set to indicate UDH
	}

	for _, msg := range multiMsg {
		multiSubSM = append(multiSubSM, &SubmitSm{
			base:                 c.base,
			ServiceType:          c.ServiceType,
			SourceAddr:           c.SourceAddr,
			DestAddr:             c.DestAddr,
			EsmClass:             esmClass,
			ProtocolID:           c.ProtocolID,
			PriorityFlag:         c.PriorityFlag,
			ScheduleDeliveryTime: c.ScheduleDeliveryTime,
			ValidityPeriod:       c.ValidityPeriod,
			RegisteredDelivery:   c.RegisteredDelivery,
			ReplaceIfPresentFlag: c.ReplaceIfPresentFlag,
			Message:              *msg,
		})
	}
	return
}

// Marshal implements PDU interface.
func (c *SubmitSm) Marshal(b *ByteBuffer) {
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
func (c *SubmitSm) Unmarshal(b *ByteBuffer) error {
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
