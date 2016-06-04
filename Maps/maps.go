package main

import "fmt"

func main() {

	fmt.Println("Enter a city abbreviation to check our database. Ex. SD")
	var name string
	fmt.Scanln(&name)

	cities := map[string]string{
		"SD":  "San Diego",
		"NYC": "New York City",
		"SF":  "San Francisco",
		"LA":  "Los Angelas",
		"MI":  "Miami",
	}

	if c, ok := cities[name]; ok {
		fmt.Println(c, ok)
	}
}
