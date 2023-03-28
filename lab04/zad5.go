package main

import "fmt"

func main() {
	macierzA := [][]int{
		{1, 2, 3, 4, 5},
		{5, 1, 2, 8, 4},
		{5, 1, 3, 6, 5},
		{7, 8, 5, 6, 1},
		{1, 5, 3, 6, 8},
	}

	macierzB := [][]int{
		{1, 2, 3, 4, 5},
		{5, 1, 2, 8, 4},
		{5, 1, 3, 6, 5},
		{7, 8, 5, 6, 1},
		{1, 5, 3, 6, 8},
	}
	// liczba kolumn w pierwszej macierzy musi być równa liczbie wierszy w drugiej macierzy
	if len(macierzA[0]) != len(macierzB) {
		fmt.Println("Nie można pomnożyć tych macierzy!")
		return
	}

	rzedy := len(macierzA)
	kolumny := len(macierzB[0])

	wynik := make([][]int, rzedy)
	for i := range wynik {
		wynik[i] = make([]int, kolumny)
	}

	for i := 0; i < rzedy; i++ {
		for j := 0; j < kolumny; j++ {
			for k := 0; k < len(macierzA[0]); k++ {
				wynik[i][j] += macierzA[i][k] * macierzB[k][j]
			}
		}
	}

	fmt.Println("Wynik mnożenia macierzy:")
	for i := 0; i < rzedy; i++ {
		for j := 0; j < kolumny; j++ {
			fmt.Print(wynik[i][j], " ")
		}
		fmt.Println()
	}

}
