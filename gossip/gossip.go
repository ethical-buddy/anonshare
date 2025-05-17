package gossip

import (
    "context"
    "encoding/json"
    "fmt"
    "log"

    "github.com/libp2p/go-libp2p"
    "github.com/libp2p/go-libp2p-core/host"
    "github.com/libp2p/go-libp2p-core/peer"
    "github.com/libp2p/go-libp2p-core/peerstore"
    pubsub "github.com/libp2p/go-libp2p-pubsub"
    "github.com/multiformats/go-multiaddr"

    "anonshare/metadata"
)

const MetadataTopic = "anonshare-meta"

// GossipNode manages pubsub and metadata distribution.
type GossipNode struct {
    Host      host.Host
    PubSub    *pubsub.PubSub
    Sub       *pubsub.Subscription
    Topic     *pubsub.Topic
    SelfID    peer.ID
    MsgChan   chan metadata.Metadata
}

// NewGossipNode initializes libp2p, Gossipsub, and topic.
func NewGossipNode(ctx context.Context, listenPort int) (*GossipNode, error) {
    // Multiaddr
    listenAddr, _ := multiaddr.NewMultiaddr(fmt.Sprintf("/ip4/0.0.0.0/tcp/%d", listenPort))

    h, err := libp2p.New(
        libp2p.ListenAddrs(listenAddr),
    )
    if err != nil {
        return nil, err
    }

    // Init Gossipsub
    ps, err := pubsub.NewGossipSub(ctx, h)
    if err != nil {
        return nil, err
    }

    topic, err := ps.Join(MetadataTopic)
    if err != nil {
        return nil, err
    }

    sub, err := topic.Subscribe()
    if err != nil {
        return nil, err
    }

    node := &GossipNode{
        Host:    h,
        PubSub:  ps,
        Sub:     sub,
        Topic:   topic,
        SelfID:  h.ID(),
        MsgChan: make(chan metadata.Metadata, 32),
    }

    go node.readLoop(ctx)

    log.Printf("[GOSSIP] Node started with ID: %s", node.SelfID.Pretty())
    return node, nil
}

// readLoop listens to incoming gossip messages.
func (g *GossipNode) readLoop(ctx context.Context) {
    for {
        msg, err := g.Sub.Next(ctx)
        if err != nil {
            log.Printf("[GOSSIP] Sub error: %v", err)
            return
        }

        if msg.ReceivedFrom == g.SelfID {
            continue // Ignore own messages
        }

        var m metadata.Metadata
        if err := json.Unmarshal(msg.Data, &m); err != nil {
            log.Printf("[GOSSIP] Invalid metadata: %v", err)
            continue
        }

        log.Printf("[GOSSIP] Received metadata: %s", m.Name)
        g.MsgChan <- m
    }
}

// PublishMetadata broadcasts metadata to the network.
func (g *GossipNode) PublishMetadata(meta metadata.Metadata) error {
    data, err := json.Marshal(meta)
    if err != nil {
        return err
    }
    return g.Topic.Publish(context.Background(), data)
}

// AddPeer adds a peer address to the peerstore manually.
func (g *GossipNode) AddPeer(peerID peer.ID, addr multiaddr.Multiaddr) {
    g.Host.Peerstore().AddAddr(peerID, addr, peerstore.PermanentAddrTTL)
}

