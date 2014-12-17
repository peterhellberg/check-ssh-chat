package main

import (
	"bytes"
	"errors"
	"fmt"
	"strings"
	"time"

	"golang.org/x/crypto/ssh"
)

func Check(addr string) error {
	conn, err := dial(addr)
	if err != nil {
		return err
	}

	session, err := conn.NewSession()
	if err != nil {
		return err
	}
	defer session.Close()

	in, err := session.StdinPipe()
	if err != nil {
		return err
	}

	var out bytes.Buffer
	session.Stdout = &out

	if err := session.Shell(); err != nil {
		return err
	}

	msg := fmt.Sprintf("check-msg-%d", time.Now().Unix())

	in.Write([]byte(fmt.Sprintf("/msg %s %s\r\n", *user, msg)))
	time.Sleep(*timeout)

	if !strings.Contains(out.String(), msg) {
		return errors.New("did not receive the check message")
	}

	return nil
}

func dial(addr string) (*ssh.Client, error) {
	key, err := MakeKey()
	if err != nil {
		return nil, err
	}

	return ssh.Dial("tcp", addr, &ssh.ClientConfig{
		User: *user,
		Auth: []ssh.AuthMethod{
			ssh.PublicKeys(key),
		},
	})
}
