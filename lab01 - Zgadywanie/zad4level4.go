package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
)

func main() {

	var players map[string]int
	players = make(map[string]int)
	gra(players)
}

func gra(players map[string]int) {
	min := 1
	max := 100
	szukana := rand.Intn(max-min) + min
	podejscia := 1
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
			podejscia++
		} else if odp > szukana {
			fmt.Println("### Za duża ###")
			podejscia++
		} else {
			fmt.Println("Zgadł_ś, Gratulacje!")
			fmt.Println("Ilość podejść: ", podejscia)

			var imie string
			fmt.Print("Podaj swoje imie:")
			fmt.Scan(&imie)
			players[imie] = podejscia

			var powtorka string
			fmt.Println("Czy gramy jeszcze raz? [tak/nie]")
			fmt.Scan(&powtorka)
			if powtorka == "tak" {
				gra(players)
			} else if powtorka == "nie" {
				leaderboard(players)
				os.Exit(1)
			} else {
				fmt.Println("Nie zrozumiałem, więc wychodzisz")
				os.Exit(1)
			}

		}
	}
}

func leaderboard(players map[string]int) {
	fmt.Println("Tak prezentuje się tablica wyników")
	fmt.Println("-----------")
	for player, score := range players {
		fmt.Println(player, score)

	}
	fmt.Println("-----------")
}
