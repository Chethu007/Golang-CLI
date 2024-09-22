package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	wordLen := 0
	for scanner.Scan() {
		words := strings.Fields(scanner.Text())
		wordLen += len(words)
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "reading standard input: %v\n", err.Error())
		os.Exit(1)
	}
	fmt.Println("Total words:", wordLen)
}
