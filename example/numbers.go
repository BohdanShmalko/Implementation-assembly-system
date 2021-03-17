package main

import (
	"errors"
	"log"
	"math"
)

type Numbers struct {
	Num1 int
	Num2 int
}

func NewNumbers(num1, num2 int) *Numbers {
	return &Numbers{
		Num1: num1,
		Num2: num2,
	}
}

func (tn *Numbers) MulSqrt() (float64, error) {
	result := math.Sqrt(float64(tn.Num1 * tn.Num2))
	if math.IsNaN(result) {
		return result, errors.New("one of the numbers is negative and the other is positive")
	}
	return result, nil
}

func main() {
	numbers := NewNumbers(10, 20)
	mulSqrt, _ := numbers.MulSqrt()
	log.Println(mulSqrt)
}
