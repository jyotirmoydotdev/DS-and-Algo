package main

import (
	"sort"
)

// func main() {

// 	target := 12
// 	position := []int{5, 3, 8}
// 	speed := []int{2, 5, 3}

// 	result := carFleet(target, position, speed)
// 	fmt.Println("result:", result)

// }

func carFleet(target int, position []int, speed []int) int {
	pair := make([]carInfo, 0, len(position))
	stack := make([]float32, 0, len(position))
	for i, _ := range position {
		pair = append(pair, carInfo{position[i], speed[i]})
	}
	sort.Slice(pair, func(i, j int) bool {
		return pair[i].pos < pair[j].pos
	})
	for i := len(pair) - 1; i >= 0; i-- {
		stack = append(stack, float32(target-pair[i].pos)/float32(pair[i].spd))
		if len(stack) >= 2 && stack[len(stack)-1] <= stack[len(stack)-2] {
			stack = stack[:len(stack)-1]
		}
	}
	return len(stack)
}

type carInfo struct {
	pos int
	spd int
}
