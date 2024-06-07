package exponentiation

import "math"

type Exponentiation struct{}

func (e *Exponentiation) Exp(x int, y int) int {
	// Имплементировать метод
	res := math.Pow(float64(x), float64(y))
	return int(res)
}
