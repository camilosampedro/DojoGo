package main

import (
	"fmt"
)

func Sqrt(x float64) float64 {
	return SqrtRec(x, x, 0)
}

func SqrtRec(first, actual float64, i int) float64 {
	var next = actual - ((actual * actual - first) / (2* actual))
	if actual - next < 1e-10 {
		return actual
	} else {
		return SqrtRec(first, next, i+1)
	}
}

func main() {
	fmt.Println(Sqrt(16))
}
