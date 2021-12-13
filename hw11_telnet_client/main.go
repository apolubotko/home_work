package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"time"
)

const (
	defaulTimeout = time.Second * 10
	// msgEOF        = "...EOF"
	// msgConClosed  = "...Connection was closed by peer".
)

func main() {
	var timeout time.Duration
	flag.DurationVar(&timeout, "timeout", defaulTimeout, "Timeout to connect to server")
	flag.Parse()

	if len(flag.Args()) != 2 {
		log.Fatal("not enought args")
	}

	host := flag.Arg(0)
	port := flag.Arg(1)

	client := NewTelnetClient(net.JoinHostPort(host, port), timeout, os.Stdin, os.Stdout)
	if err := client.Connect(); err != nil {
		log.Fatal("Can't connect to host")
	}
	defer client.Close()

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	go func() error {
		err := client.Send()
		if err != nil {
			return fmt.Errorf("some send errors: %w", err)
		}
		cancel()

		return nil
	}()

	go func() error {
		err := client.Receive()
		if err != nil {
			return fmt.Errorf("some receive errors: %w", err)
		}
		cancel()

		return nil
	}()

	<-ctx.Done()
}
