package protocol

import (
	"fmt"

	"github.com/Hedwig7s/Burrowing-Classic/lib/networking/codec"
)

// writeError is a helper to chain multiple writes and return the first error
func writeError(errs ...error) error {
	for _, err := range errs {
		if err != nil {
			return err
		}
	}
	return nil
}

type Protocol7 struct{}

func (p *Protocol7) Version() int {
	return 7
}

var packetConstructors = map[PacketID]func() Packet{
	PacketID_Identification:               func() Packet { return &IdentificationPacket{} },
	PacketID_Ping:                         func() Packet { return &PingPacket{} },
	PacketID_LevelInitialize:              func() Packet { return &LevelInitializePacket{} },
	PacketID_LevelDataChunk:               func() Packet { return &LevelDataChunkPacket{} },
	PacketID_LevelFinalize:                func() Packet { return &LevelFinalizePacket{} },
	PacketID_SetBlockServerbound:          func() Packet { return &SetBlockServerboundPacket{} },
	PacketID_SetBlockClientbound:          func() Packet { return &SetBlockClientboundPacket{} },
	PacketID_SpawnPlayer:                  func() Packet { return &SpawnPlayerPacket{} },
	PacketID_SetPositionAndOrientation:    func() Packet { return &SetPositionAndOrientationPacket{} },
	PacketID_PositionAndOrientationUpdate: func() Packet { return &PositionAndOrientationUpdatePacket{} },
	PacketID_PositionUpdate:               func() Packet { return &PositionUpdatePacket{} },
	PacketID_OrientationUpdate:            func() Packet { return &OrientationUpdatePacket{} },
	PacketID_DespawnPlayer:                func() Packet { return &DespawnPlayerPacket{} },
	PacketID_Message:                      func() Packet { return &MessagePacket{} },
	PacketID_DisconnectPlayer:             func() Packet { return &DisconnectPlayerPacket{} },
	PacketID_UpdateUserType:               func() Packet { return &UpdateUserTypePacket{} },
}

func (p *Protocol7) NewPacket(id PacketID) (Packet, error) {
	constructor, ok := packetConstructors[id]
	if !ok {
		return nil, fmt.Errorf("Packet %d not found", id)
	}

	return constructor(), nil
}

// Identification Packet (0x00)
type IdentificationPacket struct {
	data IdentificationData
}

func (p *IdentificationPacket) ID() PacketID {
	return PacketID_Identification
}

func (p *IdentificationPacket) Size() int {
	return 130
}

func (p *IdentificationPacket) Data() any {
	return p.data
}

func (p *IdentificationPacket) EncodeTo(writer *codec.PacketWriter) error {
	return writeError(
		writer.Byte(byte(p.ID())),
		writer.Byte(p.data.ProtocolVersion),
		writer.String64(p.data.Name),
		writer.String64(p.data.MotdOrKey),
		writer.Byte(p.data.UserType),
	)
}

func (p *IdentificationPacket) DecodeFrom(reader *codec.PacketReader) error {
	var err error
	p.data.ProtocolVersion, err = reader.Byte()
	if err != nil {
		return err
	}
	p.data.Name, err = reader.String64()
	if err != nil {
		return err
	}
	p.data.MotdOrKey, err = reader.String64()
	if err != nil {
		return err
	}
	p.data.UserType, err = reader.Byte()
	if err != nil {
		return err
	}
	return nil
}

// Ping Packet (0x01)
type PingPacket struct {
	data PingPacketData
}

func (p *PingPacket) ID() PacketID {
	return PacketID_Ping
}

func (p *PingPacket) Size() int {
	return 0
}

func (p *PingPacket) Data() any {
	return p.data
}

func (p *PingPacket) EncodeTo(writer *codec.PacketWriter) error {
	return writer.Byte(byte(p.ID()))
}

func (p *PingPacket) DecodeFrom(reader *codec.PacketReader) error {
	return nil
}

