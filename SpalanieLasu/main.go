package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
)

// 0 - puste
// 1 - drzewo
// 2 - ogień
// porusza się tylko prawo lewo góra dół

func policzDrzewa(las [][]int) int {
	drzewa := 0
	for i := 0; i < len(las); i++ {
		for j := 0; j < len(las); j++ {
			if las[j][i] == 1 {
				drzewa++
			}
		}
	}
	return drzewa
}
func pokazPlansze(las [][]int) {
	for i := 0; i < len(las); i++ {
		for j := 0; j < len(las[i]); j++ {
			switch las[i][j] {
			case 0:
				fmt.Print("\033[0;37m0\033[0m") // szary
			case 1:
				fmt.Print("\033[0;32m1\033[0m") // zielony
			case 2:
				fmt.Print("\033[0;31m2\033[0m") // czerwony
			}
		}
		fmt.Println()
	}
}
func rozprzestrzenianieOgnia(las [][]int, pozycjaX int, pozycjaY int) {
	if pozycjaY < 0 || pozycjaY >= len(las) ||
		pozycjaX < 0 || pozycjaX >= len(las[pozycjaY]) {
		return
	}

	// Sprawdzanie górnego sąsiada
	if pozycjaY > 0 && las[pozycjaY-1][pozycjaX] == 1 {
		las[pozycjaY-1][pozycjaX] = 2
		rozprzestrzenianieOgnia(las, pozycjaX, pozycjaY-1)
	}

	// Sprawdzanie dolnego sąsiada
	if pozycjaY < len(las)-1 && las[pozycjaY+1][pozycjaX] == 1 {
		las[pozycjaY+1][pozycjaX] = 2
		rozprzestrzenianieOgnia(las, pozycjaX, pozycjaY+1)
	}

	// Sprawdzanie lewego sąsiada
	if pozycjaX > 0 && las[pozycjaY][pozycjaX-1] == 1 {
		las[pozycjaY][pozycjaX-1] = 2
		rozprzestrzenianieOgnia(las, pozycjaX-1, pozycjaY)
	}

	// Sprawdzanie prawego sąsiada
	if pozycjaX < len(las[pozycjaY])-1 && las[pozycjaY][pozycjaX+1] == 1 {
		las[pozycjaY][pozycjaX+1] = 2
		rozprzestrzenianieOgnia(las, pozycjaX+1, pozycjaY)
	}
}

func ogien(las [][]int) {
	pozycjaStrzaluX := rand.Intn(len(las))
	pozycjaStrzaluY := rand.Intn(len(las[0]))
	fmt.Println("Pozycja strzału", pozycjaStrzaluX, pozycjaStrzaluY)

	if las[pozycjaStrzaluY][pozycjaStrzaluX] == 1 {
		las[pozycjaStrzaluY][pozycjaStrzaluX] = 2
		fmt.Println("Strzał trafił na drzewo!")
		rozprzestrzenianieOgnia(las, pozycjaStrzaluX, pozycjaStrzaluY)
	} else {
		fmt.Println("Strzał nie trafił w drzewo!")
	}

}

func zapiszDoCSV(rozmiarPlanszy string, iloscDrzew, pozostaleDrzewa, spaloneDrzewa int) {
	file, err := os.Create("dane.csv")
	if err != nil {
		log.Fatal("Nie można utworzyć pliku CSV:", err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	row := []string{rozmiarPlanszy, strconv.Itoa(iloscDrzew), strconv.Itoa(pozostaleDrzewa), strconv.Itoa(spaloneDrzewa)}
	err = writer.Write(row)
	if err != nil {
		log.Fatal("Błąd podczas zapisu danych do pliku CSV:", err)
	}
}

func sadzenieDrzew(iloscDrzew int, las [][]int) {
	posadzoneDrzewa := 0
	for posadzoneDrzewa < iloscDrzew {
		for i := 0; i < 10; i++ {
			for j := 0; j < 10; j++ {
				if posadzoneDrzewa == iloscDrzew {
					break
				}

				czyPosadzic := rand.Intn(2)
				if czyPosadzic == 0 && las[j][i] == 0 {
					las[j][i] = 1
					posadzoneDrzewa++
				}
			}
		}

	}
}

func main() {
	las := [][]int{ //10x10
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	}
	iloscDrzew := 30

	sadzenieDrzew(iloscDrzew, las)
	pokazPlansze(las)
	ogien(las)
	pokazPlansze(las)
	pozostaleDrzewa := policzDrzewa(las)
	spaloneDrzewa := iloscDrzew - pozostaleDrzewa
	fmt.Println("Ilość drzew: ", pozostaleDrzewa)
	fmt.Println("Spalone drzewa: ", spaloneDrzewa)
	zapiszDoCSV("10x10", iloscDrzew, pozostaleDrzewa, spaloneDrzewa)
}
