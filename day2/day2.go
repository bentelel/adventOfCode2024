package day2

import (
	"fmt"
	"sync"

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
	levels := utils.GetRowsFromInput(puzzleInput, "\n")
	//for _, l := range levels {
	//	fmt.Printf("level: %s\n", l)
	//}
	levelsAndRooms := utils.GetRowsAndElements(levels, "\n")
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
	// what we will do is: we will run the 2 checks for all rows individually.
	// we pass each row of the input into the function.
	// the function will then run the 2 checks for all sub-rows (which are formed by omitting 1 element of the row)
	// we will do that as go-routines to at least not completely trash performance.
	// we will get the results into a channel. if we get at least 1 positive result in a channel per row, that means that the row is safe. in that case we will pass a positive result up to the top level func which then will increment the count of safe levels.
	// func B()
	//    rowProcessor()
	//        open channel
	//        go bothChecks()
	//            return true/false into channel
	var saveLevels []bool
	for _, l := range levelsAndRooms {
		result, err := rowProcessor(l)
		if err != nil {
			panic(err)
		}
		saveLevels = append(saveLevels, result)
	}
	countOfSaveLevels := 0
	for _, l := range saveLevels {
		if l {
			countOfSaveLevels += 1
		}
	}
	fmt.Printf("The count of save levels is: %d\n", countOfSaveLevels)
}

func rowProcessor(row []string) (bool, error) {
	// fmt.Printf("starting rowprocessor for: %v\n", row)
	errors := make(chan error, len(row))
	results := make(chan bool, len(row))
	wg := &sync.WaitGroup{}
	for i := 0; i < len(row); i++ {
		wg.Add(1)
		partialRow := utils.DropElementAtIndex(row, i)
		// fmt.Printf("\tchecking part row: %v\n", partialRow)
		go func(partialRow []string) {
			defer wg.Done()
			r, e := utils.OrderAndDistanceCheck(partialRow, lowerBound, upperBound)
			results <- r
			errors <- e
		}(partialRow)
	}

	go func() {
		wg.Wait()
		close(errors)
		close(results)
	}()
	for e := range errors {
		// fmt.Printf("\t\t error: %e\n", e)
		if e != nil {
			return false, nil
		}
	}
	for r := range results {
		// fmt.Printf("\t\t res: %t\n", r)
		if r {
			// fmt.Printf("\trow: %v, is safe.\n", row)
			return true, nil
		}
	}
	return false, nil
}
