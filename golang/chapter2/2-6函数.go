package main

import (
	"fmt"
	"math"
	"reflect"
	"runtime"
)

func calc(a, b int, op string) (float64, error) {
	switch op {
	case "+":
		return float64(a + b), nil
	case "-":
		return float64(a - b), nil
	case "*":
		return float64(a * b), nil
	case "/":
		return float64(a / b), nil
	default:
		return 0, fmt.Errorf("invalid Operator")
	}
}

func eval(a, b int) (p, r int) {
	return a / b, a % b
}

func apply(op func(a, b int) float64, a, b int) float64 {
	p := reflect.ValueOf(op).Pointer()
	opName := runtime.FuncForPC(p).Name()
	fmt.Printf("Calling function %s with args "+"(%d, %d)\n", opName, a, b)
	return op(a, b)
}

func main() {
	num, err := calc(6, 3, "^")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(num)
	}
	p, r := eval(6, 3)
	fmt.Printf("a/b: %d, a%%b: %d\n", p, r)
	fmt.Println(apply(func(a, b int) float64 {
		return math.Sqrt(float64(a*a) + float64(b*b))
	}, 4, 4))
}
