package main

import "fmt"

func main() {
	t := []int32{4, 2, 3, 1}
	bubbleSort(t)
	fmt.Println(t)
}

func bubbleSort(t []int32) {
	swaps := 0
	for k := 0; k < len(t); k++ {
		for i := 0; i < len(t)-1; i++ {
			if t[i+1] < t[i] {
				t[i], t[i+1] = t[i+1], t[i]
				swaps += 1
			}
		}
	}
	fmt.Printf("Array is sorted in %d swaps.\n", swaps)
	fmt.Printf("First Element: %d\n", t[0])
	fmt.Printf("Last Element: %d\n", t[len(t)-1])
}
