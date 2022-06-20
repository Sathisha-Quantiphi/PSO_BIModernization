//Insertion Sort

package main

import "fmt"

//In Insertion Sort check from end to begin and insert the value in correct position
// InsertionSort : ""
func InsertionSort(arr []int) []int {
	if len(arr) == 0 {
		return nil
	}
	var j int
	for i := 1; i < len(arr); i++ {
		currentvalue := arr[i]
		for j = (i - 1); j >= 0 && arr[j] > currentvalue; j-- {
			arr[j+1] = arr[j]
		}
		arr[j+1] = currentvalue
	}
	return arr
}

func main() {
	arr := []int{25, 4, 22, 53, 2, 17}
	sortedlist := InsertionSort(arr)
	fmt.Println("\n sortedlist is", sortedlist)
}