// Level Initialize Packet (0x02)
type LevelInitializePacket struct {
	data LevelInitializeData
}

func (p *LevelInitializePacket) ID() PacketID {
	return PacketID_LevelInitialize
}

func (p *LevelInitializePacket) Size() int {
	return 0
}

func (p *LevelInitializePacket) Data() any {
	return p.data
}

func (p *LevelInitializePacket) EncodeTo(writer *codec.PacketWriter) error {
	return writer.Byte(byte(p.ID()))
}

func (p *LevelInitializePacket) DecodeFrom(reader *codec.PacketReader) error {
	return nil
}

// Level Data Chunk Packet (0x03)
type LevelDataChunkPacket struct {
	data LevelDataChunkData
}

func (p *LevelDataChunkPacket) ID() PacketID {
	return PacketID_LevelDataChunk
}

func (p *LevelDataChunkPacket) Size() int {
	return 1027
}

func (p *LevelDataChunkPacket) Data() any {
	return p.data
}

func (p *LevelDataChunkPacket) EncodeTo(writer *codec.PacketWriter) error {
	return writeError(
		writer.Byte(byte(p.ID())),
		writer.Short(p.data.ChunkLength),
		writer.Bytes(p.data.ChunkData[:]),
		writer.Byte(p.data.PercentComplete),
	)
}

func (p *LevelDataChunkPacket) DecodeFrom(reader *codec.PacketReader) error {
	var err error
	p.data.ChunkLength, err = reader.Short()
	if err != nil {
		return err
	}
	chunkBytes, err := reader.Bytes(1024)
	if err != nil {
		return err
	}
	copy(p.data.ChunkData[:], chunkBytes)
	p.data.PercentComplete, err = reader.Byte()
	if err != nil {
		return err
	}
	return nil
}

// Level Finalize Packet (0x04)
type LevelFinalizePacket struct {
	data LevelFinalizeData
}

func (p *LevelFinalizePacket) ID() PacketID {
	return PacketID_LevelFinalize
}

func (p *LevelFinalizePacket) Size() int {
	return 6
}

func (p *LevelFinalizePacket) Data() any {
	return p.data
}

func (p *LevelFinalizePacket) EncodeTo(writer *codec.PacketWriter) error {
	return writeError(
		writer.Byte(byte(p.ID())),
		writer.Short(p.data.XSize),
		writer.Short(p.data.YSize),
		writer.Short(p.data.ZSize),
	)
}

func (p *LevelFinalizePacket) DecodeFrom(reader *codec.PacketReader) error {
	var err error
	p.data.XSize, err = reader.Short()
	if err != nil {
		return err
	}
	p.data.YSize, err = reader.Short()
	if err != nil {
		return err
	}
	p.data.ZSize, err = reader.Short()
	if err != nil {
		return err
	}
	return nil
}

// Set Block Serverbound Packet (0x05)
type SetBlockServerboundPacket struct {
	data SetBlockServerboundData
}

func (p *SetBlockServerboundPacket) ID() PacketID {
	return PacketID_SetBlockServerbound
}

func (p *SetBlockServerboundPacket) Size() int {
	return 8
}

func (p *SetBlockServerboundPacket) Data() any {
	return p.data
}

func (p *SetBlockServerboundPacket) EncodeTo(writer *codec.PacketWriter) error {
	return writeError(
		writer.Byte(byte(p.ID())),
		writer.Short(p.data.X),
		writer.Short(p.data.Y),
		writer.Short(p.data.Z),
		writer.Byte(p.data.Mode),
		writer.Byte(p.data.BlockType),
	)
}

func (p *SetBlockServerboundPacket) DecodeFrom(reader *codec.PacketReader) error {
	var err error
	p.data.X, err = reader.Short()
	if err != nil {
		return err
	}
	p.data.Y, err = reader.Short()
	if err != nil {
		return err
	}
	p.data.Z, err = reader.Short()
	if err != nil {
		return err
	}
	p.data.Mode, err = reader.Byte()
	if err != nil {
		return err
	}
	p.data.BlockType, err = reader.Byte()
	if err != nil {
		return err
	}
	return nil
}

