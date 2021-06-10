package nwaku

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
