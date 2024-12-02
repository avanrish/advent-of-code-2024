package day2

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"path"
	"strconv"
	"strings"
)

const (
	fileName = "02/input.txt"
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
	defer file.Close()

	scanner := bufio.NewScanner(file)
	validCount := 0
	for scanner.Scan() {
		line := scanner.Text()
		nums := strings.Split(line, " ")
		isIncreasing := false
		validReport := true
		for i := 0; i < len(nums)-1; i++ {
			a, err := strconv.Atoi(nums[i])
			if err != nil {
				break
			}
			b, err := strconv.Atoi(nums[i+1])
			if err != nil {
				break
			}
			diff := a - b
			diffAbs := math.Abs(float64(diff))
			if diffAbs < 1.0 || diffAbs > 3.0 {
				validReport = false
				break
			}
			if i == 0 {
				isIncreasing = diff < 0
			} else if isIncreasing && diff > 0 {
				validReport = false
				break
			} else if !isIncreasing && diff < 0 {
				validReport = false
				break
			}
		}
		if validReport {
			validCount++
		}
	}
	fmt.Printf("Valid reports: %d\n", validCount)
}
