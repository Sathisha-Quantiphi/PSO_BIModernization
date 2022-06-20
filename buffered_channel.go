package main

import (
	"fmt"
	"time"
)

func main() {
	names := make(chan string, 3)
	go generateName(names)

	// Simulate that a different process takes 5 seconds to run before the main goroutine
	// starts reading values from the channel.
	time.Sleep(5 * time.Second)

	for name := range names {
		fmt.Printf("Name received: %v\n", name)
	}
}

// The function indicates only write into channel not to read
func generateName(names chan<- string) {
	defer close(names)
	for _, name := range []string{"Carl", "Paul", "May", "Laura", "John"} {
		time.Sleep(1 * time.Second)
		names <- name
		fmt.Printf("Name sent: %v\n", name)
	}
}
