// Obserwacje:
// liczby utrzymują się na poziomie około 250 - 260

package main

import "fmt"

func main() {
	var minimalny int = 1000
	var maksymalny int = 11000
	for j := minimalny; j <= maksymalny-1000; j += 1000 {
		var tablica []int
		var maksymalnaDlugosc int = 0
		var maksymalnyNumer int = 0

		for i := j; i <= j+1000; i++ {
			var colatzWynik = colatz(i, tablica)
			if colatzWynik > maksymalnaDlugosc {
				maksymalnaDlugosc = colatzWynik
				maksymalnyNumer = i
			}
		}
		fmt.Println("Przedział ", j, "-", j+1000, ", Numer: ", maksymalnyNumer, ", Długość: ", maksymalnaDlugosc)

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
