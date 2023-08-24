package main

// func main() {
// 	number := []int{2, 7, 11, 15}
// 	target := 9
// 	fmt.Println("twoSum:", twoSum2_2(number, target))
// }

func twoSum2_1(numbers []int, target int) []int {
	m := make(map[int]int)
	for i, num := range numbers {
		if _, ok := m[target-num]; ok {
			return []int{m[target-num] + 1, i + 1}
		}
		m[num] = i
	}
	return []int{}
}

func twoSum2_2(numbers []int, target int) []int {
	left, right := 0, len(numbers)-1
	for left < right {
		if numbers[left]+numbers[right] < target {
			left++
		} else if numbers[left]+numbers[right] > target {
			right--
		} else {
			return []int{left + 1, right + 1}
		}
	}
	return nil
}
