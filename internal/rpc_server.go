package internal

import (
	"errors"
	"fmt"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

// Args holds arguments passed to RPC methods
type Args struct {
	A, B int
}

// Arith provides an RPC service with arithmetic operations
type Arith int

// Multiply is an RPC method for multiplying two integers
func (t *Arith) Multiply(args *Args, reply *int) error {
	*reply = args.A * args.B
	return nil
}

// Divide is an RPC method for dividing two integers
func (t *Arith) Divide(args *Args, reply *float64) error {
	if args.B == 0 {
		return errors.New("divide by zero")
	}
	*reply = float64(args.A) / float64(args.B)
	return nil
}

type HelloService struct{}
type HelloArg struct {
	Name string
}

func (t *HelloService) Hello(args *HelloArg, reply *string) error {
	if args.Name == "" {
		return errors.New("name is empty")
	}

	*reply = fmt.Sprintf("Hello, %s!", args.Name)

	return nil
}

func RPCServer() {
	hello := new(HelloService)
	rpc.Register(hello)

	arith := new(Arith)
	rpc.Register(arith)
	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("Listen error:", err)
	}
	log.Println("Serving RPC handler")

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal("Accept error:", err)
		}
		go rpc.ServeCodec(jsonrpc.NewServerCodec(conn))
	}
}
