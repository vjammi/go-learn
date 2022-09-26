// server.go
package main

import (
	"fmt"
	"github.com/cenkalti/rpc2"
	"net"
)

type Args struct{ A, B int }
type Reply int

func main() {
	srv := rpc2.NewServer()

	handlerFunc := func(client *rpc2.Client, args *Args, reply *Reply) error {
		// Reversed call (server to client)
		var rep Reply
		client.Call("mult", args, &rep)
		fmt.Println("mult result:", rep)

		*reply = Reply(args.A + args.B)
		return nil
	}
	srv.Handle("add", handlerFunc)

	lis, _ := net.Listen("tcp", "127.0.0.1:5000")
	srv.Accept(lis)
}
