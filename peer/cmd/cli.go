package cli

import (
	"bufio"
	"fmt"
	"os"

	// "path/filepath"
	"strings"

	"github.com/muskiteer/anonshare/internal"
	"github.com/muskiteer/anonshare/peer/functions"
)

func Cli() {
	peerPort := os.Getenv("PEER_PORT")
    reader := bufio.NewReader(os.Stdin)
    
    for {
        showMenu()
        choice, _ := reader.ReadString('\n')
        choice = strings.TrimSpace(choice)
        
        switch choice {
        case "1":
            handleUpload(reader, peerPort)
        case "2":
            handleDownload(reader)
        case "3":
            handleListFiles()
        case "4":
            fmt.Println("👋 Goodbye!")
            return
        default:
            fmt.Println("❌ Invalid choice. Please select 1, 2, 3, or 4.")
        }
        
        fmt.Println() // Add spacing between operations
    }
}

func showMenu() {
    fmt.Println("╔══════════════════════════════════════╗")
    fmt.Println("║           ANONSHARE P2P CLI          ║")
    fmt.Println("╠══════════════════════════════════════╣")
    fmt.Println("║  1. 📤 Upload File                   ║")
    fmt.Println("║  2. 📥 Download File                 ║")
    fmt.Println("║  3. 📋 List All Files                ║")
    fmt.Println("║  4. 🚪 Exit                          ║")
    fmt.Println("╚══════════════════════════════════════╝")
    fmt.Print("👉 Please choose an option (1-4): ")
}

func handleUpload(reader *bufio.Reader, peerPort string) {
    fmt.Println("\n📤 UPLOAD FILE")
    fmt.Println("═══════════════")
    
    // Get file path
    fmt.Print("📁 Enter file path: ")
    filePath, _ := reader.ReadString('\n')
    filePath = strings.TrimSpace(filePath)
    
    if filePath == "" {
        fmt.Println("❌ File path cannot be empty!")
        return
    }
    
    // Convert to absolute path
    
    
    // Check if file exists
    absolutePath, err := internal.GetAbsolutePath(filePath)
	if err != nil {
		fmt.Printf("❌ Error getting absolute path: %v\n", err)
		return
	}

    // Get description
    fmt.Print("📝 Enter file description: ")
    description, _ := reader.ReadString('\n')
    description = strings.TrimSpace(description)
    
    if description == "" {
        description = "No description provided"
    }
    
    // Upload the file
    fmt.Println("⏳ Uploading file...")
    functions.Uploading_the_file(absolutePath, description, peerPort)
}

func handleDownload(reader *bufio.Reader) {
    fmt.Println("\n📥 DOWNLOAD FILE")
    fmt.Println("════════════════")
    
    // Get file hash
    fmt.Print("🔍 Enter file hash: ")
    hash, _ := reader.ReadString('\n')
    hash = strings.TrimSpace(hash)
    
    if hash == "" {
        fmt.Println("❌ File hash cannot be empty!")
        return
    }
    
    // Download the file
    fmt.Printf("⏳ Searching for peers with hash: %s\n", hash)
    functions.File_download(hash)
}

func handleListFiles() {
    fmt.Println("\n📋 LIST ALL FILES")
    fmt.Println("═════════════════")
    fmt.Println("⏳ Fetching files from backend...")
    functions.Gettings_files_info()
}