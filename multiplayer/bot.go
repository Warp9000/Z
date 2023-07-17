package multiplayer

import (
	"example.com/Quaver/Z/chat"
	"example.com/Quaver/Z/sessions"
	"fmt"
)

func InitializeBot() {
	chat.AddPublicMessageHandler(handleMultiplayerCommands)
}

// Handles commands made for multiplayer
func handleMultiplayerCommands(user *sessions.User, channel *chat.Channel, args []string) string {
	if !channel.IsMultiplayer {
		return ""
	}

	if args[0] != "!mp" {
		return ""
	}

	game := GetGameById(user.GetMultiplayerGameId())

	if game == nil {
		return ""
	}

	return fmt.Sprintf("You executed the multiplayer command: %v", args)
}
