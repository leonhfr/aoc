package main

import (
	_ "embed"
	"fmt"

	"github.com/leonhfr/aoc/2021/16/binstr"
	"github.com/leonhfr/aoc/2021/16/packet"
	sh "github.com/leonhfr/aoc/shared"
)

//go:embed input
var input string

var parsed packet.Packet

func main() {
	fmt.Printf("Part 1: %v\n", part1())
	fmt.Printf("Part 2: %v\n", part2())
}

func part1() int {
	return parsed.VersionSum()
}

func part2() int {
	return parsed.Value()
}

func init() {
	lines := sh.Lines(input)
	binaryString := binstr.New(lines[0])
	parsed, _ = parse(binaryString)
}

func parse(bs binstr.BinaryString) (packet.Packet, binstr.BinaryString) {
	if bs.Peek(3, 3).Decimal() == 4 {
		return parseLiteral(bs)
	}

	return parseOperator(bs)
}

func parseLiteral(bs binstr.BinaryString) (packet.Packet, binstr.BinaryString) {
	var value binstr.BinaryString
	version := bs.Eat(3)
	typeId := bs.Eat(3)
	for end := false; !end; {
		prefix := bs.Eat(1)
		part := bs.Eat(4)
		value += part
		end = prefix == "0"
	}
	return packet.NewLiteral(version, typeId, value), bs
}

func parseOperator(bs binstr.BinaryString) (packet.Packet, binstr.BinaryString) {
	var packets []packet.Packet
	version := bs.Eat(3)
	typeId := bs.Eat(3)
	lengthType := bs.Eat(1)
	if lengthType == "0" {
		length := bs.Eat(15).Decimal()
		packets, bs = parsePacketsByLength(bs, length)
	} else {
		length := bs.Eat(11).Decimal()
		packets, bs = parsePacketsByN(bs, length)
	}
	return packet.NewOperator(version, typeId, packets), bs
}

func parsePacketsByLength(bs binstr.BinaryString, length int) ([]packet.Packet, binstr.BinaryString) {
	var packets []packet.Packet
	for length > 0 {
		packet, b := parse(bs)
		length -= len(bs) - len(b)
		packets, bs = append(packets, packet), b
	}
	return packets, bs
}

func parsePacketsByN(bs binstr.BinaryString, n int) ([]packet.Packet, binstr.BinaryString) {
	var packets []packet.Packet
	for i := 0; i < n; i++ {
		packet, b := parse(bs)
		packets, bs = append(packets, packet), b
	}
	return packets, bs
}
