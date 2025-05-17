package metadata

import (
    "crypto/sha256"
    "encoding/hex"
    "fmt"
    "os"
    "path/filepath"
    "time"
)

// Metadata represents one shared resource in the network.
type Metadata struct {
    ID        string    `json:"id"`         // Deterministic hash of file
    Name      string    `json:"name"`       // Filename
    Size      int64     `json:"size"`       // Bytes
    Type      string    `json:"type"`       // MIME type or extension
    Uploader  string    `json:"uploader"`   // PeerID
    Timestamp time.Time `json:"timestamp"`  // Obfuscated timestamp (optionally fuzzed)
    Tags      []string  `json:"tags"`       // Optional categories or keywords
    Hash      string    `json:"hash"`       // SHA-256 of file contents
    Path      string    `json:"-"`          // Local path (not broadcast)
}

// NewFromFile generates metadata from a file path.
func NewFromFile(path string, uploader string, tags []string) (*Metadata, error) {
    file, err := os.Open(path)
    if err != nil {
        return nil, err
    }
    defer file.Close()

    stat, err := file.Stat()
    if err != nil {
        return nil, err
    }

    hashBytes := sha256.New()
    if _, err := hashBytes.Write([]byte(uploader)); err != nil {
        return nil, err
    }

    contentHash := sha256.New()
    if _, err := contentHash.ReadFrom(file); err != nil {
        return nil, err
    }
    sum := contentHash.Sum(nil)

    meta := &Metadata{
        Name:      filepath.Base(path),
        Size:      stat.Size(),
        Type:      detectFileType(path),
        Uploader:  uploader,
        Timestamp: time.Now(),
        Tags:      tags,
        Hash:      hex.EncodeToString(sum),
        Path:      path,
    }

    meta.ID = meta.ComputeID()
    return meta, nil
}

// ComputeID returns a unique ID derived from hash + uploader.
func (m *Metadata) ComputeID() string {
    sum := sha256.Sum256([]byte(m.Hash + m.Uploader))
    return hex.EncodeToString(sum[:])
}

// detectFileType guesses file type from extension.
func detectFileType(path string) string {
    ext := filepath.Ext(path)
    switch ext {
    case ".jpg", ".jpeg", ".png", ".gif":
        return "image"
    case ".pdf":
        return "document"
    case ".mp4", ".mkv":
        return "video"
    case ".mp3":
        return "audio"
    default:
        return "file"
    }
}

// Summary returns a human-readable description.
func (m *Metadata) Summary() string {
    return fmt.Sprintf("[%s] %s (%d bytes) from %s", m.Type, m.Name, m.Size, m.Uploader)
}

