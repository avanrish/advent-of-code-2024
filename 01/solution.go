package day1

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"path"
	"slices"
	"strconv"
	"strings"
)

const (
	fileName = "01/input.txt"
)

func Solution() {
	pwd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		return
	}

	filePath := path.Join(pwd, fileName)
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println(err)
		return
	}

	leftOccurences := map[int]int{}
	rightOccurences := map[int]int{}
	var leftList []int
	var rightList []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		values := strings.Split(line, "   ")
		if len(values) != 2 {
			continue
		}
		leftInt, err := strconv.Atoi(values[0])
		if err != nil {
			fmt.Println(err)
		}
		rightInt, err := strconv.Atoi(values[1])
		if err != nil {
			fmt.Println(err)
		}
		leftOccurences[leftInt] += 1
		rightOccurences[rightInt] += 1
		leftList = append(leftList, leftInt)
		rightList = append(rightList, rightInt)
	}
	slices.Sort(leftList)
	slices.Sort(rightList)

	totalDistance := 0.0
	for i := 0; i < len(leftList); i++ {
		totalDistance += math.Abs(float64(leftList[i] - rightList[i]))
	}

	similarityScore := 0
	for num, occurs := range leftOccurences {
		rightOccur, ok := rightOccurences[num]
		if ok {
			similarityScore += occurs * num * rightOccur
		}
	}

	fmt.Printf("Total sum: %d\n", int(totalDistance))
	fmt.Printf("Similarity score: %d\n", similarityScore)
}
