package main

import (
	"fmt"
	"time"

	sdk "github.com/saveio/themis-go-sdk"

	"github.com/saveio/themis-go-sdk/client"
)

var singleRpcServers1 = []string{
	"http://127.0.0.1:20336",
	"http://127.0.0.1:50336",
}

func main() {
	testChain := sdk.NewChain()
	testChain.NewRpcClient().SetAddress(singleRpcServers1)
	rpcCl := testChain.NewRpcClient().SetAddress(singleRpcServers1)
	rpcCl.SetCallMode(client.ServerScore)

	t1 := time.Now()
	for ;; {
		height, err := testChain.GetCurrentBlockHeight()
		if err != nil {
			fmt.Println("[ERROR] LoopTestRpc BlockHeight Error: ", err.Error())
		} else {
			fmt.Println("        LoopTestRpc BlockHeight Height: ", height)
		}
		time.Sleep(50 * time.Millisecond)
	}
	fmt.Println(time.Since(t1).String())
}
