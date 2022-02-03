package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func readSudoku() []byte {
	filename := "sudoku.csv"

	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	sudokuBoard := make([]byte, 0)

	reader := bufio.NewReader(f)

	for {
		if ch, err := reader.ReadByte(); err != nil {
			if err == io.EOF {
				break
			} else {
				log.Fatal(err)
			}
		} else {
			if ch != ',' && ch != '\n' && ch != '\r' {
				if ch < '1' || ch > '9' {
					if ch == '0' {
						ch = ' '
					} else {
						log.Fatal(fmt.Sprintf("Invalid number: \"%c\"asdf", ch))
					}
				}

				sudokuBoard = append(sudokuBoard, ch)
			}
		}
	}

	if len(sudokuBoard) != 81 {
		return nil
	}

	return sudokuBoard
}

func main() {

	board := NewSudokuBoard(readSudoku())

	solvedBoard := board
	solvedBoard.solve()

	boardString := strings.Split(board.toString(), "\n")
	solvedString := strings.Split(solvedBoard.toString(), "\n")
	gap := "  "

	boardWidth := len(boardString[0])

	fmt.Printf("%-*s", boardWidth, "Original:")
	fmt.Printf("%sSolved:\n", gap)

	for i := range boardString {
		fmt.Printf("%s%s%s\n", boardString[i], gap, solvedString[i])
	}
}
