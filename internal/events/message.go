package events

import (

	"github.com/bwmarrin/discordgo"
)

type MessageHandler struct{}

func NewMessageHandler() *MessageHandler {
	return &MessageHandler{}
}

func (h *MessageHandler) Handler(s *discordgo.Session, e *discordgo.MessageCreate) {
    if e.Content == "ping" {
      s.ChannelMessageSend(e.ChannelID, "Pong")
    }
}
