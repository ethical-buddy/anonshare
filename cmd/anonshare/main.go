package main

import (
    "context"
    "fmt"
    "os"
    "os/signal"
    "syscall"

    "anonshare/config"
    "anonshare/core"
    "anonshare/frontend/tui"
    "anonshare/gossip"
    "anonshare/metadata"
    "anonshare/node"
)

func main() {
    cfg := config.Load()

    logger := core.InitLogger(cfg.LogLevel)

    identityPath := node.IdentityPath(cfg.DataDir)
    ident, err := node.LoadIdentityOrCreate(identityPath)
    if err != nil {
        logger.Fatalf("Failed to load identity: %v", err)
    }

    logger.Infof("Starting AnonShare as %s", ident.String())

    db, err := metadata.NewMetadataDB(cfg.MetadataDB)
    if err != nil {
        logger.Fatalf("Metadata DB error: %v", err)
    }

    gossiper := gossip.NewGossipNode(ident, cfg.GossipPort)
    go gossiper.Start()

    ctx, cancel := context.WithCancel(context.Background())
    syncer := gossip.NewSyncEngine(ctx, gossiper, db)
    syncer.Start()

    // Graceful shutdown
    sig := make(chan os.Signal, 1)
    signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
    go func() {
        <-sig
        logger.Info("Shutting down...")
        cancel()
        os.Exit(0)
    }()

    tui.StartTUI()
}

