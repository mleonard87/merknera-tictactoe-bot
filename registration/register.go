package registration

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

const (
	BOT_NAME             = "[Merknera] Tic-Tac-Toe"
	BOT_VERSION          = "0.0.8"
	GAME_MNEMONIC        = "TICTACTOE"
	PROGRAMMING_LANGUAGE = "Go"
	WEBSITE              = "https://github.com/mleonard87/merknera-tictactoe-bot"
	DESCRIPTION          = `
		This bot is provided by Merknera and should always be available online. This bot has been used primarily
		for testing of the Tic-Tac-Toe game within Merknera and as such does not implement a sophisticated
		algorithm - it simply selects the first available space on the board. If your bot implements a reasonable
		algorithm it should always win against this bot.
		`
	RPC_REGISTRATION_METHOD_NAME = "RegistrationService.Register"
)

type RPCClientRequest struct {
	JsonRpcVersion string `json:"jsonrpc,omitempty"`
	Method         string `json:"method"`
	Params         struct {
		Token               string `json:"token"`
		BotName             string `json:"botname"`
		BotVersion          string `json:"botversion"`
		Game                string `json:"game"`
		RpcEndpoint         string `json:"rpcendpoint"`
		ProgrammingLanguage string `json:"programminglanguage"`
		Website             string `json:"website"`
		Description         string `json:"description"`
	} `json:"params"`
	Id int `json:"id"`
}

type RPCServerResponse struct {
	JsonRpcVersion string `json:"jsonrpc,omitempty"`
	Result         struct {
		Message string `json:"message"`
	} `json:"result,omitempty"`
	Error string `json:"error,omitempty"`
	Id    int    `json:"id"`
}

func Register() {
	merkneraEndpoint := os.Getenv("TTT_BOT_MERKNERA_URL")
	token := os.Getenv("TTT_BOT_TOKEN")
	port := os.Getenv("TTT_BOT_PORT")
	botEndpoint := os.Getenv("TTT_BOT_ENDPOINT_URL")

	rcr := new(RPCClientRequest)
	rcr.JsonRpcVersion = "2.0"
	rcr.Id = 1
	rcr.Method = RPC_REGISTRATION_METHOD_NAME
	rcr.Params.Token = token
	rcr.Params.BotName = BOT_NAME + " (" + port + ")"
	rcr.Params.BotVersion = BOT_VERSION
	rcr.Params.Game = GAME_MNEMONIC
	rcr.Params.RpcEndpoint = botEndpoint
	rcr.Params.ProgrammingLanguage = PROGRAMMING_LANGUAGE
	rcr.Params.Website = WEBSITE
	rcr.Params.Description = DESCRIPTION

	jsonBody, err := json.Marshal(*rcr)
	if err != nil {
		log.Fatal(err)
	}

	client := &http.Client{}

	req, err := http.NewRequest("POST", merkneraEndpoint, bytes.NewBuffer(jsonBody))
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")

	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()
	nextMoveResponse, err := ioutil.ReadAll(res.Body)
	res.Body.Close()

	reply := RPCServerResponse{}
	err = json.Unmarshal(nextMoveResponse, &reply)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%+v%+v\n", reply.Result.Message, reply.Error)
}
