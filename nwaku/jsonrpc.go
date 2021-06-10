package nwaku

import (
	"github.com/ethereum/go-ethereum/rpc"
)

func GetWakuDebugInfo(client *rpc.Client) WakuInfo {
	var wakuInfo WakuInfo

	if err := client.Call(&wakuInfo, "get_waku_v2_debug_v1_info"); err != nil {
		panic(err)
	}

	return wakuInfo
}

func GetWakuStoreMessages(client *rpc.Client, contentTopic string) (StoreResponse, error) {
	var storeResponse StoreResponse
	var contentFilter = ContentFilter{contentTopic}
	var contentFilters []ContentFilter

	contentFilters = append(contentFilters, contentFilter)
	if err := client.Call(&storeResponse, "get_waku_v2_store_v1_messages", "", contentFilters); err != nil {
		return storeResponse, err
	}

	return storeResponse, nil

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
// For now, just use query and publish

// General things that can be improved:
// - Generalized with different transports (HTTP, IPC, WS etc), see
// https://github.com/ethereum/go-ethereum/blob/master/rpc/client.go#L169
// - Generalize with CallContext
// - Exposing higher level methods as API
// - Consider using methods scoped to rpc.Client instead
// - Support more args in store rpc call
// - Uniform error handling
