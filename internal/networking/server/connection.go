package server

import (
	"bytes"
	"context"
	"errors"
	"io"
	"log"
	"net"

	"github.com/Hedwig7s/Burrowing-Classic/internal/cerror"
	"github.com/Hedwig7s/Burrowing-Classic/internal/networking/encoding"
	"github.com/Hedwig7s/Burrowing-Classic/internal/networking/protocol"
	"github.com/Hedwig7s/Burrowing-Classic/internal/networking/protocol_impls"
)

const BUFFER_SIZE = 8096

const (
	CON_READ_ERROR = iota
	CON_PROTOCOL_NOT_FOUND
	CON_PACKET_WITHOUT_PROTOCOL
	CON_INVALID_PACKET_ID
)

var PROTOCOLS = map[byte]protocol.Protocol{
	0x07: &protocol_impls.Protocol7{},
}

type Connection struct {
	conn     net.Conn
	closed   bool
	buffer   []byte
	protocol protocol.Protocol
	writer   *encoding.PacketWriter
	id       uint
}

func (connection *Connection) Id() uint {
	return connection.id
}

func (connection *Connection) Protocol() protocol.Protocol {
	return connection.protocol
}

func readData(connection *Connection, buffer []byte, ctx context.Context) error {
	_, err := io.ReadFull(connection.conn, buffer)
	if err != nil {
		select {
		case <-ctx.Done():
			return nil
		default:
			if errors.Is(err, io.EOF) || err.Error() == "unexpected EOF" { // Connection closed
				connection.Close()
				return nil
			}
			return cerror.NewErrorf(CON_READ_ERROR, "Error reading data: %v", err)
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
		if connection.closed {
			return nil
		}
		buffer := connection.buffer
		if err := readData(connection, buffer[:1], ctx); err != nil {
			return err
		}
		setProtocol := false
		packetId := buffer[0]
		if packetId == 0 && connection.protocol == nil {
			if err := readData(connection, buffer[1:2], ctx); err != nil {
				return err
			}
			proto := PROTOCOLS[buffer[1]]
			if proto == nil {
				return cerror.NewErrorf(CON_PROTOCOL_NOT_FOUND, "Protocol %d not found", buffer[1])
			}
			connection.protocol = proto
			setProtocol = true
		} else if connection.protocol == nil {
			return cerror.NewError(CON_PACKET_WITHOUT_PROTOCOL, "Non-identification packet sent despite no protocol being set")
		}

		builder, err := connection.protocol.CreatePacketBuilder(protocol.PacketID(packetId))
		if err != nil {
			return err
		}

		length := builder.GetSize()
		var packetSlice []byte
		if setProtocol {
			packetSlice = buffer[1:length]
		} else {
			packetSlice = buffer[:length]
		}
		if err := readData(connection, packetSlice, ctx); err != nil {
			return err
		}

		reader := encoding.NewPacketReader(bytes.NewReader(buffer[:length]))
		packet, err := builder.BuildFromReader(reader)
		if err != nil {
			return err
		}

		if err := HandlePacket(connection, packet); err != nil {
			return err
		}
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

func (connection *Connection) Write(packet protocol.Packet) error {
	if err := packet.EncodeToWriter(connection.writer); err != nil {
		return err
	}

	return nil
}

func NewConnection(conn net.Conn) *Connection {
	connection := &Connection{
		conn:   conn,
		closed: false,
		buffer: make([]byte, BUFFER_SIZE),
		writer: encoding.NewPacketWriter(conn),
	}
	return connection
}
