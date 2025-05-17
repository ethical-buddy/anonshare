package network

import (
    "log"
    "net"
    "strings"
    "time"
)

// Discovery announces the node on LAN for other peers to discover.
func StartLANDiscovery(port string) {
    addr := "255.255.255.255:" + port
    conn, err := net.Dial("udp4", addr)
    if err != nil {
        log.Println("Discovery broadcast failed:", err)
        return
    }
    defer conn.Close()

    ticker := time.NewTicker(5 * time.Second)
    for range ticker.C {
        msg := "anonshare:hello"
        conn.Write([]byte(msg))
    }
}

// ListenLANDiscovery listens for broadcast announcements from other peers.
func ListenLANDiscovery(port string) {
    addr, _ := net.ResolveUDPAddr("udp4", ":"+port)
    conn, err := net.ListenUDP("udp4", addr)
    if err != nil {
        log.Println("Discovery listener failed:", err)
        return
    }
    defer conn.Close()

    buf := make([]byte, 1024)
    for {
        n, remote, err := conn.ReadFromUDP(buf)
        if err == nil {
            msg := string(buf[:n])
            if strings.HasPrefix(msg, "anonshare:hello") {
                log.Printf("[DISCOVERY] Peer announced from %s", remote.IP.String())
            }
        }
    }
}

