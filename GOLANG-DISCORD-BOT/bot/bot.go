package bot

import (
	"fmt"
	"golang-discord-bot/config"

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
func messageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	//hiermee negeerd de bot zichzelf
	if m.Author.ID == BotId {
		return
	}

	if m.Content == "ping" {
		_, _ = s.ChannelMessageSend(m.ChannelID, "pong")
	}
	if m.Content == "pong" {
		_, _ = s.ChannelMessageSend(m.ChannelID, "ping")
	}

}
