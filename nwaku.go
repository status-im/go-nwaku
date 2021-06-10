package main

import (
	"context"
	"fmt"
	"time"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/status-im/go-nwaku/nwaku"
)

func main() {
	fmt.Println("Starting node...")
	nwaku.StartNode()

	fmt.Println("JSON RPC request...")

	_, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Assumes node started
	client, _ := rpc.Dial("http://127.0.0.1:8545")

	// Get node info
	var wakuInfo = nwaku.GetWakuDebugInfo(client)
	fmt.Println("WakuInfo ListenStr", wakuInfo.ListenStr)

	// Query messages
	var contentTopic = "/toy-chat/2/huilong/proto"
	var storeResponse = nwaku.GetWakuStoreMessages(client, contentTopic)
	fmt.Println("Fetched", len(storeResponse.Messages), "messages")

	// Publish
	var message = nwaku.WakuRelayMessage{Payload: "0x1a2b3c4d5e6f", ContentTopic: contentTopic}
	var res = nwaku.PostWakuRelayMessage(client, message)
	fmt.Println("Publish", res)
}
