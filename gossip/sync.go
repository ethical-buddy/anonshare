package gossip

import (
    "context"
    "encoding/json"
    "log"
    "time"

    "anonshare/metadata"
)

// SyncEngine pulls new metadata from the gossip channel and updates DB.
type SyncEngine struct {
    Gossip  *GossipNode
    DB      *metadata.MetadataDB
    Context context.Context
}

// NewSyncEngine creates and links the sync engine.
func NewSyncEngine(ctx context.Context, g *GossipNode, db *metadata.MetadataDB) *SyncEngine {
    return &SyncEngine{
        Gossip:  g,
        DB:      db,
        Context: ctx,
    }
}

// Start begins the sync loop, listening for new messages and inserting into DB.
func (s *SyncEngine) Start() {
    go func() {
        for {
            select {
            case <-s.Context.Done():
                return
            case meta := <-s.Gossip.MsgChan:
                if err := s.DB.InsertOrUpdate(meta); err != nil {
                    log.Printf("[SYNC] Failed to insert metadata: %v", err)
                } else {
                    log.Printf("[SYNC] Synced metadata: %s", meta.Summary())
                }
            }
        }
    }()
}

// BroadcastAll re-publishes known metadata (optional bootstrapping).
func (s *SyncEngine) BroadcastAll() error {
    metas, err := s.DB.GetAll()
    if err != nil {
        return err
    }

    for _, m := range metas {
        if err := s.Gossip.PublishMetadata(m); err != nil {
            log.Printf("[SYNC] Failed to publish %s: %v", m.ID, err)
        } else {
            time.Sleep(100 * time.Millisecond) // throttle
        }
    }
    return nil
}

// ExportAll serializes DB entries to JSON for inspection/debug.
func (s *SyncEngine) ExportAll() ([]byte, error) {
    metas, err := s.DB.GetAll()
    if err != nil {
        return nil, err
    }
    return json.MarshalIndent(metas, "", "  ")
}

