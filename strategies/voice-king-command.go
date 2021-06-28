package strategies

import (
	"github.com/bwmarrin/discordgo"
	"github.com/camvaz/go-discord-bot/utils"
)

var VoiceKingCommand map[string]func(s *discordgo.Session, guildID, channelID string, setAudioSession func(string), audioSession map[string]bool) = map[string]func(s *discordgo.Session, guildID, channelID string, setAudioSession func(string), audioSession map[string]bool){
	"voice-mimir": func(s *discordgo.Session, guildID, channelID string, setAudioSession func(string), audioSession map[string]bool) {
		utils.PlayAudio(s, guildID, channelID, "./media/webos.m4a", setAudioSession, audioSession)
	},
	"voice-pollo": func(s *discordgo.Session, guildID, channelID string, setAudioSession func(string), audioSession map[string]bool) {
		utils.PlayAudio(s, guildID, channelID, "./media/pollo-greet.ogg", setAudioSession, audioSession)
	},
}
