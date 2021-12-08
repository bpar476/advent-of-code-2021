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
	"log"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

// day4Cmd represents the day4 command
var day4Cmd = &cobra.Command{
	Use:   "day4",
	Short: "Advent of Code 2021 Day 4 (Bingo)",
	Run: func(cmd *cobra.Command, args []string) {
		day4()
	},
}

type BingoTile struct {
	value   int
	checked bool
}

type Board [][]BingoTile

type BoardReference struct {
	board int
	row   int
	col   int
}

func (b Board) String() string {
	var builder strings.Builder
	for i := range b {
		for j := range b[i] {
			builder.WriteString(fmt.Sprintf("%2d", b[i][j].value))
			if b[i][j].checked {
				builder.WriteString("✅ ")
			} else {
				builder.WriteString("❌ ")
			}
		}
		builder.WriteString("\n")
	}

	return builder.String()
}

func init() {
	rootCmd.AddCommand(day4Cmd)
}

func day4() {
	bingoTable := make(map[int][]BoardReference)

	calledNumbers := make(map[int]struct{})

	scanner := bufio.NewScanner(input)
	scanner.Split(bufio.ScanLines)

	scanner.Scan()
	sequenceRaw := scanner.Text()
	splitSequence := strings.Split(sequenceRaw, ",")
	sequence := make([]int, len(splitSequence))

	for i, elementRaw := range splitSequence {
		number, _ := strconv.Atoi(elementRaw)
		calledNumbers[number] = struct{}{}
		sequence[i] = number
	}

	// Skip the empty line
	scanner.Scan()

	// Read the first board
	board := initialiseBoard()
	boards := []Board{}
	boardIndex := 0

	row := 0

	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			boards = append(boards, board)
			board = initialiseBoard()
			row = 0
			boardIndex++

			continue
		}

		for i, rawNumber := range strings.Fields(line) {
			number, err := strconv.Atoi(rawNumber)
			if err != nil {
				log.Fatalf("Error parsing board number %d: %s", i, err.Error())
			}

			ref := BoardReference{
				board: boardIndex,
				row:   row,
				col:   i,
			}
			_, ok := bingoTable[number]
			if !ok {
				bingoTable[number] = []BoardReference{ref}
			} else {
				bingoTable[number] = append(bingoTable[number], ref)
			}
			board[row][i] = BingoTile{
				value:   number,
				checked: false,
			}
		}
		row++
	}
	boards = append(boards, board)

	fmt.Printf("Sequence: %v\n", sequence)
	fmt.Println()
	fmt.Println("Boards")
	for _, board := range boards {
		fmt.Printf("%s\n", board.String())
	}

	refs := bingoTable[15]
	fmt.Printf("There are %d %ds across all boards\n", len(refs), 15)
	for i, ref := range refs {
		fmt.Printf("Ref %d has 15 in position %d,%d\n", i, ref.row, ref.col)
	}

	boardsTranspose := make([]Board, len(boards))
	for i := range boardsTranspose {
		boardsTranspose[i] = initialiseBoard()
		for row := range boardsTranspose[i] {
			for col := range boardsTranspose[i][row] {
				boardsTranspose[i][row][col] = boards[i][col][row]
			}
		}
	}
	fmt.Println("Boards transpose")
	for _, board := range boardsTranspose {
		fmt.Printf("%s\n", board.String())
	}

	winningScores := []int{}
	boardsWithWins := make(map[int]struct{})

	for position, number := range sequence {
		refs, ok := bingoTable[number]
		if ok {
			for _, ref := range refs {
				boards[ref.board][ref.row][ref.col].checked = true
				boardsTranspose[ref.board][ref.col][ref.row].checked = true
			}

			for board := range boards {
				if _, ok := boardsWithWins[board]; !ok {
					for i := 0; i < 5; i++ {
						if allChecked(boards[board][i]) {
							fmt.Printf("Bingo! Board %d - row %d after drawing %d (draw number %d)\n", board, i, number, position)
							fmt.Println(boards[board])
							winningScores = append(winningScores, computeBoardScore(boards[board], number))
							boardsWithWins[board] = struct{}{}
							break
						}

						if allChecked(boardsTranspose[board][i]) {
							fmt.Printf("Bingo! Board %d - column %d after drawing %d (draw number %d)\n", board, i, number, position)
							fmt.Println(boardsTranspose[board])
							winningScores = append(winningScores, computeBoardScore(boards[board], number))
							boardsWithWins[board] = struct{}{}
							break
						}
					}
				}
			}
		}
	}

	fmt.Println("=== Part 1 ===")
	fmt.Printf("Answer is: %d\n", winningScores[0])
	fmt.Println()

	fmt.Println("=== Part 2 ===")
	fmt.Printf("Last winning board is: %d\n", winningScores[len(winningScores)-1])
}

func allChecked(tiles []BingoTile) bool {
	result := true
	for _, tile := range tiles {
		result = result && tile.checked
	}

	return result
}

func computeBoardScore(board Board, winningDraw int) int {
	sum := 0
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			tile := board[i][j]
			if !tile.checked {
				sum += tile.value
			}
		}
	}

	return sum * winningDraw
}

func initialiseBoard() Board {
	board := make([][]BingoTile, 5)
	for i := range board {
		board[i] = make([]BingoTile, 5)
	}

	return board
}
