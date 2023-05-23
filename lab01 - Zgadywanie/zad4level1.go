package main

import (
	"fmt"
	"math/rand"
)

func main() {
	min := 1
	max := 100
	szukana := rand.Intn(max-min) + min
	var odp int
	fmt.Println("Teraz będziesz zgadywać liczbę, którą wylosowałem")

	for {
		//fmt.Println("szukamy liczby", szukana)
		fmt.Print("Podaj liczbę: ")
		_, err := fmt.Scan(&odp)
		if err != nil {
			fmt.Println("Wystąpił błąd podczas wczytywania danych.")
			continue
		}

		if odp < szukana {
			fmt.Println("### Za mała ###")
		} else if odp > szukana {
			fmt.Println("### Za duża ###")
		} else {
			fmt.Println("Zgadł_ś, Gratulacje!")
			break
		}
	}
}
