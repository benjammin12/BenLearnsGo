package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	/*
		fmt.Print("Enter file name: ") //stores file name into a variable
		var fileName string
		fmt.Scanln(&fileName)

		f, err := os.Open(fileName) //os.Open returns an associated file and an error
	*/

	f, err := os.Open(os.Args[1])
	if err != nil { //standard error handling function, can usually get by using this
		log.Fatalln("File doesn't exist", err.Error())
	} //if there is no error, then you can use the open file

	defer f.Close()

	bs, err := ioutil.ReadAll(f) // a sucesful call returns error == nil
	if err != nil {              // if there is no error then....
		log.Fatalln("File doesn't exist here")
	}

	convert := string(bs) //Readall is returning a silce of byes and an error
	//so we need to convert the bytes into a string

	fmt.Println(convert) //we can print variable we stored read all in
}
