package intcode

import (
	_ "embed"
	"strings"
	"testing"

	sh "github.com/leonhfr/aoc/shared"
)

//go:embed sample_02
var sample string

func TestProcess(t *testing.T) {
	want := 3500
	code := Intcode(sh.ToInts(strings.Split(sh.Lines(sample)[0], ",")))
	code.Process()
	if got := code[0]; got != want {
		t.Errorf("Process: got %v, want %v", got, want)
	}
}
