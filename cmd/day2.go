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
	"strings"

	"github.com/spf13/cobra"
)

type submarineCommand string

const (
	downCommand    submarineCommand = "down"
	upCommand      submarineCommand = "up"
	forwardCommand submarineCommand = "forward"
)

var (
	// day2Cmd represents the day2 command
	day2Cmd = &cobra.Command{
		Use:   "day2",
		Short: "Advent of code 2021 Day 1 Challenges",
		Run: func(cmd *cobra.Command, args []string) {
			day2()
		},
	}
)

func init() {
	rootCmd.AddCommand(day2Cmd)
}

func day2() {
	scanner := bufio.NewScanner(input)
	scanner.Split(bufio.ScanLines)

	part1Depth := 0
	part1Distance := 0

	part2Aim := 0
	part2Depth := 0
	part2Distance := 0

	for scanner.Scan() {
		line := scanner.Text()

		command, value := readInstruction(line)

		switch command {
		case downCommand:
			part1Depth += value
			part2Aim += value
		case upCommand:
			part1Depth -= value
			part2Aim -= value
		case forwardCommand:
			part1Distance += value
			part2Distance += value
			part2Depth += value * part2Aim
		}
	}

	fmt.Println("=== Part 1 ===")
	fmt.Printf("Depth: %d, Distance: %d, Answer: %d\n", part1Depth, part1Distance, part1Depth*part1Distance)

	fmt.Println()

	fmt.Println("=== Part 2 ===")
	fmt.Printf("Depth: %d, Distance: %d, Answer: %d\n", part2Depth, part2Distance, part2Depth*part2Distance)
}

func readInstruction(line string) (submarineCommand, int) {
	tokens := strings.Split(line, " ")
	value, _ := strconv.Atoi(tokens[1])

	return submarineCommand(tokens[0]), value
}
