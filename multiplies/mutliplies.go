package multiplies

import "fmt"

func Multiply(x, y int, s string) int {
	var sgy strategy

	switch s {
	case "egyptian":
		sgy = &egyptian{}
	case "normal":
		sgy = &normal{}
	case "russian":
		sgy = &russianPeasant{}
	case "long":
		sgy = &long{}
	case "addition":
		sgy = &addition{}
	case "lattice":
		sgy = &lattice{}
	default:
		str, _ := fmt.Printf("not implemented: %s", s)
		panic(str)
	}

	return sgy.Multiply(x, y)
}

func MultiplyGivenStrategy(x, y int, sgy strategy) int {
	return sgy.Multiply(x, y)
}
