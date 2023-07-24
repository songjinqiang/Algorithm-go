package main

import (
	"fmt"
	"log"
	"net/rpc"
)

type Args struct {
	A, B int
}

type Reply struct {
	Result int
}

func main() {
	client, err := rpc.Dial("tcp", "localhost:1234")
	if err != nil {
		log.Fatal("Dial error:", err)
	}
	args := &Args{A: 10, B: 5}
	reply := new(Reply)
	err = client.Call("Arith.Multiply", args, reply)
	if err != nil {
		log.Fatal("Call error:", err)
	}
	fmt.Println("Result:", reply.Result)

}
