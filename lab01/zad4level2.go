package main

import (
	"fmt"
	"math/rand"
	"strconv"
)

func main() {
	min := 1
	max := 100
	szukana := rand.Intn(max-min) + min
	var odp string
	fmt.Println("Teraz będziesz zgadywać liczbę, którą wylosowałem")
	fmt.Println("Po wpisaniu \"koniec\" nastąpi zakończenie zabawy")

	for {
		//fmt.Println("szukamy liczby", szukana)
		fmt.Print("Podaj liczbę: ")
		fmt.Scan(&odp)
		if odp == "koniec" {
			fmt.Print("Żegnaj :( ")
			break
		}

		odp, err := strconv.Atoi(odp)
		if err != nil {
			fmt.Println("Złe dane wejściowe")
			break
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
