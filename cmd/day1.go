/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"bufio"
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

// day1Cmd represents the day1 command
var day1Cmd = &cobra.Command{
	Use:   "day1",
	Short: "Advent of code 2021 Day 1 Challenges",
	Run: func(cmd *cobra.Command, args []string) {
		day1()
	},
}

func init() {
	rootCmd.AddCommand(day1Cmd)
}

func day1() {
	defer input.Close()

	scanner := bufio.NewScanner(input)
	scanner.Split(bufio.ScanLines)

	numIncreases := 0

	windows := make([][]int, 0)
	headWindowIndex := 0
	numWindowIncreases := 0

	previousMeasurement := -1
	for hasNext, nextMeasurement := readNextMeasurement(scanner); hasNext; hasNext, nextMeasurement = readNextMeasurement(scanner) {
		// Part 1
		if nextMeasurement > previousMeasurement && previousMeasurement != -1 {
			numIncreases++
		}

		// Part 2
		middleWindowIndex := headWindowIndex - 1
		tailWindowIndex := headWindowIndex - 2
		windows = append(windows, make([]int, 3))
		windows[headWindowIndex][0] = nextMeasurement

		if middleWindowIndex >= 0 {
			windows[middleWindowIndex][1] = nextMeasurement
		}
		if tailWindowIndex >= 0 {
			windows[tailWindowIndex][2] = nextMeasurement
		}
		if tailWindowIndex > 0 && sumWindow(windows[tailWindowIndex]) > sumWindow(windows[tailWindowIndex-1]) {
			numWindowIncreases++
		}
		headWindowIndex++

		previousMeasurement = nextMeasurement
	}

	fmt.Println("--- Part 1 ---")
	fmt.Printf("Measurement increased %d times\n", numIncreases)

	fmt.Println()
	fmt.Println("--- Part 2 ---")
	fmt.Printf("Windows increased %d times\n", numWindowIncreases)

}

func readNextMeasurement(scanner *bufio.Scanner) (bool, int) {
	result := scanner.Scan()

	measurement, _ := strconv.Atoi(scanner.Text())
	return result, measurement
}

func sumWindow(window []int) int {
	sum := 0
	for _, measurement := range window {
		sum += measurement
	}

	return sum
}
