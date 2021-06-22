package handlers

import (
	"log"
	"time"

	"github.com/bwmarrin/dgvoice"
	"github.com/bwmarrin/discordgo"
)

type Bot struct {
	l           *log.Logger
	kingID      string
	victimID    string
	channelID   string
	guildID     string
	kingState   bool
	victimState bool
}

func NewBot(l *log.Logger, kingID string, victimID string, channelID string, guildID string) *Bot {
	victimState := false
	kingState := false
	return &Bot{l, kingID, victimID, channelID, guildID, victimState, kingState}
}

func (b *Bot) Log(s string) {
	b.l.Println(s)
}

func (b *Bot) FatalLog(s string, v error) {
	b.l.Fatalf(s, v)
}

func (b *Bot) VoiceUpdateHandler(s *discordgo.Session, m *discordgo.VoiceStateUpdate) {
	b.l.Printf("User: %s\n", m.UserID)
	isVictim := m.UserID == b.victimID
	isKing := m.UserID == b.kingID
	if !isVictim && !isKing {
		return
	}

	var message string
	var err error
	var msg *discordgo.Message
	if isVictim {
		if !b.victimState {
			message = "ola mimir webos mimir \n\nhttps://tenor.com/view/tuca-wevos-huevos-gif-8577692"
			dgv, err := s.ChannelVoiceJoin(b.guildID, m.ChannelID, false, true)
			if err != nil {
				b.l.Printf("Error: %v", err)
				return
			}
			dgvoice.PlayAudioFile(dgv, "./media/webos.m4a", make(chan bool))
			defer dgv.Close()
		} else {
			message = "adios mimir webos mimir \n\nhttps://tenor.com/view/huevos-eggs-gif-10539909"
		}
		b.victimState = !b.victimState

		msg, err = s.ChannelMessageSend(
			b.channelID,
			message,
		)
		if err != nil {
			b.l.Printf("Error de mensaje: %v\n", err)
		}

		go func() {
			time.Sleep(15 * time.Second)
			err = s.ChannelMessageDelete(b.channelID, msg.ID)
			if err != nil {
				b.l.Printf("Error al eliminar mensaje: %v\n", err)

			}
		}()
	}

	if isKing {
		if !b.kingState {
			message = "Llego el rey bips. \n\nhttps://tenor.com/view/clapping-drake-applause-proud-gif-9919565"
		} else {
			message = "El rey bips se retira, larga vida al rey bips.\n\nhttps://tenor.com/view/mic-drop-im-out-king-minion-gif-10937564"
		}
		b.kingState = !b.kingState

		msg, err = s.ChannelMessageSend(
			b.channelID,
			message,
		)
		if err != nil {
			b.l.Printf("Error de mensaje: %v\n", err)
		}

		go func() {
			time.Sleep(15 * time.Second)
			err = s.ChannelMessageDelete(b.channelID, msg.ID)
			if err != nil {
				b.l.Printf("Error al eliminar mensaje: %v\n", err)

			}
		}()
	}
}

func (b *Bot) PresenceHandler(s *discordgo.Session, m *discordgo.PresenceUpdate) {
	b.l.Printf("PresenceHandler - User: %s", m.User.Username)
	isMimir := m.User.ID == b.victimID

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
