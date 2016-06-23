package main

import (
	"log"
	"os"
)

func main() {
	f, err := os.Create(os.Args[1]) //second argument
	if err != nil {
		log.Fatalln("Error here")
	}

	f.WriteString(os.Args[2]) // now takes third argument, you can pass a string

	//example would be in terminal
	//$Create something.txt "Here are some words"
	//this will create a local file called osmething.txt with a string here are some words inside
}
