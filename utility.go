package main

import (
	"log"
	"os"
	"os/exec"
	"regexp"
	"strings"

	"github.com/pborman/uuid"
)

func instance() (string, error) {
	if len(os.Getenv("NETCORE_NAME")) > 0 {
		return os.Getenv("NETCORE_NAME"), nil
	}
	if len(os.Getenv("ETCD_NAME")) > 0 {
		re := regexp.MustCompile(`^/([^/]+)/`)
		hostnameParts := re.FindStringSubmatch(os.Getenv("ETCD_NAME"))
		if len(hostnameParts) > 1 && len(hostnameParts[1]) > 0 {
			return hostnameParts[1], nil
		}
	}
	return getHostname()
}

func getHostname() (string, error) {
	fqdn, err := exec.Command("hostname", "-f").Output()
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(fqdn)), nil
}

func getUUID() string {
	return uuid.New()
}

// logAfterSuccess will start a goroutine that will log.Print the given message
// if a true value is received on the channel.
func logAfterSuccess(c chan bool, message string) {
	go func() {
		if success := <-c; success {
			log.Print(message)
		}
	}()
}