package main

import "fmt"

type SudokuBoard struct {
	board [9 * 9]byte
}

func NewSudokuBoard(board []byte) SudokuBoard {
	b := SudokuBoard{}

	for i := range b.board {
		b.board[i] = board[i]
	}

	return b
}

func index(x int, y int) int {
	return y*9 + x
}

func (b *SudokuBoard) get(x int, y int) byte {
	return b.board[index(x, y)]
}

func (b *SudokuBoard) set(x int, y int, val byte) {
	b.board[index(x, y)] = val
}

func (b *SudokuBoard) toString() string {

	hline := "+---+---+---+---+---+---+---+---+---+\n"
	str := hline

	i := 0
	for i < 81 {
		end := i + 8

		str += "| "

		for i < end {

			str += fmt.Sprintf("%c | ", b.board[i])
			i++
		}

		str += fmt.Sprintf("%c |\n", b.board[i])
		str += hline
		i++
	}

	return str
}

func (b *SudokuBoard) solve() {
	solve(b, 0, 0)
}

func solve(b *SudokuBoard, x int, y int) bool {

	if y == 9 {
		return true
	}

	if x == 9 {
		return solve(b, 0, y+1)
	}

	if b.get(x, y) != ' ' {
		return solve(b, x+1, y)
	}

	values := validValues(b, x, y)

	if len(values) == 0 {
		return false
	}

	for _, val := range values {
		b.set(x, y, val)
		if solve(b, x, y) {
			return true
		}
	}

	b.set(x, y, ' ')

	return false
}

func validValues(b *SudokuBoard, x int, y int) []byte {
	set := NewSet()

	my_pos := index(x, y)

	for y_traverser := index(x, 0); y_traverser < 81; y_traverser += 9 {
		if y_traverser != my_pos {
			set.add(b.board[y_traverser])
		}
	}

	x_traverser := index(0, y)
	next_row := x_traverser + 9
	for x_traverser < next_row {
		if x_traverser != next_row {
			set.add(b.board[x_traverser])
		}

		x_traverser++
	}

	values := make([]byte, 0)

	for ch := '1'; ch <= '9'; ch++ {
		if set.exists(byte(ch)) {
			continue
		}

		values = append(values, byte(ch))
	}

	return values
}
