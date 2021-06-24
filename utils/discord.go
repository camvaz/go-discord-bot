package utils

import (
	"log"
	"time"

	"github.com/bwmarrin/dgvoice"
	"github.com/bwmarrin/discordgo"
)

func SendMessage(s *discordgo.Session, channelID string, message string) {
	var err error
	var msg *discordgo.Message
	msg, err = s.ChannelMessageSend(
		channelID,
		message,
	)
	if err != nil {
		log.Printf("Error de mensaje: %v\n", err)
	}

	go func() {
		time.Sleep(15 * time.Second)
		err = s.ChannelMessageDelete(channelID, msg.ID)
		if err != nil {
			log.Printf("Error al eliminar mensaje: %v\n", err)
		}
	}()
}

func PlayAudio(s *discordgo.Session, guildID string, channelID string, track string) {

	dgv, err := s.ChannelVoiceJoin(guildID, channelID, false, true)
	if err != nil {
		log.Printf("Error: %v", err)
		return
	}
	dgvoice.PlayAudioFile(dgv, track, make(chan bool))
	dgv.Disconnect()
	dgv.Close()
}
