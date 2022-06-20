//RGB

package main

import (
	"fmt"
)

// RGB : ""
func RGB(byteval []byte) bool {
	Rcount, Gcount, Bcount := 0, 0, 0
	for i := 0; i < len(byteval); i++ {
		if string(byteval[i]) == "R" {
			Rcount++
		}
		if string(byteval[i]) == "G" {
			Gcount++
		}
		if string(byteval[i]) == "B" {
			Bcount++
		}
		fmt.Println("\n Rcount..", Rcount, "Gcount...", Gcount, "Bcount..", Bcount)
		if (Rcount == Bcount) && (Bcount == Gcount) && (Gcount == Rcount) {
			return true
		}
	}
	return false
}

func main() {
	fmt.Println("\n Inside main func")
	var str string
	fmt.Scanln(&str)
	byteval := []byte(str)
	fmt.Println("\n byteval is", byteval)
	fmt.Println(RGB(byteval))
}
