package main

import (
	"bufio"
	"fmt"
	"github.com/brunokarpo-codings-kata/jogodavelha/board"
	"github.com/brunokarpo-codings-kata/jogodavelha/game"
	"github.com/brunokarpo-codings-kata/jogodavelha/player"
	"os"
)

var player1 player.Player
var player2 player.Player

func main() {
	fmt.Print("Write the name of player 1: ")
	reader := bufio.NewReader(os.Stdin)
	if name, err := reader.ReadString('\n'); err == nil {
		player1 = player.Player{
			Name: name,
			Mark: "X",
		}
	}
	fmt.Print("Write the name of player 2: ")
	if name, err := reader.ReadString('\n'); err == nil {
		player2 = player.Player{
			Name: name,
			Mark: "O",
		}
	}

	b := board.Init()
	g := game.Init(&player1, &player2, b)

	for {
		b.PrintBoard()
		var x, y int
		turnPlayer := g.GetTheTurnPlayer()
		fmt.Printf("%v type the board position you wanna mark: ", turnPlayer.Name)
		fmt.Scanf("%d %d", &x, &y)
		winner, err := g.Play(x, y)
		if err != nil {
			fmt.Println(err.Error())
			if err == board.DeuVelhaError {
				break
			}
		}
		if winner != nil {
			fmt.Printf("%v win the game\n", winner.Name)
			break
		}
	}
	b.PrintBoard()
}
