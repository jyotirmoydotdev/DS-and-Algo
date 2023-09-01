/*
 * Complete the 'lonelyinteger' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts INTEGER_ARRAY a as parameter.
 */

package main

func lonelyinteger(a []int32) int32 {
	// Write your code here
	m := make(map[int32]int)
	for i := 0; i < len(a); i++ {
		m[a[i]]++
	}
	for key, value := range m {
		if value != 2 {
			return key
		}
	}
	return 0
}
