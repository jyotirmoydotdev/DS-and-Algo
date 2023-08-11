package main

// func main() {
// 	nums := []int{1, 1, 1, 2, 2, 3}
// 	k := 2
// }

func topKFrequent(nums []int, k int) []int {

	countMap := map[int]int{}
	for _, num := range nums {
		if count, ok := countMap[num]; ok {
			countMap[num] = count + 1
		} else {
			countMap[num] = 1
		}
	}

	countSlice := make([][]int, len(nums)+1)
	for num, count := range countMap {
		countSlice[count] = append(countSlice[count], num)
	}

	res := []int{}
	for i := len(countSlice) - 1; i > 0; i-- {
		res = append(res, countSlice[i]...)
		if len(res) == k {
			return res
		}
	}
	return res
}
