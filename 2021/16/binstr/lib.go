package binstr

import (
	"fmt"
	"strconv"
)

type BinaryString string

func New(hex string) BinaryString {
	var bs string
	for i := range hex {
		unsigned, err := strconv.ParseUint(hex[i:i+1], 16, 4)
		if err != nil {
			panic(err)
		}
		bs += fmt.Sprintf("%04b", unsigned)
	}
	return BinaryString(bs)
}

func (bs BinaryString) Decimal() int {
	decimal, err := strconv.ParseInt(string(bs), 2, 64)
	if err != nil {
		panic(err)
	}
	return int(decimal)
}

func (bs *BinaryString) Eat(n int) (chunk BinaryString) {
	if n > len(*bs) {
		chunk, *bs = *bs, ""
		return
	}

	s := *bs
	chunk, s = s[:n], s[n:]
	*bs = s
	return
}

func (bs BinaryString) Peek(i, n int) BinaryString {
	return bs[i : i+n]
}
