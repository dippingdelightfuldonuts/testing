package main

import "fmt"

func Sum(x, y int) int {
	return x + y
}

func Multiply(x, y int) int {
	return x * y
}

func main() {
	fmt.Println("sum: ", Sum(5, 5))
}
