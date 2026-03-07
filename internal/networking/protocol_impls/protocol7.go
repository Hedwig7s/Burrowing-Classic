package protocol_impls

import (
	"github.com/Hedwig7s/Burrowing-Classic/internal/cerror"
	"github.com/Hedwig7s/Burrowing-Classic/internal/networking/encoding"
	"github.com/Hedwig7s/Burrowing-Classic/internal/networking/protocol"
)

const (
	BUILDER_DATA_TYPE_MISMATCH = iota
	PROTOCOL_PACKET_NOT_FOUND
)

func writeError(errs ...error) error {
	for _, err := range errs {
		if err != nil {
			return err
		}
	}
	return nil
}

func buildPacket[T any](data any, constructor func(T) protocol.Packet) (protocol.Packet, error) {
	d, ok := data.(T)
	if !ok {
		return nil, cerror.NewErrorf(BUILDER_DATA_TYPE_MISMATCH, "Build: expected %T, got %T", new(T), data)
	}
	return constructor(d), nil
}

type Protocol7 struct{}

func (p *Protocol7) Version() int {
	return 7
}

func (p *Protocol7) CreatePacketBuilder(id protocol.PacketID) (protocol.PacketBuilder, error) {
	switch id {
	case protocol.PacketID_Identification:
		return &identificationBuilder7{}, nil
	case protocol.PacketID_Ping:
		return &pingBuilder7{}, nil
	case protocol.PacketID_LevelInitialize:
		return &levelInitializeBuilder7{}, nil
	case protocol.PacketID_LevelDataChunk:
		return &levelDataChunkBuilder7{}, nil
	case protocol.PacketID_LevelFinalize:
		return &levelFinalizeBuilder7{}, nil
	case protocol.PacketID_SetBlockServerbound:
		return &setBlockServerboundBuilder7{}, nil
	case protocol.PacketID_SetBlockClientbound:
		return &setBlockClientboundBuilder7{}, nil
	case protocol.PacketID_SpawnPlayer:
		return &spawnPlayerBuilder7{}, nil
	case protocol.PacketID_SetPositionAndOrientation:
		return &setPositionAndOrientationBuilder7{}, nil
	case protocol.PacketID_PositionAndOrientationUpdate:
		return &positionAndOrientationUpdateBuilder7{}, nil
	case protocol.PacketID_PositionUpdate:
		return &positionUpdateBuilder7{}, nil
	case protocol.PacketID_OrientationUpdate:
		return &orientationUpdateBuilder7{}, nil
	case protocol.PacketID_DespawnPlayer:
		return &despawnPlayerBuilder7{}, nil
	case protocol.PacketID_Message:
		return &messageBuilder7{}, nil
	case protocol.PacketID_DisconnectPlayer:
		return &disconnectPlayerBuilder7{}, nil
	case protocol.PacketID_UpdateUserType:
		return &updateUserTypeBuilder7{}, nil
	default:
		return nil, cerror.NewErrorf(protocol.PROTOCOL_PACKET_NOT_FOUND, "Packet %d not found", id)
	}
}

type IdentificationPacket7 struct {
	id   protocol.PacketID
	data encoding.IdentificationData
}

func (p *IdentificationPacket7) ID() protocol.PacketID {
	return p.id
}

func (p *IdentificationPacket7) Size() int {
	return 131
}

func (p *IdentificationPacket7) Data() any {
	return p.data
}

func (p *IdentificationPacket7) EncodeToWriter(writer *encoding.PacketWriter) error {
	return writeError(
		writer.Byte(byte(p.ID())),
		writer.Byte(p.data.ProtocolVersion),
		writer.String64(p.data.Name),
		writer.String64(p.data.MotdOrKey),
		writer.Byte(p.data.UserType),
	)
}

type identificationBuilder7 struct{}

func (b *identificationBuilder7) GetSize() int {
	return 130
}

func (b *identificationBuilder7) BuildFromReader(reader *encoding.PacketReader) (protocol.Packet, error) {

	var data encoding.IdentificationData
	var err error

	data.ProtocolVersion, err = reader.Byte()
	if err != nil {
		return nil, err
	}

	data.Name, err = reader.String64()
	if err != nil {
		return nil, err
	}

	data.MotdOrKey, err = reader.String64()
	if err != nil {
		return nil, err
	}

	data.UserType, err = reader.Byte()
	if err != nil {
		return nil, err
	}

	return &IdentificationPacket7{
		id:   protocol.PacketID_Identification,
		data: data,
	}, nil
}

