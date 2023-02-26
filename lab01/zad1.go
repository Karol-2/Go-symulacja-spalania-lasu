package main

import (
	"fmt"
	"strconv"
)

func main() {

	fmt.Println("Ile masz lat?: ")
	var wiek string
	fmt.Scanln(&wiek)
	fmt.Println("masz " + wiek + " lat")
	lat, err := strconv.ParseInt(wiek, 10, 0)

	if err != nil {
		fmt.Println("Error during conversion")
		return
	}

	fmt.Println("W miesiącach:")
	fmt.Println(lat * 12)
	fmt.Println("W dniach:")
	fmt.Println(lat * 12 * 365)

}
