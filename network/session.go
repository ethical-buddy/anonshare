package network

import (
    "crypto/rand"
    "fmt"
)

// GenerateSessionKey returns a random 32-byte key (placeholder).
func GenerateSessionKey() ([]byte, error) {
    key := make([]byte, 32)
    _, err := rand.Read(key)
    if err != nil {
        return nil, err
    }
    return key, nil
}

// StartSession is a stub for future secure channel setup.
func StartSession(peer string) {
    fmt.Println("[SESSION] Starting session with", peer)
}

