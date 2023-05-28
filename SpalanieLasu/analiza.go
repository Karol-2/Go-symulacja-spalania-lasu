package main

import (
	"encoding/csv"
	"log"
	"os"
	"strconv"
)

func generujWykres() {
	file, err := os.Open("data.csv")
	if err != nil {
		log.Fatal("Nie można otworzyć pliku CSV:", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	rows, err := reader.ReadAll()
	if err != nil {
		log.Fatal("Błąd podczas odczytu danych z pliku CSV:", err)
	}

	var rozmiary []string
	var ilosciDrzew []float64

	for _, row := range rows {
		rozmiar := row[0]
		iloscDrzew, err := strconv.ParseFloat(row[1], 64)
		if err != nil {
			log.Fatal("Błąd podczas przetwarzania danych z pliku CSV:", err)
		}

		rozmiary = append(rozmiary, rozmiar)
		ilosciDrzew = append(ilosciDrzew, iloscDrzew)
	}
}

func main() {
	generujWykres()
}