// Set Block Clientbound Packet (0x06)
type SetBlockClientboundPacket struct {
	data SetBlockClientboundData
}

func (p *SetBlockClientboundPacket) ID() PacketID {
	return PacketID_SetBlockClientbound
}

func (p *SetBlockClientboundPacket) Size() int {
	return 7
}

func (p *SetBlockClientboundPacket) Data() any {
	return p.data
}

func (p *SetBlockClientboundPacket) EncodeTo(writer *codec.PacketWriter) error {
	return writeError(
		writer.Byte(byte(p.ID())),
		writer.Short(p.data.X),
		writer.Short(p.data.Y),
		writer.Short(p.data.Z),
		writer.Byte(p.data.BlockType),
	)
}

func (p *SetBlockClientboundPacket) DecodeFrom(reader *codec.PacketReader) error {
	var err error
	p.data.X, err = reader.Short()
	if err != nil {
		return err
	}
	p.data.Y, err = reader.Short()
	if err != nil {
		return err
	}
	p.data.Z, err = reader.Short()
	if err != nil {
		return err
	}
	p.data.BlockType, err = reader.Byte()
	if err != nil {
		return err
	}
	return nil
}

// Spawn Player Packet (0x07)
type SpawnPlayerPacket struct {
	data SpawnPlayerData
}

func (p *SpawnPlayerPacket) ID() PacketID {
	return PacketID_SpawnPlayer
}

func (p *SpawnPlayerPacket) Size() int {
	return 73
}

func (p *SpawnPlayerPacket) Data() any {
	return p.data
}

func (p *SpawnPlayerPacket) EncodeTo(writer *codec.PacketWriter) error {
	return writeError(
		writer.Byte(byte(p.ID())),
		writer.SByte(p.data.PlayerID),
		writer.String64(p.data.PlayerName),
		writer.FShort(p.data.X),
		writer.FShort(p.data.Y),
		writer.FShort(p.data.Z),
		writer.Byte(p.data.Yaw),
		writer.Byte(p.data.Pitch),
	)
}

func (p *SpawnPlayerPacket) DecodeFrom(reader *codec.PacketReader) error {
	var err error
	p.data.PlayerID, err = reader.SByte()
	if err != nil {
		return err
	}
	p.data.PlayerName, err = reader.String64()
	if err != nil {
		return err
	}
	p.data.X, err = reader.FShort()
	if err != nil {
		return err
	}
	p.data.Y, err = reader.FShort()
	if err != nil {
		return err
	}
	p.data.Z, err = reader.FShort()
	if err != nil {
		return err
	}
	p.data.Yaw, err = reader.Byte()
	if err != nil {
		return err
	}
	p.data.Pitch, err = reader.Byte()
	if err != nil {
		return err
	}
	return nil
}

// Set Position and Orientation Packet (0x08)
type SetPositionAndOrientationPacket struct {
	data SetPositionAndOrientationData
}

func (p *SetPositionAndOrientationPacket) ID() PacketID {
	return PacketID_SetPositionAndOrientation
}

func (p *SetPositionAndOrientationPacket) Size() int {
	return 9
}

func (p *SetPositionAndOrientationPacket) Data() any {
	return p.data
}

func (p *SetPositionAndOrientationPacket) EncodeTo(writer *codec.PacketWriter) error {
	return writeError(
		writer.Byte(byte(p.ID())),
		writer.SByte(p.data.PlayerID),
		writer.FShort(p.data.X),
		writer.FShort(p.data.Y),
		writer.FShort(p.data.Z),
		writer.Byte(p.data.Yaw),
		writer.Byte(p.data.Pitch),
	)
}

