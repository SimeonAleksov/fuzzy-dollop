package events

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

type UserHandler struct{}

func NewUserHandler() *UserHandler {
	return &UserHandler{}
}

func (h *UserHandler) Handler(s *discordgo.Session, e *discordgo.Ready) {
  roleID := "954475354326700052"
  for _, guild := range s.State.Guilds {
    fmt.Printf("Guild: %s, %s \n", guild.ID, guild.Name)
    members, err := s.GuildMembers(guild.ID, "", 100)
    if err != nil {
      panic(err)
    }
    for _, member := range members {
      fmt.Printf("User(id=%s, username=%s) \n", member.User.Username, member.User.ID)
      s.GuildMemberRoleAdd(guild.ID, member.User.ID, roleID)
    }
  }
}
