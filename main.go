package main

import (
	"GoVelha/game"
	"fmt"
	"math/rand"
	"time"
)

func main() {

	game.ClearScreen()

	player1 := game.NewPlayer("", "X") // jogador 1
	fmt.Print("Player 1 (X), entre com o seu nome: ")
	fmt.Scanf("%s", &player1.Name)

	player2 := game.NewPlayer("", "O") // jogador 1
	fmt.Print("Player 2 (O), entre com o seu nome: ")
	fmt.Scanf("%s", &player2.Name)

	gameBoard := game.Board{Players: []game.Player{*player1, *player2}} // tabuleiro
	gameBoard.LoadBoard()

	rand.Seed(time.Now().UnixNano())

	for playerIndex := rand.Intn(2); gameBoard.FindWinner() != true; playerIndex++ {

		if playerIndex >= len(gameBoard.Players) {
			playerIndex = 0
		}

		player := gameBoard.Players[playerIndex]
		game.GetChoice(gameBoard, player)
	}

	gameBoard.PrintBoard()
	for _, p := range gameBoard.Players {
		if p.Winner {
			fmt.Printf("Parabéns, %s!! Você ganhou!\n", p.Name)
		}
	}

}
