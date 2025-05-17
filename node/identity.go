package node

import (
    "crypto/ed25519"
    "crypto/rand"
    "encoding/base64"
    "errors"
    "fmt"
    "os"
    "path/filepath"
)

// Identity represents the ephemeral identity of a node.
type Identity struct {
    Private ed25519.PrivateKey
    Public  ed25519.PublicKey
    ID      string // base64-encoded public key
}

// GenerateIdentity creates a new random Ed25519 keypair.
func GenerateIdentity() (*Identity, error) {
    pub, priv, err := ed25519.GenerateKey(rand.Reader)
    if err != nil {
        return nil, err
    }

    id := base64.StdEncoding.EncodeToString(pub)

    return &Identity{
        Private: priv,
        Public:  pub,
        ID:      id,
    }, nil
}

// Save writes the identity to a file.
func (id *Identity) Save(path string) error {
    if id == nil {
        return errors.New("identity is nil")
    }

    data := base64.StdEncoding.EncodeToString(id.Private)
    return os.WriteFile(path, []byte(data), 0600)
}

// LoadIdentity reads an identity from a file or creates a new one.
func LoadIdentityOrCreate(path string) (*Identity, error) {
    if _, err := os.Stat(path); os.IsNotExist(err) {
        id, err := GenerateIdentity()
        if err != nil {
            return nil, err
        }
        if err := id.Save(path); err != nil {
            return nil, err
        }
        return id, nil
    }

    data, err := os.ReadFile(path)
    if err != nil {
        return nil, err
    }

    privBytes, err := base64.StdEncoding.DecodeString(string(data))
    if err != nil {
        return nil, err
    }

    priv := ed25519.PrivateKey(privBytes)
    pub := priv.Public().(ed25519.PublicKey)
    idStr := base64.StdEncoding.EncodeToString(pub)

    return &Identity{
        Private: priv,
        Public:  pub,
        ID:      idStr,
    }, nil
}

// IdentityPath returns a suggested path for identity key file.
func IdentityPath(baseDir string) string {
    return filepath.Join(baseDir, "identity.key")
}

// String prints a short version of identity for logs.
func (id *Identity) String() string {
    return fmt.Sprintf("anonshare://%s", id.ID[:12])
}

