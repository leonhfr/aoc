package main

import (
	_ "embed"
	"fmt"
	"regexp"
	"sort"

	sh "github.com/leonhfr/aoc/shared"
	set "github.com/leonhfr/aoc/shared/setpoint"
)

//go:embed input
var input string

type (
	reading  struct{ sensor, beacon set.Point }
	interval [2]int
	quad     struct{ topLeft, bottomRight set.Point }
)

var readings []reading

func main() {
	fmt.Printf("Part 1: %v\n", part1())
	fmt.Printf("Part 2: %v\n", part2())
}

func part1() int {
	line := 2_000_000
	var intervals []interval
	beacons := set.New()
	for _, r := range readings {
		if r.beacon.Y == line {
			beacons.Add(r.beacon)
		}

		distance := set.ManhattanDistance(r.sensor, r.beacon)
		deltaY := sh.Abs(line - r.sensor.Y)
		if deltaY < distance {
			deltaX := distance - deltaY
			intervals = append(intervals, [2]int{
				r.sensor.X - deltaX,
				r.sensor.X + deltaX,
			})
		}
	}

	return count(merge(intervals)) - beacons.Len()
}

func part2() int {
	max := 4_000_000
	var q quad
	root := quad{set.Point{X: 0, Y: 0}, set.Point{X: max, Y: max}}
	queue := []quad{root}

QUAD:
	for len(queue) > 0 {
		q, queue = queue[0], queue[1:]

		for _, r := range readings {
			d0 := set.ManhattanDistance(r.sensor, r.beacon)
			d1 := set.ManhattanDistance(r.sensor, q.topLeft)
			d2 := set.ManhattanDistance(r.sensor, set.Point{X: q.bottomRight.X, Y: q.topLeft.Y})
			d3 := set.ManhattanDistance(r.sensor, q.bottomRight)
			d4 := set.ManhattanDistance(r.sensor, set.Point{X: q.topLeft.X, Y: q.bottomRight.Y})

			// quad fully inside sensor field
			if sh.Max(d1, d2, d3, d4) <= d0 {
				goto QUAD
			}
		}

		dx := q.topLeft.X + (q.bottomRight.X-q.topLeft.X)/2
		dy := q.topLeft.Y + (q.bottomRight.Y-q.topLeft.Y)/2

		if q.bottomRight.X-q.topLeft.X == 0 && q.bottomRight.Y-q.topLeft.Y == 0 {
			return max*q.topLeft.X + q.topLeft.Y
		}

		queue = append(queue,
			quad{q.topLeft, set.Point{X: dx, Y: dy}},
			quad{set.Point{X: dx + 1, Y: q.topLeft.Y}, set.Point{X: q.bottomRight.X, Y: dy}},
			quad{set.Point{X: q.topLeft.X, Y: dy + 1}, set.Point{X: dx, Y: q.bottomRight.Y}},
			quad{set.Point{X: dx + 1, Y: dy + 1}, q.bottomRight},
		)
	}

	panic("NOT FOUND")
}

func init() {
	regex := regexp.MustCompile(`^Sensor at x=(-?\d+), y=(-?\d+): closest beacon is at x=(-?\d+), y=(-?\d+)$`)
	for _, line := range sh.Lines(input) {
		if regex.MatchString(line) {
			fields := regex.FindStringSubmatch(line)
			sensor := set.Point{X: sh.ToInt(fields[1]), Y: sh.ToInt(fields[2])}
			beacon := set.Point{X: sh.ToInt(fields[3]), Y: sh.ToInt(fields[4])}
			readings = append(readings, reading{sensor, beacon})
		}
	}
}

func merge(intervals []interval) []interval {
	sort.Slice(intervals, func(i, j int) bool { return intervals[i][0] < intervals[j][0] })
	stack := []interval{intervals[0]}
	for _, i := range intervals {
		if stack[len(stack)-1][0] <= i[0] && i[0] <= stack[len(stack)-1][1] {
			stack[len(stack)-1][1] = sh.Max(stack[len(stack)-1][1], i[1])
		} else {
			stack = append(stack, i)
		}
	}
	return stack
}

func count(intervals []interval) int {
	var c int
	for _, i := range intervals {
		c += i[1] - i[0] + 1
	}
	return c
}
