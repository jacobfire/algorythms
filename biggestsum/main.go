package main

import (
	"fmt"
	"math"
)

func main() {
	numbers := []int{
		3, 1, 6, 7, 4, 5, 1, 8, 9,
	}

	fmt.Println(maxSum(numbers))
}

func maxSum(numbers []int) int {
	var max int
	temp := make([]int, len(numbers)+2)

	for i := range numbers {
		max = int(math.Max(float64(temp[i]+numbers[i]), float64(temp[i+1])))
		temp[i+2] = max
	}
	// fmt.Printf("%+v", temp)

	return max
}
