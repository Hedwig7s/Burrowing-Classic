package protocol

import (
	"fmt"

	"github.com/Hedwig7s/Burrowing-Classic/lib/networking/codec"
)

type Protocol interface {
	Version() int
	NewPacket(id PacketID) (Packet, error)
}

func EncodePacket(packet Packet) (*codec.PacketWriter, error) {
	writer := codec.NewPacketWriter(nil)
	err := packet.EncodeTo(writer)
	if err != nil {
		return nil, err
	}
	return writer, nil
}

func DecodePacket(protocol Protocol, id PacketID, reader *codec.PacketReader) (Packet, error) {
	packet, err := protocol.NewPacket(id)
	if packet == nil {
		return nil, fmt.Errorf("Couldn't find packet id %d", id)
	}
	err = packet.DecodeFrom(reader)
	if err != nil {
		return nil, err
	}
	return packet, nil
}
