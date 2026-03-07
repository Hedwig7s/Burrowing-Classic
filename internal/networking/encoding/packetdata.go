package encoding

type IdentificationData struct {
	ProtocolVersion byte
	Name            string
	MotdOrKey       string
	UserType        byte
}

type PingPacketData struct {
}

type LevelInitializeData struct {
}

type LevelDataChunkData struct {
	ChunkLength     int16
	ChunkData       [1024]byte
	PercentComplete byte
}

type LevelFinalizeData struct {
	XSize int16
	YSize int16
	ZSize int16
}

type SetBlockServerboundData struct {
	X         int16
	Y         int16
	Z         int16
	Mode      byte
	BlockType byte
}

type SetBlockClientboundData struct {
	X         int16
	Y         int16
	Z         int16
	BlockType byte
}

type SpawnPlayerData struct {
	PlayerID   int8
	PlayerName string
	X          float32
	Y          float32
	Z          float32
	Yaw        byte
	Pitch      byte
}

type SetPositionAndOrientationData struct {
	PlayerID int8
	X        float32
	Y        float32
	Z        float32
	Yaw      byte
	Pitch    byte
}

type PositionAndOrientationUpdateData struct {
	PlayerID int8
	ChangeX  float32
	ChangeY  float32
	ChangeZ  float32
	Yaw      byte
	Pitch    byte
}

type PositionUpdateData struct {
	PlayerID int8
	ChangeX  float32
	ChangeY  float32
	ChangeZ  float32
}

type OrientationUpdateData struct {
	PlayerID int8
	Yaw      byte
	Pitch    byte
}

type DespawnPlayerData struct {
	PlayerID int8
}

type MessageData struct {
	PlayerID int8
	Message  string
}

type DisconnectPlayerData struct {
	DisconnectReason string
}

type UpdateUserTypeData struct {
	UserType byte
}
