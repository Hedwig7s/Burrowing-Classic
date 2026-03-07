package protocol

import (
	"bytes"

	"github.com/Hedwig7s/Burrowing-Classic/internal/networking/encoding"
)

type PacketID byte

const (
	PacketID_Identification = iota
	PacketID_Ping
	PacketID_LevelInitialize
	PacketID_LevelDataChunk
	PacketID_LevelFinalize
	PacketID_SetBlockServerbound
	PacketID_SetBlockClientbound
	PacketID_SpawnPlayer
	PacketID_SetPositionAndOrientation
	PacketID_PositionAndOrientationUpdate
	PacketID_PositionUpdate
	PacketID_OrientationUpdate
	PacketID_DespawnPlayer
	PacketID_Message
	PacketID_DisconnectPlayer
	PacketID_UpdateUserType
)

type Packet interface {
	ID() PacketID
	EncodeToWriter(writer *encoding.PacketWriter) error
	Data() any
	Size() int
}

type PacketBuilder interface {
	GetSize() int
	BuildFromReader(reader *encoding.PacketReader) (Packet, error)
	Build(data any) (Packet, error)
}

func DecodePacketFromBytes(builder PacketBuilder, buffer []byte) (Packet, *encoding.PacketReader, error) {
	reader := encoding.NewPacketReader(bytes.NewBuffer(buffer))
	packet, err := builder.BuildFromReader(reader)
	return packet, reader, err
}
