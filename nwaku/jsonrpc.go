package main

import (
	"context"
	"fmt"
	"time"
	"github.com/ethereum/go-ethereum/rpc"
)

// NOTE Can be generalized with different transports (HTTP, IPC, WS etc)
// https://github.com/ethereum/go-ethereum/blob/master/rpc/client.go#L169

type WakuInfo struct {
	ListenStr string `json:"listenStr"`
}

func main() {
	fmt.Println("JSON RPC request...")

	_, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Assumes node started
	client, _ := rpc.Dial("http://127.0.0.1:8545")
	var wakuInfo WakuInfo
	// TODO id, params?
	if err := client.Call(&wakuInfo, "get_waku_v2_debug_v1_info"); err != nil {
		panic(err)
	}
	fmt.Println("WakuInfo ListenStr", wakuInfo.ListenStr)
}
