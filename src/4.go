



package main

import "fmt"

func main() {
	x := 5.0
	y := 6.0
	result := calculate(x, y)
	fmt.Println(result)
}

func calculate(x float64, y float64) float64 {
	return x + y
}