package main

import (
	"fmt"
	"strconv"
	"strings"
)

// horizontalByDepth sums the total value of each directional move.
// For forward movement this is just sum of mf, for horizontal postition
// this needs to be the total displacement: sum of mD values minus sum of
// mU values. (Since submarine down is positive up is negative)
func getDepths(mF, mU, mD []int, aD int) (int, int) {
	var h, d, u int
	for _, v := range mF {
		h += v
	}
	for _, v := range mD {
		d += v
	}
	for _, v := range mU {
		u += v
	}

	return h * (d - u), h * aD
}

// getDirections takes moves and returns int slices for the power
// of the movement.
func getDirections(m []string) (mF, mU, mD []int, aim, aimDepth int) {
	for _, v := range m {
		switch {
		case strings.Contains(v, "forward"):
			n, _ := strconv.Atoi(strings.Split(v, " ")[1])
			mF = append(mF, n)
			aimDepth += (aim * n)
		case strings.Contains(v, "up"):
			n, _ := strconv.Atoi(strings.Split(v, " ")[1])
			mU = append(mU, n)
			aim -= n
		case strings.Contains(v, "down"):
			n, _ := strconv.Atoi(strings.Split(v, " ")[1])
			mD = append(mD, n)
			aim += n
		}
	}
	return mF, mU, mD, aim, aimDepth
}

// getMovements takes the raw input and determins the move per line
func getMovements(input string) (m []string) {
	m = strings.Split(input, "\n")
	return m
}

func main() {
	moves := getMovements(rawInput)
	mvsFwd, mvsUp, mvsDwn, a, aD := getDirections(moves)
	n, n2 := getDepths(mvsFwd, mvsUp, mvsDwn, aD)

	fmt.Printf("Horizontal movement muliplyed by depth:'%d'\n", n)
	fmt.Printf("Aim:'%d' aimDepth:'%d'\n", a, aD)
	fmt.Printf("Horizontal movement muliplyed by aimDepth:'%d'\n", n2)
}
