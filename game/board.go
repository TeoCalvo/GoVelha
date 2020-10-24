package game

import (
	"fmt"
	"io/ioutil"
	"strconv"
)

// Board é a estrutura do tabuleiro
type Board struct {
	BoardString  string
	BoardDigital [][]int
	Players      []Player
}

// LoadBoard carrega os dados em string do tabuleiro
func (b *Board) LoadBoard() {
	data, _ := ioutil.ReadFile("board.txt")
	b.BoardString = string(data)

	b.BoardDigital = [][]int{
		[]int{0, 0, 0},
		[]int{0, 0, 0},
		[]int{0, 0, 0},
	}
}

// PrintBoard exibe board com base nos dados inseridos
func (b *Board) PrintBoard() {
	data := []string{}
	count := 1
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if b.BoardDigital[i][j] == 0 {
				data = append(data, strconv.Itoa(count))
			} else if b.BoardDigital[i][j] == 1 {
				data = append(data, "X")
			} else {
				data = append(data, "O")
			}
			count++
		}
	}
	fmt.Printf(b.BoardString,
		data[0],
		data[1],
		data[2],
		data[3],
		data[4],
		data[5],
		data[6],
		data[7],
		data[8])
}

// UpdateBoard atualiza os dados do tabuleiro
func (b *Board) UpdateBoard(p Player, pos int) {
	count := 1
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if pos == count {
				b.BoardDigital[i][j] = p.Number
				i, j = 3, 3
			}
			count++
		}
	}
}

func (b *Board) FindWinner() bool {

	checks := []int{
		CheckCols(b.BoardDigital),
		CheckRows(b.BoardDigital),
		CheckDiag(b.BoardDigital),
		CheckDiagSec(b.BoardDigital)}

	for _, c := range checks {
		result := c
		if result == 3 {
			b.Players[0].Winner = true
			return true
		} else if result == 30 {
			b.Players[1].Winner = true
			return true
		}
	}
	return false
}
