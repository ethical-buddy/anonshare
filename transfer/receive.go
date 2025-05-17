package transfer

import (
    "fmt"

    "anonshare/network"
)

// StartReceiver runs the TCP server to accept incoming file transfers.
func StartReceiver(port string) {
    fmt.Println("[RECEIVE] Ready to receive files")
    network.StartTCPServer(port)
}

