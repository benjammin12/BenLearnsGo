package main

import "fmt"

func main() {
	nums := []int{3, 3, 5, 4, 4, 9, 1, 3, 8, 8, 9, 9, 9, 9, 7, 10, 10}
	fmt.Println(countClumps(nums))
}

func countClumps(slc []int) int {
	count := 0

	for i := 2; i < len(slc); i++ {

		if i == (len(slc) - 1) {
			if slc[i] == slc[i-1] {
				count++
			}
		}

		if slc[i] != slc[i-1] && slc[i-2] == slc[i-1] {
			count++
		}
	}

	return count

}
