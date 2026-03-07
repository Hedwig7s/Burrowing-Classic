package protocol

import (
	"github.com/Hedwig7s/Burrowing-Classic/internal/cerror"
	"github.com/Hedwig7s/Burrowing-Classic/internal/networking/encoding"
)

type Protocol interface {
	Version() int

	CreatePacketBuilder(id PacketID) (PacketBuilder, error)
}

func EncodePacket(packet Packet) (*encoding.PacketWriter, error) {
	writer := encoding.NewPacketWriter(nil)
	err := packet.EncodeToWriter(writer)
	if err != nil {
		return nil, err
	}
	return writer, nil
}

func DecodePacket(protocol Protocol, id PacketID, reader *encoding.PacketReader) (Packet, error) {
	builder, err := protocol.CreatePacketBuilder(id)
	if err != nil {
		return nil, cerror.NewErrorf(PROTOCOL_PACKET_NOT_FOUND, "Couldn't find packet id %d", id)
	}

	packet, err := builder.BuildFromReader(reader)
	if err != nil {
		return nil, err
	}

	return packet, nil
}

const (
	PROTOCOL_PACKET_NOT_FOUND = iota
)
