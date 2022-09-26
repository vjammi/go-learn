// client.go
package main

import (
	"fmt"
	"github.com/cenkalti/rpc2"
	"net"
)

type Args struct{ A, B int }
type Reply int

func main() {
	conn, _ := net.Dial("tcp", "127.0.0.1:5000")
	clt := rpc2.NewClient(conn)

	handlerFunc := func(client *rpc2.Client, args *Args, reply *Reply) error {
		*reply = Reply(args.A * args.B)
		return nil
	}

	clt.Handle("mult", handlerFunc)
	go clt.Run()

	var rep Reply
	clt.Call("add", Args{5, 2}, &rep)
	fmt.Println("add result:", rep)
}
