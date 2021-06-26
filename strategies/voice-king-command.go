package strategies

import (
	"github.com/bwmarrin/discordgo"
	"github.com/camvaz/go-discord-bot/utils"
)

var VoiceKingCommand map[string]func(s *discordgo.Session, guildID, channelID string) = map[string]func(s *discordgo.Session, guildID, channelID string){
	"voice-mimir": func(s *discordgo.Session, guildID, channelID string) {
		utils.PlayAudio(s, guildID, channelID, "./media/webos.m4a")
	},
	"voice-pollo": func(s *discordgo.Session, guildID, channelID string) {
		utils.PlayAudio(s, guildID, channelID, "./media/pollo-greet.ogg")
	},
} 