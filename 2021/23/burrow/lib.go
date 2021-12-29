package burrow

import (
	"sort"
	"strings"

	sh "github.com/leonhfr/aoc/shared"
)

const A = int('A')

// State represents one state of the amphipods migration
type State struct {
	Energy  int
	hallway []int // 0..10, 11 positions
	rooms   []int // 0..15, room1=0..3 bottom to top
}

// Step is the signature of the functions that determines
// all possible next states for a given state
type Step func(State) []State

// New creates a new State from a multiline string diagram
func New(diagram string) State {
	lines := strings.Split(diagram, "\n")
	hallway, rooms := make([]int, 11), make([]int, 16)
	for i := 0; i < 11; i++ {
		v := int(lines[1][i+1])
		if isAmphipod(v) {
			hallway[i] = v
		}
	}
	for i := 0; i < 4; i++ {
		x := 2*i + 3
		for j := 0; j < 4; j++ {
			v := int(lines[5-j][x])
			if isAmphipod(v) {
				rooms[4*i+j] = v
			}
		}
	}
	return State{0, hallway, rooms}
}

// Copy creates a copy of a State
func Copy(s State) State {
	hallway, rooms := make([]int, 11), make([]int, 16)
	copy(hallway, s.hallway)
	copy(rooms, s.rooms)
	return State{s.Energy, hallway, rooms}
}

// Done determines whether the state represents a finished migration
func (s State) Done() bool {
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			if s.rooms[4*i+j] != A+i {
				return false
			}
		}
	}
	return true
}

func (s State) Next() (next []State) {
	steps := []Step{lowToTop, topToLow, moveFromHallway, moveToHallway}
	for _, step := range steps {
		next = append(next, step(s)...)
		states := step(s)
		if len(states) > 0 {
			return states
		}
	}
	return nil
}

// lowToTop moves amphipods from lower spaces to the top of the room
func lowToTop(s State) (next []State) {
	// If:
	//    is an amphipod
	// && (not in target room || in target room && blocking spaces below not in target room)
	// && all spaces on top are empty
	// => move from highest occupied space below room top to room top
	for i := 0; i < 4; i++ {
		for j := 0; j < 3; j++ {
			value := s.rooms[4*i+j]
			target := value == A+i
			blocking := j > 0 && !allOf(s.rooms[4*i:4*i+j], A+i)
			empty := allOf(s.rooms[4*i+j+1:4*(i+1)], 0)
			if isAmphipod(value) && (!target || target && blocking) && empty {
				c := Copy(s)
				c.rooms[4*i+j], c.rooms[4*(i+1)-1] = c.rooms[4*(i+1)-1], c.rooms[4*i+j]
				c.Energy += energy(value, 4-j-1)
				next = append(next, c)
			}
		}
	}
	return
}

// topToLow moves amphipods from the top of the room to lower spaces
func topToLow(s State) (next []State) {
	// If:
	//    is an amphipod
	// && in target room
	// && target room contains only target amphipods
	// && spaces below are empty
	// => move from room top to highest empty space
	for i := 0; i < 4; i++ {
		value := s.rooms[4*(i+1)-1]
		target := value == A+i
		if isAmphipod(value) && target {
			for j := 0; j < 3; j++ {
				empty := allOf(s.rooms[4*i+j:4*i+3], 0)
				blocking := !allOf(s.rooms[4*i:4*i+3], value, 0)
				if empty && !blocking {
					c := Copy(s)
					c.rooms[4*i+j], c.rooms[4*i+3] = c.rooms[4*i+3], c.rooms[4*i+j]
					c.Energy += energy(value, 3-j)
					next = append(next, c)
					break
				}
			}
		}
	}
	return
}

// moveToHallway moves amphipods from room tops to reachable and allowed spaces in the hallway
func moveToHallway(s State) (next []State) {
	// If:
	//    is an amphipod && not target room
	// || is an amphipod && target room && room unsolved
	// => move to all reachable and allowed spaces in the hallway
	for i := 0; i < 4; i++ {
		value := s.rooms[4*(i+1)-1]
		target := value == A+i
		unsolved := !allOf(s.rooms[4*i:4*(i+1)-1], value, 0)
		if isAmphipod(value) && (!target || target && unsolved) {
			reachable := reachableToHallway(s.hallway, i)
			for _, r := range reachable {
				c := Copy(s)
				c.rooms[4*(i+1)-1], c.hallway[r] = 0, value
				c.Energy += energy(value, sh.Abs(2*i+2-r)+1)
				next = append(next, c)
			}
		}
	}
	return
}

// moveFromHallway moves amphipods from the hallway to reachable room tops
func moveFromHallway(s State) (next []State) {
	// If:
	//    is an amphipod
	// && target room top free
	// && hallway path clear to target room
	// && room solved
	// => move to room top
	// fmt.Println(s)
	for i := 0; i < 11; i++ {
		value := s.hallway[i]
		if isAmphipod(value) {
			room := value - A
			free := s.rooms[4*(room+1)-1] == 0
			clearLeft := i < 2*room+2 && allOf(s.hallway[i+1:2*room+3], 0)
			clearRight := i > 2*room+2 && allOf(s.hallway[2*room+2:i], 0)
			solved := allOf(s.rooms[4*room:4*(room+1)], value, 0)
			if free && (clearLeft || clearRight) && solved {
				c := Copy(s)
				c.hallway[i], c.rooms[4*(room+1)-1] = 0, value
				c.Energy += energy(value, sh.Abs(2*room+2-i)+1)
				next = append(next, c)
			}
		}
	}
	// fmt.Println(next)
	return
}

func reachableToHallway(hallway []int, room int) (spaces []int) {
	for i := 2*room + 2; i >= 0 && hallway[i] == 0; i-- {
		if !shyness(i) && hallway[i] == 0 {
			spaces = append(spaces, i)
		}
	}
	for i := 2*room + 2; i <= 10 && hallway[i] == 0; i++ {
		if !shyness(i) && hallway[i] == 0 {
			spaces = append(spaces, i)
		}
	}
	sort.Ints(spaces)
	return
}

func shyness(i int) bool {
	return i == 2 || i == 4 || i == 6 || i == 8
}

// energy computes the energy spent by an amphipod pod to cover the distance d
func energy(pod, d int) int {
	return d * pow10(pod-A)
}

// isAmphipod checks that the given value is a valid amphipod
func isAmphipod(a int) bool {
	return A <= a && a <= A+3
}

// allOf checks whether the array contains only one of values
func allOf(array []int, values ...int) bool {
	for _, v := range array {
		if !oneOf(v, values...) {
			return false
		}
	}
	return true
}

// oneOf checks whether the value is one of the values
func oneOf(value int, values ...int) bool {
	for _, v := range values {
		if v == value {
			return true
		}
	}
	return false
}

// pow10 raises 10 to the power of n
func pow10(n int) int {
	r := 1
	for i := 0; i < n; i++ {
		r *= 10
	}
	return r
}
