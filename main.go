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

const repoURL = "https://github.com/peterhellberg/check-ssh-chat"

var buildCommit string

func main() {
	flag.Usage = usage
	flag.Parse()

	addr := fmt.Sprintf("%s:%d", *host, *port)

	l("Checking: %s", addr)
	if err := Check(addr); err != nil {
		l("Error: %v", err)
		os.Exit(1)
	}

	l("The ssh-chat server seems to be working")
}

func usage() {
	fmt.Fprintf(os.Stderr, "usage: ./check-ssh-chat [-h hostname] [-v]\n\n")

	if buildCommit != "" {
		fmt.Fprintf(os.Stderr, "build: "+repoURL+"/commit/"+buildCommit+"\n\n")
	}

	fmt.Fprintf(os.Stderr, "flags:\n")
	flag.PrintDefaults()
	fmt.Fprintf(os.Stderr, "\n")
	os.Exit(2)
}

func l(format string, args ...interface{}) {
	if *verbose {
		fmt.Printf(format+"\n", args...)
	}
}
