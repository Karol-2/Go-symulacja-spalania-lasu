package main

import "fmt"

func main() {
	//var liczba int
	var tablica []int

	var maksymalnaDlugosc int = 0
	var maksymalnyNumer int = 0

	for i := 1000; i <= 2000; i++ {

		//fmt.Println("Dla liczby", i)
		//fmt.Println(colatz(i, tablica))
		if colatz(i, tablica) > maksymalnaDlugosc {
			maksymalnaDlugosc = colatz(i, tablica)
			maksymalnyNumer = i

		}
	}
	fmt.Println("n dla największej długości: ", maksymalnyNumer, ", Długość: ", maksymalnaDlugosc)

}

func colatz(x int, tablica []int) int {
	var arr = append(tablica, x)
	if x == 1 {
		//fmt.Println("Długość:", len(arr))
		return len(arr)
	} else if x%2 == 0 {
		var nowyX int = x / 2
		//fmt.Print(nowyX, ", ")
		return colatz(nowyX, arr)

	} else {
		var nowyX int = (x * 3) + 1
		//fmt.Print(nowyX, ", ")
		return colatz(nowyX, arr)
	}
}
