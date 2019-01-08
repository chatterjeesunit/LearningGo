package main

import (
	"fmt"
	"math/rand"
	"time"
)

/**
Simulate a single round of Ping Pong between two teams of players..
1. Each Team plays alternatively. First Team1 and then Team 2
2. Within each Team, each player plays alternatively

3. So the order of playing  will be
Team 1 : Player 1
Team 2 : Player 1
Team 1 : Player 2
Team 2 : Player 2

4. Each player sleeps for a random time interval when they play (upto a max of 750 ms)
- to simulate time taken by a player to play.

5. Print each player's name and the time taken to play (the random sleep time)

6. Create 4 go routines for each player playing the game.

7. Make use of channel to make sure each player and team plays as per their turn.

8. If any player takes more than 500 ms to play his shot, then that player and team loses the game and game is finished

9. After a game is finished, print the name of the team that has won the game.
*/

var ()

const (
	MAX_TIME_PER_PLAYER_TO_PLAY = time.Duration(500 * time.Millisecond)
	MAX_RANDOM_SLEEP_TIME       = 750
)

type Player struct {
	name    string
	canPlay chan bool
}

type Team struct {
	name               string
	players            [2]Player
	currentPlayerIndex int
}

type Game struct {
	teams            [2]Team
	finishGame       chan bool
	currentTeamIndex int
}

func (game *Game) playNextRound() {

	//Current Team plays
	currentTeam := game.teams[game.currentTeamIndex]
	currentTeam.play(game)

	//Change the next team to play in the game
	game.currentTeamIndex = flipIndex(game.currentTeamIndex)
}

func (team *Team) play(game *Game) {
	//Allow current player to play
	currentPlayer := team.players[team.currentPlayerIndex]
	currentPlayer.canPlay <- true

	//Change the next player for the team
	team.currentPlayerIndex = flipIndex(team.currentPlayerIndex)

	//Update the current team
	game.teams[game.currentTeamIndex] = *team
}

func (player *Player) play(game *Game) {
	fmt.Printf("\t%s - Ready\n", player.name)
	for {
		<-player.canPlay

		sleepTime := fetchRandomSleepTime()

		if sleepTime < MAX_TIME_PER_PLAYER_TO_PLAY {
			fmt.Printf("\t\t%s \t: played in %v\n", player.name, sleepTime)
			time.Sleep(sleepTime)
			game.finishGame <- false
		} else {
			fmt.Printf("\t\t%s \t: took more than %v to play his shot, and hence their team lost the match.\n", player.name, MAX_TIME_PER_PLAYER_TO_PLAY)
			game.finishGame <- true
		}
	}
}

func main() {
	fmt.Println("Starting a new Game")

	game := startGame()

	for {
		finishGame := <-game.finishGame

		if finishGame {
			break
		} else {
			game.playNextRound()
		}

	}

	fmt.Println("Game is finished....")

	fmt.Printf("Team %s has WON\n", game.teams[game.currentTeamIndex].name)

}

func startGame() Game {
	player1 := Player{"John", make(chan bool, 1)}
	player2 := Player{"Joe", make(chan bool, 1)}
	player3 := Player{"Anna", make(chan bool, 1)}
	player4 := Player{"Sara", make(chan bool, 1)}

	team1 := Team{"Boys", [2]Player{player1, player2}, 0}
	team2 := Team{"Girls", [2]Player{player3, player4}, 0}

	game := Game{[2]Team{team1, team2}, make(chan bool, 1), 0}

	go game.teams[0].players[0].play(&game)
	go game.teams[0].players[1].play(&game)
	go game.teams[1].players[0].play(&game)
	go game.teams[1].players[1].play(&game)

	game.finishGame <- false

	return game
}

func flipIndex(currentIndex int) int {
	if currentIndex == 0 {
		return 1
	}
	return 0
}

func fetchRandomSleepTime() time.Duration {
	rand.Seed(time.Now().UnixNano())
	randomNumber := rand.Int63n(MAX_RANDOM_SLEEP_TIME)
	return time.Duration(randomNumber) * time.Millisecond
}
