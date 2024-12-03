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

func sliceStartsAscending(s []string) (bool, error) {
	/// returns true if the supplied slice starts out ascending
	/// if multiple entries of the same numbers follow at the start, the func keeps checking until it hits the first non-same number.
	/// a slice of all the same numbers is considered ascending.
	/// this might not be sufficient for day2B because day2B allows for a single error in the ordering
	/// so if we have a slice like [2 1 3 4 5], this func would return that it is ordered descending, while it actually is ordered ascending, but with 1 error.
	for i := 0; i < len(s)-1; i++ {
		left, err := strconv.Atoi(s[i])
		if err != nil {
			return false, err
		}
		right, err := strconv.Atoi(s[i+1])
		if err != nil {
			return false, err
		}

		if left == right {
			continue
		} else if left < right {
			return true, nil
		} else {
			return false, nil
		}
	}
	return true, nil
}

func IsSliceAllAscendingOrDescending(s []string) (bool, error) {
	startsAscending, err := sliceStartsAscending(s)
	if err != nil {
		return false, nil
	}
	for i := 0; i < len(s)-1; i++ {
		left, err := strconv.Atoi(s[i])
		if err != nil {
			return false, err
		}
		right, err := strconv.Atoi(s[i+1])
		if err != nil {
			return false, err
		}
		if left == right {
			return false, nil
		}
		if startsAscending {
			if left < right {
				continue
			} else {
				return false, nil
			}
		} else {
			if left > right {
				continue
			} else {
				return false, nil
			}
		}
	}
	return true, nil
}

func IsSliceAllAscendingOrDescendingWithDampening(s []string) (bool, error) {
	// does the same as the version without "WithDampening", but only returns false if our check condition is not met twice. This should be the same as if we ignore 1 wrong element.
	// the way in which we test the slice needs to be different, because we can not rely on sliceStartsAscending to get the slices ordering.
	// we check the slice for both ascending and descending order at the same time and keep track of how many elements are not in correct ordering for both orderings. if any more than 1 are wrong for each ordering kind, we return false.
	// iterate over the slice from the frint
	orderErrorCounterAscending := 0
	orderErrorCounterDescending := 0
	for i := 0; i < len(s)-1; i++ {
		left, err := strconv.Atoi(s[i])
		if err != nil {
			return false, err
		}
		right, err := strconv.Atoi(s[i+1])
		if err != nil {
			return false, err
		}
		// if left < right, then we have at least 1 pair of elements which are not ordered descending.
		if left == right {
			orderErrorCounterAscending += 1
			orderErrorCounterDescending += 1
		} else if left < right {
			orderErrorCounterDescending += 1
		} else if right > left {
			// if left > right, then we have at least 1 pair of elements which are not ordered ascending.
			orderErrorCounterAscending += 1
		}
		if orderErrorCounterAscending > 2 && orderErrorCounterDescending > 2 {
			return false, nil
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

func AreDistancesOkWithDampening(input []string, lowerBound int, upperBound int) (bool, error) {
	// returns true (and nil) if the distance between any 2 sequential elements of the input are within (including) the lower and upper bound
	// tolerates 1 of these errors and only return false if the second error is found.
	// we can not use the errorCounter approach here because:
	// slice like 9 7 6 2 1 > we get the first error when looking at pairing 6 2.
	// the dampening removes 1 level. that would be as if the slice is 9 7 2 1
	// which still would be wrong because of the distnace between 7 2.
	// our errorCounter approach would flag this as valid though!
	// hang on!
	errorCounter := 0
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
			// to catch errors like 9 7 6 2 1, we check one index further to the right against the left side as well. if that also fails the distance check, then return false
			sliceDropLeft := input[i+1:]
			sliceDropRight := append(input[:i+1], input[i+2:]...)
			leftOk, err := AreDistancesOk(sliceDropLeft, lowerBound, upperBound)
			if err != nil {
				return false, err
			}
			rightOk, err := AreDistancesOk(sliceDropRight, lowerBound, upperBound)
			if err != nil {
				return false, err
			}
			if !leftOk && !rightOk {
				return false, nil
			}
			errorCounter += 1
			if errorCounter > 1 {
				return false, nil
			}
		}
	}
	return true, nil
}
