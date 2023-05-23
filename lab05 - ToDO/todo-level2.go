package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	file, err := os.Open("./lab05 - ToDO/todo.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		if line[0:1] == "#" {
			continue
		}

		switch line[0:3] {
		case "[x]":
			fmt.Println("\033[32m", line, "\033[0m") // zielony
		case "[ ]":
			fmt.Println("\033[31m", line, "\033[0m") // czerwony
		case "[-]":
			fmt.Println("\033[90m", line, "\033[0m") // szary
		case "[+]":
			fmt.Println("\033[33m", line, "\033[0m") // żółty
		default:
			fmt.Println(line)
		}

	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error:", err)
	}

}
