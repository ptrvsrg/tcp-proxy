package proxy

import (
	"net"

	"tcp-proxy/internal/log"
)

func Start(port int) {
	// Create listener
	listener, err := net.ListenTCP("tcp4", &net.TCPAddr{Port: port})
	if err != nil {
		log.Log.Errorf("Listener creation error: %v", err)
		return
	}
	log.Log.Infof("Proxy server is listening on %v", listener.Addr())

	// Defer closing listener
	defer func() {
		listener.Close()
		log.Log.Debugf("Listener is closed")
	}()

	// Accept clients
	for {
		client, err := listener.AcceptTCP()
		if err != nil {
			log.Log.Errorf("Client accepting error: %v", err)
			return
		}
		log.Log.Infof("Accepted client %v", client.RemoteAddr())

		go handleClient(client)
	}
}

func handleClient(client *net.TCPConn) {
	// Defer closing client connection
	defer func() {
		client.Close()
		log.Log.Debugf("%v: Client connection is closed", client.RemoteAddr())
	}()

	// Authenticate
	authReply, err := authenticate(client)
	if err != nil {
		log.Log.Errorf("%v: %v", client.RemoteAddr(), err)
		if err := sendAuthReply(client, authReply); err != nil {
			log.Log.Errorf("%v: %v", client.RemoteAddr(), err)
		}
		return
	}
	if err := sendAuthReply(client, authReply); err != nil {
		log.Log.Errorf("%v: %v", client.RemoteAddr(), err)
		return
	}
	log.Log.Debugf("%v: Client is authenticated", client.RemoteAddr())

	// Execute command
	peer, commandReply, err := connectCommand(client)
	if err != nil {
		log.Log.Errorf("%v: %v", client.RemoteAddr(), err)
		if err := sendCommandReply(client, commandReply); err != nil {
			log.Log.Errorf("%v: %v", client.RemoteAddr(), err)
		}
		return
	}
	if err := sendCommandReply(client, commandReply); err != nil {
		log.Log.Errorf("%v: %v", client.RemoteAddr(), err)
		return
	}
	log.Log.Debugf("%v: Proxy server is connected to peer", client.RemoteAddr())

	// Transfer data
	transferData(client, peer)
}
