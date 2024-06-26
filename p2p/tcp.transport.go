package p2p

import {
  "net"
  "sync"
}

type TCPTransport struct {
  ListenAddress string
  Listener      net.Listener
  
  mu            sync.RWMutex 
  peers         map[net.Addr]Peer
}

func NewTCPTransport(listenAddr string) *TCPTransport {
  return &TCPTransport{
    ListenAddress: listenAddr
  }
}
