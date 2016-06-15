package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

//States representing a map of information for a our csv file
type States struct {
	id     int
	name   string
	abbr   string
	region string
}

/*
func parseState() (States, error) {

}
*/

func main() {
	r, err := os.Open("state_table.csv")
	if err != nil {
		log.Fatalln("Error opening", err)
	}

	read := csv.NewReader(r)

	stateInfo, err := read.ReadAll()
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(stateInfo)

}