func (p *SetPositionAndOrientationPacket) DecodeFrom(reader *codec.PacketReader) error {
	var err error
	p.data.PlayerID, err = reader.SByte()
	if err != nil {
		return err
	}
	p.data.X, err = reader.FShort()
	if err != nil {
		return err
	}
	p.data.Y, err = reader.FShort()
	if err != nil {
		return err
	}
	p.data.Z, err = reader.FShort()
	if err != nil {
		return err
	}
	p.data.Yaw, err = reader.Byte()
	if err != nil {
		return err
	}
	p.data.Pitch, err = reader.Byte()
	if err != nil {
		return err
	}
	return nil
}

// Position and Orientation Update Packet (0x09)
type PositionAndOrientationUpdatePacket struct {
	data PositionAndOrientationUpdateData
}

func (p *PositionAndOrientationUpdatePacket) ID() PacketID {
	return PacketID_PositionAndOrientationUpdate
}

func (p *PositionAndOrientationUpdatePacket) Size() int {
	return 6
}

func (p *PositionAndOrientationUpdatePacket) Data() any {
	return p.data
}

func (p *PositionAndOrientationUpdatePacket) EncodeTo(writer *codec.PacketWriter) error {
	return writeError(
		writer.Byte(byte(p.ID())),
		writer.SByte(p.data.PlayerID),
		writer.FByte(p.data.ChangeX),
		writer.FByte(p.data.ChangeY),
		writer.FByte(p.data.ChangeZ),
		writer.Byte(p.data.Yaw),
		writer.Byte(p.data.Pitch),
	)
}

func (p *PositionAndOrientationUpdatePacket) DecodeFrom(reader *codec.PacketReader) error {
	var err error
	p.data.PlayerID, err = reader.SByte()
	if err != nil {
		return err
	}
	p.data.ChangeX, err = reader.FByte()
	if err != nil {
		return err
	}
	p.data.ChangeY, err = reader.FByte()
	if err != nil {
		return err
	}
	p.data.ChangeZ, err = reader.FByte()
	if err != nil {
		return err
	}
	p.data.Yaw, err = reader.Byte()
	if err != nil {
		return err
	}
	p.data.Pitch, err = reader.Byte()
	if err != nil {
		return err
	}
	return nil
}

// Position Update Packet (0x0a)
type PositionUpdatePacket struct {
	data PositionUpdateData
}

func (p *PositionUpdatePacket) ID() PacketID {
	return PacketID_PositionUpdate
}

func (p *PositionUpdatePacket) Size() int {
	return 4
}

func (p *PositionUpdatePacket) Data() any {
	return p.data
}

func (p *PositionUpdatePacket) EncodeTo(writer *codec.PacketWriter) error {
	return writeError(
		writer.Byte(byte(p.ID())),
		writer.SByte(p.data.PlayerID),
		writer.FByte(p.data.ChangeX),
		writer.FByte(p.data.ChangeY),
		writer.FByte(p.data.ChangeZ),
	)
}

func (p *PositionUpdatePacket) DecodeFrom(reader *codec.PacketReader) error {
	var err error
	p.data.PlayerID, err = reader.SByte()
	if err != nil {
		return err
	}
	p.data.ChangeX, err = reader.FByte()
	if err != nil {
		return err
	}
	p.data.ChangeY, err = reader.FByte()
	if err != nil {
		return err
	}
	p.data.ChangeZ, err = reader.FByte()
	if err != nil {
		return err
	}
	return nil
}

// Orientation Update Packet (0x0b)
type OrientationUpdatePacket struct {
	data OrientationUpdateData
}

func (p *OrientationUpdatePacket) ID() PacketID {
	return PacketID_OrientationUpdate
}

func (p *OrientationUpdatePacket) Size() int {
	return 3
}

func (p *OrientationUpdatePacket) Data() any {
	return p.data
}

