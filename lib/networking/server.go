package networking

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net"
)

type Server struct {
	bind_address string
	port         uint16
	listener     net.Listener
	started      bool
	connections  []*Connection
}

var (
	ServerAlreadyStarted = errors.New("Server already started!")
	ListenerWhileStopped = errors.New("Listener present despite server being stopped!")
)

func (server *Server) Start(ctx context.Context) error {
	if server.started {
		return ServerAlreadyStarted
	}
	if server.listener != nil {
		return ListenerWhileStopped
	}
	server.started = true
	var err error
	server.listener, err = net.Listen("tcp", fmt.Sprintf("%s:%d", server.bind_address, server.port))
	if err != nil {
		return err
	}
	go func() {
		<-ctx.Done()
		server.Close()
	}()
	for {
		conn, err := server.listener.Accept()
		if err != nil {
			select {
			case <-ctx.Done():
				return nil
			default:
				return err
			}
		}
		connection := NewConnection(conn)
		// FIXME: Wait group perhaps?
		go func() {
			err := connection.Start(ctx)
			if err != nil {
				log.Printf("Error in connection: %v", err)
				connection.Close()
			}
		}()
		server.connections = append(server.connections, connection) // TODO: Removal

	}
}

func (server *Server) Close() error {
	if !server.started {
		return nil
	}
	server.started = false
	if server.listener != nil {
		err := server.listener.Close()
		if err != nil {
			log.Printf("Warning: Error in closing server listener: %v\n", err) // TODO: Log better
		}
	}
	server.listener = nil
	return nil
}

func NewServer(bind_address string, port uint16) *Server {
	return &Server{bind_address: bind_address, port: port, started: false}
}
