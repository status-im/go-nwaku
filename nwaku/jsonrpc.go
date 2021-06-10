package nwaku

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
	// TODO Should be hex encoded string here
	Payload []byte `json:"payload"`
	ContentTopic string `json:"contentTopic"`
	Version int `json:"version"`
	Timestamp float64 `json:"timestamp"`
}

type WakuRelayMessage struct {
	Payload string `json:"payload"`
	ContentTopic string `json:"contentTopic"`
	//	Version int `json:"version"`
	//	Timestamp float64 `json:"timestamp"`
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

// TODO Support more args
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

func postWakuRelayMessage(client *rpc.Client, message WakuRelayMessage) bool {
	var topic = "/waku/2/default-waku/proto"
	var res bool

	if err := client.Call(&res, "post_waku_v2_relay_v1_message", topic, message); err != nil {
		panic(err)
	}

	return res
}

// TODO Subscribe, then poll for getting messages
// https://rfc.vac.dev/spec/16/#post_waku_v2_relay_v1_subscriptions
// https://rfc.vac.dev/spec/16/#get_waku_v2_relay_v1_messages
// For now, just do query and publish

func main() {
	fmt.Println("JSON RPC request...")

	_, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Assumes node started
	client, _ := rpc.Dial("http://127.0.0.1:8545")

	// Get node info
	var wakuInfo = getWakuDebugInfo(client)
	fmt.Println("WakuInfo ListenStr", wakuInfo.ListenStr)

	// Query messages
	var contentTopic = "/toy-chat/2/huilong/proto"
	var storeResponse = getWakuStoreMessages(client, contentTopic)
	fmt.Println("Fetched", len(storeResponse.Messages), "messages")

	// Publish
	var message = WakuRelayMessage{Payload: "0x1a2b3c4d5e6f", ContentTopic: contentTopic}
	var res = postWakuRelayMessage(client, message)
	fmt.Println("Publish", res)
}
