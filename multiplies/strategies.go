package multiplies

import (
	"fmt"
	"math"
	"strconv"
)

type strategy interface {
	Multiply(int, int) int
}

type egyptian struct {
}

func (this *egyptian) Multiply(x, y int) int {
	keys := make([]int, 0)
	lookup := make(map[int]int)
	currentValue := y

	for binary := 1; binary <= x; binary *= 2 {
		lookup[binary] = currentValue
		keys = append(keys, binary)

		currentValue = currentValue * 2 // double y each iteration
	}

	reversedKeys := reverse(keys)
	remainer := x

	solution := 0

	for _, key := range reversedKeys {
		value, _ := lookup[key]

		if key <= remainer {
			remainer -= key
			solution += value
		}
	}

	return solution
}

func reverse(a []int) []int {
	for left, right := 0, len(a)-1; left < right; left, right = left+1, right-1 {
		a[left], a[right] = a[right], a[left]
	}

	return a
}

type normal struct {
}

func (this *normal) Multiply(x, y int) int {
	return x * y
}

type russianPeasant struct {
}

func (this *russianPeasant) Multiply(x, y int) int {
	binary := y
	result := 0

	for left := x; left > 0; left /= 2 {
		if left%2 == 1 {
			result += binary
		}

		binary *= 2
	}

	return result
}

type long struct {
}

func (this *long) Multiply(x, y int) int {
	xStr := reverseStr(strconv.Itoa(x))
	yStr := reverseStr(strconv.Itoa(y))
	result := 0
	carry := 0

	for yMagnitude, yChar := range yStr {
		yc := string(yChar)
		yValue, _ := strconv.Atoi(yc)
		magnitude := 0

		for xMagnitude, xChar := range xStr {
			xc := string(xChar)
			xValue, _ := strconv.Atoi(xc)

			magnitude = int(math.Pow(10, float64(xMagnitude))) * int(math.Pow(10, float64(yMagnitude)))

			subMultiply := xValue * yValue
			result += (subMultiply%10 + carry) * magnitude
			carry = subMultiply / 10
		}

		if carry > 0 {
			result += carry * magnitude * 10
			carry = 0
		}
	}

	return result
}

func reverseStr(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

type addition struct {
}

func (this *addition) Multiply(x, y int) int {
	result := 0

	for i := 0; i < y; i++ {
		result += x
	}

	return result
}

type lattice struct {
}

type cell struct {
	label int
	upper int
	lower int
}

func (this cell) String() string {
	if this.label != 0 {
		return strconv.Itoa(this.label)
	}

	return fmt.Sprintf("%d/%d", this.upper, this.lower)
}

func (this *lattice) Multiply(x, y int) int {
	xStr := strconv.Itoa(x)
	yStr := strconv.Itoa(y)

	xSize := len(xStr) + 1
	ySize := len(yStr) + 1

	matrix := make([][]cell, xSize)
	rows := make([]cell, xSize*ySize)
	for j := 0; j < ySize; j++ {
		matrix[j] = rows[j*xSize : (j+1)*xSize]
	}

	for i, xChar := range xStr {
		value, _ := strconv.Atoi(string(xChar))
		matrix[0][i] = cell{
			label: value,
		}
	}

	for j, yChar := range yStr {
		value, _ := strconv.Atoi(string(yChar))
		matrix[1+j][xSize-1] = cell{
			label: value,
		}
	}

	// fill in cells
	for j := 1; j < ySize; j++ {
		for i := 0; i < xSize-1; i++ {
			grabX := matrix[0][i]
			grabY := matrix[j][xSize-1]

			mult := grabX.label * grabY.label
			matrix[j][i] = cell{
				upper: mult / 10,
				lower: mult % 10,
			}
		}
	}

	fmt.Println("matrix: ", matrix)
	return 0
}
