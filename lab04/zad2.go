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
	var macierzC = make([][]int, 5)

	for i := 0; i < len(macierzA); i++ {
		macierzC[i] = make([]int, len(macierzA))

		for j := 0; j < len(macierzA); j++ {

			for k := 0; k < len(macierzB); k++ {
				macierzC[i][j] = macierzA[i][k] * macierzB[j][k]
			}
		}
	}

	for i := 0; i < len(macierzC); i++ {
		fmt.Println(macierzC[i])
	}

}
