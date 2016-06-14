package main

/*
write a program to assign a number 1-7 through each of
3 departments
-fire
-police != odd number
-sanitation
can't give gthe same number to two departments
sum must equal 12
*/

import "fmt"

func main() {
	permutations := 0
	for fd := 1; fd <= 7; fd++ {
		for pd := 2; pd <= 6; pd += 2 {
			for sd := 1; sd <= 7; sd++ {

				if isValidNumber(fd, pd, sd) {
					fmt.Println(fd, pd, sd)
					permutations++
				}

			} // end sd
		} //end pd
	} //end fd

	fmt.Println("Total permutations =", permutations)

}

func isValidNumber(fd, pd, sd int) bool {
	return (fd != pd && pd != sd && fd != sd && fd+pd+sd == 12)
}