func (b *identificationBuilder7) Build(data any) (protocol.Packet, error) {
	return buildPacket[encoding.IdentificationData](data, func(d encoding.IdentificationData) protocol.Packet {
		return &IdentificationPacket7{
			id:   protocol.PacketID_Identification,
			data: d,
		}
	})
}

type PingPacket7 struct {
	id   protocol.PacketID
	data encoding.PingPacketData
}

func (p *PingPacket7) ID() protocol.PacketID {
	return p.id
}

func (p *PingPacket7) Size() int {
	return 1
}

func (p *PingPacket7) Data() any {
	return p.data
}

func (p *PingPacket7) EncodeToWriter(writer *encoding.PacketWriter) error {
	return writer.Byte(byte(p.ID()))
}

type pingBuilder7 struct{}

func (b *pingBuilder7) GetSize() int {
	return 0
}

func (b *pingBuilder7) BuildFromReader(reader *encoding.PacketReader) (protocol.Packet, error) {
	return &PingPacket7{
		id:   protocol.PacketID_Ping,
		data: encoding.PingPacketData{},
	}, nil
}

func (b *pingBuilder7) Build(data any) (protocol.Packet, error) {
	return buildPacket[encoding.PingPacketData](data, func(d encoding.PingPacketData) protocol.Packet {
		return &PingPacket7{
			id:   protocol.PacketID_Ping,
			data: d,
		}
	})
}

type LevelInitializePacket7 struct {
	id   protocol.PacketID
	data encoding.LevelInitializeData
}

func (p *LevelInitializePacket7) ID() protocol.PacketID {
	return p.id
}

func (p *LevelInitializePacket7) Size() int {
	return 1
}

func (p *LevelInitializePacket7) Data() any {
	return p.data
}

func (p *LevelInitializePacket7) EncodeToWriter(writer *encoding.PacketWriter) error {
	return writer.Byte(byte(p.ID()))
}

type levelInitializeBuilder7 struct{}

func (b *levelInitializeBuilder7) GetSize() int {
	return 0
}

func (b *levelInitializeBuilder7) BuildFromReader(reader *encoding.PacketReader) (protocol.Packet, error) {
	return &LevelInitializePacket7{
		id:   protocol.PacketID_LevelInitialize,
		data: encoding.LevelInitializeData{},
	}, nil
}

func (b *levelInitializeBuilder7) Build(data any) (protocol.Packet, error) {
	return buildPacket[encoding.LevelInitializeData](data, func(d encoding.LevelInitializeData) protocol.Packet {
		return &LevelInitializePacket7{
			id:   protocol.PacketID_LevelInitialize,
			data: d,
		}
	})
}

type LevelDataChunkPacket7 struct {
	id   protocol.PacketID
	data encoding.LevelDataChunkData
}

func (p *LevelDataChunkPacket7) ID() protocol.PacketID {
	return p.id
}

func (p *LevelDataChunkPacket7) Size() int {
	return 1028
}

func (p *LevelDataChunkPacket7) Data() any {
	return p.data
}

func (p *LevelDataChunkPacket7) EncodeToWriter(writer *encoding.PacketWriter) error {
	return writeError(
		writer.Byte(byte(p.ID())),
		writer.Short(p.data.ChunkLength),
		writer.Bytes(p.data.ChunkData[:]),
		writer.Byte(p.data.PercentComplete),
	)
}

type levelDataChunkBuilder7 struct{}

func (b *levelDataChunkBuilder7) GetSize() int {
	return 1027
}

func (b *levelDataChunkBuilder7) BuildFromReader(reader *encoding.PacketReader) (protocol.Packet, error) {
	var data encoding.LevelDataChunkData
	var err error

	data.ChunkLength, err = reader.Short()
	if err != nil {
		return nil, err
	}

	chunkBytes, err := reader.Bytes(1024)
	if err != nil {
		return nil, err
	}
	copy(data.ChunkData[:], chunkBytes)

	data.PercentComplete, err = reader.Byte()
	if err != nil {
		return nil, err
	}

	return &LevelDataChunkPacket7{
		id:   protocol.PacketID_LevelDataChunk,
		data: data,
	}, nil
}

