package node

import (
    "anonshare/core/logger"
    "fmt"
)

func BootNode() {
    logger.LogInfo("Bootstrapping anonshare node...")
    fmt.Println("Node is live. Ready to participate.")
}

