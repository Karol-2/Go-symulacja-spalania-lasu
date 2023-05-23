package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
)

func main() {
	gra()
}

func gra() {
	min := 1
	max := 100
	szukana := rand.Intn(max-min) + min
	var odp string
	fmt.Println("Teraz będziesz zgadywać liczbę, którą wylosowałem")
	fmt.Println("Po wpisaniu \"koniec\" nastąpi zakończenie zabawy")

	for {
		fmt.Println("szukamy liczby", szukana)
		fmt.Print("Podaj liczbę: ")
		fmt.Scan(&odp)
		if odp == "koniec" {
			fmt.Println("Żegnaj :( ")
			os.Exit(1)
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

			var powtorka string
			fmt.Println("Czy gramy jeszcze raz? [tak/nie]")
			fmt.Scan(&powtorka)
			if powtorka == "tak" {
				gra()
			} else if powtorka == "nie" {
				os.Exit(1)
			} else {
				fmt.Println("Nie zrozumiałem, więc wychodzisz")
				os.Exit(1)
			}

		}
	}
}
