package udd_ssh

import (
	"fmt"
	"log"
	"os"

	"golang.org/x/crypto/ssh"
)

func NewClient(hostname string, user string, privateKey string) *ssh.Client {
	key, err := os.ReadFile(privateKey)
	if err != nil {
		log.Fatal(err)
	}
	signer, err := ssh.ParsePrivateKey(key)
	if err != nil {
		log.Fatal(err)
	}
	config := &ssh.ClientConfig{
		User: user,
		Auth: []ssh.AuthMethod{
			ssh.PublicKeys(signer),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}
	client, err := ssh.Dial("tcp", fmt.Sprintf("%s:%s", hostname, "22"), config)
	if err != nil {
		log.Fatal(err)
	}
	return client
}
