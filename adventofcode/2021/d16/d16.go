package d16

import (
	"adventofcode2021/lib"
	"fmt"
	"math"
)

const (
	realInput = "A20D5080210CE4BB9BAFB001BD14A4574C014C004AE46A9B2E27297EECF0C013F00564776D7E3A825CAB8CD47B6C537DB99CD746674C1000D29BBC5AC80442966FB004C401F8771B61D8803D0B22E4682010EE7E59ACE5BC086003E3270AE4024E15C8010073B2FAD98E004333F9957BCB602E7024C01197AD452C01295CE2DC9934928B005DD258A6637F534CB3D89A944230043801A596B234B7E58509E88798029600BCF5B3BA114F5B3BA10C9E77BAF20FA4016FCDD13340118B929DD4FD54E60327C00BEB7002080AA850031400D002369400B10034400F30021400F20157D804AD400FE00034E000A6D001EB2004E5C00B9AE3AC3C300470029091ACADBFA048D656DFD126792187008635CD736B3231A51BA5EBDF42D4D299804F26B33C872E213C840022EC9C21FFB34EDE7C559C8964B43F8AD77570200FC66697AFEB6C757AC0179AB641E6AD9022006065CEA714A4D24C0179F8E795D3078026200FC118EB1B40010A8D11EA27100990200C45A83F12C401A8611D60A0803B1723542889537EFB24D6E0844004248B1980292D608D00423F49F9908049798B4452C0131006230C14868200FC668B50650043196A7F95569CF6B663341535DCFE919C464400A96DCE1C6B96D5EEFE60096006A400087C1E8610A4401887D1863AC99F9802DC00D34B5BCD72D6F36CB6E7D95EBC600013A88010A8271B6281803B12E124633006A2AC3A8AC600BCD07C9851008712DEAE83A802929DC51EE5EF5AE61BCD0648028596129C3B98129E5A9A329ADD62CCE0164DDF2F9343135CCE2137094A620E53FACF37299F0007392A0B2A7F0BA5F61B3349F3DFAEDE8C01797BD3F8BC48740140004322246A8A2200CC678651AA46F09AEB80191940029A9A9546E79764F7C9D608EA0174B63F815922999A84CE7F95C954D7FD9E0890047D2DC13B0042488259F4C0159922B0046565833828A00ACCD63D189D4983E800AFC955F211C700"
)

func Run()         { lib.Run(16, part1, part2) }
func part1() int64 { return sumVersions(decodePacket(fromHexString(realInput))) }
func part2() int64 { return eval(decodePacket(fromHexString(realInput))) }

func sumVersions(p *Packet) int64 {
	sum := int64(p.Version)
	for _, sp := range p.Subpackets {
		sum += sumVersions(sp)
	}
	return sum
}

func eval(p *Packet) int64 {
	switch p.TypeID {
	case SumPacket:
		return lib.Reduce(func(x int64, p *Packet) int64 { return x + eval(p) }, 0, p.Subpackets)
	case ProductPacket:
		return lib.Reduce(func(x int64, p *Packet) int64 { return x * eval(p) }, 1, p.Subpackets)
	case MinimumPacket:
		return lib.Reduce(func(x int64, p *Packet) int64 { return lib.Min(x, eval(p)) }, math.MaxInt64, p.Subpackets)
	case MaximumPacket:
		return lib.Reduce(func(x int64, p *Packet) int64 { return lib.Max(x, eval(p)) }, math.MinInt64, p.Subpackets)
	case LiteralPacket:
		return p.Literal
	case GreaterThanPacket:
		return oneIfTrue(eval(p.Subpackets[0]) > eval(p.Subpackets[1]))
	case LessThanPacket:
		return oneIfTrue(eval(p.Subpackets[0]) < eval(p.Subpackets[1]))
	case EqualToPacket:
		return oneIfTrue(eval(p.Subpackets[0]) == eval(p.Subpackets[1]))
	}
	panic("invalid packet type")
}

func oneIfTrue(b bool) int64 {
	if b {
		return 1
	}
	return 0
}

type PacketTypeID uint8

const (
	SumPacket         PacketTypeID = 0
	ProductPacket     PacketTypeID = 1
	MinimumPacket     PacketTypeID = 2
	MaximumPacket     PacketTypeID = 3
	LiteralPacket     PacketTypeID = 4
	GreaterThanPacket PacketTypeID = 5
	LessThanPacket    PacketTypeID = 6
	EqualToPacket     PacketTypeID = 7
)

type Packet struct {
	Version          uint8
	TypeID           PacketTypeID
	LengthTypeID     uint8
	SubpacketsLength int64
	NumSubpackets    int64
	Literal          int64
	Subpackets       []*Packet
}

func decodePacket(b *Bits) *Packet {
	p := &Packet{
		Version: uint8(b.getNum(3)),
		TypeID:  PacketTypeID(b.getNum(3)),
	}
	if p.TypeID == LiteralPacket {
		p.Literal = decodeLiteral(b)
	} else {
		p.LengthTypeID = uint8(b.getNum(1))
		if p.LengthTypeID == 0 {
			p.SubpacketsLength = b.getNum(15)
			startPos := b.pos
			for int64(b.pos-startPos) < p.SubpacketsLength {
				p.Subpackets = append(p.Subpackets, decodePacket(b))
			}
		} else {
			p.NumSubpackets = b.getNum(11)
			for i := int64(0); i < p.NumSubpackets; i++ {
				p.Subpackets = append(p.Subpackets, decodePacket(b))
			}
		}
	}
	return p
}

func decodeLiteral(b *Bits) int64 {
	var x int64
	for {
		prefix := b.getNum(1)
		x <<= 4
		x |= b.getNum(4)
		if prefix == 0 {
			break
		}
	}
	return x
}

type Bits struct {
	bits []byte
	pos  int
}

func fromHexString(s string) *Bits {
	var a Bits
	for i := 0; i < len(s); i++ {
		var v byte
		switch {
		case '0' <= s[i] && s[i] <= '9':
			v = s[i] - '0'
		case 'A' <= s[i] && s[i] <= 'F':
			v = s[i] - 'A' + 10
		default:
			panic(fmt.Sprintf("invalid hex digit: %c", s[i]))
		}
		a.bits = append(a.bits, v&0b1000>>3, v&0b100>>2, v&0b10>>1, v&0b1)
	}
	return &a
}

func (b Bits) length() int {
	return len(b.bits) - b.pos
}

func (b *Bits) getNum(numBits int) int64 {
	if b.length() < numBits {
		panic(fmt.Sprintf("Bits.getNum(%d): %d bit(s) remain", numBits, b.length()))
	}
	var x int64
	for i := numBits - 1; i >= 0; i-- {
		x |= int64(b.bits[i+b.pos]) << (numBits - i - 1)
	}
	b.pos += numBits
	return x
}
