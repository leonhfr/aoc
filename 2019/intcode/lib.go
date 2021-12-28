package intcode

type Intcode []int

func New(code Intcode, noun, verb int) Intcode {
	c := make([]int, len(code))
	copy(c, code)
	c[1], c[2] = noun, verb
	return c
}

func (ic *Intcode) Process() {
	code := *ic
	for i := 0; i < len(code); i += 4 {
		switch code[i] {
		case 1:
			code[code[i+3]] = code[code[i+1]] + code[code[i+2]]
		case 2:
			code[code[i+3]] = code[code[i+1]] * code[code[i+2]]
		case 99:
			return
		}
	}
}
