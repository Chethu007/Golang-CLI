package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"
)

func main() {
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	go prompter()
	for {
		select {
		case res := <-sig:
			signal.Stop(sig)
			fmt.Println(" Got signal:", res)
			fmt.Println("Performing cleanup and shutting down...")
			os.Exit(0)
		}
	}
}

func prompter() {
	fmt.Print(">> ")
	scanner := bufio.NewScanner(os.Stdin)
	file, err := os.Create("log.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		fmt.Println("<-", line)
		//io.Copy(file, os.Stdin)
		file.WriteString(line + "\n")
		fmt.Print(">> ")
	}
	if err = scanner.Err(); err != nil {
		fmt.Println("Error:", err.Error())
		os.Exit(1)
	}
}
