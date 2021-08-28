package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/alftirta/dice-game/dg"
)

func init() {
	log.SetPrefix("Error: ")
	log.SetFlags(0)
}

func main() {
	// Get totalPlayers from the user's input.
	fmt.Print("Jumlah pemain: ")
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n'); if err != nil {log.Fatal(err)}
	input = strings.TrimSuffix(input, "\n")
	totalPlayers, err := strconv.Atoi(input); if err != nil {log.Fatal(err)}
	if totalPlayers < 1 {
		log.Fatal("You need at least 1 player to play the game.")
	}

	// Get totalDice from the user's input.
	fmt.Print("Jumlah dadu: ")
	reader = bufio.NewReader(os.Stdin)
	input, err = reader.ReadString('\n'); if err != nil {log.Fatal(err)}
	input = strings.TrimSuffix(input, "\n")
	totalDice, err := strconv.Atoi(input); if err != nil {log.Fatal(err)}
	if totalDice < 1 {
		log.Fatal("You need at least 1 dice for each player to play the game.")
	}

	// Get timeDelay from the user's input.
	fmt.Print("Waktu delay (dalam detik): ")
	reader = bufio.NewReader(os.Stdin)
	input, err = reader.ReadString('\n'); if err != nil {log.Fatal(err)}
	input = strings.TrimSuffix(input, "\n")
	inputInt, err := strconv.Atoi(input); if err != nil {log.Fatal(err)}
	if inputInt < 0 {
		log.Fatal("You can only input non-negative delay time.")
	}
	timeDelay := time.Duration(inputInt)

	// Create players.
	players, err := dg.CreatePlayers(totalPlayers, totalDice); if err != nil {log.Fatal(err)}

	// Create game.
	game, err := dg.CreateGame(players, timeDelay); if err != nil {log.Fatal(err)}

	// Play the game.
	game.Play()
}
