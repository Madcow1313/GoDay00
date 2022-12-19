package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

type sFlags struct {
	mean, median, mode, sd bool
}

func calculateStandardDeviation(sequence []int, mean float64) float64 {
	if len(sequence) < 2 {
		return 0
	}
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
		median = float64(sequence[length/2-1]+sequence[length/2]) / 2.0
	} else {
		median = float64(sequence[length/2])
	}
	return median
}

func handleSequence(sequence []int, flags *sFlags) {
	sort.Slice(sequence, func(i, j int) bool {
		return sequence[i] < sequence[j]
	})
	if flags.mean {
		fmt.Printf("Mean: %.2f\n", calculateMean(sequence))
	}
	if flags.median {
		fmt.Printf("Median: %.2f\n", calculateMedian(sequence))
	}
	if flags.mode {
		fmt.Printf("Mode: %d\n", calculateMode(sequence))
	}
	if flags.sd {
		fmt.Printf("SD: %.2f\n", calculateStandardDeviation(sequence, calculateMean(sequence)))
	}

}

func checkAllFlags(flags *sFlags) {
	if !flags.mean && !flags.median && !flags.mode && !flags.sd {
		flags.mean = true
		flags.median = true
		flags.mode = true
		flags.sd = true
	}
}

func main() {
	flags := new(sFlags)
	flag.BoolVar(&flags.mean, "mean", false, "show mean")
	flag.BoolVar(&flags.median, "median", false, "show median")
	flag.BoolVar(&flags.mode, "mode", false, "show mode")
	flag.BoolVar(&flags.sd, "sd", false, "show standard deviation")
	flag.Parse()
	checkAllFlags(flags)
	fmt.Println("Enter sequence of integers")
	sequence := []int{}
	for {
		var _err error
		input, err := bufio.NewReader(os.Stdin).ReadString('\n')
		if len(input) == 0 {
			break
		}
		if err != nil {
			fmt.Println(err.Error())
			continue
		}
		splittedInput := strings.Fields(input)
		for _, integer := range splittedInput {
			number, err := strconv.Atoi(integer)
			_err = err
			if number <= 100000 && number >= -100000 && _err == nil {
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
	}
	handleSequence(sequence, flags)
}
