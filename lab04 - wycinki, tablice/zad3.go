package main

import "fmt"

func main() {

	var macierz = make([][]int, 3)
	liczba := 0

	for i := range macierz {
		macierz[i] = make([]int, 3)
		for j := range macierz[i] {
			macierz[i][j] = liczba
			liczba++
		}
	}

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			fmt.Print(macierz[i][j], " ")
		}
		fmt.Println()
	}

}
