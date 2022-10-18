package main

import "fmt"

func main() {
	x,y := 1,2
	fmt.Println("X:", x,"y:",y)
	Swap(&x,&y)
	fmt.Println("X:", x,"y:",y)
}
func Avarege(grades []float64) float64 {
	var total float64
	for _, grade := range grades {
		total += grade
	}
	return total / float64(len(grades))
}

func Half(number int) (int, bool) {
	half := number / 2
	isOdd := number%2 == 0
	return half, isOdd
}

func GetBiggestNumber(numbers ...int) int {
	var biggestNumber int

	for i, number := range numbers {
		if i == 0 || number > biggestNumber {
			biggestNumber = number
		}
	}

	return biggestNumber
}

func MakeOddGenerator() func() uint {
	i := uint(1)
	return func() (number uint) {
		number = i
		i += 2
		return
	}
}

func Fibonacci(n int) int {
	if n < 2 {
		return n
	}
	return Fibonacci(n-1) + Fibonacci(n-2)
}

func Swap(x,y *int) {
	*x, *y  = *y, *x
}