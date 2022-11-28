package main

import (
	"bufio"
	"errors"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func calculateStandardDeviation(sequence []int, mean float64) float64 {
	var quadraticSum float64
	for _, number := range sequence {
		quadraticSum += math.Pow(float64(number)-mean, 2)
	}
	return math.Sqrt(quadraticSum / float64(len(sequence)-1))
}

func calculateMode(sequence []int) int {
	var max, modeIdx int
	modeMap := make(map[int]int, len(sequence))
	for _, number := range sequence {
		modeMap[number] += 1
	}
	max = modeMap[0]
	for idx, number := range modeMap {
		if number > max {
			max = number
			modeIdx = idx
		}
	}
	return modeIdx
}

func calculateMean(sequence []int) float64 {
	var sum int
	var mean float64

	for _, number := range sequence {
		sum += number
	}
	mean = float64(sum) / float64(len(sequence))
	return mean
}

func calculateMedian(sequence []int) float64 {
	var median float64
	length := len(sequence)
	if length%2 == 0 {
		median = float64((sequence[length/2-1] + sequence[length/2]) / 2.0)
	} else {
		median = float64(sequence[length/2])
	}
	return median
}

func handleSequence(sequence []int) {
	sort.Slice(sequence, func(i, j int) bool {
		return sequence[i] < sequence[j]
	})
	fmt.Printf("Mean: %.1f\n", calculateMean(sequence))
	fmt.Printf("Median: %.1f\n", calculateMedian(sequence))
	fmt.Printf("Mode: %d\n", calculateMode(sequence))
	fmt.Printf("SD: %.2f\n", calculateStandardDeviation(sequence, calculateMean(sequence)))
}

func main() {
	for {
		var _err error
		sequence := []int{}
		fmt.Println("Enter sequence of integers")
		input, err := bufio.NewReader(os.Stdin).ReadString('\n')
		if err != nil {
			fmt.Println(err.Error())
			continue
		}
		splittedInput := strings.Fields(input)
		for _, integer := range splittedInput {
			number, err := strconv.Atoi(integer)
			_err = err
			if number <= 10000 && number >= -10000 && _err == nil {
				sequence = append(sequence, number)
			} else {
				if err == nil {
					_err = errors.New("wrong input")
				} else {
					_err = err
				}
				fmt.Println(_err.Error())
				break
			}
		}
		if _err != nil {
			continue
		}
		handleSequence(sequence)
	}
}
