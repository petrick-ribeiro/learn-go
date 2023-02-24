package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/spf13/viper"
)

const prefix string = "go-bot"

func main()  {
  // Define Settings
  vp := viper.New()
  vp.SetConfigName("config")
  vp.SetConfigType("yaml")
  vp.AddConfigPath(".")
  err := vp.ReadInConfig()
  if err != nil {
    panic(fmt.Errorf("fatal error config file: %w", err))
  }

  // Creates a new Discord Session.
  dg, err := discordgo.New("Bot " + vp.GetString("discord.token"))
  if err != nil {
    log.Fatal(err)
  }

  // Handle the commands
  dg.AddHandler(func (s *discordgo.Session, m *discordgo.MessageCreate)  {
    if m.Author.ID == s.State.User.ID {
      return
    }

    args := strings.Split(m.Content, ".")

    if args[0] != prefix {
      return
    }

    if args[1] == "ping" {
      s.ChannelMessageSend(m.ChannelID, "Pong!")
    }

    if strings.ToLower(args[1]) == "hello" || strings.ToLower(args[1]) == "hi" {
      s.ChannelMessageSend(m.ChannelID, "Hello, i'm glad to see you!")
    }
  })

  dg.Identify.Intents = discordgo.IntentsAllWithoutPrivileged

  err = dg.Open()
  if err != nil {
    log.Fatal(err)
  }

  defer dg.Close()

  fmt.Println("The Bot is up and running!")

  sc := make(chan os.Signal, 1)
  signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
  <-sc
}
