package main

import "fmt"

func main() {
    n := 3

    for n <= 100{                           //goes for loop is similar to a while loop {} is necessary
        if n % 5 == 0 && n % 3 ==0 {        //Parenthesis around conditions of if statement is optional
            fmt.Println("FizzBuzz")         //Curly brackets are REQUIRED for if statements in Go
        }else if n % 5 == 0 {
            fmt.Println("Buzz")
        }else if n % 3 == 0 {
            fmt.Println("Fizz")
        }else {
            fmt.Println(n)
        }
        n++                                 // you have to increment your counter otherwise you'll have an infinite loop
    }

}
