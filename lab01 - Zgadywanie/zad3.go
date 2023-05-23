package main

import (
	"fmt"
	"math"
)

func main() {
	var a float64
	var b float64
	var c float64

	fmt.Println("podaj a: ")
	fmt.Scanln(&a)
	fmt.Println("podaj b: ")
	fmt.Scanln(&b)
	fmt.Println("podaj c: ")
	fmt.Scanln(&c)

	var delta float64 = (b * b) - (4 * a * c)
	fmt.Println("Delta równa:", delta)

	if delta < 0 {
		fmt.Println("Delta mniejsza od zera, zero pierwiastków")

	} else if delta == 0 {
		fmt.Println("Delta równa zero, jeden pierwiastek")
		var x0 float64 = -b / 2 * a
		fmt.Println("x0:", x0)

	} else {

		fmt.Println("Delta większa od zera, dwa pierwiastki")

		var x1 float64 = (-b - math.Sqrt(delta)) / 2 * a
		var x2 float64 = (-b + math.Sqrt(delta)) / 2 * a

		fmt.Println("x1:", x1)
		fmt.Println("x2:", x2)

	}
}