func (b *levelDataChunkBuilder7) Build(data any) (protocol.Packet, error) {
	return buildPacket[encoding.LevelDataChunkData](data, func(d encoding.LevelDataChunkData) protocol.Packet {
		return &LevelDataChunkPacket7{
			id:   protocol.PacketID_LevelDataChunk,
			data: d,
		}
	})
}

type LevelFinalizePacket7 struct {
	id   protocol.PacketID
	data encoding.LevelFinalizeData
}

func (p *LevelFinalizePacket7) ID() protocol.PacketID {
	return p.id
}

func (p *LevelFinalizePacket7) Size() int {
	return 7
}

func (p *LevelFinalizePacket7) Data() any {
	return p.data
}

func (p *LevelFinalizePacket7) EncodeToWriter(writer *encoding.PacketWriter) error {
	return writeError(
		writer.Byte(byte(p.ID())),
		writer.Short(p.data.XSize),
		writer.Short(p.data.YSize),
		writer.Short(p.data.ZSize),
	)
}

type levelFinalizeBuilder7 struct{}

func (b *levelFinalizeBuilder7) GetSize() int {
	return 6
}

func (b *levelFinalizeBuilder7) BuildFromReader(reader *encoding.PacketReader) (protocol.Packet, error) {
	var data encoding.LevelFinalizeData
	var err error

	data.XSize, err = reader.Short()
	if err != nil {
		return nil, err
	}

	data.YSize, err = reader.Short()
	if err != nil {
		return nil, err
	}

	data.ZSize, err = reader.Short()
	if err != nil {
		return nil, err
	}

	return &LevelFinalizePacket7{
		id:   protocol.PacketID_LevelFinalize,
		data: data,
	}, nil
}

func (b *levelFinalizeBuilder7) Build(data any) (protocol.Packet, error) {
	return buildPacket[encoding.LevelFinalizeData](data, func(d encoding.LevelFinalizeData) protocol.Packet {
		return &LevelFinalizePacket7{
			id:   protocol.PacketID_LevelFinalize,
			data: d,
		}
	})
}

type SetBlockServerboundPacket7 struct {
	id   protocol.PacketID
	data encoding.SetBlockServerboundData
}

func (p *SetBlockServerboundPacket7) ID() protocol.PacketID {
	return p.id
}

func (p *SetBlockServerboundPacket7) Size() int {
	return 9
}

func (p *SetBlockServerboundPacket7) Data() any {
	return p.data
}

func (p *SetBlockServerboundPacket7) EncodeToWriter(writer *encoding.PacketWriter) error {
	return writeError(
		writer.Byte(byte(p.ID())),
		writer.Short(p.data.X),
		writer.Short(p.data.Y),
		writer.Short(p.data.Z),
		writer.Byte(p.data.Mode),
		writer.Byte(p.data.BlockType),
	)
}

type setBlockServerboundBuilder7 struct{}

func (b *setBlockServerboundBuilder7) GetSize() int {
	return 8
}

func (b *setBlockServerboundBuilder7) BuildFromReader(reader *encoding.PacketReader) (protocol.Packet, error) {
	var data encoding.SetBlockServerboundData
	var err error

	data.X, err = reader.Short()
	if err != nil {
		return nil, err
	}

	data.Y, err = reader.Short()
	if err != nil {
		return nil, err
	}

	data.Z, err = reader.Short()
	if err != nil {
		return nil, err
	}

	data.Mode, err = reader.Byte()
	if err != nil {
		return nil, err
	}

	data.BlockType, err = reader.Byte()
	if err != nil {
		return nil, err
	}

	return &SetBlockServerboundPacket7{
		id:   protocol.PacketID_SetBlockServerbound,
		data: data,
	}, nil
}

func (b *setBlockServerboundBuilder7) Build(data any) (protocol.Packet, error) {
	return buildPacket[encoding.SetBlockServerboundData](data, func(d encoding.SetBlockServerboundData) protocol.Packet {
		return &SetBlockServerboundPacket7{
			id:   protocol.PacketID_SetBlockServerbound,
			data: d,
		}
	})
}

type SetBlockClientboundPacket7 struct {
	id   protocol.PacketID
	data encoding.SetBlockClientboundData
}

