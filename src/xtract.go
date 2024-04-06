package main

import (
	"bufio"
	"fmt"
	"os"
	"tools/src/search"
)

/*
func help(exit bool) {
	fmt.Println("Parseo de log barcode-DNI")
	fmt.Println("")
	fmt.Println(" parseLog <file> ")
	fmt.Println("")
	if exit {
		os.Exit(1)
	}
}*/

func main() {

	//status := search.GetValueFromKey(buff, "status=", " ")
	scanner := bufio.NewScanner(os.Stdin)
	var buffer string
	//buffer := scanner.Text()
	//fmt.Println(buffer) // Println will add back the final '\n'

	for scanner.Scan() {
		buffer = scanner.Text()
		//fmt.Println(buffer) // Println will add back the final '\n'
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}

	//fmt.Println(buffer)

	key := os.Args[1]

	valor := search.GetValueFromKeyWithDelimiters(buffer, key, "[", "]", true)
	fmt.Println(valor)
}
