package main

import (
	"flag"
	"fmt"
	"math"
)

type F struct{ a, b, c float64 }

func main() {
	var a, b, c float64
	flag.Float64Var(&a, "a", 0, "a")
	flag.Float64Var(&b, "b", 0, "b")
	flag.Float64Var(&c, "c", 0, "c")
	flag.Parse()
	f := F{a, b, c}
	w, p := d(&f)
	if !p {
		fmt.Println(p)
	} else if len(w) == 2 {
		fmt.Println(p)
		fmt.Println("(", w[0], ", ", w[1], ")")
	} else {
		fmt.Println("(", w[0], ")")
	}
}

func d(f *F) ([]float64, bool) {
	d := f.b*f.b - 4*f.a*f.c

	if d < 0 {
		return nil, false
	}
	s := math.Sqrt(d)
	return []float64{(-f.b - s) / (2 * f.a), (-f.b + s) / (2 * f.a)}, true
}
