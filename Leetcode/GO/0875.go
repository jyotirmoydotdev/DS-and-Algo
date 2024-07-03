package main

// func main() {
// 	slice := []int{3, 6, 7, 11}
// 	hour := 8
// 	news := minEatingSpeed(slice, hour)
// 	fmt.Println("Ans:", news)
// }

func canEat(piles []int, timeLimit, speed int) bool {
	timeNeed := 0
	for _, banana := range piles {
		timeNeed = timeNeed + (banana+speed-1)/speed
		if timeLimit < timeNeed {
			return false
		}
	}
	return true
}

func minEatingSpeed(piles []int, h int) int {
	low, high, ans := 1, 1000000000, 1
	for low <= high {
		mid := (low + high) / 2
		if canEat(piles, h, mid) {
			ans = mid
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return ans
}

// 1 ,2 ,3 ,4 ,5 ,6 ,7 ,8 ,9, 10, 11
// ^	 ^     ^
