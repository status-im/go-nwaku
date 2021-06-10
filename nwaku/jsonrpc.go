package nwaku

import (
	"github.com/ethereum/go-ethereum/rpc"
)

// NOTE Can be generalized with different transports (HTTP, IPC, WS etc)
// https://github.com/ethereum/go-ethereum/blob/master/rpc/client.go#L169

// Later: Generalize with CallContext (not prio)

// NOTE Exposing these methods publicly directly, might be wrapped for higher
// level a la Node API

func GetWakuDebugInfo(client *rpc.Client) WakuInfo {
	var wakuInfo WakuInfo

	if err := client.Call(&wakuInfo, "get_waku_v2_debug_v1_info"); err != nil {
		panic(err)
	}

	return wakuInfo
}

// TODO Support more args
func GetWakuStoreMessages(client *rpc.Client, contentTopic string) StoreResponse {
	var storeResponse StoreResponse
	var contentFilter = ContentFilter{contentTopic}
	var contentFilters []ContentFilter

	contentFilters = append(contentFilters, contentFilter)
	if err := client.Call(&storeResponse, "get_waku_v2_store_v1_messages", "", contentFilters); err != nil {
		panic(err)
	}

	return storeResponse

}

func PostWakuRelayMessage(client *rpc.Client, message WakuRelayMessage) bool {
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
