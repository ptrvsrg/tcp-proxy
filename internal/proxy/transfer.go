package proxy

import (
	"io"
	"net"
	"sync"

	"tcp-proxy/internal/log"
)

func transferData(client *net.TCPConn, peer *net.TCPConn) {
	var wg sync.WaitGroup
	wg.Add(2)

	go copyData(client, peer, &wg)
	go copyData(peer, client, &wg)

	wg.Wait()
}

func copyData(dest *net.TCPConn, src *net.TCPConn, wg *sync.WaitGroup) {
	defer wg.Done()
	defer dest.Close()

	_, err := io.Copy(dest, src)
	if err != nil {
		log.Log.Errorf("%v: Reading error: %v", dest.RemoteAddr(), err)
	}
}
