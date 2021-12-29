package main

import (
	_ "embed"
	"fmt"
	"math"

	sh "github.com/leonhfr/aoc/shared"
)

var ta = targetArea{124, 174, -123, -86}

func main() {
	maxh, lowestX, highestX := part1()
	fmt.Printf("Part 1: %v\n", maxh)
	fmt.Printf("Part 2: %v\n", part2(lowestX, highestX))
}

type targetArea struct {
	xmin, xmax, ymin, ymax int
}

type velocity struct {
	x, y int
}

func part1() (int, int, int) {
	// Check that we can pick x so that velocity_x slows down to 0 when it hits the target,
	// therefore we need an that complies with :
	//   xmin <= x(x+1)/2 <=xmax
	//   2*xmin+0.25 <= (x+0.5)^2 <= 2*xmax+0.25

	initialX := func(b int) float64 {
		return math.Sqrt(2*float64(b)+0.25) - 0.5
	}
	lowestX, highestX := int(math.Ceil(initialX(ta.xmin))), int(math.Floor(initialX(ta.xmax)))

	// Check that at least one x exists
	if !(lowestX <= highestX) {
		panic("No solution for x to reach velocity 0 in the target window.")
	}

	// We toss with a set x so that the velocity_x becomes null relatively quickly
	// We also toss with a y as high as possible, so we need to check the maximum x that still hits the target
	// After 2*y steps, the next step would reach the row -y-1, so we pick y so that ymin=-y-1
	// In which case, our max height would be maxh=-ymin*(-ymin-1)/2

	maxh := -ta.ymin * (-ta.ymin - 1) / 2

	return maxh, lowestX, highestX
}

func part2(lowestX, highestX int) int {
	solutions := make(map[velocity]int)
	// We do this in two parts
	// 1: solutions where x reach a velocity of 0 in the target window

	// First we compute the maximum number of steps that it'd take to reach ymin starting with y=0
	initialY := func(b int) float64 {
		return math.Sqrt(2*float64(b)+0.25) - 0.5
	}
	maxSteps := int(math.Ceil(initialY(-ta.ymin)))

	for step := 1; step <= maxSteps; step++ {
		minFallY := sh.Min(0, ta.ymin+step*(step+1)/2)
		maxFallY := sh.Min(0, ta.ymax+step*(step+1)/2)
		longPitchMin := -int(math.Floor(float64(maxFallY) / float64(step)))
		longPitchMax := -int(math.Ceil(float64(minFallY) / float64(step)))
		longPitchMin = sh.Max(longPitchMin, int(math.Ceil(float64(lowestX)/2.0-(float64(step)+1.0)/2.0)))

		if longPitchMin <= longPitchMax {
			for x := lowestX; x <= highestX; x++ {
				for y := longPitchMin; y <= longPitchMax; y++ {
					solutions[velocity{x, y}] = step
				}
			}
		}
	}

	// 2: solutions where x goes past the target
	for step := 1; step <= lowestX; step++ {
		xAimMin := ta.xmin + step*(step-1)/2
		xAimMax := ta.xmax + step*(step-1)/2
		yAimMin := ta.ymin + step*(step-1)/2
		yAimMax := ta.ymax + step*(step-1)/2

		xPastMin := sh.Min(ceilDiv(xAimMin, step), ceilDiv(xAimMax, step))
		xPastMax := sh.Max(floorDiv(xAimMin, step), floorDiv(xAimMax, step))
		yPastMin := sh.Min(ceilDiv(yAimMin, step), ceilDiv(yAimMax, step))
		yPastMax := sh.Max(floorDiv(yAimMin, step), floorDiv(yAimMax, step))

		for x := xPastMin; x <= xPastMax; x++ {
			for y := yPastMin; y <= yPastMax; y++ {
				solutions[velocity{x, y}] = step
			}
		}
	}

	return len(solutions)
}

func ceilDiv(a, b int) int {
	return int(math.Ceil(float64(a) / float64(b)))
}

func floorDiv(a, b int) int {
	return int(math.Floor(float64(a) / float64(b)))
}
