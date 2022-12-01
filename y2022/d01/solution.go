package d01

import (
	"fmt"
	"io"
	"sort"
	"strconv"

	"github.com/busser/adventofcode/helpers"
	"golang.org/x/exp/constraints"
)

// PartOne solves the first problem of day 1 of Advent of Code 2022.
func PartOne(r io.Reader, w io.Writer) error {
	// TODO: Read the input. For example:
	lines, err := helpers.LinesFromReader(r)
	if err != nil {
		return fmt.Errorf("could not read input: %w", err)
	}


  top, _, err := getMaximumTotalCalories(lines)
  if err != nil {
	  return fmt.Errorf("problem in line iterator");
  }

	// TODO: Write the answer. For example:
	_, err = fmt.Fprintf(w, "%d", top)
	if err != nil {
		return fmt.Errorf("could not write answer: %w", err)
	}

	return nil
} 

// PartTwo solves the second problem of day 1 of Advent of Code 2022.
func PartTwo(r io.Reader, w io.Writer) error {
	// TODO: Read the input. For example:
	lines, err := helpers.LinesFromReader(r)
	if err != nil {
		return fmt.Errorf("could not read input: %w", err)
	}

  _, topThree ,err := getMaximumTotalCalories(lines)
  if err != nil {
	  return fmt.Errorf("problem in line iterator");
  }
	// TODO: Write the answer. For example:
	_, err = fmt.Fprintf(w, "%d", topThree)
	if err != nil {
		return fmt.Errorf("could not write answer: %w", err)
	}

	return nil
}

func getMaximumTotalCalories(lines []string) (int, int, error) {

  var currentMax int = 0
  var localMax int = 0
  allSums := []int{}

  for _, value := range lines {

    iValue, err := strconv.Atoi(value)

    if err != nil {
      currentMax = max(currentMax, localMax)
	  allSums = append(allSums, localMax);
      localMax = 0
    }

    localMax += iValue
  }

	sort.Slice(allSums,  func(a,b int) bool { return allSums[a] > allSums[b]})

	topThree := 0;

	for _, value := range allSums[0:3] {
		helpers.Println(value)
		topThree += value
	}

	helpers.Println(topThree)
  

  return currentMax, topThree, nil
}


func max[T constraints.Ordered](a, b T) T {
	if a > b {
		return a
	}
	return b
}

