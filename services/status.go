package services

import (
	"fmt"
	"net/http"
)

type StatusPingArgs struct{}

type StatusReplyArgs struct {
	Ping string `json:"ping"`
}

type Status struct{}

func (ss *Status) Ping(r *http.Request, args *StatusPingArgs, reply *StatusReplyArgs) error {
	fmt.Println("Ping!")
	reply.Ping = "OK"
	return nil
}
