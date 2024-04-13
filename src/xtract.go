package main

import (
	"bufio"
	"fmt"
	"os"
	"tools/src/search"

	"github.com/jessevdk/go-flags"
)

type Options struct {
	Color bool `short:"c" long:"color" description:"Colour key and value"`
}

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

	var buffer string

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		buffer = scanner.Text()
		//fmt.Println(buffer) // Println will add back the final '\n'
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "sin datos en stdin!!", err)
	}

	//buffer = "timeFake=\"yo\" time=\"2024-04-05 11:39:27.864\" level=INFO module=helpers head_rest=\"User-Agent:[Mozilla/4.0 [en] (WinNT; I)] X-Forwarded-For:[76.249.26.144, 172.29.99.68, 172.29.98.125, 10.131.0.1] X-Forwarded-Port:[443]\""

	key := os.Args[1]

	var opts Options
	_, err := flags.Parse(&opts)
	if err != nil {
		// Handle error
		fmt.Println(err)
		return
	}

	valor, err := search.GetSmartValue(buffer, key)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	} else {
		fmt.Printf("%s => %s\n", key, valor)
	}

}
