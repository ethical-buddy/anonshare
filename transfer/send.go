package transfer

import (
    "fmt"
    "io"
    "net"
    "os"
    "path/filepath"
)

// SendFile sends a file to the given peer (ip:port).
func SendFile(filePath string, peerAddr string) error {
    conn, err := net.Dial("tcp", peerAddr)
    if err != nil {
        return fmt.Errorf("failed to connect: %w", err)
    }
    defer conn.Close()

    file, err := os.Open(filePath)
    if err != nil {
        return fmt.Errorf("failed to open file: %w", err)
    }
    defer file.Close()

    _, fileName := filepath.Split(filePath)
    fmt.Fprintf(conn, "%s\n", fileName) // optional: send name (ignored on receiver currently)

    _, err = io.Copy(conn, file)
    if err != nil {
        return fmt.Errorf("failed to send file: %w", err)
    }

    fmt.Println("[SEND] File sent:", filePath)
    return nil
}

