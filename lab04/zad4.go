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
	fmt.Println("Macierz orginalny: ")
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			fmt.Print(macierz[i][j], " ")
		}
		fmt.Println()
	}

	var macierzOdwrocony = make([][]int, 3)
	liczba = 8

	for i := range macierzOdwrocony {
		macierzOdwrocony[i] = make([]int, 3)
		for j := range macierzOdwrocony[i] {
			macierzOdwrocony[i][j] = liczba
			liczba--
		}
	}
	fmt.Println("Macierz odwrocony: ")
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			fmt.Print(macierzOdwrocony[i][j], " ")
		}
		fmt.Println()
	}

	result := make([][]int, 3)
	for i := range result {
		result[i] = make([]int, 3)
	}

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			result[i][j] = macierz[i][j] + macierzOdwrocony[i][j]
		}
	}

	fmt.Println("Wynik dodawania: ")
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			fmt.Print(result[i][j], " ")
		}
		fmt.Println()
	}

}
