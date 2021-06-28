package strategies

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/camvaz/go-discord-bot/utils"
)

var ArgCommandStrategy map[string]func(s *discordgo.Session, channelID string, arg string) = map[string]func(s *discordgo.Session, channelID string, arg string){
	"webos": func(s *discordgo.Session, channelID, arg string) {
		utils.SendMessage(s, channelID, fmt.Sprintf("webos %s \n\nhttps://tenor.com/view/tuca-wevos-huevos-gif-8577692\n", arg))
	},
	"funa": func(s *discordgo.Session, channelID, arg string) {
		utils.SendMessage(s, channelID, fmt.Sprintf("webos %s \n\nhttps://tenor.com/view/tuca-wevos-huevos-gif-8577692\n", arg))
	},
}
