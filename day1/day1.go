package day1

import (
	"fmt"
	"strconv"

	"bentelel/adventOfCode2024/utils"
)

const inputDirectory string = "day1"

func A() {
	fmt.Print("running day 1A.\n")
	puzzleInput, err := utils.ReadFileToString("inputA.txt", inputDirectory)
	if err != nil {
		panic(err)
	}
	// fmt.Printf("puzzle input:\n%s", puzzleInput)
	leftInput, rightInput := utils.GetListsFromInput(puzzleInput)
	//for _, v := range leftInput {
	//	fmt.Printf("left input: %s\n", v)
	//}
	//for _, v := range rightInput {
	//	fmt.Printf("right input: %s\n", v)
	//}
	leftInputSorted, err := utils.SortSliceAscending(leftInput)
	rightInputSorted, err := utils.SortSliceAscending(rightInput)
	diffs, err := utils.FindDifferencesBetweenSlices(leftInputSorted, rightInputSorted)
	if err != nil {
		panic(err)
	}
	var sum int = 0
	for _, d := range diffs {
		//	v := strconv.Itoa(d)
		//	fmt.Printf("dif: %s\n", v)
		sum += d
	}
	sumString := strconv.Itoa(sum)
	fmt.Printf("The overall sum of differences is: %s\n", sumString)
}

func B() {
	fmt.Print("running day 1B.\n")
	puzzleInput, err := utils.ReadFileToString("inputA.txt", inputDirectory)
	if err != nil {
		panic(err)
	}
	leftInput, rightInput := utils.GetListsFromInput(puzzleInput)

	countsOfNumsInRight, err := utils.GetCountOfNumbers(rightInput)
	if err != nil {
		panic(err)
	}
	var totalScore int = 0
	for _, l := range leftInput {
		score, err := utils.GetSimilarityScore(utils.StripTrailingNewlines(l), countsOfNumsInRight)
		if err != nil {
			panic(err)
		}
		totalScore += score
	}
	totalScoreString := strconv.Itoa(totalScore)
	fmt.Printf("The overall sum of similarity scores is: %s\n", totalScoreString)
}
