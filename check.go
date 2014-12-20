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

	now := time.Now().Unix()
	msg := fmt.Sprintf("check-msg-%d", now)
	username := fmt.Sprintf("%s-%d", *user, now)

	conn, err := dial(addr, username)
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

	in.Write([]byte(fmt.Sprintf("/msg %s %s\r\n", username, msg)))

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

func dial(addr, username string) (*ssh.Client, error) {
	key, err := MakeKey()
	if err != nil {
		return nil, err
	}

	return ssh.Dial("tcp", addr, &ssh.ClientConfig{
		User: username,
		Auth: []ssh.AuthMethod{
			ssh.PublicKeys(key),
		},
	})
}
