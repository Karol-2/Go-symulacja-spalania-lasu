package las

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

func ObliczProecnty(data map[int][]int) map[int]float64 {
	averages := make(map[int]float64)

	for numTrees, values := range data {
		totalTrees := float64(values[0])
		burnedTrees := float64(values[1])

		percentage := (burnedTrees / totalTrees) * 100
		averages[numTrees] = percentage
	}

	return averages
}

func Analiza() {
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

	averages := ObliczProecnty(data)

	for numTrees, percentage := range averages {
		fmt.Printf("Ilość drzew: %d, Procent drzew które przetrwały: %.2f%%\n", numTrees, percentage)
	}
}
