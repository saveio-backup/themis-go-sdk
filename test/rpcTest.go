package main

import (
	"fmt"
	"time"

	sdk "github.com/saveio/themis-go-sdk"

)

var singleRpcServers1 = []string{
	"http://127.0.0.1:20336",
	"http://127.0.0.1:50336",
}

func main() {
	testChain := sdk.NewChain()
	testChain.NewRpcClient().SetAddress(singleRpcServers1)
	testChain.NewRpcClient().SetAddress(singleRpcServers1)
	//rpcCl := testChain.NewRpcClient().SetAddress(singleRpcServers1)
	//rpcCl.SetCallMode(client.ServerScore)

	t1 := time.Now()
	for ;; {
		hash, err := testChain.GetBlockHash(5)
		if err != nil {
			fmt.Println("[ERROR] LoopTestRpc GetBlockHash Error: ", err.Error())
		} else {
			fmt.Println("        LoopTestRpc GetBlockHash Height: ", hash.ToHexString())
		}


		hashString := hash.ToHexString()
		block, err := testChain.GetBlockByHash(hashString)
		if err != nil {
			fmt.Println("[ERROR] LoopTestRpc GetBlockByHash Error: ", err.Error())
		} else {
			blockHash := block.Hash()
			fmt.Println("        LoopTestRpc GetBlockByHash Height: ", blockHash.ToHexString())
		}

		//height, err := testChain.GetCurrentBlockHeight()
		//if err != nil {
		//	fmt.Println("[ERROR] LoopTestRpc BlockHeight Error: ", err.Error())
		//} else {
		//	fmt.Println("        LoopTestRpc BlockHeight Height: ", height)
		//}
		time.Sleep(10 * time.Millisecond)
	}
	fmt.Println(time.Since(t1).String())
}
