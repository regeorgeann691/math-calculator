package main

import (
	"fmt"
)

func main() {
	var s1, s2, s3, s4 string

	fmt.Print("Enter three numbers: ")
	fmt.Scan(&s1)
	fmt.Scan(&s2)
	fmt.Scan(&s3)

	if s1 == "5" && s2 == "9" && s3 == "8" {
		s4 = "success"
	} else if s1 == "6" && s2 == "8" && s3 == "7" {
		s4 = "error"
	}
	fmt.Println(s4)
}
