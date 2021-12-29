package burrow

import (
	"fmt"
	"testing"
)

var test_start = `#############
#...........#
###B#C#B#D###
  #D#C#B#A#
  #D#B#A#C#
  #A#D#C#A#
  #########`

var test_end = `#############
#...........#
###A#B#C#D###
  #A#B#C#D#
  #A#B#C#D#
  #A#B#C#D#
  #########`

var empty_hallway = []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

func Test_New(t *testing.T) {
	tests := []struct {
		name string
		args string
		want State
	}{
		{
			"Start", test_start, State{
				0,
				[]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
				[]int{65, 68, 68, 66, 68, 66, 67, 67, 67, 65, 66, 66, 65, 67, 65, 68},
			},
		},
		{
			"End", test_end, State{
				0,
				[]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
				[]int{65, 65, 65, 65, 66, 66, 66, 66, 67, 67, 67, 67, 68, 68, 68, 68},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := New(tt.args)
			if ok, str := checkState(got, tt.want); !ok {
				t.Errorf(str)
			}
		})
	}
}

func Test_Done(t *testing.T) {
	tests := []struct {
		name string
		args string
		want bool
	}{
		{"Start", test_start, false},
		{"End", test_end, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(tt.args).Done(); got != tt.want {
				t.Errorf("state.Won() = %v", got)
			}
		})
	}
}

func Test_lowToTop(t *testing.T) {
	a := `#############
#...........#
###A#.#.#.###
  #A#.#B#D#
  #A#C#A#A#
  #A#C#C#D#
  #########`

	tests := []struct {
		name string
		args State
		want []State
	}{
		{"a", New(a), []State{
			{200, empty_hallway, []int{65, 65, 65, 65, 67, 0, 0, 67, 67, 65, 66, 0, 68, 65, 68, 0}},
			{10, empty_hallway, []int{65, 65, 65, 65, 67, 67, 0, 0, 67, 65, 0, 66, 68, 65, 68, 0}},
			{1000, empty_hallway, []int{65, 65, 65, 65, 67, 67, 0, 0, 67, 65, 66, 0, 68, 65, 0, 68}},
		}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := lowToTop(tt.args)
			if ok, str := checkStates(got, tt.want); !ok {
				t.Errorf(str)
			}
		})
	}
}

func Test_topToLow(t *testing.T) {
	a := `#############
#...........#
###A#B#C#.###
  #A#.#.#D#
  #A#.#.#A#
  #A#C#C#D#
  #########`

	tests := []struct {
		name string
		args State
		want []State
	}{
		{"a", New(a), []State{
			{200, empty_hallway, []int{65, 65, 65, 65, 67, 0, 0, 66, 67, 67, 0, 0, 68, 65, 68, 0}},
		}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := topToLow(tt.args)
			if ok, str := checkStates(got, tt.want); !ok {
				t.Errorf(str)
			}
		})
	}
}

func Test_moveToHallway(t *testing.T) {
	a := `#############
#.......C...#
###A#B#C#C###
  #A#.#.#D#
  #A#.#.#A#
  #A#C#C#D#
  #########`

	tests := []struct {
		name string
		args State
		want []State
	}{
		{"a", New(a), []State{
			{50, []int{66, 0, 0, 0, 0, 0, 0, 67, 0, 0, 0}, []int{65, 65, 65, 65, 67, 0, 0, 0, 67, 0, 0, 67, 68, 65, 68, 67}},
			{40, []int{0, 66, 0, 0, 0, 0, 0, 67, 0, 0, 0}, []int{65, 65, 65, 65, 67, 0, 0, 0, 67, 0, 0, 67, 68, 65, 68, 67}},
			{20, []int{0, 0, 0, 66, 0, 0, 0, 67, 0, 0, 0}, []int{65, 65, 65, 65, 67, 0, 0, 0, 67, 0, 0, 67, 68, 65, 68, 67}},
			{20, []int{0, 0, 0, 0, 0, 66, 0, 67, 0, 0, 0}, []int{65, 65, 65, 65, 67, 0, 0, 0, 67, 0, 0, 67, 68, 65, 68, 67}},
			{200, []int{0, 0, 0, 0, 0, 0, 0, 67, 0, 67, 0}, []int{65, 65, 65, 65, 67, 0, 0, 66, 67, 0, 0, 67, 68, 65, 68, 0}},
			{300, []int{0, 0, 0, 0, 0, 0, 0, 67, 0, 0, 67}, []int{65, 65, 65, 65, 67, 0, 0, 66, 67, 0, 0, 67, 68, 65, 68, 0}},
		}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := moveToHallway(tt.args)
			if ok, str := checkStates(got, tt.want); !ok {
				t.Errorf(str)
			}
		})
	}
}

