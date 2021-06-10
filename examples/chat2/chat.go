package main

import (
	"chat2/pb"
	"context"
	"crypto/sha256"
	"encoding/hex"
	"time"

	"github.com/ethereum/go-ethereum/rpc"
	"github.com/golang/protobuf/proto"
	"github.com/libp2p/go-libp2p-core/peer"
	//"github.com/status-im/go-waku/waku/v2/node"
	//wpb "github.com/status-im/go-waku/waku/v2/protocol/pb"
	"golang.org/x/crypto/pbkdf2"

	"github.com/status-im/go-nwaku/nwaku"
)

// Chat represents a subscription to a single PubSub topic. Messages
// can be published to the topic with Chat.Publish, and received
// messages are pushed to the Messages channel.
type Chat struct {
	// Messages is a channel of messages received from other peers in the chat room
	Messages chan *pb.Chat2Message

	// TODO Replace this
	//sub  *node.Subscription
	pubsubTopic string
	// TODO Replace with wrapper
	//node *node.WakuNode
	client *rpc.Client

	self         peer.ID
	contentTopic string
	useV1Payload bool
	nick         string
}

// NewChat tries to subscribe to the PubSub topic for the room name, returning
// a ChatRoom on success.
func NewChat(client *rpc.Client, selfID peer.ID, contentTopic string, useV1Payload bool, nickname string) (*Chat, error) {
	var defaultTopic = "/waku/2/default-waku/proto"
	// join the default waku topic and subscribe to it
	_, err := nwaku.PostWakuRelaySubscriptions(client, []string{defaultTopic})
	if err != nil {
		return nil, err
	}

	c := &Chat{
		//node:         n,
		client:         client,
		// XXX Not used directly anymore
		//sub:          sub,
		pubsubTopic:  defaultTopic,
		self:         selfID,
		contentTopic: contentTopic,
		nick:         nickname,
		useV1Payload: useV1Payload,
		Messages:     make(chan *pb.Chat2Message, 1024),
	}

	// start reading messages from the subscription in a loop
	go c.readLoop()

	return c, nil
}

func generateSymKey(password string) []byte {
	// AesKeyLength represents the length (in bytes) of an private key
	AESKeyLength := 256 / 8
	return pbkdf2.Key([]byte(password), nil, 65356, AESKeyLength, sha256.New)
}

// Publish sends a message to the pubsub topic.
func (cr *Chat) Publish(ctx context.Context, message string) error {

	msg := &pb.Chat2Message{
		Timestamp: uint64(time.Now().Unix()),
		Nick:      cr.nick,
		Payload:   []byte(message),
	}

	msgBytes, err := proto.Marshal(msg)
	if err != nil {
		return err
	}

	//log.Println("Publish", msg)

	//var version uint32
	// TODO Add support
	//var timestamp float64 = float64(time.Now().UnixNano())
	// var keyInfo *node.KeyInfo = &node.KeyInfo{}

	// if cr.useV1Payload { // Use WakuV1 encryption
	// 	keyInfo.Kind = node.Symmetric
	// 	keyInfo.SymKey = generateSymKey(cr.contentTopic)
	// 	version = 1
	// } else {
	// 	keyInfo.Kind = node.None
	// 	version = 0
	// }

	// TODO Implement, see what makes sense here vs private API
	// p := new(node.Payload)
	// p.Data = msgBytes
	// p.Key = keyInfo

	// // XXX Is this right?
	// payload, err := p.Encode(version)
	// if err != nil {
	// 	return err
	// }

	// For version 0, should get payload.Data, []byte
	// TODO Want it hex encoded though
	var payload = msgBytes

	// wakuMsg := &wpb.WakuMessage{
	// 	Payload:      payload,
	// 	Version:      version,
	// 	ContentTopic: cr.contentTopic,
	// 	Timestamp:    timestamp,
	// }

	var hexEncoded = make([]byte, hex.EncodedLen(len(payload)))
	hex.Encode(hexEncoded, payload)
	//fmt.Println("%s\n", hexEncoded)

	// TODO Replace with jSON RPC
	//_, err = cr.node.Publish(ctx, wakuMsg, nil)
	// NOTE version field support https://rfc.vac.dev/spec/16/#wakurelaymessage
	var wakuMsg = nwaku.WakuRelayMessage{
		Payload: string(hexEncoded), // "0x1a2b3c4d5e6f",
		ContentTopic: cr.contentTopic,
		//Timestamp: timestamp,
	}
	// TODO Error handling
	var _, _ = nwaku.PostWakuRelayMessage(cr.client, wakuMsg)

	return nil
}

func (cr *Chat) decodeMessage(wakumsg nwaku.WakuMessage) {
	// TODO Re-enable
	// var keyInfo *node.KeyInfo = &node.KeyInfo{}
	// if cr.useV1Payload { // Use WakuV1 encryption
	// 	keyInfo.Kind = node.Symmetric
	// 	keyInfo.SymKey = generateSymKey(cr.contentTopic)
	// } else {
	// 	keyInfo.Kind = node.None
	// }

	var payload = wakumsg.Payload

	msg := &pb.Chat2Message{}
	if err := proto.Unmarshal(payload, msg); err != nil {
		return
	}

	// send valid messages onto the Messages channel
	cr.Messages <- msg
}

// readLoop pulls messages from the pubsub topic and pushes them onto the Messages channel.
// TODO Improve polling with channels etc here
// XXX This also means that we don't see our message straight away currently
func (cr *Chat) readLoop() {
	for {
		var wakuMessages, _ = nwaku.GetWakuRelayMessages(cr.client, cr.pubsubTopic)
		// for value := range cr.sub.C {
		for _, msg := range wakuMessages {
			cr.decodeMessage(msg)
		}

		time.Sleep(2 * time.Second)
	}
}

func (cr *Chat) displayMessages(messages []nwaku.WakuMessage) {
	for _, msg := range messages {
		cr.decodeMessage(msg)
	}
}
