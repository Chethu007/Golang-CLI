package main

import (
	"flag"
	"fmt"
)

func main() {
	//fmt.Println("Exiting with code 123")
	//os.Exit(123)
	//echo $? - To check last exit status
	userName := flag.String("username", "", "username to authenticate")
	var password, proxy string
	flag.StringVar(&password, "password", "", "password to authenticate")
	flag.StringVar(&proxy, "proxy-url", "", "Go proxy URL")
	var port int
	flag.IntVar(&port, "port", 8080, "Port to listen")
	flag.Parse()
	fmt.Println("userName:", *userName)
	fmt.Println("password:", password)
	fmt.Println("proxy url:", proxy)
	fmt.Println("port:", port)
}
