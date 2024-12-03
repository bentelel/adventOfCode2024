package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"bentelel/adventOfCode2024/day1"
	"bentelel/adventOfCode2024/day2"
)

var registry = map[string]func(){"1A": day1.A, "1B": day1.B, "2A": day2.A}

func main() {
	// main entrypoint for advent of code 2024 puzzles
	// pass the day and A or B in to run the appropriate function
	// funcs will be named Day1A, Day1B in their respective packages

	if len(os.Args) < 3 {
		fmt.Printf("Program expects 2 inputs (day and exercise (A or B)).\nA clean call looks like this \"adventOfCode2024 1 A\".\n%s inputs where provided.\n", strconv.Itoa(len(os.Args)-1))
		return
	}

	// get args
	day := os.Args[1]
	exercise := strings.ToUpper(os.Args[2])
	// get func
	fun := registry[day+exercise]
	if fun == nil {
		fmt.Printf("Entry day %s, exercise %s not found. Please retry.\n", day, exercise)
		return
	}
	// the funcs will not return anything and just print their respective answers themselves.
	fun()
	return
}
