package webhooks

import (
	"example.com/Quaver/Z/config"
	"fmt"
	"github.com/disgoorg/disgo/discord"
	"github.com/disgoorg/disgo/webhook"
	"log"
	"time"
)

var (
	antiCheat webhook.Client
)

const quaverLogo string = "https://i.imgur.com/DkJhqvT.jpg"
const antiCheatDescription string = "**❌ Anti-cheat Triggered!**"

func Initialize() {
	if antiCheat != nil {
		panic("webhooks already initialized")
	}

	var err error

	antiCheat, err = webhook.NewWithURL(config.Instance.DiscordWebhooks.AntiCheat)

	if err != nil {
		panic(err)
	}

	log.Printf("Initialized anti-cheat webhook: %v\n", antiCheat.ID().String())
}

func SendAntiCheat(username string, url string, icon string, reason string, text string) {
	embed := discord.NewEmbedBuilder().
		SetAuthor(username, url, icon).
		SetDescription(antiCheatDescription).
		SetFields(discord.EmbedField{
			Name:  reason,
			Value: text,
		}).
		SetThumbnail(quaverLogo).
		SetFooter("Quaver", quaverLogo).
		SetTimestamp(time.Now()).
		SetColor(0xFF0000).
		Build()

	_, err := antiCheat.CreateEmbeds([]discord.Embed{embed})

	if err != nil {
		log.Printf("Failed to send anti-cheat webhook: %v\n", err)
	}
}

func SendAntiCheatProcessLog(username string, url string, icon string, processes []string) {
	formatted := ""

	for i, proc := range processes {
		formatted += fmt.Sprintf("**%v. %v**\n", i+1, proc)
	}

	SendAntiCheat(username, url, icon, "Detected Processes", formatted)
}