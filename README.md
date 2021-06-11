# go-nwaku

Go wrapper for running nim-waku as a subprocess.

## Rationale

1. Provide a friendly interface to use Waku for Go environments.
2. More wood behind fewer arrows; promote code reuse.

## Direction

For similar projects, see: https://github.com/ethereum/py-geth

JSON RPC spec that nim-waku exposes: https://rfc.vac.dev/spec/16/

## Running

- (Temp) Ensure you have `wakunode2` nim-waku in the `bin` directory

## API calls used by chat2

- DialPeer
- AddStorePeer
x Query
x Subscribe
x Publish
- ListPeers

## Caveats

Assumes we can spawn a child process. In some environments, such as on iOS, this
may not be permitted.
