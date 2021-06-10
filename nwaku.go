package main

import (
	"context"
	"fmt"
	"log"
	"time"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/status-im/go-nwaku/nwaku"
)

func main() {
	fmt.Println("Starting node...")

	nodeStopped := make(chan bool, 1)
	go nwaku.StartNode(nodeStopped)

	// NOTE: RPC server needs time to start, this can be improved
	time.Sleep(2 * time.Second)

	fmt.Println("JSON RPC request...")

	_, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Assumes node started
	client, _ := rpc.Dial("http://127.0.0.1:8545")

	var defaultTopic = "/waku/2/default-waku/proto"
	var contentTopic = "/toy-chat/2/huilong/proto"

	// Get node info
	var wakuInfo, _ = nwaku.GetWakuDebugInfo(client)
	fmt.Println("WakuInfo ListenStr", wakuInfo.ListenStr)

	// Query messages
	var storeResponse, _ = nwaku.GetWakuStoreMessages(client, contentTopic)
	fmt.Println("Fetched", len(storeResponse.Messages), "messages")

	// Subscribe
	var res, _ = nwaku.PostWakuRelaySubscriptions(client, []string{defaultTopic})
	fmt.Println("Subscribe", res)

	// Publish
	var message = nwaku.WakuRelayMessage{Payload: "0x1a2b3c4d5e6f", ContentTopic: contentTopic}
	var res2, _ = nwaku.PostWakuRelayMessage(client, message)
	fmt.Println("Publish", res2)

	// Get messages
	var wakuMessages, _ = nwaku.GetWakuRelayMessages(client, defaultTopic)
	fmt.Println("Get messages", wakuMessages)

    <-nodeStopped
    log.Printf("exiting main")
}
