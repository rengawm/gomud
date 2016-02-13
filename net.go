package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"sync"
)

type NetManager struct {
	port            int
	connections     []*Connection
	connectionsLock *sync.Mutex
}

func NewNetManager(port int) *NetManager {
	return &NetManager{
		port:            port,
		connectionsLock: &sync.Mutex{},
	}
}

func (self *NetManager) Start() (err error) {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%v", self.port))
	if err != nil {
		return fmt.Errorf("Error creating listener: %v", err)
	}
	log.Printf("Accepting connections on port %v", self.port)

	for {
		// Wait for a connection.
		netConn, err := listener.Accept()
		if err != nil {
			log.Printf("Error accepting connection: %v", err)
			continue
		}

		connection := NewConnection(netConn)

		self.connectionsLock.Lock()
		self.connections = append(self.connections, connection)
		go connection.Run()
		self.connectionsLock.Unlock()
	}
}

type Connection struct {
	netConn net.Conn
	In      chan string
	Out     chan string
}

func NewConnection(netConn net.Conn) *Connection {
	return &Connection{
		netConn: netConn,
		In:      make(chan string, 10000),
		Out:     make(chan string, 10000),
	}
}

func (self *Connection) Run() {
	scanner := bufio.NewScanner(self.netConn)
	for scanner.Scan() {
		log.Printf("RECV: %s", scanner.Text())
	}
}
