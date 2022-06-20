//Selection Sort

package main

import "fmt"

//Find which is smallest and place that is first position
//after found small value, swap between first position and that position (where you found small element)

// SelectionSort : ""
func SelectionSort(arr []int) []int {
	var small int
	for i := 0; i < len(arr); i++ {
		small = i
		for j := (i + 1); j < len(arr); j++ {
			if arr[small] > arr[j] {
				small = j
			}
		}
		temp := arr[i]
		arr[i] = arr[small]
		arr[small] = temp
	}
	return arr
}

func main() {
	arr := []int{25, 4, 22, 53, 2, 17}
	sortedlist := SelectionSort(arr)
	fmt.Println("\n sortedlist is", sortedlist)
}
