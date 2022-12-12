package main

import (
	_ "embed"
	"fmt"
	"strings"

	sh "github.com/leonhfr/aoc/shared"
)

//go:embed input
var input string

var lines []string

func main() {
	fmt.Printf("Part 1: %v\n", part1())
	fmt.Printf("Part 2: %v\n", part2())
}

type directory struct {
	name string
	size int
}

const (
	CUTOFF = 100_000
	UNUSED = 30_000_000
	TOTAL  = 70_000_000
	TARGET = TOTAL - UNUSED
)

func part1() int {
	var sum int
	for _, dir := range dirList(lines) {
		if dir.size <= CUTOFF {
			sum += dir.size
		}
	}
	return sum
}

func part2() int {
	list := dirList(lines)
	sum := list[len(list)-1].size
	min := sum
	for _, dir := range list {
		if sum-dir.size <= TARGET && dir.size < min {
			min = dir.size
		}
	}
	return min
}

func dirList(lines []string) []directory {
	var list, stack []directory
	for _, line := range lines {
		switch {
		case line == "$ cd /":
			stack = append(stack, directory{name: "/"})
		case line == "$ cd ..":
			dir := stack[len(stack)-1]
			list = append(list, dir)
			stack = stack[:len(stack)-1]
			stack[len(stack)-1].size += dir.size
		case strings.HasPrefix(line, "$ cd "):
			dir := strings.TrimPrefix(line, "$ cd ")
			stack = append(stack, directory{name: dir})
		case line == "$ ls":
			// do nothing
		case strings.HasPrefix(line, "dir "):
			// do nothing
		default:
			// 000 filename.ext
			fields := strings.Split(line, " ")
			size := sh.ToInt(fields[0])
			stack[len(stack)-1].size += size
		}
	}

	for len(stack) > 0 {
		dir := stack[len(stack)-1]
		list = append(list, dir)
		stack = stack[:len(stack)-1]
		if len(stack) > 0 {
			stack[len(stack)-1].size += dir.size
		}
	}

	return list
}

func init() {
	lines = sh.Lines(input)
}
