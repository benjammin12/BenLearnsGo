package main

import (
	"fmt"
	"strings"
)

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
	} else {
		fmt.Print("Not in database")
	}

	str := make(map[string]int)
	str["cats"] = 4000000
	str["dogs"] = 6000000
	str["birds"] = 34531

	fmt.Println(str)

	for key, value := range str {
		fmt.Println("Animal:", key, ",Amount:", value)
	}

	sx := "Hello, from macbook pro pro pro"

	total := wordCount(sx)

	fmt.Println("Total amount of the word pro:", total["pro"])

	//USING WORDCOUNT func

	//

}

func wordCount(str string) map[string]int { //word count function takes in a string
	//and returns a map of words in the string and their ammounts map[words(string)]amount(int)
	map1 := map[string]int{} //create a string

	for _, word := range strings.Fields(str) { //strings.Fields turns the string into a alice of words
		map1[word]++ //for every string, increment the count
	}

	return map1
}
