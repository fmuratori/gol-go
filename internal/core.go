package core

type Coordinates struct {
	I int
	J int
}

func NewCoordinates(i int, j int) Coordinates {
	return Coordinates{
		I: i,
		J: j,
	}
}

type Board struct {
	width  int
	height int
	board  [][]int
}

func NewBoard(width int, height int) *Board {
	return &Board{
		width:  width,
		height: height,
		board:  CreateMatrix(width, height),
	}
}

func (b *Board) AddCell(i int, j int) {
	b.board[i][j] = 1
}

func (b *Board) Tick() {
	newBoard := MatrixDeepCopy(b.board)

	for i := 0; i < b.height; i++ {
		for j := 0; j < b.width; j++ {

			count := b.countCloseCells(i, j)
			if newBoard[i][j] == 1 && (count <= 1 || count >= 4) {
				newBoard[i][j] = 0

			} else if newBoard[i][j] == 0 && count == 3 {
				newBoard[i][j] = 1
			}
		}
	}

	b.board = newBoard
}

func (b *Board) countCloseCells(i int, j int) int {
	count := 0
	for di := -1; di < 2; di++ {
		for dj := -1; dj < 2; dj++ {
			ni := i + di
			nj := j + dj

			if di == 0 && dj == 0 {
				continue
			}

			if ni < 0 || ni >= b.height || nj < 0 || nj >= b.width {
				continue
			}

			if b.board[ni][nj] == 1 {
				count += 1
			}
		}
	}
	return count
}

func (b *Board) String() string {
	return MatrixToString(b.board)
}
