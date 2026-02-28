package networking

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"log"
	"net"

	"github.com/Hedwig7s/Burrowing-Classic/lib/networking/codec"
	"github.com/Hedwig7s/Burrowing-Classic/lib/networking/protocol"
)

const BUFFER_SIZE = 8096

var PROTOCOLS = map[byte]protocol.Protocol{
	0x07: &protocol.Protocol7{},
}

type Connection struct {
	conn     net.Conn
	closed   bool
	buffer   []byte
	protocol protocol.Protocol
}

func readData(connection *Connection, buffer []byte, ctx context.Context) error {
	_, err := io.ReadFull(connection.conn, buffer)
	if err != nil {
		select {
		case <-ctx.Done():
			return nil
		default:
			if errors.Is(err, io.EOF) { // Connection closed
				connection.Close()
				return nil
			}
			return fmt.Errorf("Error reading data: %v", err)
		}
	}
	return nil
}

// TODO: More detailed errors
func (connection *Connection) Start(ctx context.Context) error {
	go func() {
		<-ctx.Done()
		connection.Close()
	}()
	for {
		buffer := connection.buffer
		if err := readData(connection, buffer[:1], ctx); err != nil {
			return err
		}
		setProtocol := false
		packetId := buffer[0]
		if packetId == 0 && connection.protocol == nil {
			if err := readData(connection, buffer[:1], ctx); err != nil {
				return err
			}
			protocol := PROTOCOLS[buffer[0]]
			if protocol == nil {
				return fmt.Errorf("Protocol %d not found", buffer[0])
			}
			connection.protocol = protocol
			setProtocol = true
		} else if connection.protocol == nil {
			return fmt.Errorf("Non-identification packet sent despite no protocol being set")
		}
		packet, err := connection.protocol.NewPacket(protocol.PacketID(packetId))
		if err != nil {
			return err
		}
		length := packet.Size()
		var packetSlice []byte
		if setProtocol {
			packetSlice = buffer[1:length]
		} else {
			packetSlice = buffer[:length]
		}
		if err := readData(connection, packetSlice, ctx); err != nil {
			return err
		}
		// FIXME: Probably shouldn't define the reader like this
		if err := packet.DecodeFrom(codec.NewPacketReader(bytes.NewReader(buffer[:length]))); err != nil {
			return err
		}
		log.Printf("Got packet %d: %v\n", packetId, packet.Data())
	}
}

func (connection *Connection) Close() {
	if connection.closed {
		return
	}
	connection.closed = true

	if err := connection.conn.Close(); err != nil {
		log.Printf("Warning: Error in closing connection: %v\n", err) // TODO: Log better
	}
}

func NewConnection(conn net.Conn) *Connection {
	connection := &Connection{
		conn:   conn,
		closed: false,
		buffer: make([]byte, BUFFER_SIZE),
	}
	return connection
}
