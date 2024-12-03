package day2

import (
	"fmt"

	"bentelel/adventOfCode2024/utils"
)

const inputDirectory string = "day2"

func A() {
	fmt.Print("running day 1A.\n")
	puzzleInput, err := utils.ReadFileToString("input.txt", inputDirectory)
	if err != nil {
		panic(err)
	}
	levels := utils.GetRowsFromInput
	fmt.Printf("The overall sum of differences is: %d\n", 0)
}
