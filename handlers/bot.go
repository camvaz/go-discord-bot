// Bot handler
// TODO: refactor ascii and messages
package handlers

import (
	"log"
	"strings"

	"github.com/bwmarrin/dgvoice"
	"github.com/bwmarrin/discordgo"
	"github.com/camvaz/go-discord-bot/strategies"
	"github.com/camvaz/go-discord-bot/utils"
)

type Bot struct {
	l           *log.Logger
	kingID      string
	victimID    string
	channelID   string
	guildID     string
	commandFlag string
	kingState   bool
	victimState bool
}

func NewBot(l *log.Logger, kingID string, victimID string, channelID string, guildID string, commandFlag string) *Bot {
	victimState := false
	kingState := false
	return &Bot{l, kingID, victimID, channelID, guildID, commandFlag, victimState, kingState}
}

func (b *Bot) Log(s string) {
	b.l.Println(s)
}

func (b *Bot) FatalLog(s string, v error) {
	b.l.Fatalf(s, v)
}

func (b *Bot) MessageCreationHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.ChannelID != b.channelID {
		return
	}
	chatCommand := m.Content[0:1]
	if chatCommand != b.commandFlag {
		return
	}
	if len(m.Content) <= 1 {
		utils.SendMessage(s, b.channelID, "ke kieres we ? \n\n")
		return
	}
	chatMessage := m.Content[1:]
	splittedCommand := strings.Split(chatMessage, " ")

	switch len(splittedCommand) {
	case 1:
		if val, ok := strategies.SimpleCommandStrategy[splittedCommand[0]]; ok {
			val(s, b.channelID)
		}
	case 2:
		if val, ok := strategies.ArgCommandStrategy[splittedCommand[0]]; ok {
			val(s, b.channelID, splittedCommand[1])
		}
	default:
		utils.SendMessage(s, b.channelID, "a")
	}
}

func (b *Bot) VoiceUpdateHandler(s *discordgo.Session, m *discordgo.VoiceStateUpdate) {
	b.l.Printf("User: %s\n", m.UserID)
	isVictim := m.UserID == b.victimID
	isKing := m.UserID == b.kingID
	if !isVictim && !isKing {
		return
	}

	var message string
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
	}

	if isKing {
		if !b.kingState {
			message = "Llego el rey bips. \n\nhttps://tenor.com/view/clapping-drake-applause-proud-gif-9919565"
		} else {
			message = "El rey bips se retira, larga vida al rey bips.\n\nhttps://tenor.com/view/mic-drop-im-out-king-minion-gif-10937564"
		}
		b.kingState = !b.kingState
	}

	utils.SendMessage(s, b.channelID, message)
}
