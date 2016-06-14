package main

import (
	"log"
	"os"
)

func main() {
	f, err := os.Create(os.Args[1])
	if err != nil {
		log.Fatalln("Error here")
	}

	f.WriteString("Something will go here\n")
}
