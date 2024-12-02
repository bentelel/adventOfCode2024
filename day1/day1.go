package day1

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

const inputDirectory string = "day1"

func buildPathToFile(filename string) (string, error) {
	// get path of currently running program
	exe, err := os.Executable()
	if err != nil {
		return "", err
	}
	exePath := filepath.Dir(exe)
	targetPath := filepath.Join(exePath, inputDirectory, filename)
	return targetPath, nil
}

func readFileToString(filename string) (string, error) {
	path, err := buildPathToFile(filename)
	if err != nil {
		return "", err
	}
	bytes, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

func popFromSlice(s []string) []string {
	if len(s) == 0 {
		return s
	}
	return s[:len(s)-1]
}

func getListsFromInput(input string) ([]string, []string) {
	rows := strings.Split(input, "\n")
	// pop off last row if nil
	if rows[len(rows)-1] == "" {
		rows = popFromSlice(rows)
	}
	var leftList, rightList []string
	for _, r := range rows {
		values := strings.Split(r, "   ")
		leftList = append(leftList, values[0])
		rightList = append(rightList, values[1])
	}

	return leftList, rightList
}

func sortSliceAscending(s []string) ([]string, error) {
	// lets use bubble sort
	// maximum number of passes we need to make is len(s)-1
	for j := 0; j < len(s)-1; j++ {
		for i := 0; i < len(s)-1; i++ {
			first, err := strconv.Atoi(stripTrailingNewlines(s[i]))
			if err != nil {
				panic(err)
			}
			second, err := strconv.Atoi(stripTrailingNewlines(s[i+1]))
			if err != nil {
				panic(err)
			}
			if second < first {
				s[i+1], s[i] = s[i], s[i+1]
			}
		}
		//for _, v := range s {
		//	fmt.Printf("current sorting: %s\n", v)
		//}
		if isSliceSortedAsc(s) {
			fmt.Print("slice is sorted\n")

			fmt.Printf("first element: %s\n", s[0])
			fmt.Printf("last element: %s\n", s[len(s)-1])
			return s, nil
		}
		// fmt.Print("\n")
	}
	return []string{}, errors.New("Slice Not Sortable.\n")
}

func isSliceSortedAsc(s []string) bool {
	for i := 0; i < len(s)-1; i++ {
		first, err := strconv.Atoi(stripTrailingNewlines(s[i]))
		if err != nil {
			panic(err)
		}
		second, err := strconv.Atoi(stripTrailingNewlines(s[i+1]))
		if err != nil {
			panic(err)
		}
		if second < first {
			return false
		}
	}
	return true
}

func stripTrailingNewlines(s string) string {
	s = strings.Replace(s, "\r", "", -1)
	s = strings.Replace(s, "\n", "", -1)
	s = strings.Replace(s, "\n\r", "", -1)
	return s
}

func findDifferencesBetweenSlices(l []string, r []string) ([]int, error) {
	if len(l) != len(r) {
		return []int{}, errors.New("Slices must have the same length.\n")
	}

	diff := make([]int, len(l))
	for i := 0; i < len(l); i++ {
		lInt, err := strconv.Atoi(stripTrailingNewlines(l[i]))
		if err != nil {
			return []int{}, err
		}
		rInt, err := strconv.Atoi(stripTrailingNewlines(r[i]))
		if err != nil {
			return []int{}, err
		}
		diff[i] = rInt - lInt
	}
	return diff, nil
}

func A() {
	fmt.Print("running day 1A.\n")
	puzzleInput, err := readFileToString("inputA.txt")
	if err != nil {
		panic(err)
	}
	// fmt.Printf("puzzle input:\n%s", puzzleInput)
	leftInput, rightInput := getListsFromInput(puzzleInput)
	//for _, v := range leftInput {
	//	fmt.Printf("left input: %s\n", v)
	//}
	//for _, v := range rightInput {
	//	fmt.Printf("right input: %s\n", v)
	//}
	leftInputSorted, err := sortSliceAscending(leftInput)
	rightInputSorted, err := sortSliceAscending(rightInput)
	diffs, err := findDifferencesBetweenSlices(leftInputSorted, rightInputSorted)
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
