package pdu

import "github.com/kashari/draupnir/constants"

// UnsuccessSme indicates submission was unsuccessful and the respective errors.
type UnsuccessSme struct {
	Address
	errorStatusCode constants.CommandStatusType
}

// NewUnsuccessSme returns new UnsuccessSme
func NewUnsuccessSme() (c UnsuccessSme) {
	c = UnsuccessSme{
		Address:         NewAddress(),
		errorStatusCode: constants.ESME_ROK,
	}
	return
}

// NewUnsuccessSmeWithAddr returns new UnsuccessSme with address.
func NewUnsuccessSmeWithAddr(addr string, status constants.CommandStatusType) (c UnsuccessSme, err error) {
	c = NewUnsuccessSme()
	if err = c.SetAddress(addr); err == nil {
		c.SetErrorStatusCode(status)
	}
	return
}

// NewUnsuccessSmeWithTonNpi create new address with ton, npi and error code.
func NewUnsuccessSmeWithTonNpi(ton, npi byte, status constants.CommandStatusType) UnsuccessSme {
	return UnsuccessSme{
		Address:         NewAddressWithTonNpi(ton, npi),
		errorStatusCode: status,
	}
}

// Unmarshal from buffer.
func (c *UnsuccessSme) Unmarshal(b *ByteBuffer) (err error) {
	var st int32
	if err = c.Address.Unmarshal(b); err == nil {
		st, err = b.ReadInt()
		if err == nil {
			c.errorStatusCode = constants.CommandStatusType(st)
		}
	}
	return
}

// Marshal to buffer.
func (c *UnsuccessSme) Marshal(b *ByteBuffer) {
	c.Address.Marshal(b)
	b.WriteInt(int32(c.errorStatusCode))
}

// SetErrorStatusCode sets error status code.
func (c *UnsuccessSme) SetErrorStatusCode(v constants.CommandStatusType) {
	c.errorStatusCode = v
}

// ErrorStatusCode returns assigned status code.
func (c *UnsuccessSme) ErrorStatusCode() constants.CommandStatusType {
	return c.errorStatusCode
}

// UnsuccessSmes represents list of UnsuccessSme.
type UnsuccessSmes struct {
	l []UnsuccessSme
}

// NewUnsuccessSmes returns list of UnsuccessSme.
func NewUnsuccessSmes() (u UnsuccessSmes) {
	u.l = make([]UnsuccessSme, 0, 8)
	return
}

// Add to list.
func (c *UnsuccessSmes) Add(us ...UnsuccessSme) {
	c.l = append(c.l, us...)
}

// Get list.
func (c *UnsuccessSmes) Get() []UnsuccessSme {
	return c.l
}

// Unmarshal from buffer.
func (c *UnsuccessSmes) Unmarshal(b *ByteBuffer) (err error) {
	var n byte
	if n, err = b.ReadByte(); err == nil {
		c.l = make([]UnsuccessSme, n)

		var i byte
		for ; i < n; i++ {
			if err = c.l[i].Unmarshal(b); err != nil {
				return
			}
		}
	}
	return
}

// Marshal to buffer.
func (c *UnsuccessSmes) Marshal(b *ByteBuffer) {
	n := byte(len(c.l))
	_ = b.WriteByte(n)

	var i byte
	for ; i < n; i++ {
		c.l[i].Marshal(b)
	}
}
