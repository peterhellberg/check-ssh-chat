package main

import (
	"flag"
	"fmt"
	"os"
	"time"
)

var (
	user = flag.String("n", "check-ssh-chat", "Username")
	host = flag.String("h", "localhost", "Hostname")
	port = flag.Int("p", 22, "Port")

	timeout = flag.Duration("t", 5*time.Second, "Timeout for the check")
	verbose = flag.Bool("v", false, "Verbose output")
)

func main() {
	flag.Parse()

	addr := fmt.Sprintf("%s:%d", *host, *port)

	l("Checking: %s", addr)
	if err := Check(addr); err != nil {
		l("Error: %v", err)
		os.Exit(1)
	}

	l("The ssh-chat server seems to be working")
}

func l(format string, args ...interface{}) {
	if *verbose {
		fmt.Printf(format+"\n", args...)
	}
}
