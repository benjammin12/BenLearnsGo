package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	r, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatalln("Error opening file", err)
	}

	numOf := wordCount(r)
	fmt.Println("Number of whales is:", numOf["whales"])
}

func wordCount(rd io.Reader) map[string]int { //word count function takes in a string
	//and returns a map of words in the string and their ammounts map[words(string)]amount(int)
	counter := map[string]int{} //create a string

	moby := bufio.NewScanner(rd)
	moby.Split(bufio.ScanWords)

	for moby.Scan() {
		line := moby.Text()
		line = strings.ToLower(line)
		counter[line]++
	}

	return counter
}
