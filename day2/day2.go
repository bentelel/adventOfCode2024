package day2

import (
	"fmt"

	"bentelel/adventOfCode2024/utils"
)

const inputDirectory string = "day2"

func A() {
	fmt.Print("running day 2A.\n")
	puzzleInput, err := utils.ReadFileToString("input.txt", inputDirectory)
	if err != nil {
		panic(err)
	}
	levels := utils.GetRowsFromInput(puzzleInput)
	for _, l := range levels {
		fmt.Printf("level: %s\n", l)
	}
	levelsAndRooms := utils.GetRowsAndElements(levels)
	for _, l := range levelsAndRooms {
		fmt.Printf("level: %s\n", l)
	}

	for _, l := range levelsAndRooms {
		allAscOrDesc := utils.IsSliceAllAscendingOrDescending(l)
		fmt.Printf("level: %v, is desc/asc: %t\n", l, allAscOrDesc)

	}

	fmt.Printf("The overall sum of differences is: %d\n", 0)
}
