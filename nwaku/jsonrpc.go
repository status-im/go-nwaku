package main

import (
	"context"
	"fmt"
	"time"
	"github.com/ethereum/go-ethereum/rpc"
)

// NOTE Can be generalized with different transports (HTTP, IPC, WS etc)
// https://github.com/ethereum/go-ethereum/blob/master/rpc/client.go#L169

// Later: Generalize with CallContext (not prio)

// TODO Move to types
type WakuInfo struct {
	ListenStr string `json:"listenStr"`
}

// XXX in this case Payload isn't string but something else
// panic: json: cannot unmarshal array into Go struct field WakuMessage.messages.payload of type string
// TODO This should be toy-chat protobuf probably
type WakuMessage struct {
	Payload []byte `json:"payload"`
	ContentTopic string `json:"contentTopic"`
	Version int `json:"version"`
	Timestamp float64 `json:"timestamp"`
}

type StoreResponse struct {
	Messages []WakuMessage  `json:"messages"`
}

type ContentFilter struct {
	ContentTopic string `json:"contentTopic"`
}

func getWakuDebugInfo(client *rpc.Client) WakuInfo {
	var wakuInfo WakuInfo

	if err := client.Call(&wakuInfo, "get_waku_v2_debug_v1_info"); err != nil {
		panic(err)
	}

	return wakuInfo
}

func getWakuStoreMessages(client *rpc.Client, contentTopic string) StoreResponse {
	var storeResponse StoreResponse
	var contentFilter = ContentFilter{contentTopic}
	var contentFilters []ContentFilter

	contentFilters = append(contentFilters, contentFilter)
	if err := client.Call(&storeResponse, "get_waku_v2_store_v1_messages", "", contentFilters); err != nil {
		panic(err)
	}

	return storeResponse

}

// TODO Publish
// TODO Subscribe

func main() {
	fmt.Println("JSON RPC request...")

	_, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Assumes node started
	client, _ := rpc.Dial("http://127.0.0.1:8545")

	var wakuInfo = getWakuDebugInfo(client)
	fmt.Println("WakuInfo ListenStr", wakuInfo.ListenStr)

	// TODO More args
	var contentTopic = "/toy-chat/2/huilong/proto"
	var storeResponse = getWakuStoreMessages(client, contentTopic)
	fmt.Println("Fetched", len(storeResponse.Messages), "messages")
}

