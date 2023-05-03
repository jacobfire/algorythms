package main

import "fmt"

func main() {
	var in = []int{
		3, 4, 5, 1, 8,
	}
	out := selectionSort(in)

	fmt.Println(out)
}

func selectionSort(in []int) []int {
	length := len(in)
	for i := 0; i < length-1; i++ {
		minIndex := i
		for j := i + 1; j < length; j++ {
			if in[minIndex] > in[j] {
				minIndex = j
			}
		}
		if minIndex != i {
			t := in[minIndex]
			in[minIndex] = in[i]
			in[i] = t
		}
	}

	return in
}