func (p *SetBlockClientboundPacket7) ID() protocol.PacketID {
	return p.id
}

func (p *SetBlockClientboundPacket7) Size() int {
	return 8
}

func (p *SetBlockClientboundPacket7) Data() any {
	return p.data
}

func (p *SetBlockClientboundPacket7) EncodeToWriter(writer *encoding.PacketWriter) error {
	return writeError(
		writer.Byte(byte(p.ID())),
		writer.Short(p.data.X),
		writer.Short(p.data.Y),
		writer.Short(p.data.Z),
		writer.Byte(p.data.BlockType),
	)
}

type setBlockClientboundBuilder7 struct{}

func (b *setBlockClientboundBuilder7) GetSize() int {
	return 7
}

func (b *setBlockClientboundBuilder7) BuildFromReader(reader *encoding.PacketReader) (protocol.Packet, error) {
	var data encoding.SetBlockClientboundData
	var err error

	data.X, err = reader.Short()
	if err != nil {
		return nil, err
	}

	data.Y, err = reader.Short()
	if err != nil {
		return nil, err
	}

	data.Z, err = reader.Short()
	if err != nil {
		return nil, err
	}

	data.BlockType, err = reader.Byte()
	if err != nil {
		return nil, err
	}

	return &SetBlockClientboundPacket7{
		id:   protocol.PacketID_SetBlockClientbound,
		data: data,
	}, nil
}

func (b *setBlockClientboundBuilder7) Build(data any) (protocol.Packet, error) {
	return buildPacket[encoding.SetBlockClientboundData](data, func(d encoding.SetBlockClientboundData) protocol.Packet {
		return &SetBlockClientboundPacket7{
			id:   protocol.PacketID_SetBlockClientbound,
			data: d,
		}
	})
}

type SpawnPlayerPacket7 struct {
	id   protocol.PacketID
	data encoding.SpawnPlayerData
}

func (p *SpawnPlayerPacket7) ID() protocol.PacketID {
	return p.id
}

func (p *SpawnPlayerPacket7) Size() int {
	return 74
}

func (p *SpawnPlayerPacket7) Data() any {
	return p.data
}

