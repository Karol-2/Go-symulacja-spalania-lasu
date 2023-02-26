package main

import (
	"fmt"
) 

func main() { 

fmt.Println("Ile masz lat?: ") 
var wiek float64 
fmt.Scanln(&wiek) 

 
fmt.Println("Na Merkurym:") 
fmt.Println(wiek / 0.24085) 
fmt.Println("Na Wenus:") 
fmt.Println(wiek / 0.61521) 
fmt.Println("Na Mars:") 
fmt.Println(wiek / 1.88089) 
fmt.Println("Na Jowisz:") 
fmt.Println(wiek / 11.8622) 
fmt.Println("Na Saturn:") 
fmt.Println(wiek / 29.4577) 
fmt.Println("Na Uran:") 
fmt.Println(wiek / 84.0153) 
fmt.Println("Na Neptun:") 
fmt.Println(wiek / 164.788) 
} 