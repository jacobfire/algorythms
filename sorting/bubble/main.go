package main

import "fmt"

func main() {
	var input = []int{3, 4, 1, 5, 8}
	fmt.Println(input)

	out := bubbleSort(input)
	fmt.Println(out)
}

func bubbleSort(in []int) []int {
	length := len(in)
	maxIndex := 0
	for {
		if length == 1 {
			return in
		}
		for i := 1; i < length; i++ {
			fmt.Println(length)
			if in[i-1] > in[i] {
				t := in[i-1]
				in[i-1] = in[i]
				in[i] = t
				maxIndex = i
			}
		}
		length = maxIndex
	}
}
