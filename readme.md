# libp2p Example

This is a simple Go application that demonstrates how to use the libp2p library to create a peer-to-peer network. The application initializes a libp2p host, sets up handlers for incoming streams, connects to a specified peer, and sends a message to that peer.

## Features

- Peer Initialization: Create a new libp2p host that listens for incoming connections.
- Stream Handling: Handle incoming streams and print received messages.
- Peer Notification: Notify when peers connect or disconnect.
- Message Sending: Connect to a peer and send a message if an address is provided.

## Prerequisites

- Go 1.18 or higher

## Installation

- Clone the repository: `git clone https://github.com/abuzaforfagun/p2p-network-basic.git`
- `cd p2p-network-basic`
- `go mod tidy`

## Usage

1. Build the application, `go build -o libp2p-example` (Windows: `go build -o libp2p-example.exe`)
2. Run the application, `./libp2p-example` (Windows: `./libp2p-example.exe`)
3. Copy the listening address (Ex. `/ip4/127.0.0.1/tcp/53854/p2p/12D3KooWNjiwSApw2FfY3vVuLvtSFzvyAZHXHPasxqRm9tRLdS31`)
4. Connect another Peer: `./libp2p-example <listening-address>`
