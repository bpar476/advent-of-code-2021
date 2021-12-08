/*
Copyright © 2021 Ben Partridge

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

	"github.com/spf13/cobra"
)

type bitCount struct {
	Zeros int
	Ones  int
}

type bitCriteria func(count bitCount, bit, bitstring int) bool

// day3Cmd represents the day3 command
var day3Cmd = &cobra.Command{
	Use:   "day3",
	Short: "Advent of code 2021 Day 3 Challenges",
	Run: func(cmd *cobra.Command, args []string) {
		day3()
	},
}

func init() {
	rootCmd.AddCommand(day3Cmd)
}

func day3() {
	scanner := bufio.NewScanner(input)
	scanner.Split(bufio.ScanLines)

	// Read in bit strings as integers of binary value
	bitstrings := []int{}
	length := 0
	for scanner.Scan() {
		length = len(scanner.Text())
		bitstring := 0
		bitStringText := scanner.Text()
		for i := range bitStringText {
			char := bitStringText[length-1-i]
			var bit int = 0
			if char == '1' {
				bit = 1
			}
			bitstring |= (bit << i)
		}
		bitstrings = append(bitstrings, bitstring)
	}

	// Count positions in entire set
	counts := getCountsForBitstrings(bitstrings, length)

	// Compute gamma and epsilon
	gamma := 0
	epsilon := 0

	for i, count := range counts {
		if count.Zeros > count.Ones {
			epsilon |= 1 << (length - i - 1)
		} else {
			gamma |= 1 << (length - i - 1)
		}
	}

	fmt.Println("=== Part 1 ===")
	fmt.Printf("Gamma rate is %b, Epsilon rate is %b. Answer is: %d\n", gamma, epsilon, gamma*epsilon)
	fmt.Println()

	// Compute oxygen rating
	oxygenRate := reduceBitstringsUsingCriteria(bitstrings, length, counts, oxygenCriteria)

	carbonRate := reduceBitstringsUsingCriteria(bitstrings, length, counts, co2Rating)

	fmt.Println("=== Part 2 ===")
	fmt.Printf("Oxygen rate is: %d\n", oxygenRate)
	fmt.Printf("Carbon rate is: %d\n", carbonRate)
	fmt.Printf("Answer is %d\n", oxygenRate*carbonRate)
}

func getCountsForBitstrings(bitstrings []int, length int) []bitCount {
	counts := make([]bitCount, length)
	for i := 0; i < length; i++ {
		count := bitCount{}
		for _, bitstring := range bitstrings {
			if hasOneInPosition(length-i-1, bitstring) {
				count.Ones++
			} else {
				count.Zeros++
			}
		}
		counts[i] = count
	}

	return counts
}

func hasOneInPosition(position, bitstring int) bool {
	mask := 1 << position
	return mask&bitstring == mask
}

func reduceBitstringsUsingCriteria(bitstrings []int, binaryLength int, initialCounts []bitCount, selectionCriteria bitCriteria) int {
	candidates := bitstrings
	counts := initialCounts

	for i := 0; i < binaryLength && len(candidates) > 1; i++ {
		newCandidates := make([]int, 0, len(candidates))
		for _, candidate := range candidates {
			if selectionCriteria(counts[i], binaryLength-i-1, candidate) {
				newCandidates = append(newCandidates, candidate)
			}
		}
		candidates = newCandidates
		counts = getCountsForBitstrings(candidates, binaryLength)
	}

	return candidates[0]
}

func oxygenCriteria(count bitCount, bit int, bitstring int) bool {
	bitstringHasOneInPosition := hasOneInPosition(bit, bitstring)
	return count.Ones >= count.Zeros == bitstringHasOneInPosition
}

func co2Rating(count bitCount, bit, bitstring int) bool {
	bitstringHasOneInPosition := hasOneInPosition(bit, bitstring)
	return count.Ones < count.Zeros == bitstringHasOneInPosition
}
