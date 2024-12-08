package day3

import (
	"fmt"
	"strings"

	"bentelel/adventOfCode2024/utils"
)

// It seems like the goal of the program is just to multiply some numbers.
// It does that with instructions like mul(X,Y), where X and Y are each 1-3 digit numbers.
// For instance, mul(44,46) multiplies 44 by 46 to get a result of 2024.
// Similarly, mul(123,4) would multiply 123 by 4.
// However, because the program's memory has been corrupted,
// there are also many invalid characters that should be ignored,
// even if they look like part of a mul instruction.
// Sequences like mul(4*, mul(6,9!, ?(12,34), or mul ( 2 , 4 ) do nothing.

// nested mul() is not a thing! mul( mul(15,3),mul(1,1)) does not multiply the 2 results of the inner mul() only the inner mul are evaluated
// whitespace within the mul also makes the mul invalid. p.e. mul( 1, 2) is invalid and to be ignored!
// inputs are only 1 to 3 digits long!

// naive plan:
//
//	split input by "mul(", p.e. input xmul(1,2)_mul(§3,3)__mul( 2, 3)__7
//	[x 1,2)_ $3,3)__  2,3)__7]
//	check each entry if it contains ")", if not: discard
//	then split each element by ) and take the first element
//	[1,2 $3,3  2,3]
//	check input if it contains a ",", if not, discard
//	[1,2 $3,3  2,3]
//	split input by ","
//	[[1 2] [$3 3] [ 2 3]]
//	check each subslice for len()==2 if it has more, discard
//	check each element in each slice if it is int with at max 3 digits and at least 1 digit.
const (
	LOWERBOUND    int    = 1
	UPPERBOUND    int    = 999
	MUL           string = "mul("
	RIGHT_BRACKET string = ")"
	COMMA         string = ","
	DONT          string = "don't()"
	DO            string = "do()"
)

func A() {
	input, err := utils.ReadFileToString("input.txt", "day3")
	if err != nil {
		panic(err)
	}
	overallResult, err := getSumOfMuls(input, LOWERBOUND, UPPERBOUND)
	if err != nil {
		panic(err)
	}
	fmt.Printf("The overall result is: %d\n", overallResult)
}

func B() {
	// new: there are do() and don't() occurances within the input.
	// do() enables mul() while don't() diables them
	// all mul() after a don't() are not counted until the next do() happens
	// muls are enabled at the start
	// plan:
	//  keep the logic from a, but split the input by do() and don't() before doing As logic.
	//  first split by don't()
	//    input: mul(1,2)__mul(2,3)__don't()mul(5,6)__mul(2,4)_do()_mul(1,2)
	//    output [mul(1,2)__mul(2,3)__ mul(5,6)__mul(2,4)_do()_mul(1,2)]
	//    get out the 0th element > those are enabled by default!
	//    rest: split by do()
	//    keep right side, discard left side
	//    assumption: do() and don't() always alternate!
	//    input [mul(1,2)__mul(2,3)__ mul(5,6)__mul(2,4)_do()_mul(1,2)]
	//    output 1: mul(1,2)__mul(2,3)__
	//    output 2: [mul(5,6)__mul(2,4)_ _mul(1,2)] > discard left > _mul(1,2)
	//  then run logic from A for all remaining input parts
	input, err := utils.ReadFileToString("input.txt", "day3")
	if err != nil {
		panic(err)
	}
	var input_split_by_dont []string = strings.Split(input, DONT)
	if len(input_split_by_dont) == 0 {
		panic("Input is empty.. what")
	}
	firstBlock := input_split_by_dont[0]
	var input_only_dos []string
	input_only_dos = append(input_only_dos, firstBlock)
	for _, inp := range input_split_by_dont[1:] {
		input_split_by_do := strings.Split(inp, DO)
		if len(input_split_by_do) > 2 {
			panic("assumption that do and dont alternate seems to be wrong")
		}
		// discard first part and only use second part
		input_only_dos = append(input_only_dos, input_split_by_do[1])
	}
	var overallSum int = 0
	for _, inp := range input_only_dos {
		result, err := getSumOfMuls(inp, LOWERBOUND, UPPERBOUND)
		if err != nil {
			panic(err)
		}
		overallSum += result
	}

	fmt.Printf("The overall result is: %d\n", overallSum)
}

func getSumOfMuls(input string, lowerBound int, upperBound int) (int, error) {
	var muls []string = strings.Split(input, MUL)
	muls_temp := []string{}
	for _, m := range muls {
		if strings.Contains(m, RIGHT_BRACKET) {
			muls_temp = append(muls_temp, strings.Split(m, RIGHT_BRACKET)[0])
		}
	}
	muls = muls_temp
	var factors [][]string
	for _, m := range muls {
		if strings.Contains(m, COMMA) {
			f := strings.Split(m, COMMA)
			factors = append(factors, f)
		}
	}
	// fmt.Printf("%v\n", factors)
	var overallResult int = 0
	for _, fs := range factors {
		if len(fs) > 2 {
			continue
		}
		ok, left, err := utils.RepresentsIntegerWithinBounds(fs[0], lowerBound, upperBound)
		if !ok {
			continue
		}
		if err != nil {
			return 0, err
		}
		ok, right, err := utils.RepresentsIntegerWithinBounds(fs[1], lowerBound, upperBound)
		if !ok {
			continue
		}
		if err != nil {
			return 0, err
		}
		result := left * right
		overallResult += result
	}
	return overallResult, nil
}
