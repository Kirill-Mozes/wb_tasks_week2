package funcs

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"log"
	"net"
	"time"
)

var (
	errBuildDialWithTimeout   = errors.New("mytelnet: can not build dial connection")
	errCloseConnection        = errors.New("mytelnet: can not close connection")
	errRequestToBadConnection = errors.New("mytelnet: can not get/send. probalby connection was closed by peer")
)

type Client struct {
	addr            string
	timeout         time.Duration
	conn            net.Conn
	inputDataReader *bufio.Reader
	connDataReader  *bufio.Reader
	writer          io.Writer
}

func NewClient(addr string, timeout time.Duration, inputDataReader io.Reader, writer io.Writer) *Client {
	return &Client{
		addr:            addr,
		timeout:         timeout,
		inputDataReader: bufio.NewReader(inputDataReader),
		writer:          writer,
	}
}

func (c *Client) BuildConnection() error {
	conn, err := net.DialTimeout("tcp", c.addr, c.timeout)
	if err != nil {
		return errBuildDialWithTimeout
	}
	c.conn = conn
	c.connDataReader = bufio.NewReader(conn)
	log.Println("mytelnet: connection build successfully")
	return nil
}

func (c *Client) Receive() error {
	text, err := c.connDataReader.ReadString('\n')
	if err != nil {
		if err == io.EOF {
			err = errRequestToBadConnection
		}
		return err
	}
	if _, err := fmt.Fprint(c.writer, text); err != nil {
		return err
	}
	return nil
}

func (c *Client) Send() error {
	text, err := c.inputDataReader.ReadString('\n')
	if err != nil {
		return err
	}
	if _, err := c.conn.Write([]byte(text)); err != nil {
		return errRequestToBadConnection
	}
	return nil

}

func (c *Client) Close() error {
	if err := c.conn.Close(); err != nil {
		return errCloseConnection
	}
	return nil
}
