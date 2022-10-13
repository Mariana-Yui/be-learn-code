package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

func mathCalc() {
	var a, b int = 3, 4
	var c int
	c = int(math.Sqrt(float64(a*a + b*b)))
	fmt.Println(c)
}

func complexCalc() {
	var a complex128 = cmplx.Pow(math.E, math.Pi*1i) + 1
	fmt.Println(a)
	fmt.Printf("%.3f\n", a)
}

func main() {
	fmt.Println()
	mathCalc()
	complexCalc()
}
