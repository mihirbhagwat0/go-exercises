package main

import (
	"fmt"
	"math/rand"
)

const (
	targetScore = 100
	rounds      = 10
)

func playGame(hold1, hold2 int) bool {
	player1Score, player2Score := 0, 0

	for {
		turnTotal := 0
		for turnTotal < hold1 {
			roll := rand.Intn(6) + 1
			if roll == 1 {
				turnTotal = 0
				break
			}
			turnTotal += roll
		}
		player1Score += turnTotal
		if player1Score >= targetScore {
			return true
		}

		turnTotal = 0
		for turnTotal < hold2 {
			roll := rand.Intn(6) + 1
			if roll == 1 {
				turnTotal = 0
				break
			}
			turnTotal += roll
		}
		player2Score += turnTotal
		if player2Score >= targetScore {
			return false
		}
	}
}

func simulateGames(hold1, hold2 int) (wins, losses int) {
	for i := 0; i < rounds; i++ {
		if playGame(hold1, hold2) {
			wins++
		} else {
			losses++
		}
	}
	return wins, losses
}

func story1() {
	hold1, hold2 := 10, 15
	wins, losses := simulateGames(hold1, hold2)
	winRate := float64(wins) * 100 / float64(rounds)
	lossRate := float64(losses) * 100 / float64(rounds)

	fmt.Printf("Hold %d vs Hold %d: Wins: %d/%d (%.1f%%), Losses: %d/%d (%.1f%%)\n",
		hold1, hold2, wins, rounds, winRate, losses, rounds, lossRate)

}

func story2(player1Hold int) {
	for hold2 := 1; hold2 <= 20; hold2++ {
		wins, losses := simulateGames(player1Hold, hold2)
		fmt.Printf("Holding at %d vs Holding at %d: wins: %d/%d (%.1f%%), losses: %d/%d (%.1f%%)\n",
			player1Hold, hold2, wins, rounds, float64(wins)*100/float64(rounds), losses, rounds, float64(losses)*100/float64(rounds))
	}

	for hold2 := 22; hold2 <= 100; hold2++ {
		wins, losses := simulateGames(player1Hold, hold2)
		fmt.Printf("Holding at %d vs Holding at %d: wins: %d/%d (%.1f%%), losses: %d/%d (%.1f%%)\n",
			player1Hold, hold2, wins, rounds, float64(wins)*100/float64(rounds), losses, rounds, float64(losses)*100/float64(rounds))
	}
}

func main() {
	fmt.Println("Story 1: Fixed Strategies (Hold at 10 vs 15)")
	story1()

	fmt.Println("\nStory 2: Player 1 fixed strategy and Player 2 Changes Strategy (Hold at 1-20, 22-100)")
	story2(21)
}
