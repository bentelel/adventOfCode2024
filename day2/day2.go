package day2

import (
	"fmt"
	"strconv"

	"bentelel/adventOfCode2024/utils"
)

const (
	inputDirectory string = "day2"
	lowerBound     int    = 1
	upperBound     int    = 3
)

func A() {
	fmt.Print("running day 2A.\n")
	puzzleInput, err := utils.ReadFileToString("input.txt", inputDirectory)
	if err != nil {
		panic(err)
	}
	levels := utils.GetRowsFromInput(puzzleInput)
	//for _, l := range levels {
	//	fmt.Printf("level: %s\n", l)
	//}
	levelsAndRooms := utils.GetRowsAndElements(levels)
	levelsAndRooms_temp := [][]string{}
	// first check -- is a row all ascending or all descending? if yes, we press on. if no, we ignore it.
	for _, l := range levelsAndRooms {
		allAscOrDesc, err := utils.IsSliceAllAscendingOrDescending(l)
		if err != nil {
			panic(err)
		}
		// fmt.Printf("level: %v, is desc/asc: %t\n", l, allAscOrDesc)
		if allAscOrDesc {
			levelsAndRooms_temp = append(levelsAndRooms_temp, l)
		} // else {
		//	fmt.Printf("unsafe level (not asc/desc): %v\n", l)
		// [9 11 13 14 17 24]
		//}
	}
	levelsAndRooms = levelsAndRooms_temp
	levelsAndRooms_temp = [][]string{}
	// second check -- is the distance between all elements in each room ok?
	for _, l := range levelsAndRooms {
		distancesAreOk, err := utils.AreDistancesOk(l, lowerBound, upperBound)
		if err != nil {
			panic(err)
		}
		// fmt.Printf("level: %v, distance is ok: %t\n", l, distancesAreOk)
		if distancesAreOk {
			levelsAndRooms_temp = append(levelsAndRooms_temp, l)
		} // else {
		//fmt.Printf("unsafe level (bounds): %v\n", l)
		//}
	}
	levelsAndRooms = levelsAndRooms_temp
	// fmt.Printf("final levels: %v\n", levelsAndRooms)
	countOfSaveLevels := len(levelsAndRooms)
	fmt.Printf("The count of save levels is: %d\n", countOfSaveLevels)
}

func B() {
	fmt.Print("running day 2B.\n")
	puzzleInput, err := utils.ReadFileToString("input.txt", inputDirectory)
	if err != nil {
		panic(err)
	}
	levels := utils.GetRowsFromInput(puzzleInput, "\n")
	levelsAndRooms := utils.GetRowsAndElements(levels, " ")
	levelsAndRooms_temp := [][]string{}
	// first check -- is a row all ascending or all descending? if yes, we press on. if no, we ignore it.
	fmt.Printf("The count of save levels is: %d\n", 0)
}
