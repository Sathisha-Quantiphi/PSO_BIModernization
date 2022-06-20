//Bubble Sort

package main

import "fmt"

// BubbleSort : ""
func BubbleSort(arr []int) []int {
	if len(arr) != 0 {
		for i := len(arr); i > 0; i-- {
			isRearrange := false
			for j := 0; j < (i - 1); j++ {
				if arr[j] > arr[j+1] {
					arr[j], arr[j+1] = arr[j+1], arr[j]
					isRearrange = true
				}
			}
			if !isRearrange {
				break
			}
		}
	}
	return arr
}

func main() {
	arrayvalue := []int{4, 89, 2, 44, 666, 1}
	sortedValue := BubbleSort(arrayvalue)
	fmt.Println("\n Bubble Sorted values are ", sortedValue)
}
