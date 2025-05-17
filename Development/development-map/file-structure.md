anonshare/
├── cmd/
│   └── anonshare/           # Main entry point (main.go)
├── config/
│   └── config.go            # App configuration structs & loading
├── core/
│   ├── app.go               # App lifecycle, signal handling
│   └── logger.go            # Logging setup
├── node/
│   ├── identity.go          # Node ephemeral key management
│   └── node.go              # Node boot, network lifecycle
├── metadata/
│   ├── metadata.go          # Structs, metadata creation/parse
│   └── db.go                # SQLite local DB handler
├── gossip/
│   ├── gossip.go            # Metadata gossip protocol
│   └── sync.go              # Sync, merge, push-pull logic
├── network/
│   ├── discovery.go         # DHT or LAN discovery
│   ├── transport.go         # Connection logic (TCP/QUIC/WebRTC)
│   └── session.go           # Secure session establishment
├── transfer/
│   ├── send.go              # File sending
│   └── receive.go           # File receiving
├── privacy/
│   ├── crypto.go            # Encryption utils (NaCl or libsodium)
│   └── obfuscation.go       # Timestamp + metadata fuzzing
├── frontend/
│   ├── tui/                 # TUI frontend (tview)
│   │   ├── ui.go
│   │   └── pages/
│   └── gui/                 # (Optional) GUI frontend (fyne/webview)
├── assets/                  # Thumbnails, default icons
├── internal/                # Internal utils
│   └── utils.go
├── go.mod
└── README.md

