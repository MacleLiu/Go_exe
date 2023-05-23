package main

import "fmt"

func max(vals ...int) int {
	if len(vals) == 0 {
		return 0
	}
	max := vals[0]
	for _, val := range vals {
		if val > max {
			max = val
		}
	}
	return max
}

func min(vals ...int) int {
	if len(vals) == 0 {
		return 0
	}
	min := vals[0]
	for _, val := range vals {
		if val < min {
			min = val
		}
	}
	return min
}

func max1(first int, vals ...int) int {
	max := first
	for _, val := range vals {
		if val > max {
			max = val
		}
	}
	return max
}

func min1(first int, vals ...int) int {
	min := first
	for _, val := range vals {
		if val < min {
			min = val
		}
	}
	return min
}

func main() {
	values := []int{2, 4, 7, 1, 5, 8, 3, 3, 1}
	fmt.Println(max(values...))
	fmt.Println(min(values...))
	fmt.Println(max1(7, values...))
	fmt.Println(min1(5, values...))
}
