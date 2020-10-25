package game

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

// SumInt realiza a soma de uma fatia (slice) do tipo Int
func SumInt(x []int) int {
	total := 0
	for _, v := range x {
		total += v
	}
	return total
}

// CheckRows realiza a checagem do vencedor linha a linha
func CheckRows(data [][]int) int {
	for _, row := range data {
		result := SumInt(row)
		if result == 3 {
			return 3
		} else if result == 30 {
			return 30
		}
	}
	return 0
}

// CheckCols realiza a checagem do vencedor coluna a coluna
func CheckCols(data [][]int) int {
	for col := 0; col < 3; col++ {
		values := []int{}
		for row := 0; row < 3; row++ {
			values = append(values, data[row][col])
		}
		result := SumInt(values)
		if result == 3 {
			return 3
		} else if result == 30 {
			return 30
		}
	}
	return 0
}

// CheckDiag realiza a checagem do vencedor na diagonal principal
func CheckDiag(data [][]int) int {
	values := []int{}
	for i := 0; i < 3; i++ {
		values = append(values, data[i][i])
	}
	result := SumInt(values)
	if result == 3 {
		return 3
	} else if result == 30 {
		return 30
	}
	return 0
}

// CheckDiagSec realiza a checagem do vencedor na diagonal secundária
func CheckDiagSec(data [][]int) int {
	values := []int{}
	j := 2
	for i := 0; i < 3; i++ {
		values = append(values, data[i][j-i])
	}
	result := SumInt(values)
	if result == 3 {
		return 3
	} else if result == 30 {
		return 30
	}
	return 0
}

// ClearScreen limpa tela
func ClearScreen() {
	cmd := exec.Command("")
	if runtime.GOOS == "linux" {
		cmd = exec.Command("clear")
	} else if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls")
	}
	cmd.Stdout = os.Stdout
	cmd.Run()
}

// GetChoice realiza a escolha do usuário com validação
func GetChoice(board Board, player Player) {
	for result := false; result != true; {
		board.PrintBoard()
		fmt.Printf("Escolha sua jogada, %s (%s): ", player.Name, player.Symbol)
		choice := 0
		fmt.Scanf("%d", &choice)
		result = board.UpdateBoard(player, choice)
	}

}
