package main

import (
	"log"
	"os"
	"os/signal"

	"github.com/bwmarrin/discordgo"
	"github.com/camvaz/go-discord-bot/handlers"
	"github.com/joho/godotenv"
)

var s *discordgo.Session

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func init() {
	var err error
	BotToken := os.Getenv("DISCORD_API_TOKEN")
	s, err = discordgo.New("Bot " + BotToken)
	if err != nil {
		log.Fatalf("Invalid bot parameters: %v", err)
	}
}

func main() {
	GeneralChannelID := os.Getenv("DISCORD_GENERAL_CHANNEL")
	VictimID := os.Getenv("DISCORD_MIMIR_ID")
	ChannelID := os.Getenv("DISCORD_GENERAL_CHANNEL")
	GuildID := os.Getenv("DISCORD_GUILD_ID")
	l := log.New(os.Stdout, "ndejous-bot", log.LstdFlags)
	bot := handlers.NewBot(l, VictimID, ChannelID, GuildID)
	generalChannel := &discordgo.Channel{
		ID: GeneralChannelID,
	}

	s.State.ChannelAdd(generalChannel)

	s.AddHandler(func(s *discordgo.Session, r *discordgo.Ready) {
		bot.Log("Session ready")
	})

	s.AddHandler(bot.PresenceHandler)
	s.AddHandler(bot.VoiceUpdateHandler)

	err := s.Open()
	if err != nil {
		bot.FatalLog("Cannot open the session: %v", err)
	}

	defer s.Close()

	stop := make(chan os.Signal)
	signal.Notify(stop, os.Interrupt)
	<-stop
	bot.Log("Gracefully shutdowning")
}
