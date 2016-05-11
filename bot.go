package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"

	"os"

	"github.com/mleonard87/merknera-tictactoe-bot/registration"
	"github.com/mleonard87/merknera-tictactoe-bot/services"
	"github.com/mleonard87/rpc"
	"github.com/mleonard87/rpc/json"
)

func Init() {
	s := rpc.NewServer()
	s.RegisterCodec(json.NewCodec(), "application/json")
	s.RegisterService(new(services.Status), "")
	s.RegisterService(new(services.TicTacToe), "")
	http.Handle("/rpc", s)

	port := os.Getenv("TTT_BOT_PORT")
	portStr := fmt.Sprintf(":%s", port)

	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func() {
		fmt.Printf("Merknera Tic-Tac-Toe bot is now listening on localhost:%s\n", port)
		log.Fatal(http.ListenAndServe(portStr, nil))
		wg.Done()
	}()

	registration.Register()

	wg.Wait()
}

func main() {
	Init()
}
