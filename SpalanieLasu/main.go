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
// ogień porusza się tylko prawo lewo góra dół

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
	pozycjaStrzaluX := rand.Intn(len(las[0]))
	pozycjaStrzaluY := rand.Intn(len(las))

	fmt.Println("Pozycja strzału", pozycjaStrzaluX, pozycjaStrzaluY)

	if las[pozycjaStrzaluY][pozycjaStrzaluX] == 1 {
		las[pozycjaStrzaluY][pozycjaStrzaluX] = 2
		fmt.Println("Strzał trafił na drzewo!")
		rozprzestrzenianieOgnia(las, pozycjaStrzaluX, pozycjaStrzaluY)
	} else {
		fmt.Println("Strzał nie trafił w drzewo!")
	}

}

func zapiszDoCSV(rozmiar string, iloscDrzew, pozostaleDrzewa, spaloneDrzewa int) {
	file, err := os.OpenFile("data.csv", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal("Nie można otworzyć pliku CSV:", err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	row := []string{
		rozmiar,
		strconv.Itoa(iloscDrzew),
		strconv.Itoa(pozostaleDrzewa),
		strconv.Itoa(spaloneDrzewa),
	}

	err = writer.Write(row)
	if err != nil {
		log.Fatal("Błąd podczas zapisu do pliku CSV:", err)
	}

	writer.Flush()

	log.Println("Dane zostały zapisane do pliku CSV.")
}

func sadzenieDrzew(iloscDrzew int, las [][]int) {
	posadzoneDrzewa := 0
	for posadzoneDrzewa < iloscDrzew {
		for i := 0; i < len(las[0]); i++ {
			for j := 0; j < len(las); j++ {
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
func generujDzialke(szerokosc, wysokosc int) [][]int {
	las := make([][]int, wysokosc)
	for i := 0; i < wysokosc; i++ {
		las[i] = make([]int, szerokosc)
	}
	return las
}
func symulacja(szerokosc, wysokosc, iloscDrzew int) {
	if szerokosc*wysokosc <= iloscDrzew {
		fmt.Println("ERROR!, ilość drzew przekracza maksymalną ilość: ", szerokosc*wysokosc)
		return
	}

	las := generujDzialke(szerokosc, wysokosc)

	sadzenieDrzew(iloscDrzew, las)
	drzewaPrzed := policzDrzewa(las)
	pokazPlansze(las)
	ogien(las)
	pokazPlansze(las)
	drzewaPo := policzDrzewa(las)
	spaloneDrzewa := drzewaPrzed - drzewaPo
	fmt.Println("Ilość drzew:", drzewaPo)
	fmt.Println("Spalone drzewa:", spaloneDrzewa)
	rozmiar := fmt.Sprintf("%dx%d", szerokosc, wysokosc)
	zapiszDoCSV(rozmiar, iloscDrzew, drzewaPo, spaloneDrzewa)
}

func obliczProecnty(data map[int][]int) map[int]float64 {
	averages := make(map[int]float64)

	for numTrees, values := range data {
		totalTrees := float64(values[0])
		burnedTrees := float64(values[1])

		percentage := (burnedTrees / totalTrees) * 100
		averages[numTrees] = percentage
	}

	return averages
}

func analiza() {
	file, err := os.Open("data.csv")
	if err != nil {
		fmt.Println("Błąd podczas otwierania pliku:", err)
		return
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.Comma = ','

	data := make(map[int][]int)

	lines, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Błąd podczas odczytywania danych z pliku:", err)
		return
	}

	for _, line := range lines {

		numTrees, err := strconv.Atoi(line[1])
		if err != nil {
			fmt.Println("Błąd podczas konwersji ilości drzew:", err)
			continue
		}

		treesPo, err := strconv.Atoi(line[2])
		if err != nil {
			fmt.Println("Błąd podczas konwersji drzew po spaleniu:", err)
			continue
		}

		burnedTrees, err := strconv.Atoi(line[3])
		if err != nil {
			fmt.Println("Błąd podczas konwersji spalonych drzew:", err)
			continue
		}

		data[numTrees] = []int{numTrees, treesPo, burnedTrees}
	}

	averages := obliczProecnty(data)

	for numTrees, percentage := range averages {
		fmt.Printf("Ilość drzew: %d, Procent drzew które przetrwały: %.2f%%\n", numTrees, percentage)
	}
}
func pokazOpcje() {
	fmt.Println("==========================================")
	fmt.Println("Podaj numer wybranej opcji.")
	fmt.Println("1. Pojedyncza symulacja 15x10")
	fmt.Println("2. Symulacja 15x10 dla podanej ilości drzew i ilości eksperymentów (testy po 10)")
	fmt.Println("3. Symulacja na wszystkie dowolne parametry")
	fmt.Println("4. Analiza danych z pliku data.csv")
	fmt.Println("5. Wyczyszczenie pliku data.csv")
	fmt.Println("6. Informacje")
	fmt.Println("7. Wyjście")
	fmt.Println("==========================================")
}
func opcjaPierwsza() {
	szerokosc := 15
	wysokosc := 10
	var iloscDrzew int

	fmt.Print("Podaj ilość drzew: ")
	fmt.Scan(&iloscDrzew)
	symulacja(szerokosc, wysokosc, iloscDrzew)
}
func opcjaDruga() {
	szerokosc := 15
	wysokosc := 10

	var iloscDrzew int
	fmt.Print("Podaj ilość drzew: ")
	fmt.Scan(&iloscDrzew)

	var iloscEksperymentow int
	fmt.Print("Podaj ilość eksperymentów: ")
	fmt.Scan(&iloscEksperymentow)

	for ilosc := 10; ilosc < 150; ilosc += 10 {
		for i := 0; i < iloscEksperymentow; i++ {
			symulacja(szerokosc, wysokosc, ilosc)
		}
	}
	symulacja(szerokosc, wysokosc, iloscDrzew)
}
func opcjaTrzecia() {
	var szerokosc int
	fmt.Print("Podaj szerokość lasu: ")
	fmt.Scan(&szerokosc)

	var wysokosc int
	fmt.Print("Podaj wysokość lasu: ")
	fmt.Scan(&wysokosc)

	var iloscDrzew int
	fmt.Print("Podaj ilość drzew: ")
	fmt.Scan(&iloscDrzew)

	var iloscEksperymentow int
	fmt.Print("Podaj ilość eksperymentów: ")
	fmt.Scan(&iloscEksperymentow)

	var wartoscStartowaEksperymentu int
	fmt.Print("Podaj najmniejszą ilość drzew (wartość startowa): ")
	fmt.Scan(&wartoscStartowaEksperymentu)

	var zwiekszenie int
	fmt.Print("Podaj o jaką wartość mają się różnić eksperymenty: ")
	fmt.Scan(&zwiekszenie)

	for ilosc := wartoscStartowaEksperymentu; ilosc < (wysokosc * szerokosc); ilosc += zwiekszenie {
		for i := 0; i < iloscEksperymentow; i++ {
			symulacja(szerokosc, wysokosc, ilosc)
		}
	}
	symulacja(szerokosc, wysokosc, iloscDrzew)
}
func wyczyscPlik() {
	filePath := "data.csv"

	file, err := os.OpenFile(filePath, os.O_TRUNC|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("błąd podczas otwierania pliku: %v", err)
		return
	}
	defer file.Close()

	err = file.Truncate(0)
	if err != nil {
		fmt.Println("błąd podczas czyszczenia zawartości pliku: %v", err)
		return
	}

	fmt.Println("Wyczyszczono plik data.csv!")
}
func informacje() {
	fmt.Println("Ten program symuluje spalanie się lasu.")
	fmt.Println("Zakładamy że piorun trafia w losowe miejsce w lesie, w zależności od tego czy jest na tym polu drzewo, może on wzniecić ogień lub nie.")
	fmt.Println("Jeżeli jest jakieś drzewo z prawej/lewej/pod/nad to ogień się tam rozprzestrzeni.")
	fmt.Println("Oznaczenia:")
	fmt.Println("0 - puste pole (brak drzewa)")
	fmt.Println("1 - drzewo (nie spalone)")
	fmt.Println("2 - ogień (drzewo spalone)")
	fmt.Println("Analiza:")
	fmt.Println("Dane z pliku data.csv są grupowane według ilości drzew i następnie obliczany jest średni procent drzew które przetrwały.")
	fmt.Println("Wyniki są podawane w losowej kolejności.")
	fmt.Println("Autor: Karol Krawczykiewicz - Gdańsk 2023")
}
func main() {
	dzialanie := true
	for dzialanie {
		pokazOpcje()
		var opcja int

		fmt.Print("Wybrana opcja: ")
		fmt.Scan(&opcja)

		switch opcja {
		case 1:
			opcjaPierwsza()
		case 2:
			opcjaDruga()
		case 3:
			opcjaTrzecia()
		case 4:
			analiza()
		case 5:
			wyczyscPlik()
		case 6:
			informacje()
		case 7:
			dzialanie = false
			os.Exit(0)
		default:
			fmt.Println("Podaj inną wartość!")
		}
	}
}
