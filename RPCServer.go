package main

import (
	"log"
	"net"
	"net/rpc"
)

type Arith int

type Args struct {
	A, B int
}

type Reply struct {
	Result int
}

func (t *Arith) Multiply(args *Args, reply *Reply) error {
	reply.Result = args.A * args.B
	return nil
}

func main() {
	arith := new(Arith)
	rpc.Register(arith)

	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("ListenTCP error:", err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal("Accept error:", err)
		}
		go rpc.ServeConn(conn)
	}

}
