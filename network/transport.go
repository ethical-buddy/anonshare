package network

import (
    "fmt"
    "io"
    "net"
    "os"
)

// StartTCPServer listens for incoming file transfers.
func StartTCPServer(port string) {
    ln, err := net.Listen("tcp", ":"+port)
    if err != nil {
        fmt.Println("TCP server error:", err)
        return
    }
    defer ln.Close()

    fmt.Println("[TCP] Listening on", port)

    for {
        conn, err := ln.Accept()
        if err != nil {
            continue
        }
        go handleConn(conn)
    }
}

func handleConn(conn net.Conn) {
    defer conn.Close()

    file, err := os.CreateTemp("", "received_*")
    if err != nil {
        fmt.Println("Create temp file failed:", err)
        return
    }
    defer file.Close()

    io.Copy(file, conn)
    fmt.Printf("[TCP] Received file: %s\n", file.Name())
}

