# First steps

The program will support arguments to retrieve the information about the target host and some other connection details
as

`command.exe hostname`

## Connecting with SSH

One of the hardest part (implementing the SSH protocol) is already available in Go as
a [package](https://pkg.go.dev/golang.org/x/crypto/ssh).

`go get golang.org/x/crypto/ssh`

This package provides the relevant interfaces and functions to initialize, and interact with an SSH connections.

The first [_feature_](../README.md#features) for this repo is implemented in `udd_ssh/sample_client.go`. It initializes the connection it supports initiating a connection with a key instead of a password.