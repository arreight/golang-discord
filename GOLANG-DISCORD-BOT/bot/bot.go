package bot

import (
	"fmt"
	"golang-discord-bot/config"
	"io"
	"log"

	"github.com/bwmarrin/discordgo"
)

var BotId string
var goBot *discordgo.Session

func Start() {

	goBot, err := discordgo.New("Bot " + config.Token)

	if err != nil {
		fmt.Println(err.Error())
		return
	}
	u, err := goBot.User("@me")

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	BotId = u.ID

	goBot.AddHandler(messageHandler)

	err = goBot.Open()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("Bot is running !")
}

// leest bericht en geeft reactie op bericht
// hier moet de variable BotID gecalled worden

//help func command wordt in een appart bestand gedaan
func messageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	//hiermee negeerd de bot zichzelf
	if m.Author.ID == BotId {
		return
	}

	if m.Content == "ping" {
		_, _ = s.ChannelMessageSend(m.ChannelID, "pong")
		log.Println("ping call")
	}
	if m.Content == "pong" {
		_, _ = s.ChannelMessageSend(m.ChannelID, "ping")
		log.Println("pong call")
	}
	if m.Content == "help" {
		_, _ = s.ChannelMessageSend(m.ChannelID, "Dit is de help functie aanroeping. Voor alle vragen ga naar de programeur van deze code.")
		log.Println("help call")
	}
	if m.Content == "log" {
		log.Println("log call")
		ChannelFileSend(ChannelID, "log", r, io.Reader)(*Message, error)
		return s.ChannelMessageSendComplex(channelID, &MessageSend{File: &File{Name: name, Reader: r}})
	}
}