func (p *SpawnPlayerPacket7) EncodeToWriter(writer *encoding.PacketWriter) error {
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

type spawnPlayerBuilder7 struct{}

func (b *spawnPlayerBuilder7) GetSize() int {
	return 73
}

func (b *spawnPlayerBuilder7) BuildFromReader(reader *encoding.PacketReader) (protocol.Packet, error) {
	var data encoding.SpawnPlayerData
	var err error

	data.PlayerID, err = reader.SByte()
	if err != nil {
		return nil, err
	}

	data.PlayerName, err = reader.String64()
	if err != nil {
		return nil, err
	}

	data.X, err = reader.FShort()
	if err != nil {
		return nil, err
	}

	data.Y, err = reader.FShort()
	if err != nil {
		return nil, err
	}

	data.Z, err = reader.FShort()
	if err != nil {
		return nil, err
	}

	data.Yaw, err = reader.Byte()
	if err != nil {
		return nil, err
	}

	data.Pitch, err = reader.Byte()
	if err != nil {
		return nil, err
	}

	return &SpawnPlayerPacket7{
		id:   protocol.PacketID_SpawnPlayer,
		data: data,
	}, nil
}

func (b *spawnPlayerBuilder7) Build(data any) (protocol.Packet, error) {
	return buildPacket[encoding.SpawnPlayerData](data, func(d encoding.SpawnPlayerData) protocol.Packet {
		return &SpawnPlayerPacket7{
			id:   protocol.PacketID_SpawnPlayer,
			data: d,
		}
	})
}

type SetPositionAndOrientationPacket7 struct {
	id   protocol.PacketID
	data encoding.SetPositionAndOrientationData
}

func (p *SetPositionAndOrientationPacket7) ID() protocol.PacketID {
	return p.id
}

func (p *SetPositionAndOrientationPacket7) Size() int {
	return 10
}

func (p *SetPositionAndOrientationPacket7) Data() any {
	return p.data
}

func (p *SetPositionAndOrientationPacket7) EncodeToWriter(writer *encoding.PacketWriter) error {
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

type setPositionAndOrientationBuilder7 struct{}

func (b *setPositionAndOrientationBuilder7) GetSize() int {
	return 9
}

func (b *setPositionAndOrientationBuilder7) BuildFromReader(reader *encoding.PacketReader) (protocol.Packet, error) {
	var data encoding.SetPositionAndOrientationData
	var err error

	data.PlayerID, err = reader.SByte()
	if err != nil {
		return nil, err
	}

	data.X, err = reader.FShort()
	if err != nil {
		return nil, err
	}

	data.Y, err = reader.FShort()
	if err != nil {
		return nil, err
	}

	data.Z, err = reader.FShort()
	if err != nil {
		return nil, err
	}

	data.Yaw, err = reader.Byte()
	if err != nil {
		return nil, err
	}

	data.Pitch, err = reader.Byte()
	if err != nil {
		return nil, err
	}

	return &SetPositionAndOrientationPacket7{
		id:   protocol.PacketID_SetPositionAndOrientation,
		data: data,
	}, nil
}

func (b *setPositionAndOrientationBuilder7) Build(data any) (protocol.Packet, error) {
	return buildPacket[encoding.SetPositionAndOrientationData](data, func(d encoding.SetPositionAndOrientationData) protocol.Packet {
		return &SetPositionAndOrientationPacket7{
			id:   protocol.PacketID_SetPositionAndOrientation,
			data: d,
		}
	})
}

type PositionAndOrientationUpdatePacket7 struct {
	id   protocol.PacketID
	data encoding.PositionAndOrientationUpdateData
}

func (p *PositionAndOrientationUpdatePacket7) ID() protocol.PacketID {
	return p.id
}

func (p *PositionAndOrientationUpdatePacket7) Size() int {
	return 8
}

func (p *PositionAndOrientationUpdatePacket7) Data() any {
	return p.data
}

func (p *PositionAndOrientationUpdatePacket7) EncodeToWriter(writer *encoding.PacketWriter) error {
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

type positionAndOrientationUpdateBuilder7 struct{}

func (b *positionAndOrientationUpdateBuilder7) GetSize() int {
	return 7
}

func (b *positionAndOrientationUpdateBuilder7) BuildFromReader(reader *encoding.PacketReader) (protocol.Packet, error) {
	var data encoding.PositionAndOrientationUpdateData
	var err error

	data.PlayerID, err = reader.SByte()
	if err != nil {
		return nil, err
	}

	data.ChangeX, err = reader.FByte()
	if err != nil {
		return nil, err
	}

	data.ChangeY, err = reader.FByte()
	if err != nil {
		return nil, err
	}

	data.ChangeZ, err = reader.FByte()
	if err != nil {
		return nil, err
	}

	data.Yaw, err = reader.Byte()
	if err != nil {
		return nil, err
	}

	data.Pitch, err = reader.Byte()
	if err != nil {
		return nil, err
	}

	return &PositionAndOrientationUpdatePacket7{
		id:   protocol.PacketID_PositionAndOrientationUpdate,
		data: data,
	}, nil
}

func (b *positionAndOrientationUpdateBuilder7) Build(data any) (protocol.Packet, error) {
	return buildPacket[encoding.PositionAndOrientationUpdateData](data, func(d encoding.PositionAndOrientationUpdateData) protocol.Packet {
		return &PositionAndOrientationUpdatePacket7{
			id:   protocol.PacketID_PositionAndOrientationUpdate,
			data: d,
		}
	})
}

type PositionUpdatePacket7 struct {
	id   protocol.PacketID
	data encoding.PositionUpdateData
}

func (p *PositionUpdatePacket7) ID() protocol.PacketID {
	return p.id
}

func (p *PositionUpdatePacket7) Size() int {
	return 5
}

func (p *PositionUpdatePacket7) Data() any {
	return p.data
}

func (p *PositionUpdatePacket7) EncodeToWriter(writer *encoding.PacketWriter) error {
	return writeError(
		writer.Byte(byte(p.ID())),
		writer.SByte(p.data.PlayerID),
		writer.FByte(p.data.ChangeX),
		writer.FByte(p.data.ChangeY),
		writer.FByte(p.data.ChangeZ),
	)
}

type positionUpdateBuilder7 struct{}

func (b *positionUpdateBuilder7) GetSize() int {
	return 4
}

func (b *positionUpdateBuilder7) BuildFromReader(reader *encoding.PacketReader) (protocol.Packet, error) {
	var data encoding.PositionUpdateData
	var err error

	data.PlayerID, err = reader.SByte()
	if err != nil {
		return nil, err
	}

	data.ChangeX, err = reader.FByte()
	if err != nil {
		return nil, err
	}

	data.ChangeY, err = reader.FByte()
	if err != nil {
		return nil, err
	}

	data.ChangeZ, err = reader.FByte()
	if err != nil {
		return nil, err
	}

	return &PositionUpdatePacket7{
		id:   protocol.PacketID_PositionUpdate,
		data: data,
	}, nil
}

func (b *positionUpdateBuilder7) Build(data any) (protocol.Packet, error) {
	return buildPacket[encoding.PositionUpdateData](data, func(d encoding.PositionUpdateData) protocol.Packet {
		return &PositionUpdatePacket7{
			id:   protocol.PacketID_PositionUpdate,
			data: d,
		}
	})
}

type OrientationUpdatePacket7 struct {
	id   protocol.PacketID
	data encoding.OrientationUpdateData
}

func (p *OrientationUpdatePacket7) ID() protocol.PacketID {
	return p.id
}

func (p *OrientationUpdatePacket7) Size() int {
	return 4
}

func (p *OrientationUpdatePacket7) Data() any {
	return p.data
}

func (p *OrientationUpdatePacket7) EncodeToWriter(writer *encoding.PacketWriter) error {
	return writeError(
		writer.Byte(byte(p.ID())),
		writer.SByte(p.data.PlayerID),
		writer.Byte(p.data.Yaw),
		writer.Byte(p.data.Pitch),
	)
}

type orientationUpdateBuilder7 struct{}

func (b *orientationUpdateBuilder7) GetSize() int {
	return 3
}

func (b *orientationUpdateBuilder7) BuildFromReader(reader *encoding.PacketReader) (protocol.Packet, error) {
	var data encoding.OrientationUpdateData
	var err error

	data.PlayerID, err = reader.SByte()
	if err != nil {
		return nil, err
	}

	data.Yaw, err = reader.Byte()
	if err != nil {
		return nil, err
	}

	data.Pitch, err = reader.Byte()
	if err != nil {
		return nil, err
	}

	return &OrientationUpdatePacket7{
		id:   protocol.PacketID_OrientationUpdate,
		data: data,
	}, nil
}

func (b *orientationUpdateBuilder7) Build(data any) (protocol.Packet, error) {
	return buildPacket[encoding.OrientationUpdateData](data, func(d encoding.OrientationUpdateData) protocol.Packet {
		return &OrientationUpdatePacket7{
			id:   protocol.PacketID_OrientationUpdate,
			data: d,
		}
	})
}

type DespawnPlayerPacket7 struct {
	id   protocol.PacketID
	data encoding.DespawnPlayerData
}

func (p *DespawnPlayerPacket7) ID() protocol.PacketID {
	return p.id
}

func (p *DespawnPlayerPacket7) Size() int {
	return 2
}

func (p *DespawnPlayerPacket7) Data() any {
	return p.data
}

func (p *DespawnPlayerPacket7) EncodeToWriter(writer *encoding.PacketWriter) error {
	return writeError(
		writer.Byte(byte(p.ID())),
		writer.SByte(p.data.PlayerID),
	)
}

type despawnPlayerBuilder7 struct{}

func (b *despawnPlayerBuilder7) GetSize() int {
	return 1
}

func (b *despawnPlayerBuilder7) BuildFromReader(reader *encoding.PacketReader) (protocol.Packet, error) {
	var data encoding.DespawnPlayerData
	var err error

	data.PlayerID, err = reader.SByte()
	if err != nil {
		return nil, err
	}

	return &DespawnPlayerPacket7{
		id:   protocol.PacketID_DespawnPlayer,
		data: data,
	}, nil
}

func (b *despawnPlayerBuilder7) Build(data any) (protocol.Packet, error) {
	return buildPacket[encoding.DespawnPlayerData](data, func(d encoding.DespawnPlayerData) protocol.Packet {
		return &DespawnPlayerPacket7{
			id:   protocol.PacketID_DespawnPlayer,
			data: d,
		}
	})
}

type MessagePacket7 struct {
	id   protocol.PacketID
	data encoding.MessageData
}

func (p *MessagePacket7) ID() protocol.PacketID {
	return p.id
}

func (p *MessagePacket7) Size() int {
	return 66
}

func (p *MessagePacket7) Data() any {
	return p.data
}

func (p *MessagePacket7) EncodeToWriter(writer *encoding.PacketWriter) error {
	return writeError(
		writer.Byte(byte(p.ID())),
		writer.SByte(p.data.PlayerID),
		writer.String64(p.data.Message),
	)
}

type messageBuilder7 struct{}

func (b *messageBuilder7) GetSize() int {
	return 64
}

func (b *messageBuilder7) BuildFromReader(reader *encoding.PacketReader) (protocol.Packet, error) {
	var data encoding.MessageData
	var err error

	data.PlayerID, err = reader.SByte()
	if err != nil {
		return nil, err
	}

	data.Message, err = reader.String64()
	if err != nil {
		return nil, err
	}

	return &MessagePacket7{
		id:   protocol.PacketID_Message,
		data: data,
	}, nil
}

func (b *messageBuilder7) Build(data any) (protocol.Packet, error) {
	return buildPacket[encoding.MessageData](data, func(d encoding.MessageData) protocol.Packet {
		return &MessagePacket7{
			id:   protocol.PacketID_Message,
			data: d,
		}
	})
}

type DisconnectPlayerPacket7 struct {
	id   protocol.PacketID
	data encoding.DisconnectPlayerData
}

func (p *DisconnectPlayerPacket7) ID() protocol.PacketID {
	return p.id
}

func (p *DisconnectPlayerPacket7) Size() int {
	return 66
}

func (p *DisconnectPlayerPacket7) Data() any {
	return p.data
}

func (p *DisconnectPlayerPacket7) EncodeToWriter(writer *encoding.PacketWriter) error {
	return writeError(
		writer.Byte(byte(p.ID())),
		writer.String64(p.data.DisconnectReason),
	)
}

type disconnectPlayerBuilder7 struct{}

func (b *disconnectPlayerBuilder7) GetSize() int {
	return 65
}

func (b *disconnectPlayerBuilder7) BuildFromReader(reader *encoding.PacketReader) (protocol.Packet, error) {
	var data encoding.DisconnectPlayerData
	var err error

	data.DisconnectReason, err = reader.String64()
	if err != nil {
		return nil, err
	}

	return &DisconnectPlayerPacket7{
		id:   protocol.PacketID_DisconnectPlayer,
		data: data,
	}, nil
}

func (b *disconnectPlayerBuilder7) Build(data any) (protocol.Packet, error) {
	return buildPacket[encoding.DisconnectPlayerData](data, func(d encoding.DisconnectPlayerData) protocol.Packet {
		return &DisconnectPlayerPacket7{
			id:   protocol.PacketID_DisconnectPlayer,
			data: d,
		}
	})
}

type UpdateUserTypePacket7 struct {
	id   protocol.PacketID
	data encoding.UpdateUserTypeData
}

func (p *UpdateUserTypePacket7) ID() protocol.PacketID {
	return p.id
}

func (p *UpdateUserTypePacket7) Size() int {
	return 2
}

func (p *UpdateUserTypePacket7) Data() any {
	return p.data
}

func (p *UpdateUserTypePacket7) EncodeToWriter(writer *encoding.PacketWriter) error {
	return writeError(
		writer.Byte(byte(p.ID())),
		writer.Byte(p.data.UserType),
	)
}

type updateUserTypeBuilder7 struct{}

func (b *updateUserTypeBuilder7) GetSize() int {
	return 1
}

func (b *updateUserTypeBuilder7) BuildFromReader(reader *encoding.PacketReader) (protocol.Packet, error) {
	var data encoding.UpdateUserTypeData
	var err error

	data.UserType, err = reader.Byte()
	if err != nil {
		return nil, err
	}

	return &UpdateUserTypePacket7{
		id:   protocol.PacketID_UpdateUserType,
		data: data,
	}, nil
}

func (b *updateUserTypeBuilder7) Build(data any) (protocol.Packet, error) {
	return buildPacket[encoding.UpdateUserTypeData](data, func(d encoding.UpdateUserTypeData) protocol.Packet {
		return &UpdateUserTypePacket7{
			id:   protocol.PacketID_UpdateUserType,
			data: d,
		}
	})
}
