# go-ssh-cli
Sample repository for a tool written in Go to execute interactive SSH commands.

I've recently started my journey learning the Go language, and I hope thi s repo can help others too to get more familiar with the nuances of the language.

This repo will collect information found on the web in the Golang package documentation and other resources found on the web. 

## Problem statement

The code in this repository should help execute commands on remote servers using SSH.

Such commands can be interactive and in some cases could require the input of sensitive information.

A typical example of this scenario is the execution of `ssh-keygen`


### Features
The list of features I aim to use in this project are the following:

- [X] Login using keys
- [ ] Interactive commands
   - [ ] identify sensitive information to mask the input

## Resources needed

I'm going to use a Raspberry Pi as target for the SSH connection. However, there are several free shell providers that one could use to have a valid target.


--- 
[Next steps](docs/first_code.md)