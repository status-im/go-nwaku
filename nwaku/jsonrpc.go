package main

import (
	"context"
	"fmt"
	"time"
	"github.com/ethereum/go-ethereum/rpc"
)

// NOTE Can be generalized with different transports (HTTP, IPC, WS etc)
// https://github.com/ethereum/go-ethereum/blob/master/rpc/client.go#L169

// TODO Move to types
type WakuInfo struct {
	ListenStr string `json:"listenStr"`
}

type HistoricalMessageResponse struct {
	Messages string  `json:"messages"`
}

type ContentFilter struct {
	ContentTopic string `json:"contentTopic"`
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

	// TODO Lets do query of messages
	//
	//curl -d '{"jsonrpc":"2.0","id":"id","method":"get_waku_v2_store_v1_messages", "params":["", [{"contentTopic":"/waku/2/default-content/proto"}]]}' --header "Content-Type: application/json" http://localhost:8545
	var messageResponse HistoricalMessageResponse
	var contentFilter = ContentFilter{"/toy-chat/2/huilong/proto"}
	var contentFilters []ContentFilter
	contentFilters = append(contentFilters, contentFilter)
	var arg1 = ""
	var arg2 = contentFilters
	if err := client.Call(&messageResponse, "get_waku_v2_store_v1_messages", arg1, arg2); err != nil {
		panic(err)
	}
	fmt.Println("Resp", messageResponse)
}

// ./bin/wakunode2 --storenode=/ip4/188.166.135.145/tcp/30303/p2p/16Uiu2HAmL5okWopX7NqZWBUKVqW8iUxCEmd5GMHLVPwCgzYzQv3e
// curl -d '{"jsonrpc":"2.0","id":"id","method":"get_waku_v2_store_v1_messages", "params":["", [{"contentTopic":"/toy-chat/2/huilong/proto"}]]}' --header "Content-Type: application/json" http://localhost:8545

// Later: Generalize with CallContext (not prio)
