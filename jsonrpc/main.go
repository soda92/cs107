package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
)

func main() {
	fmt.Println("Starting LSP over stdin/stdout")
	for{
		reader := bufio.NewReader(os.Stdin)
		line, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("read line: %s-\n", line)
	}
}