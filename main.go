package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
    "strings"
    "log"

	"github.com/bwmarrin/discordgo"
	"github.com/SimeonAleksov/reactme/internal/config"
	"github.com/SimeonAleksov/reactme/internal/events"
)

func main() {
	const fileName = "./internal/config/config.json"

	cfg, err := config.ParseConfigFromJSONFile(fileName)
	if err != nil {
		panic(err)
	}

	s, err := discordgo.New("Bot " + cfg.Token)
	if err != nil {
		panic(err)
	}

	s.Identify.Intents = discordgo.MakeIntent(
		discordgo.IntentsGuildMembers |
			discordgo.IntentsGuildMessages)
    var (
        commands = []*discordgo.ApplicationCommand{
            {
                Name: "basic-command",
                // All commands and options must have a description
                // Commands/options without description will fail the registration
                // of the command.
                Description: "Basic command",
            },
            {
                Name:        "basic-command-with-files",
                Description: "Basic command with files",
            },
        }
        commandHandlers = map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
            "basic-command": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
                s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
                    Type: discordgo.InteractionResponseChannelMessageWithSource,
                    Data: &discordgo.InteractionResponseData{
                        Content: "Hey there! Congratulations, you just executed your first slash command",
                    },
                })
            },
            "basic-command-with-files": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
                s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
                    Type: discordgo.InteractionResponseChannelMessageWithSource,
                    Data: &discordgo.InteractionResponseData{
                        Content: "Hey there! Congratulations, you just executed your first slash command with a file in the response",
                        Files: []*discordgo.File{
                            {
                                ContentType: "text/plain",
                                Name:        "test.txt",
                                Reader:      strings.NewReader("Hello Discord!!"),
                            },
                        },
                    },
                })
            },
        }
    )

    ap := &discordgo.Application{}
	ap.Name = "TestApp"
	ap.Description = "TestDesc"
	ap, err = s.ApplicationCreate(ap)
    log.Printf("ApplicationCreate: err: %+v, app: %+v\n", err, ap)

	registerEvents(s)

    s.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		if h, ok := commandHandlers[i.ApplicationCommandData().Name]; ok {
			h(s, i)
		}
	})

	for i, v := range commands {
		cmd, err := s.ApplicationCommandCreate("954468721185935450", "", v)
		if err != nil {
			log.Panicf("Cannot create '%v' command: %v", v.Name, err)
		}
        log.Println(i)
        log.Println(cmd)
	}
	fmt.Println("Bot is now running. Press CTRL-C to exit...")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	s.Close()
}

func registerEvents(s *discordgo.Session) {
	s.AddHandler(events.NewReadyHandler().Handler)
    s.AddHandler(events.NewUserHandler().Handler)
    s.AddHandler(events.NewMessageHandler().Handler)
}


