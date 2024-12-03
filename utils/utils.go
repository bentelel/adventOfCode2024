package utils

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func buildPathToFile(filename string, inputDirectory string) (string, error) {
	// get path of currently running program
	exe, err := os.Executable()
	if err != nil {
		return "", err
	}
	exePath := filepath.Dir(exe)
	targetPath := filepath.Join(exePath, inputDirectory, filename)
	return targetPath, nil
}

func ReadFileToString(filename string, inputDirectory string) (string, error) {
	path, err := buildPathToFile(filename, inputDirectory)
	if err != nil {
		return "", err
	}
	bytes, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

func PopFromSlice(s []string) []string {
	// ideally this should take in pointers and not the slice itself and then change the underlying slice and return the popped element.
	if len(s) == 0 {
		return s
	}
	return s[:len(s)-1]
}

func PopFromSliceOfSlices(s [][]string) [][]string {
	// ideally this should take in pointers and not the slice itself and then change the underlying slice and return the popped element.
	if len(s) == 0 {
		return s
	}
	return s[:len(s)-1][:]
}

func GetListsFromInput(input string) ([]string, []string) {
	rows := strings.Split(input, "\n")
	// pop off last row if nil
	if rows[len(rows)-1] == "" {
		rows = PopFromSlice(rows)
	}
	var leftList, rightList []string
	for _, r := range rows {
		values := strings.Split(r, "   ")
		leftList = append(leftList, values[0])
		rightList = append(rightList, values[1])
	}

	return leftList, rightList
}

func SortSliceAscending(s []string) ([]string, error) {
	// lets use bubble sort
	// maximum number of passes we need to make is len(s)-1
	for j := 0; j < len(s)-1; j++ {
		for i := 0; i < len(s)-1; i++ {
			first, err := strconv.Atoi(StripTrailingNewlines(s[i]))
			if err != nil {
				panic(err)
			}
			second, err := strconv.Atoi(StripTrailingNewlines(s[i+1]))
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
		if IsSliceSortedAsc(s) {
			fmt.Print("slice is sorted\n")
			return s, nil
		}
		// fmt.Print("\n")
	}
	return []string{}, errors.New("Slice Not Sortable.\n")
}

func IsSliceSortedAsc(s []string) bool {
	for i := 0; i < len(s)-1; i++ {
		first, err := strconv.Atoi(StripTrailingNewlines(s[i]))
		if err != nil {
			panic(err)
		}
		second, err := strconv.Atoi(StripTrailingNewlines(s[i+1]))
		if err != nil {
			panic(err)
		}
		if second < first {
			return false
		}
	}
	return true
}

func StripTrailingNewlines(s string) string {
	s = strings.Replace(s, "\r", "", -1)
	s = strings.Replace(s, "\n", "", -1)
	s = strings.Replace(s, "\n\r", "", -1)
	return s
}

func FindDifferencesBetweenSlices(l []string, r []string) ([]int, error) {
	if len(l) != len(r) {
		return []int{}, errors.New("Slices must have the same length.\n")
	}

	diff := make([]int, len(l))
	for i := 0; i < len(l); i++ {
		lInt, err := strconv.Atoi(StripTrailingNewlines(l[i]))
		if err != nil {
			return []int{}, err
		}
		rInt, err := strconv.Atoi(StripTrailingNewlines(r[i]))
		if err != nil {
			return []int{}, err
		}
		dist := lInt - rInt
		if dist < 0 {
			dist *= -1
		}
		diff[i] = dist
	}
	return diff, nil
}

func GetCountOfNumbers(s []string) (map[string]int, error) {
	if len(s) == 0 {
		return map[string]int{}, nil
	}
	var result map[string]int = make(map[string]int)
	for _, ss := range s {
		// clean up ss before using it, there are carriage returns and shit in there which leads to nasty bugs..
		ss = StripTrailingNewlines(ss)
		count, ok := result[ss]
		if ok {
			result[ss] = count + 1
		} else {
			result[ss]++
		}

	}
	return result, nil
}

func GetSimilarityScore(str string, counts map[string]int) (int, error) {
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

func GetRowsFromInput(input string) []string {
	rows := strings.Split(input, "\n")
	// pop off last row if nil
	if rows[len(rows)-1] == "" {
		rows = PopFromSlice(rows)
	}
	return rows
}

func GetRowsAndElements(input []string) [][]string {
	var result [][]string = make([][]string, len(input))
	for i := 0; i < len(input); i++ {
		elems := strings.Split(input[i], " ")
		elems = Map(elems, StripTrailingNewlines)
		result[i] = elems

	}
	// pop last row if nil
	if len(result[len(result)-1]) == 0 {
		result = PopFromSliceOfSlices(result)
	}
	return result
}

func Map(vs []string, f func(string) string) []string {
	vsm := make([]string, len(vs))
	for i, v := range vs {
		vsm[i] = f(v)
	}
	return vsm
}

func sliceStartsAscending(s []string) bool {
	/// returns true if the supplied slice starts out ascending
	/// if multiple entries of the same numbers follow at the start, the func keeps checking until it hits the first non-same number.
	/// a slice of all the same numbers is considered ascending.
	for i := 0; i < len(s)-1; i++ {
		if s[i] == s[i+1] {
			continue
		} else if s[i] < s[i+1] {
			return true
		} else {
			return false
		}
	}
	return true
}

func IsSliceAllAscendingOrDescending(s []string) (bool, error) {
	startsAscending := sliceStartsAscending(s)

	for i := 0; i < len(s)-1; i++ {
		left, err := strconv.Atoi(s[i])
		if err != nil {
			return false, err
		}
		right, err := strconv.Atoi(s[i+1])
		if err != nil {
			return false, err
		}
		if startsAscending {
			if left <= right {
				continue
			} else {
				return false, nil
			}
		} else {
			if left >= right {
				continue
			} else {
				return false, nil
			}
		}
	}
	return true, nil
}

func AreDistancesOk(input []string, lowerBound int, upperBound int) (bool, error) {
	// returns true (and nil) if the distance between any 2 sequential elements of the input are within (including) the lower and upper bound
	for i := 0; i < len(input)-1; i++ {
		left, err := strconv.Atoi(input[i])
		if err != nil {
			return false, err
		}
		right, err := strconv.Atoi(input[i+1])
		if err != nil {
			return false, err
		}
		distance := left - right
		if distance < 0 {
			distance *= -1
		}
		// fmt.Printf("left: %d, right: %d, distance: %d\n", left, right, distance)
		if distance < lowerBound || distance > upperBound {
			return false, nil
		}
	}
	return true, nil
}
