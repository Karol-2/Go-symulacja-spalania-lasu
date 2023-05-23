package main

import "fmt"

func main() {

	wektor1 := make([]int64, 20)
	wektor2 := make([]int64, 20)
	zsumowane := make([]int64, 20)

	for i := 0; i < 20; i++ {
		wektor1[i] = 2
		wektor2[i] = 3
		zsumowane[i] = wektor1[i] + wektor2[i]
	}

	fmt.Println("wektor1: ", wektor1)
	fmt.Println("wektor2: ", wektor2)
	fmt.Println("suma wektorów", zsumowane)

}
