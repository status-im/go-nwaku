# Using the `chat2` application

## Background

The `chat2` application is a basic command-line chat app using the [Waku v2 suite of protocols](https://specs.vac.dev/specs/waku/v2/waku-v2). It connects to a [fleet of test nodes](fleets.status.im) to provide end-to-end p2p chat capabilities. The Waku team is currently using this application for internal testing. If you want try our protocols, or join the dogfooding fun, follow the instructions below.

## Preparation
```
make
```

## Basic application usage

To start the `chat2` application in its most basic form, run the following from the project directory

```
./build/chat2
```

The app will randomly select and connect to a peer from the test fleet.

```
No static peers configured. Choosing one at random from test fleet...
```

Wait for the chat prompt (`>`) and chat away!

## Retrieving historical messages

TODO

## Specifying a static peer

In order to connect to a *specific* node as [`relay`](https://specs.vac.dev/specs/waku/v2/waku-relay) peer, define that node's `multiaddr` as a `staticnode` when starting the app:

```
./build/chat2 -staticnode=/ip4/134.209.139.210/tcp/30303/p2p/16Uiu2HAmPLe7Mzm8TsYUubgCAW1aJoeFScxrLj8ppHFivPo97bUZ
```

This will bypass the random peer selection process and connect to the specified node.

## In-chat options

| Command | Effect |
| --- | --- |
| `/help` | displays available in-chat commands |
| `/connect` | interactively connect to a new peer |
| `/nick` | change nickname for current chat session |
| `/peers` | Display the list of connected peers |

## `chat2` message protobuf format

Each `chat2` message is encoded as follows

```protobuf
message Chat2Message {
  uint64 timestamp = 1;
  string nick = 2;
  bytes payload = 3;
}
```

where `timestamp` is the Unix timestamp of the message, `nick` is the relevant `chat2` user's selected nickname and `payload` is the actual chat message being sent. The `payload` is the byte array representation of a UTF8 encoded string.