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
	leftInput, rightInput := utils.GetListsFromInput(puzzleInput, "\n")
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
	leftInput, rightInput := utils.GetListsFromInput(puzzleInput, "\n")

	countsOfNumsInRight, err := utils.GetCountOfNumbers(rightInput)
	if err != nil {
		panic(err)
	}
	var totalScore int = 0
	for _, l := range leftInput {
		score, err := getSimilarityScore(utils.StripTrailingNewlines(l), countsOfNumsInRight)
		if err != nil {
			panic(err)
		}
		totalScore += score
	}
	totalScoreString := strconv.Itoa(totalScore)
	fmt.Printf("The overall sum of similarity scores is: %s\n", totalScoreString)
}

func getSimilarityScore(str string, counts map[string]int) (int, error) {
	count, ok := counts[str]
	if !ok {
		return 0, nil
	}
	value, err := strconv.Atoi(str)
	if err != nil {
		return 0, err
	}
	return count * value, nil
}
