package main

import "math/rand"

func main() {
	rand.Seed(12345)
	result := rand.Intn(10) + 1
	fmt.Println("Result:", result)
}
