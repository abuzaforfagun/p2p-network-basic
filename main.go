package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/libp2p/go-libp2p"
	"github.com/libp2p/go-libp2p/core/host"
	"github.com/libp2p/go-libp2p/core/network"
	"github.com/libp2p/go-libp2p/core/peer"
	"github.com/libp2p/go-libp2p/p2p/protocol/ping"
	"github.com/multiformats/go-multiaddr"
)

func main() {
	host, err := createHost()
	if err != nil {
		log.Fatal("Error creating host:", err)
	}

	printNodeInfo(host)

	setUpHandlers(host)

	if len(os.Args) > 1 {
		address := os.Args[1]
		peerInfo, err := connectPeer(host, address)
		if err != nil {
			log.Fatal("Error connecting or sending message:", err)
		}

		sendMessage(host, peerInfo)
	}

	// Keep the application running
	select {}
}

func createHost() (host.Host, error) {
	host, err := libp2p.New(
		libp2p.ListenAddrStrings("/ip4/127.0.0.1/tcp/0"),
	)
	if err != nil {
		return nil, err
	}
	return host, nil
}

func printNodeInfo(host host.Host) {
	peerInfo := peer.AddrInfo{
		ID:    host.ID(),
		Addrs: host.Addrs(),
	}
	addrs, err := peer.AddrInfoToP2pAddrs(&peerInfo)
	if err != nil {
		log.Fatal("Error getting P2P addresses:", err)
	}

	fmt.Println("Node ID:", host.ID())
	fmt.Println("Listening on:", addrs[0])
}

func setUpHandlers(host host.Host) {
	host.SetStreamHandler(ping.ID, func(s network.Stream) {
		handleStream(s)
	})

	host.Network().Notify(&network.NotifyBundle{
		ConnectedF: func(n network.Network, c network.Conn) {
			fmt.Println("New peer connected:", c.RemotePeer())
		},
		DisconnectedF: func(n network.Network, c network.Conn) {
			fmt.Println("Peer disconnected:", c.RemotePeer())
		},
	})
}

func handleStream(s network.Stream) {
	fmt.Println("Got a new stream!")
	defer s.Close()

	buf := make([]byte, 1024)
	n, err := s.Read(buf)
	if err != nil && err != io.EOF {
		log.Println("Error reading from stream:", err)
		return
	}
	fmt.Println("Message received:", string(buf[:n]))
}

func connectPeer(host host.Host, address string) (*peer.AddrInfo, error) {
	addr, err := multiaddr.NewMultiaddr(address)
	if err != nil {
		return nil, fmt.Errorf("invalid multiaddress: %w", err)
	}

	peerInfo, err := peer.AddrInfoFromP2pAddr(addr)
	if err != nil {
		return nil, fmt.Errorf("failed to get peer info: %w", err)
	}

	if err := host.Connect(context.Background(), *peerInfo); err != nil {
		return nil, fmt.Errorf("failed to connect to peer: %w", err)
	}

	return peerInfo, nil
}

func sendMessage(host host.Host, peerInfo *peer.AddrInfo) error {
	stream, err := host.NewStream(context.Background(), peerInfo.ID, ping.ID)
	if err != nil {
		return fmt.Errorf("failed to create new stream: %w", err)
	}
	defer stream.Close()

	message := "Hello from sender!"
	if _, err := stream.Write([]byte(message)); err != nil {
		return fmt.Errorf("failed to send message: %w", err)
	}
	fmt.Println("Message sent:", message)
	return nil
}
