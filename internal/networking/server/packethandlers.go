package server

import (
	"github.com/Hedwig7s/Burrowing-Classic/internal/cerror"
	"github.com/Hedwig7s/Burrowing-Classic/internal/networking/encoding"
	"github.com/Hedwig7s/Burrowing-Classic/internal/networking/protocol"
)

const idMismatch = "Packet is not %s. Packet ID: %d"
const dataMismatch = "Packet data is not %s. Got %v"

const (
	PACKETHANDLER_ID_MISMATCH = iota
	PACKETHANDLER_DATA_MISMATCH
)

type PacketHandler func(connection *Connection, packet protocol.Packet) error

var PacketHandlers = map[protocol.PacketID]PacketHandler{
	protocol.PacketID_Identification: func(connection *Connection, packet protocol.Packet) error {
		if packet.ID() != protocol.PacketID_Identification {
			return cerror.NewErrorf(PACKETHANDLER_ID_MISMATCH, idMismatch, "Identification", packet.ID())
		}
		var ok bool
		var data encoding.IdentificationData
		if data, ok = packet.Data().(encoding.IdentificationData); !ok {
			return cerror.NewErrorf(PACKETHANDLER_DATA_MISMATCH, dataMismatch, "IdentificationData", packet)
		}

		identification_data := encoding.IdentificationData{
			ProtocolVersion: data.ProtocolVersion,
			Name:            "Burrowing Classic",
			MotdOrKey:       "Where we're going, we don't need a motd.", // TODO: Add config
			UserType:        data.UserType,
		}
		identification_builder, err := connection.Protocol().CreatePacketBuilder(protocol.PacketID_Identification)

		if err != nil {
			return err
		}

		identification, err := identification_builder.Build(identification_data)
		if err != nil {
			return err
		}
		connection.Write(identification)
		b, err := connection.Protocol().CreatePacketBuilder(protocol.PacketID_LevelInitialize)
		if err != nil {
			return err
		}
		level_initialize, err := b.Build(encoding.LevelInitializeData{})
		if err != nil {
			return err
		}
		connection.Write(level_initialize)
		return nil
	}, /*
		PacketID_SetBlockServerbound: func(connection *Connection, packet Packet) error {
			if packet.ID() != Protocol.PacketID_SetBlockServerbound {
				return cerror.NewErrorf(PACKETHANDLER_ID_MISMATCH, idMismatch, "SetBlockServerbound", packet.ID())
			}
			var ok bool
			var data encoding.SetBlockServerboundData
			if data, ok = packet.Data().(encoding.SetBlockServerboundData); !ok {
				return cerror.NewErrorf(PACKETHANDLER_DATA_MISMATCH, dataMismatch, "SetBlockServerboundData", packet)
			}
			return nil
		},
		Protocol.PacketID_SetPositionAndOrientation: func(connection *Connection, packet Packet) error {
			if packet.ID() != Protocol.PacketID_SetPositionAndOrientation {
				return cerror.NewErrorf(PACKETHANDLER_ID_MISMATCH, idMismatch, "SetPositionAndOrientation", packet.ID())
			}
			var ok bool
			var data encoding.SetPositionAndOrientationData
			if data, ok = packet.Data().(encoding.SetPositionAndOrientationData); !ok {
				return cerror.NewErrorf(PACKETHANDLER_DATA_MISMATCH, dataMismatch, "SetPositionAndOrientationData", packet)
			}
			return nil
		},
		Protocol.PacketID_Message: func(connection *Connection, packet Packet) error {
			if packet.ID() != Protocol.PacketID_Message {
				return cerror.NewErrorf(PACKETHANDLER_ID_MISMATCH, idMismatch, "Message", packet.ID())
			}
			var ok bool
			var data encoding.MessageData
			if data, ok = packet.Data().(encoding.MessageData); !ok {
				return cerror.NewErrorf(PACKETHANDLER_DATA_MISMATCH, dataMismatch, "MessageData", packet)
			}
			return nil
		},*/
}

func HandlePacket(connection *Connection, packet protocol.Packet) error {
	handler := PacketHandlers[packet.ID()]
	if handler == nil {
		return cerror.NewErrorf(PACKETHANDLER_ID_MISMATCH, idMismatch, "", packet.ID())
	}
	return handler(connection, packet)
}
