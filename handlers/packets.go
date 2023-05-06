package handlers

import (
	"encoding/json"
	"example.com/Quaver/Z/packets"
	"example.com/Quaver/Z/sessions"
	"fmt"
	"log"
	"net"
)

// HandleIncomingPackets Handles incoming messages from clients
func HandleIncomingPackets(conn net.Conn, msg string) {
	user := sessions.GetUserByConnection(conn)

	if user == nil {
		log.Printf("[%v] Received packet while not logged in: %v\n", conn.RemoteAddr(), msg)
		return
	}

	var p packets.Packet

	if err := json.Unmarshal([]byte(msg), &p); err != nil {
		log.Println(err)
		return
	}

	switch p.Id {
	case packets.PacketIdClientPong:
		handleClientPong(user, unmarshalPacket[packets.ClientPong](msg))
	default:
		log.Println(fmt.Errorf("unknown packet: %v", msg))
	}
}

// unmarshalPacket Unmarshal a packet of a specified type
func unmarshalPacket[T any](packet string) *T {
	var data T

	if err := json.Unmarshal([]byte(packet), &data); err != nil {
		log.Println(err)
		return nil
	}

	return &data
}