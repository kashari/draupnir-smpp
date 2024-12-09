package pdu

import (
	"fmt"

	"github.com/kashari/draupnir/constants"
)

type Address struct {
	addr_ton byte
	addr_npi byte
	address string
}

func NewAddress() Address {
	return Address{addr_ton: constants.GetDefaultTon(), addr_npi: constants.GetDefaultNpi()}
}

// NewAddressWithAddr returns new address.
func NewAddressWithAddr(addr string) (a Address, err error) {
	a = NewAddress()
	err = a.SetAddress(addr)
	return
}

// NewAddressWithTonNpi returns new address with ton, npi.
func NewAddressWithTonNpi(ton, npi byte) Address {
	return Address{addr_ton: ton, addr_npi: npi}
}

// NewAddressWithTonNpiAddr returns new address with ton, npi, addr string.
func NewAddressWithTonNpiAddr(ton, npi byte, addr string) (a Address, err error) {
	a = NewAddressWithTonNpi(ton, npi)
	err = a.SetAddress(addr)
	return
}

// Unmarshal from buffer.
func (c *Address) Unmarshal(b *ByteBuffer) (err error) {
	if c.addr_ton, err = b.ReadByte(); err == nil {
		if c.addr_npi, err = b.ReadByte(); err == nil {
			c.address, err = b.ReadCString()
		}
	}
	return
}

// Marshal to buffer.
func (c *Address) Marshal(b *ByteBuffer) {
	b.Grow(3 + len(c.address))

	_ = b.WriteByte(c.addr_ton)
	_ = b.WriteByte(c.addr_npi)
	_ = b.WriteCString(c.address)
}

// SetTon sets ton.
func (c *Address) SetTon(ton byte) {
	c.addr_ton = ton
}

// SetNpi sets npi.
func (c *Address) SetNpi(npi byte) {
	c.addr_npi = npi
}

// SetAddress sets address.
func (c *Address) SetAddress(addr string) (err error) {
	if len(addr) > constants.SM_ADDR_LEN {
		err = fmt.Errorf("address length exceeds limit. (%d > %d)", len(addr), constants.SM_ADDR_LEN)
	} else {
		c.address = addr
	}
	return
}

// Ton returns assigned ton.
func (c Address) Ton() byte {
	return c.addr_ton
}

// Npi returns assigned npi.
func (c Address) Npi() byte {
	return c.addr_npi
}

// Address returns assigned address (in string).
func (c Address) Address() string {
	return c.address
}

// String implement stringer interface
func (c Address) String() string {
	return c.address
}

