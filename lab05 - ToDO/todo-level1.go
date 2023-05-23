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
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error:", err)
	}

}
