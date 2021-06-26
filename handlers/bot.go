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
	kingChannel string
	sessionMap  map[string]string
}

func NewBot(l *log.Logger, kingID string, victimID string, polloID string, channelID string, guildID string, commandFlag string) *Bot {
	kingChannel := ""
	sessionMap := map[string]string{}
	return &Bot{l, kingID, victimID, polloID, channelID, guildID, commandFlag, kingChannel, sessionMap}
}

func (b *Bot) Log(s string) {
	b.l.Println(s)
}

func (b *Bot) FatalLog(s string, v error) {
	b.l.Fatalf(s, v)
}

func (b *Bot) setUserSession(m *discordgo.VoiceStateUpdate){
	_, ok := b.sessionMap[m.UserID] 
	if ok {
		if m.ChannelID == ""{
			delete(b.sessionMap, m.UserID)
		}	
		return
	}
	b.sessionMap[m.UserID] = m.ChannelID
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

	channelID, inChannel := b.sessionMap[m.Author.ID]
	voiceCommand, okVoiceCommand := strategies.VoiceKingCommand[splittedCommand[0]];
	if okVoiceCommand && inChannel {
		voiceCommand(s, b.guildID, channelID)
		return
	} else if okVoiceCommand && !inChannel {
		utils.SendMessage(s,b.channelID, "no estas en un canal de voz nmms we")
		return
	}

	utils.SendMessage(s, b.channelID, "a")
}

func (b *Bot) VoiceUpdateHandler(s *discordgo.Session, m *discordgo.VoiceStateUpdate) {
	b.setUserSession(m)
	utils.PrettyPrint(b.sessionMap)
	var ok bool
	if _, ok = b.sessionMap[b.polloID]; ok && m.UserID == b.polloID{
		utils.PlayAudio(s, b.guildID, m.ChannelID, "./media/pollo-greet.ogg")
		return
	}	

	var message string
	if _, ok = b.sessionMap[b.victimID]; ok && m.UserID == b.victimID{
		message = "ola mimir webos mimir \n\nhttps://tenor.com/view/tuca-wevos-huevos-gif-8577692"
		utils.PlayAudio(s, b.guildID, m.ChannelID, "./media/webos.m4a")
		utils.SendMessage(s, b.channelID,message)
		return
	} else if m.UserID == b.victimID {
		message = "adios mimir webos mimir \n\nhttps://tenor.com/view/huevos-eggs-gif-10539909"
		utils.SendMessage(s, b.channelID,message)
		return
	}

	if _, ok = b.sessionMap[b.kingID]; ok && m.UserID == b.kingID {
		message ="Llego el rey bips. \n\nhttps://tenor.com/view/clapping-drake-applause-proud-gif-9919565"
		utils.SendMessage(s, b.channelID,message)
	} else if m.UserID == b.kingID {
		message = "El rey bips se retira, larga vida al rey bips.\n\nhttps://tenor.com/view/mic-drop-im-out-king-minion-gif-10937564"
		utils.SendMessage(s, b.channelID,message)
	}
}
