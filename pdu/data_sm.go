package pdu

import "github.com/kashari/draupnir/constants"

// DataSm PDU is used to transfer data between the SMSC and the ESME.
// It may be used by both the ESME and SMSC.
type DataSm struct {
	base
	ServiceType        string
	SourceAddr         Address
	DestAddr           Address
	EsmClass           byte
	RegisteredDelivery byte
	DataCoding         byte
}

// NewDataSm returns new data sm pdu.
func NewDataSm() PDU {
	c := &DataSm{
		base:               newBase(),
		ServiceType:        constants.DFLT_SRVTYPE,
		SourceAddr:         NewAddress(),
		DestAddr:           NewAddress(),
		EsmClass:           constants.DFLT_ESM_CLASS,
		RegisteredDelivery: constants.DFLT_REG_DELIVERY,
		DataCoding:         constants.DFLT_DATA_CODING,
	}
	c.CommandID = constants.DATA_SM
	return c
}

// CanResponse implements PDU interface.
func (c *DataSm) CanResponse() bool {
	return true
}

// GetResponse implements PDU interface.
func (c *DataSm) GetResponse() PDU {
	return NewDataSmRespFromReq(c)
}

// Marshal implements PDU interface.
func (c *DataSm) Marshal(b *ByteBuffer) {
	c.base.marshal(b, func(b *ByteBuffer) {
		b.Grow(len(c.ServiceType) + 4)

		_ = b.WriteCString(c.ServiceType)
		c.SourceAddr.Marshal(b)
		c.DestAddr.Marshal(b)
		_ = b.WriteByte(c.EsmClass)
		_ = b.WriteByte(c.RegisteredDelivery)
		_ = b.WriteByte(c.DataCoding)
	})
}

// Unmarshal implements PDU interface.
func (c *DataSm) Unmarshal(b *ByteBuffer) error {
	return c.base.unmarshal(b, func(b *ByteBuffer) (err error) {
		if c.ServiceType, err = b.ReadCString(); err == nil {
			if err = c.SourceAddr.Unmarshal(b); err == nil {
				if err = c.DestAddr.Unmarshal(b); err == nil {
					if c.EsmClass, err = b.ReadByte(); err == nil {
						if c.RegisteredDelivery, err = b.ReadByte(); err == nil {
							c.DataCoding, err = b.ReadByte()
						}
					}
				}
			}
		}
		return
	})
}
