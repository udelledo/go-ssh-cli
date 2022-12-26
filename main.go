package main

import (
	"fmt"
	"os"
	"path/filepath"

	"go-ssh-cli/udd_ssh"
)

func main() {
	hostname := os.Args[1]
	homeDir, _ := os.UserHomeDir()
	client := udd_ssh.NewClient(hostname, "root", filepath.Join(homeDir, ".ssh", "id_ed25519"))
	defer client.Close()
	fmt.Printf("%s\n", client.User())
}
