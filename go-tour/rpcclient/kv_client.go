package main

import (
	"fmt"
	"log"
	"net/rpc"
)

//
// Common RPC request/reply definitions
//

//const (
//	OK       = "OK"
//	ErrNoKey = "ErrNoKey"
//)

// type Err string
type PutArgs struct {
	Key   string
	Value string
}

type Err string

type PutReply struct {
	Err Err
}

type GetArgs struct {
	Key string
}

type GetReply struct {
	Err   Err
	Value string
}

//
// Client
//

func connect() *rpc.Client {
	client, err := rpc.Dial("tcp", ":1234")
	if err != nil {
		log.Fatal("dialing:", err)
	}
	return client
}

func get(key string) string {
	client := connect()
	args := GetArgs{"subject"}
	reply := GetReply{}
	err := client.Call("KV.Get", &args, &reply)
	if err != nil {
		log.Fatal("error:", err)
	}
	client.Close()
	return reply.Value
}

func put(key string, val string) {
	client := connect()
	args := PutArgs{"subject", "6.824"}
	reply := PutReply{}
	err := client.Call("KV.Put", &args, &reply)
	if err != nil {
		log.Fatal("error:", err)
	}
	client.Close()
}

//
// main
//

func main() {
	//server()

	put("subject1", "6.824-1")
	put("subject2", "6.824-2")
	put("subject3", "6.824-3")
	put("subject4", "6.824-4")
	put("subject5", "6.824-5")
	put("subject6", "6.824-6")
	fmt.Printf("Put 6 records in kv db")

	s1 := get("subject1")
	fmt.Printf("get(subject1) -> %s\n", s1)
	s2 := get("subject2")
	fmt.Printf("get(subject2) -> %s\n", s2)
	s3 := get("subject3")
	fmt.Printf("get(subject3) -> %s\n", s3)
	s4 := get("subject4")
	fmt.Printf("get(subject4) -> %s\n", s4)
	s5 := get("subject5")
	fmt.Printf("get(subject5) -> %s\n", s5)
	s6 := get("subject6")
	fmt.Printf("get(subject6) -> %s\n", s6)

	print("shutting down the kv db...")
}
