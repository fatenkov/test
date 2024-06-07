package calculator

type Calculator struct{}

func (c *Calculator) Sum(x int, y int) int {
	// Имплементировать метод
	res := x + y
	return res
}

func (c *Calculator) Sub(x int, y int) int {
	// Имплементировать метод
	res := x - y
	return res
}

func (c *Calculator) Div(x int, y int) int {
	// Имплементировать метод
	var res int
	if y == 0 {
		res = 0
	} else {
		res = x / y
	}
	return res
}

func (c *Calculator) Mult(x int, y int) int {
	// Имплементировать метод
	res := x * y
	return res
}
