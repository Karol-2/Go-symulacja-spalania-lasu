package main

import "fmt"

func main() {
	var tablica []int

	var maksymalnaDlugosc int = 0
	var maksymalnyNumer int = 0

	for i := 1000; i <= 2000; i++ {
		var colatzWynik = colatz(i, tablica)
		if colatzWynik > maksymalnaDlugosc {
			maksymalnaDlugosc = colatzWynik
			maksymalnyNumer = i

		}
	}
	fmt.Println("n dla największej długości: ", maksymalnyNumer, ", Długość: ", maksymalnaDlugosc)

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
