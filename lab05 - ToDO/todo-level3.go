package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

func main() {

	file, err := os.Open("./lab05 - ToDO/todo.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	color := flag.Bool("nocolor", false, "disable color output")
	nowe := flag.Bool("new", false, "show only new todos")
	zrobione := flag.Bool("done", false, "show only done todos")
	odrzucone := flag.Bool("cancelled", false, "show only disabled todos")
	flag.Parse()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		if line[0:1] == "#" {
			continue
		}
		if !(*nowe || *zrobione || *odrzucone) {
			if !*color {
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
			} else {
				fmt.Println(line)
			}
		} else if *nowe {
			if !*color {
				switch line[0:3] {
				case "[ ]":
					fmt.Println("\033[31m", line, "\033[0m") // czerwony
				default:
					continue
				}
			} else {
				fmt.Println(line)
			}

		} else if *zrobione {
			if !*color {
				switch line[0:3] {
				case "[x]":
					fmt.Println("\033[32m", line, "\033[0m") // zielony
				default:
					continue
				}
			} else {
				fmt.Println(line)
			}
		} else if *odrzucone {
			if !*color {
				switch line[0:3] {
				case "[-]":
					fmt.Println("\033[90m", line, "\033[0m") // szary
				default:
					continue
				}
			} else {
				fmt.Println(line)
			}
		}

	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error:", err)
	}

}
