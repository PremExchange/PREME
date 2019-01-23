package main

import (
	"log"
	"os"

	"github.com/mit-dci/opencx/cxrpc"
)

var (
	defaultServer = "localhost"
	defaultPort   = 12345
)

// TODO figure out this, call in functions specific to method
type openCxClient struct {
	RPCClient *cxrpc.OpencxRPCClient
}

// opencx-cli is the client, opencx is the server
func main() {
	var err error

	commandArg := os.Args[1:]

	client := new(openCxClient)
	err = client.setupCxClient(defaultServer, defaultPort)

	if err != nil {
		log.Fatalf("Error setting up OpenCX RPC Client: \n%s", err)
	}

	// TODO just for now
	err = client.parseCommands(commandArg)
	if err != nil {
		log.Fatalf("Error parsing commands: \n%s", err)
	}
}

// NewOpenCxClient creates a new openCxClient for use as an RPC Client
func(cl *openCxClient) setupCxClient(server string, port int) error {
	var err error

	cl.RPCClient, err = cxrpc.NewOpencxRPCClient(server, port)
	if err != nil {
		return err
	}

	return nil
}

func(cl *openCxClient) Call(serviceMethod string, args interface{}, reply interface{}) error {
	return cl.RPCClient.Call(serviceMethod, args, reply)
}
