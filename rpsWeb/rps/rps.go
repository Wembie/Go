package rps

import (
	"math/rand"
	"strconv"
)

const (
	ROCK     = 0 //Piedra > Tijeras; (tijeras + 1) % 3 = 0
	PAPER    = 1 //Papel > Piedra; (piedra + 1) % 3 = 1
	SCISSORS = 2 //Tijeras > Papel; (papel + 1) % 3 = 2
)

type Round struct {
	Message           string `json:"message"`
	ComputerChoice    string `json:"computer_choice"`
	RoundResult       string `json:"round_result"`
	ComputerChoiceInt int    `json:"computer_choice_int"`
	ComputerScore     string `json:"computer_score"`
	PlayerScore       string `json:"player_score"`
}

var winMessages = []string{
	"Congrats",
	"Good Work",
	"Good Luck",
	"Melo-Cara-Melo-PA",
}

var loseMessages = []string{
	"F",
	"Try again :D",
	"Bad Luck",
	"Cha loka",
}

var drawMessages = []string{
	"=",
	"AI = Human brain?",
	"Draw",
	"Como le empata un IA, xd",
}

var ComputerScore, PlayerScore int

func PlayRound(playerValue int) Round {
	computerValue := rand.Intn(3)
	
	var computerChoice, roundResult string
	var computerChoiceInt int

	switch computerValue{
		case ROCK:
			computerChoiceInt = ROCK
			computerChoice = "AI chose ROCK"
		case PAPER:
			computerChoiceInt = PAPER
			computerChoice = "AI chose PAPER"
		case SCISSORS:
			computerChoiceInt = SCISSORS
			computerChoice = "AI chose SCISSORS"
	}
	messageInt := rand.Intn(4)

	var message string
	if playerValue == computerValue{
		roundResult = "Draw"
		message = drawMessages[messageInt]
	} else if playerValue == (computerValue+1)%3{
		PlayerScore++
		roundResult = "The Player Wins"
		message = winMessages[messageInt]
	} else {
		ComputerScore++
		roundResult = "AI Wins"
		message = loseMessages[messageInt]
	}
	return Round{
		Message: 			message,
		ComputerChoice: 	computerChoice,
		RoundResult: 		roundResult,
		ComputerChoiceInt: 	computerChoiceInt,
		ComputerScore: 		strconv.Itoa(ComputerScore),
		PlayerScore: 		strconv.Itoa(PlayerScore),
	}
}