package main

import (
	"fmt"
	"golang-discord-bot/bot"
	"golang-discord-bot/config"
)

func main() {
	err := config.ReadConfig()
	//!= is als er iets gebeurd wat niet verwerkt kan worden
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	bot.Start()

	<-make(chan struct{})

	return

}
