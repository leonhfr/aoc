package packet

import (
	"github.com/leonhfr/aoc/2021/16/binstr"
)

type Packet struct {
	Version int
	TypeId  int
	Literal int
	Packets []Packet
}

func NewLiteral(version, typeId, value binstr.BinaryString) Packet {
	return Packet{
		version.Decimal(),
		typeId.Decimal(),
		value.Decimal(),
		[]Packet{},
	}
}

func NewOperator(version, typeId binstr.BinaryString, packets []Packet) Packet {
	return Packet{
		version.Decimal(),
		typeId.Decimal(),
		0,
		packets,
	}
}

func (p Packet) Value() int {
	switch p.TypeId {
	case 0:
		sum := 0
		for _, k := range p.Packets {
			sum += k.Value()
		}
		return sum
	case 1:
		product := 1
		for _, k := range p.Packets {
			product *= k.Value()
		}
		return product
	case 2:
		min := p.Packets[0].Value()
		for _, k := range p.Packets {
			c := k.Value()
			if c < min {
				min = c
			}
		}
		return min
	case 3:
		max := p.Packets[0].Value()
		for _, k := range p.Packets {
			c := k.Value()
			if c > max {
				max = c
			}
		}
		return max
	case 4:
		return p.Literal
	case 5:
		if p.Packets[0].Value() > p.Packets[1].Value() {
			return 1
		}
	case 6:
		if p.Packets[0].Value() < p.Packets[1].Value() {
			return 1
		}
	case 7:
		if p.Packets[0].Value() == p.Packets[1].Value() {
			return 1
		}
	}
	return 0
}

func (p Packet) VersionSum() int {
	sum := p.Version
	for _, k := range p.Packets {
		sum += k.VersionSum()
	}
	return sum
}