func (p *OrientationUpdatePacket) EncodeTo(writer *codec.PacketWriter) error {
	return writeError(
		writer.Byte(byte(p.ID())),
		writer.SByte(p.data.PlayerID),
		writer.Byte(p.data.Yaw),
		writer.Byte(p.data.Pitch),
	)
}

func (p *OrientationUpdatePacket) DecodeFrom(reader *codec.PacketReader) error {
	var err error
	p.data.PlayerID, err = reader.SByte()
	if err != nil {
		return err
	}
	p.data.Yaw, err = reader.Byte()
	if err != nil {
		return err
	}
	p.data.Pitch, err = reader.Byte()
	if err != nil {
		return err
	}
	return nil
}

// Despawn Player Packet (0x0c)
type DespawnPlayerPacket struct {
	data DespawnPlayerData
}

func (p *DespawnPlayerPacket) ID() PacketID {
	return PacketID_DespawnPlayer
}

func (p *DespawnPlayerPacket) Size() int {
	return 1
}

func (p *DespawnPlayerPacket) Data() any {
	return p.data
}

func (p *DespawnPlayerPacket) EncodeTo(writer *codec.PacketWriter) error {
	return writeError(
		writer.Byte(byte(p.ID())),
		writer.SByte(p.data.PlayerID),
	)
}

func (p *DespawnPlayerPacket) DecodeFrom(reader *codec.PacketReader) error {
	var err error
	p.data.PlayerID, err = reader.SByte()
	if err != nil {
		return err
	}
	return nil
}

// Message Packet (0x0d)
type MessagePacket struct {
	data MessageData
}

func (p *MessagePacket) ID() PacketID {
	return PacketID_Message
}

func (p *MessagePacket) Size() int {
	return 65
}

func (p *MessagePacket) Data() any {
	return p.data
}

func (p *MessagePacket) EncodeTo(writer *codec.PacketWriter) error {
	return writeError(
		writer.Byte(byte(p.ID())),
		writer.SByte(p.data.PlayerID),
		writer.String64(p.data.Message),
	)
}

func (p *MessagePacket) DecodeFrom(reader *codec.PacketReader) error {
	var err error
	p.data.PlayerID, err = reader.SByte()
	if err != nil {
		return err
	}
	p.data.Message, err = reader.String64()
	if err != nil {
		return err
	}
	return nil
}

// Disconnect Player Packet (0x0e)
type DisconnectPlayerPacket struct {
	data DisconnectPlayerData
}

func (p *DisconnectPlayerPacket) ID() PacketID {
	return PacketID_DisconnectPlayer
}

func (p *DisconnectPlayerPacket) Size() int {
	return 64
}

func (p *DisconnectPlayerPacket) Data() any {
	return p.data
}

func (p *DisconnectPlayerPacket) EncodeTo(writer *codec.PacketWriter) error {
	return writeError(
		writer.Byte(byte(p.ID())),
		writer.String64(p.data.DisconnectReason),
	)
}

func (p *DisconnectPlayerPacket) DecodeFrom(reader *codec.PacketReader) error {
	var err error
	p.data.DisconnectReason, err = reader.String64()
	if err != nil {
		return err
	}
	return nil
}

// Update User Type Packet (0x0f)
type UpdateUserTypePacket struct {
	data UpdateUserTypeData
}

func (p *UpdateUserTypePacket) ID() PacketID {
	return PacketID_UpdateUserType
}

func (p *UpdateUserTypePacket) Size() int {
	return 1
}

func (p *UpdateUserTypePacket) Data() any {
	return p.data
}

func (p *UpdateUserTypePacket) EncodeTo(writer *codec.PacketWriter) error {
	return writeError(
		writer.Byte(byte(p.ID())),
		writer.Byte(p.data.UserType),
	)
}

func (p *UpdateUserTypePacket) DecodeFrom(reader *codec.PacketReader) error {
	var err error
	p.data.UserType, err = reader.Byte()
	if err != nil {
		return err
	}
	return nil
}
