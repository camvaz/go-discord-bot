package handlers

import (
	"log"
	"os"

	"github.com/bwmarrin/discordgo"
)

type Bot struct {
	l *log.Logger
	victimID string
	channelID string
	guildID string
	victimState bool
}

func NewBot(l *log.Logger, victimID string, channelID string, guildID string) *Bot {
	victimState := false
	return &Bot{l, victimID, channelID, guildID, victimState}
}

func (b *Bot) Log(s string) {
	b.l.Println(s)
}

func (b *Bot) FatalLog(s string, v error){
	b.l.Fatalf(s,v)
}

func (b *Bot) VoiceUpdateHandler(s *discordgo.Session, m *discordgo.VoiceStateUpdate){
	b.l.Printf("Connected: %s. Victim: %s", m.UserID, b.victimID)
	if m.UserID != b.victimID{
		return
	}
	var message string
	if !b.victimState  {
		message = "/tts ola mimir webos mimir"
	} else {
		message = "/tts adios mimir webos mimir"
	}

	b.victimState = !b.victimState
	s.ChannelMessageSend(
		b.channelID,
		message,
	)
}

func (b *Bot) PresenceHandler(s *discordgo.Session, m *discordgo.PresenceUpdate) {
	MimirID := os.Getenv("DISCORD_MIMIR_ID")
	GuildID := os.Getenv("DISCORD_GUILD_ID")
	ChannelID := os.Getenv("DISCORD_GENERAL_CHANNEL")
	b.l.Printf("PresenceHandler - User: %s", m.User.Username)

	isMimir := m.User.ID == MimirID
	if !isMimir {
		return
	}

	var message discordgo.MessageCreate
	isOffline := m.Status == discordgo.Status("offline")
	if isOffline {
		// Handle mimir not online
		message = discordgo.MessageCreate{
			Message: &discordgo.Message{
				GuildID:   GuildID,
				ChannelID: ChannelID,
				Content:   "Mimir se desconectó, no podía saberse, seguiremos vigilantes.",
			},
		}
		s.State.MessageAdd(message.Message)
		return
	}

	isOnline := m.Status == discordgo.Status("online")
	if isOnline {
	message = discordgo.MessageCreate{
		Message: &discordgo.Message{
			GuildID:   GuildID,
			ChannelID: ChannelID,
			Content:   "Se conecto el mimir, preparen sus huevos. https://tenor.com/view/breakfast-lunch-brunch-dinner-eggs-gif-9519822",
		},
	}
	s.State.MessageAdd(message.Message)
	}
}
