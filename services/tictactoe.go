package services

import (
	"fmt"
	"net/http"
)

type NextMoveArgs struct {
	GameId    int      `json:"gameid"`
	Mark      string   `json:"mark"`
	GameState []string `json:"gamestate"`
}

type NextMoveReply struct {
	Position int `json:"position"`
}

type CompleteArgs struct {
	GameId    int      `json:"gameid"`
	Winner    bool     `json:"winner"`
	Mark      string   `json:"mark"`
	GameState []string `json:"gamestate"`
}

type ErrorArgs struct {
	GameId    int    `json:"gameid"`
	Message   string `json:"message"`
	ErrorCode int    `json:"errorcode"`
}

type TicTacToe struct{}

func (ttt *TicTacToe) NextMove(r *http.Request, args *NextMoveArgs, reply *NextMoveReply) error {

	fmt.Printf("NextMove for GameId: %d - you are %s\n", args.GameId, args.Mark)

	fmt.Println(args.GameState)

	for i, v := range args.GameState {
		if v != "X" && v != "O" {
			fmt.Printf("Playing: %d\n\n", i)
			reply.Position = i
			return nil
		}
	}

	return nil
}

func (ttt *TicTacToe) Complete(r *http.Request, args *CompleteArgs, reply *interface{}) error {
	if args.Winner {
		fmt.Printf("GameID: %d - You won!\n", args.GameId)
	} else {
		fmt.Printf("GameID: %d - You lost!\n", args.GameId)
	}

	return nil
}

func (ttt *TicTacToe) Error(r *http.Request, args *ErrorArgs, reply *interface{}) error {
	fmt.Printf("Error - %d: %s\n", args.ErrorCode, args.Message)

	return nil
}
