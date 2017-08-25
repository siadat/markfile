package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
)

type opts struct {
	port, root string
}

func (o *opts) String() string {
	var buf bytes.Buffer
	buf.WriteString(fmt.Sprintf("Listener  %s\n", o.ListenAddr()))
	buf.WriteString(fmt.Sprintf("Directory %s\n", o.Root()))
	return buf.String()
}

func (o *opts) ListenAddr() string {
	return ":" + o.port
}

func (o *opts) Root() string {
	return o.root
}

func parseOpts() *opts {
	flagPort := "8080"
	if len(os.Args) >= 2 {
		flagPort = os.Args[1]
	}

	flagDir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	return &opts{
		port: flagPort,
		root: flagDir,
	}
}
