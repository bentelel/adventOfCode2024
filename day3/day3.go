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
//  split input by "mul(", p.e. input xmul(1,2)_mul(§3,3)__mul( 2, 3)__7
//  [x 1,2)_ $3,3)__  2,3)__7]
//  check each entry if it contains ")", if not: discard
//  then split each element by ) and take the first element
//  [1,2 $3,3  2,3]
//  check input if it contains a ",", if not, discard
//  [1,2 $3,3  2,3]
//  split input by ","
//  [[1 2] [$3 3] [ 2 3]]
//  check each subslice for len()==2 if it has more, discard
//  check each element in each slice if it is int with at max 3 digits and at least 1 digit.

func A() {
	lowerBound := 1
	upperBound := 999
	input, err := utils.ReadFileToString("input.txt", "day3")
	if err != nil {
		panic(err)
	}
	var muls []string = strings.Split(input, "mul(")
	muls_temp := []string{}
	for _, m := range muls {
		if strings.Contains(m, ")") {
			muls_temp = append(muls_temp, strings.Split(m, ")")[0])
		}
	}
	muls = muls_temp
	var factors [][]string
	for _, m := range muls {
		if strings.Contains(m, ",") {
			f := strings.Split(m, ",")
			factors = append(factors, f)
		}
	}
	fmt.Printf("%v\n", factors)
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
			panic(err)
		}
		ok, right, err := utils.RepresentsIntegerWithinBounds(fs[1], lowerBound, upperBound)
		if !ok {
			continue
		}
		if err != nil {
			panic(err)
		}
		result := left * right
		overallResult += result
	}
	fmt.Printf("The overall result is: %d\n", overallResult)
}
