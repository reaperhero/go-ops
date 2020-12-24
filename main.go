package main

import (
	"github.com/gliderlabs/ssh"
	"github.com/reaperhero/go-ops/config"
	"io"
	"log"
)

func init()  {
	config.InitViper()
}

func main() {
	ssh.Handle(func(s ssh.Session) {
		io.WriteString(s, "Hello world\n")
	})

	log.Fatal(ssh.ListenAndServe(":2222", nil))
}
