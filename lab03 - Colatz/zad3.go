package main

import (
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

func main() {
	var tablica []int
	var values []int

	for i := 1; i <= 100000; i++ {

		values = append(values, colatz(i, tablica))
	}

	histPlot(values)
	// go get gonum.org/v1/plot@latest

}
func histPlot(values plotter.Values) {
	p, err := plot.New()
	if err != nil {
		panic(err)
	}
	p.Title.Text = "Histogram"
	hist, err := plotter.NewHist(values, 20)
	if err != nil {
		panic(err)
	}
	p.Add(hist)

	if err := p.Save(3*vg.Inch, 3*vg.Inch, "hist.png"); err != nil {
		panic(err)
	}

}

func colatz(x int, tablica []int) int {
	var arr = append(tablica, x)
	if x == 1 {
		return len(arr)
	} else if x%2 == 0 {
		var nowyX int = x / 2
		return colatz(nowyX, arr)

	} else {
		var nowyX int = (x * 3) + 1
		return colatz(nowyX, arr)
	}
}
