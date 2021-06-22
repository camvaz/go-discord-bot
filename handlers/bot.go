package handlers

import (
	"log"
	"os"

	"github.com/bwmarrin/dgvoice"
	"github.com/bwmarrin/discordgo"
)

type Bot struct {
	l           *log.Logger
	victimID    string
	channelID   string
	guildID     string
	victimState bool
}

func NewBot(l *log.Logger, victimID string, channelID string, guildID string) *Bot {
	victimState := false
	return &Bot{l, victimID, channelID, guildID, victimState}
}

func (b *Bot) Log(s string) {
	b.l.Println(s)
}

func (b *Bot) FatalLog(s string, v error) {
	b.l.Fatalf(s, v)
}

func (b *Bot) VoiceUpdateHandler(s *discordgo.Session, m *discordgo.VoiceStateUpdate) {
	if m.UserID != b.victimID {
		return
	}
	var message string

	if !b.victimState {
		message = "ola mimir webos mimir"
		dgv, err := s.ChannelVoiceJoin(b.guildID, m.ChannelID, false, true)
		if err != nil {
			b.l.Printf("Error: %v", err)
			return
		}
		dgvoice.PlayAudioFile(dgv, "./media/webos.m4a", make(chan bool))
		dgv.Close()

	} else {
		message = "adios mimir webos mimir"
	}

	s.ChannelMessageSend(
		b.channelID,
		message,
	)

	b.victimState = !b.victimState
}

func (b *Bot) PresenceHandler(s *discordgo.Session, m *discordgo.PresenceUpdate) {
	MimirID := os.Getenv("DISCORD_MIMIR_ID")
	b.l.Printf("PresenceHandler - User: %s", m.User.Username)
	isMimir := m.User.ID == MimirID

	if !isMimir {
		return
	}

	isOffline := m.Status == discordgo.Status("offline")
	if isOffline {
		s.ChannelMessageSend(
			b.channelID,
			"Mimir se desconectó, no podía saberse, seguiremos vigilantes.",
		)
		return
	}

	isOnline := m.Status == discordgo.Status("online")
	if isOnline {
		s.ChannelMessageSend(
			b.channelID,
			"Se conecto el mimir, preparen sus huevos. https://tenor.com/view/breakfast-lunch-brunch-dinner-eggs-gif-9519822",
		)
		return
	}

}
