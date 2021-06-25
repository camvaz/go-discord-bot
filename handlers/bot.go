package handlers

import (
	"log"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/camvaz/go-discord-bot/strategies"
	"github.com/camvaz/go-discord-bot/utils"
)

type Bot struct {
	l           *log.Logger
	kingID      string
	victimID    string
	polloID     string
	channelID   string
	guildID     string
	commandFlag string
	kingState   bool
	victimState bool
	polloState  bool
	kingChannel string
}

func NewBot(l *log.Logger, kingID string, victimID string, polloID string, channelID string, guildID string, commandFlag string) *Bot {
	victimState := false
	kingState := false
	polloState := false
	kingChannel := ""
	return &Bot{l, kingID, victimID, polloID, channelID, guildID, commandFlag, victimState, kingState, polloState, kingChannel}
}

func (b *Bot) Log(s string) {
	b.l.Println(s)
}

func (b *Bot) FatalLog(s string, v error) {
	b.l.Fatalf(s, v)
}

func (b *Bot) MessageCreationHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.ChannelID != b.channelID || len(m.Content) == 0 {
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
	
	if val, ok := strategies.SimpleCommandStrategy[splittedCommand[0]]; ok {
		val(s, b.channelID)
		return
	}
	if val, ok := strategies.ArgCommandStrategy[splittedCommand[0]]; ok {
		if len(splittedCommand) <= 1{
			return;
		}
		val(s, b.channelID, splittedCommand[1])
		return
	}
	if val, ok := strategies.VoiceKingCommand[splittedCommand[0]]; ok && m.Author.ID == b.kingID  {
		if len(b.kingChannel) == 0 {
			utils.SendMessage(s,b.channelID, "Bips no estas en un canal de voz nmms")
			return
		}
		val(s, b.guildID, b.kingChannel)
		return
	}

	utils.SendMessage(s, b.channelID, "a")
}

func (b *Bot) VoiceUpdateHandler(s *discordgo.Session, m *discordgo.VoiceStateUpdate) {
	isVictim := m.UserID == b.victimID
	isKing := m.UserID == b.kingID
	isPollo := m.UserID == b.polloID
	if !isVictim && !isKing && !isPollo {
		return
	}

	if isPollo {
		if !b.polloState{
			utils.PlayAudio(s, b.guildID, m.ChannelID, "./media/pollo-greet.ogg")
		}	
		b.polloState = !b.polloState
		return
	}

	var message string
	if isVictim {
		if !b.victimState {
			message = "ola mimir webos mimir \n\nhttps://tenor.com/view/tuca-wevos-huevos-gif-8577692"
			utils.PlayAudio(s, b.guildID, m.ChannelID, "./media/webos.m4a")
		} else {
			message = "adios mimir webos mimir \n\nhttps://tenor.com/view/huevos-eggs-gif-10539909"
		}
		b.victimState = !b.victimState
	}

	if isKing {
		if !b.kingState {
			message = "Llego el rey bips. \n\nhttps://tenor.com/view/clapping-drake-applause-proud-gif-9919565"
			b.kingChannel = m.ChannelID
		} else {
			message = "El rey bips se retira, larga vida al rey bips.\n\nhttps://tenor.com/view/mic-drop-im-out-king-minion-gif-10937564"
			b.kingChannel = ""
		}
		b.kingState = !b.kingState
	}

	utils.SendMessage(s, b.channelID, message)
}
