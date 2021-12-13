package main

import (
	"context"
	"fmt"
	"io"
	"net"
	"os"
	"time"
)

type TelnetClient interface {
	Connect() error
	io.Closer
	Send() error
	Receive() error
}

type Client struct {
	context.Context
	address string
	timeout time.Duration
	in      io.ReadCloser
	out     io.Writer
	connect net.Conn
}

func NewTelnetClient(address string, timeout time.Duration, in io.ReadCloser, out io.Writer) TelnetClient {
	client := Client{
		address: address,
		timeout: timeout,
		in:      in,
		out:     out,
	}

	return &client
}

func (c *Client) Connect() error {
	con, err := net.DialTimeout("tcp", c.address, c.timeout)
	if err != nil {
		return fmt.Errorf("connect: %w", err)
	}
	c.connect = con

	return nil
}

func (c *Client) Close() error {
	if c.connect != nil {
		if err := c.connect.Close(); err != nil {
			return fmt.Errorf("close: %w", err)
		}
	}

	return nil
}

func (c *Client) Send() error {
	_, err := io.Copy(c.connect, c.in)
	if err != nil {
		return fmt.Errorf("send: %w", err)
	}
	fmt.Fprintln(os.Stderr, "...EOF")

	return nil
}

func (c *Client) Receive() error {
	_, err := io.Copy(c.out, c.connect)
	if err != nil {
		return fmt.Errorf("receive: %w", err)
	}
	fmt.Fprintln(os.Stderr, "...Connection was closed by peer")

	return nil
}
