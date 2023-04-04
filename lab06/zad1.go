package main

import (
	"flag"
	"fmt"
	"math"
)

type Funkcja struct {
	a float64
	b float64
	c float64
}

func main() {
	var a, b, c float64
	flag.Float64Var(&a, "a", 0.0, "Value of a")
	flag.Float64Var(&b, "b", 0.0, "Value of b")
	flag.Float64Var(&c, "c", 0.0, "Value of c")

	flag.Parse()

	var funkcja Funkcja
	funkcja.a = a
	funkcja.b = b
	funkcja.c = c
	wyniki, policzalne := liczenieDelty(&funkcja)
	if policzalne == false {
		fmt.Println("Nie da się policzyć delty dla tej funkcji")
	} else if len(wyniki) == 2 {
		fmt.Println("Da się policzyć deltę dla tej funkcji, dwa pierwiastki")
		fmt.Println("Wyniki: (", wyniki[0], ", ", wyniki[1], ")")
	} else {
		fmt.Println("Da się policzyć deltę dla tej funkcji, jeden pierwiastek")
		fmt.Println("Wynik: (", wyniki[0], ")")
	}
}

func liczenieDelty(funkcja *Funkcja) ([]float64, bool) {
	delta := (funkcja.b * funkcja.b) - 4*funkcja.a*funkcja.c
	fmt.Println("delta =", delta)
	wyniki := make([]float64, 0)

	if delta < 0 {
		return nil, false
	} else if delta == 0 {
		x0 := (-funkcja.b - math.Sqrt(delta)) / (2 * funkcja.a)
		wyniki = append(wyniki, x0)
		return wyniki, true
	} else {
		x1 := (-funkcja.b - math.Sqrt(delta)) / (2 * funkcja.a)
		x2 := (-funkcja.b + math.Sqrt(delta)) / (2 * funkcja.a)
		wyniki = append(wyniki, x1, x2)
		return wyniki, true
	}
}
