// Note : Need to write article and add to readme
package main

// func main() {
// 	temp := []int{73, 74, 75, 71, 69, 72, 76, 73}
// 	// 			  0   1   2   3   4   5   6   7
// 	res := dailyTemperatures(temp)
// 	fmt.Println("res:", res)
// }

// ----------- ( Brute Force O(n^2)) -----------
// func dailyTemperatures(temperatures []int) []int {
// 	res := make([]int, 0, 8)
// 	for i, _ := range temperatures {
// 		counter := 0
// 		added := false
// 		for j := i; j < len(temperatures); j++ {
// 			if temperatures[j] > temperatures[i] {
// 				res = append(res, counter)
// 				added = true
// 				break
// 			}
// 			counter++
// 		}
// 		if !added {
// 			res = append(res, 0)
// 		}
// 	}
// 	return res
// }

// ----------- ( stack O(n)) -----------
func dailyTemperatures(temperatures []int) []int {
	res := make([]int, len(temperatures))
	for i := len(temperatures) - 1; i >= 0; i-- {
		j := i + 1
		for j < len(temperatures) && temperatures[j] <= temperatures[i] {
			if res[j] <= 0 {
				break
			}
			j += res[j]
		}
		if j < len(temperatures) && temperatures[j] > temperatures[i] {
			res[i] = j - i
		}
	}
	return res
}
