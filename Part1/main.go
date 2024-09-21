package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	Greeting(os.Stdout, "Chethan V", "Welcome")
	getInputFromUser()
	Greeting1("Chethan")
}

func Greeting1(s string) {
	fmt.Println("Hi", s)
}

func Greeting(out io.Writer, s, s1 string) {
	n, err := fmt.Fprintln(out, fmt.Sprintf("Hello %s!!", s), s1)
	if err != nil {
		log.Fatal(err)
	}
	println("No of lines", n)
}

func getInputFromUser() {
	file, err := os.Create("output.txt")
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error creating the file:", err)
	}
	defer file.Close()
	fmt.Println("Please enter text (Ctrl+D to stop):")

	// Copy data from standard input (os.Stdin) to the file
	_, err = io.Copy(file, os.Stdin)
	if err != nil {
		fmt.Println("Error copying input:", err)
		return
	}
	fmt.Println("Input successfully written to output.txt")
}
