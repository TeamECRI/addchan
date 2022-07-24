package main

import (
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")

	if err != nil {
		fmt.Printf("⚠ Discord のトークンを .env から読み込めませんでした: %v\n", err)
		fmt.Printf("ℹ️ シェルからの読み込みを試行します...\n")
		if os.Getenv("TOKEN") == "" {
			fmt.Printf("⚠ トークンが指定されていません。終了します。\n")
			os.Exit(1)
		} else {
			fmt.Printf("✅ Discord のトークンを読み込みました。\n")
		}
	} else {
		fmt.Printf("✅ Discord のトークンを読み込みました。\n")
	}

	dg, err := discordgo.New("Bot " + os.Getenv("TOKEN"))
	if err != nil {
		fmt.Printf("⚠ セッションの作成に失敗しました: %v\n", err)
		return
	}

	dg.AddHandler(messageCreate)

	// In this example, we only care about receiving message events.
	dg.Identify.Intents = discordgo.IntentsGuildMessages

	// Open a websocket connection to Discord and begin listening.
	err = dg.Open()
	if err != nil {
		fmt.Printf("⚠ 接続を開始できませんでした: %v\n", err)
		return
	}

	fmt.Println("✅ Discord に接続しました。")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	dg.Close()
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {

	if m.Author.ID == s.State.User.ID {
		return
	}

	if strings.HasPrefix(m.Content, "!addch ") {
		args := strings.Split(m.Content, " ")
		if m.Content == "!addch " {
			s.ChannelMessageSend(m.ChannelID, "⚠ チャンネル名を指定してください。")
		} else {
			s.GuildChannelCreate(m.GuildID, args[1], 0)
			s.ChannelMessageSend(m.ChannelID, "✅ チャンネル **"+args[1]+"** を作成しました。")
		}
	}
}
