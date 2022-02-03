
from argparse import ArgumentError
from csv import reader

class SudokuBoard:
    def __init__(self, board: list[int]):

        if len(board) != 81:
            raise ArgumentError('board should have 81 elements')

        self.board = []

        for val in board:
            if val == '0':
                val = ' '

            self.board += str(val)

    def __repr__(self):
        hline = '+---+---+---+---+---+---+---+---+---+\n'
        string = hline

        for i in range(0, 81, 9):
            string += '| '
            string += ' | '.join(self.board[i:i+9]) + ' |\n'
            string += hline

        return string

    def solve(self):
        self._solve()

    def _solve(self, i=0):
        
        if i == 81:
            return True

        if self.board[i] != ' ':
            return self._solve(i+1)

        for val in self._find_valid_values(i):
            self.board[i] = val
            if self._solve(i):
                return True

        self.board[i] = ' '

        return False

    def _find_valid_values(self, i: int) -> set:

        sudoku_values = set(['1', '2', '3', '4', '5', '6', '7', '8', '9'])
        taken_values = set()
        
        y_traverser = i - 9
        while y_traverser >= 0:
            taken_values.add(self.board[y_traverser])
            y_traverser -= 9

        y_traverser = i + 9
        while y_traverser < 81:
            taken_values.add(self.board[y_traverser])
            y_traverser += 9

        row_start_x = (i // 9) * 9
        row_end_x = row_start_x + 9 if row_start_x + 9 < 81 else 81

        x_traverser = i - 1
        while x_traverser >= row_start_x:
            taken_values.add(self.board[x_traverser])
            x_traverser -= 1

        x_traverser = i + 1
        while x_traverser < row_end_x:
            taken_values.add(self.board[x_traverser])
            x_traverser += 1

        asdf =  set.difference(sudoku_values, taken_values)
        return asdf


def read_sudoku_file(filename: str) -> list[str]:

    board = []
    with open(filename) as f:
        sudoku_reader = reader(f)
        
        for row in sudoku_reader:
            board.extend(row)

    return board

board = SudokuBoard(read_sudoku_file('sudoku.csv'))

boardLines = str(board).split('\n')

board.solve()
solvedLines = str(board).split('\n')

gap = '  '

print(f'{{0:<{len(boardLines[0])}s}}{gap}Solved:'.format('Original:'))

for original, solved in zip(boardLines, solvedLines):
    print(f'{original}{gap}{solved}')
