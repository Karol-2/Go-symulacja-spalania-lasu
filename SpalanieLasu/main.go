package main

import "spalanie-lasu/las"
import (
	"fmt"
	"os"
)

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
	las.Symulacja(szerokosc, wysokosc, iloscDrzew)
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
			las.Symulacja(szerokosc, wysokosc, ilosc)
		}
	}
	las.Symulacja(szerokosc, wysokosc, iloscDrzew)
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
			las.Symulacja(szerokosc, wysokosc, ilosc)
		}
	}
	las.Symulacja(szerokosc, wysokosc, iloscDrzew)
}
func wyczyscPlik() {
	wd, _ := os.Getwd()
	fmt.Println(wd)
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
			las.Analiza()
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
