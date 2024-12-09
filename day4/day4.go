package day4

import (
	"fmt"
	"strings"

	"bentelel/adventOfCode2024/utils"
)

// find number of instances of word XMAS in word search
// can be found in any orientation (backwards, forwards, horizontal, vertical, diagonal)
// idea:
//
//	get function to find valid neighbours within input
//	for "field"-indices, these are all 8 neighbours
//	abc
//	def
//	hij
//
// looking from e:
// 1    i-1           > d
// 2    i+1           > f
// 3    i-len(row)-1  > a
// 4    i-len(row)    > b
// 5    i-len(row)+1  > c
// 6    i+len(row)-1  > h
// 7    i+len(row)    > i
// 8    i+len(row)+1  > j
// for indexes on column 0:
//
//	only 2, 4, 5, 7 and 8 are valid
//
// for indexes on column len(row)-1
//
//	only 1, 3, 4, 6 and 7 are valid
//
// for indixes on row 0:
//
//	only 1, 2, 6, 7, 8 are valid
//
// for indexes on row len(columns)-1
//
//	only 1, 2, 3, 4, 5 are valid
//
// possible plan:
//
//	create map from letters to indices of the letters
//	loop over all indices for X
//	 for each, get the list of indicies for Ms and compare that list against all possible neightbours
//	  we do not care for the edge-cases because indicies like -1, or over the lenght of the input etc will never be in the list anyhow
const (
	SEPARATOR string = "\n"
)

func getXMASletters() []string {
	return []string{"X", "M", "A", "S"}
}

func A() {
	input, err := utils.ReadFileToString("input.txt", "day4")
	if err != nil {
		panic(err)
	}
	// get the row length before we remove the row separator
	rowLength := len(utils.StripTrailingNewlines(strings.Split(input, SEPARATOR)[0]))
	// remove all newline characters
	input = strings.Replace(input, SEPARATOR, "", -1)
	input = utils.StripTrailingNewlines(input)
	lettersCount := make(map[string]int)
	for _, letter := range getXMASletters() {
		count := strings.Count(input, letter)
		lettersCount[letter] = count
	}
	//for key, val := range lettersCount {
	//	fmt.Printf("key: %s, val: %d\n", key, val)
	//}
	// now we need to start the neighbour-lookup starting with letter X
	// for each X-index:
	//   find the intersection of index-neighbours with the next letters indices
	//   for all those indices, start the process anew
	// if last letter is reache and we have intersects, we have a hit
	var letterIndices map[string][]int = make(map[string][]int, len(getXMASletters()))
	for _, l := range getXMASletters() {
		letterIndices[l] = utils.GetAllIndices(l, input)
		fmt.Printf("letter: %s, indices: %v\n", l, letterIndices[l])
	}
	// this holds slices like [1 2 3 5 8]
	// this DFS business screams to be made a goroutine. but i can not be arsed to understand this right now.
	indices := letterIndices[getXMASletters()[0]]
	var results []bool
	for _, ind := range indices {
		res := DFS(1, ind, rowLength, letterIndices)
		fmt.Printf("for start index %d the results are: %v\n", ind, res)
		results = append(results, res...)
	}
	total := 0
	for _, r := range results {
		if r {
			total += 1
		}
	}
	// my result is to high because I count also "broken" shapes. it needs to be strict lines!
	// this means I need to completely overhaul my solution to only check straight lines..
	fmt.Printf("Count of XMAS: %d", total)
}

func B() {
}

func DFS(letterLevel int, startIndex int, rowLength int, m map[string][]int) []bool {
	//if letterLevel > len(getXMASletters()) {
	//	return true
	//}
	letterToSearch := getXMASletters()[letterLevel]
	neighbours := neighbouringIndices(startIndex, rowLength)
	intersection := utils.Intersect(neighbours, m[letterToSearch])
	fmt.Printf("letterlevel: %d, startIndex: %d, rowLength: %d, letter: %s, intersect: %v\n", letterLevel, startIndex, rowLength, letterToSearch, intersection)
	if len(intersection) == 0 {
		return []bool{false}
	}
	if letterLevel == len(getXMASletters())-1 {
		var results []bool
		for range intersection {
			results = append(results, true)
		}
		return results
	}
	var results []bool
	for _, ind := range intersection {
		results = append(results, DFS(letterLevel+1, ind, rowLength, m)...)
	}
	return results
}

// reihe 24, einmal angenommen

func neighbouringIndices(index int, rowLength int) []int {
	// we need to "kill" some indices which are not valid.
	// these are the ones, where index-1 or index+1 would carry over to the last or next line. those are actually not true neightbours.
	// if index%rowlength == 0, then we are at the start of the line and nuke the left index
	// if index%rowLength == rowLength-1, then we are at the end of a line and nuke the right index
	// we also need to nuke some of the diagonal indices
	impossibleIndex := -999
	var topLeft int
	var left int
	var bottomLeft int
	if index%rowLength == 0 {
		topLeft = impossibleIndex
		left = impossibleIndex
		bottomLeft = impossibleIndex
	} else {
		topLeft = index - rowLength - 1
		left = index - 1
		bottomLeft = index + rowLength - 1
	}
	top := index - rowLength
	var right int
	var topRight int
	var bottomRight int
	if index%rowLength == rowLength-1 {
		right = impossibleIndex
		topRight = impossibleIndex
		bottomRight = impossibleIndex
	} else {
		right = index + 1
		topRight = index - rowLength + 1
		bottomRight = index + rowLength + 1
	}
	bottom := index + rowLength
	return []int{topLeft, top, topRight, left, right, bottomLeft, bottom, bottomRight}
}

func raysFromIndex(startIndex int, rowLength int, rayLength int) [][]int {
	// casts ray from start index in all 8 directions, returns slice of new indices ordered from nearest to startIndex to farest
	var tl, t, tr, r, br, b, bl, l []int
	for i := 1; i <= rayLength; i++ {
		if startIndex%rowLength != 0 {
			// left edge of the input
			// top left
			tl = append(tl, startIndex-i*(rowLength-1))
			// bottom left
			bl = append(bl, startIndex+i*(rowLength-1))
			// left
			l = append(l, startIndex-i*1)
		}
		if startIndex%rowLength != rowLength-1 {
			// right edge of the input
			// top right
			tr = append(tr, startIndex-i*(rowLength+1))
			// right
			r = append(r, startIndex+i*1)
			// bottom right
			br = append(br, startIndex+i*(rowLength+1))
		}
		// top
		t = append(t, startIndex-i*rowLength)
		// bottom
		b = append(b, startIndex+i*rowLength)
	}

	var rays [][]int = [][]int{tl, t, tr, r, br, b, bl, l}
	return rays
}
