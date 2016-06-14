package main

import (
	"io/ioutil"
	"log"
	"os"
)

func main() {

	srcFile, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatalln("No argument", err)
	}

	reader, err := ioutil.ReadAll(srcFile)
	if err != nil {
		log.Fatalln("Can't read file", err)
	}

	dst, err := os.Create(os.Args[2])
	if err != nil {
		log.Fatalln("Can't create file", err)
	}

	dst.Write(reader)

	defer dst.Close()
	defer srcFile.Close()

}