func Test_moveFromHallway(t *testing.T) {
	a := `#############
#A..B......C#
###.#.#.#C###
  #A#.#.#D#
  #A#B#.#A#
  #A#C#C#D#
  #########`

	tests := []struct {
		name string
		args State
		want []State
	}{
		{"a", New(a), []State{
			{3, []int{0, 0, 0, 66, 0, 0, 0, 0, 0, 0, 67}, []int{65, 65, 65, 65, 67, 66, 0, 0, 67, 0, 0, 0, 68, 65, 68, 67}},
			{500, []int{65, 0, 0, 66, 0, 0, 0, 0, 0, 0, 0}, []int{65, 65, 65, 0, 67, 66, 0, 0, 67, 0, 0, 67, 68, 65, 68, 67}},
		}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := moveFromHallway(tt.args)
			if ok, str := checkStates(got, tt.want); !ok {
				t.Errorf(str)
			}
		})
	}
}

func Test_reachableToHallway(t *testing.T) {
	type arguments struct {
		hallway []int
		room    int
	}
	tests := []struct {
		name string
		args arguments
		want []int
	}{
		{"A", arguments{[]int{65, 0, 0, 0, 0, 0, 0, 68, 0, 0, 0}, 0}, []int{1, 3, 5}},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprint(tt.args), func(t *testing.T) {
			got := reachableToHallway(tt.args.hallway, tt.args.room)
			if ok, str := checkSlice(got, tt.want); !ok {
				t.Errorf(str)
			}
		})
	}
}

func Test_allOf(t *testing.T) {
	type arguments struct {
		a []int
		v int
	}
	tests := []struct {
		name int
		args arguments
		want bool
	}{
		{1, arguments{[]int{65, 65, 65}, 65}, true},
		{2, arguments{[]int{65, 68, 65}, 65}, false},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprint(tt.args), func(t *testing.T) {
			if got := allOf(tt.args.a, tt.args.v); got != tt.want {
				t.Errorf("allOf(%v,%v) = %v", tt.args.a, tt.args.v, got)
			}
		})
	}
}

func Test_pow10(t *testing.T) {
	tests := []struct {
		args int
		want int
	}{
		{0, 1},
		{1, 10},
		{2, 100},
		{3, 1000},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprint(tt.args), func(t *testing.T) {
			if got := pow10(tt.args); got != tt.want {
				t.Errorf("pow10(%v) = %v", tt.args, got)
			}
		})
	}
}

func checkStates(a, b []State) (bool, string) {
	if len(a) != len(b) {
		return false, fmt.Sprintf("len(a)=%v != len(b)=%v", len(a), len(b))
	}

	for i := 0; i < len(a); i++ {
		if ok, str := checkState(a[i], b[i]); !ok {
			return false, str
		}
	}

	return true, ""
}

func checkState(a, b State) (bool, string) {
	if a.Energy != b.Energy {
		return false, fmt.Sprintf("a.energy=%v != b.energy=%v", a.Energy, b.Energy)
	}

	for i := 0; i < 11; i++ {
		if a.hallway[i] != b.hallway[i] {
			return false, fmt.Sprintf(
				"a.hallway[%v]=%v != b.hallway[%v]=%v",
				i, a.hallway[i],
				i, b.hallway[i],
			)
		}
	}

	for i := 0; i < 16; i++ {
		if a.rooms[i] != b.rooms[i] {
			return false, fmt.Sprintf(
				"a.rooms[%v]=%v != a.rooms[%v]=%v",
				i, a.rooms[i],
				i, b.rooms[i],
			)
		}
	}

	return true, ""
}

func checkSlice(a, b []int) (bool, string) {
	if len(a) != len(b) {
		return false, fmt.Sprintf("len(a)=%v != len(b)=%v", len(a), len(b))
	}

	for i := 0; i < len(a); i++ {
		if ok := a[i] == b[i]; !ok {
			return false, fmt.Sprintf(
				"a[%v]=%v != b[%v]=%v",
				i, a[i], i, b[i],
			)
		}
	}

	return true, ""
}
