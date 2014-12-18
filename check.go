package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"

	"golang.org/x/crypto/ssh"
)

func Check(addr string) error {
	go func() {
		time.Sleep(*timeout)
		l("Timed out.")
		os.Exit(2)
	}()

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

	out, err := session.StdoutPipe()
	if err != nil {
		return err
	}

	if err := session.Shell(); err != nil {
		return err
	}

	err = session.RequestPty("xterm", 80, 40, ssh.TerminalModes{})
	if err != nil {
		return err
	}

	msg := fmt.Sprintf("check-msg-%d", time.Now().Unix())
	in.Write([]byte(fmt.Sprintf("/msg %s %s\r\n", *user, msg)))

	scanner := bufio.NewScanner(out)
	for scanner.Scan() {
		line := scanner.Text()
		if err != nil {
			return err
		}
		if strings.Contains(line, msg) {
			return nil
		}
	}
	err = scanner.Err()
	if err != nil {
		return err
	}

	return errors.New("did not receive the check message")
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
