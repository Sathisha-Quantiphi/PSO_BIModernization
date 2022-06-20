package main

import (
	"fmt"
)

func main() {
	fmt.Println("We are executing a goroutine")
	arr := []int{2, 3, 4}
	ch := make(chan int, len(arr))
	go func(arr []int, ch chan int) {
		for _, elem := range arr {
			ch <- elem * 3
		}
	}(arr, ch)
	for i := 0; i < len(arr); i++ {
		fmt.Printf("Result: %v \n", <-ch)
	}
}
