package proxy

import (
	"fmt"
	"io"
	"net"
)

func authenticate(client *net.TCPConn) (byte, error) {
	// Check version
	version := make([]byte, 1)
	_, err := io.ReadFull(client, version)
	if err != nil {
		return SOCKS_AUTH_METHOD_NO_ACCEPTABLE_METHODS, NewErrAuthRequestParsing("No socks version")
	}
	if version[0] != SOCKS_VERSION {
		return SOCKS_AUTH_METHOD_NO_ACCEPTABLE_METHODS, NewErrAuthRequestParsing(fmt.Sprintf(
			"Socks version %v is expected, but not %v", SOCKS_VERSION, version[0]))
	}

	// Check auth method count
	methodCount := make([]byte, 1)
	_, err = io.ReadFull(client, methodCount)
	if err != nil {
		return SOCKS_AUTH_METHOD_NO_ACCEPTABLE_METHODS, NewErrAuthRequestParsing("No authentication method count")
	}

	// Check auth methods
	methods := make([]byte, methodCount[0])
	actualMethodCount, err := io.ReadFull(client, methods)
	if err != nil {
		return SOCKS_AUTH_METHOD_NO_ACCEPTABLE_METHODS, NewErrAuthRequestParsing(fmt.Sprintf(
			"Not enough authentication methods: Expected %v, received %v", methodCount, actualMethodCount))
	}

	for _, methods := range methods {
		if methods == SOCKS_AUTH_METHOD_NO_REQUIRED {
			// Found supported method
			return SOCKS_AUTH_METHOD_NO_REQUIRED, nil
		}
	}

	// Not found supported method
	return SOCKS_AUTH_METHOD_NO_ACCEPTABLE_METHODS, NewErrAuthRequestParsing(fmt.Sprintf(
		"Unsupported auth methods, %v method supported", []byte{SOCKS_AUTH_METHOD_NO_REQUIRED}))
}

func sendAuthReply(client *net.TCPConn, method byte) error {
	// Create message
	replyMsg := []byte{SOCKS_VERSION, method}

	// Send reply
	_, err := client.Write(replyMsg)
	if err != nil {
		return NewErrAuthReplySending(err.Error())
	}
	return nil
}
