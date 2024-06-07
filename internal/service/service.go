package service

type Service interface { //содержит методы, доступные в этом интерфейсе, сами методы в calculator.go
	Sum(x int, y int) int
	Sub(x int, y int) int
	Div(x int, y int) int
	Mult(x int, y int) int
	//Exp(x int, y int) int  // тогда красным зажигается Server = httpserver.New(8080, &calculator.Calculator{})
}

type Exponentiation interface {
	Exp(x int, y int) int
}
