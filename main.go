package main

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"log"
	"loot-summary/commands"
	"loot-summary/db"
	"os"
	"os/signal"
)

var s *discordgo.Session

func init() {
	var err error
	s, err = discordgo.New("Bot " + os.Getenv("DISCORD_BOT_TOKEN"))
	if err != nil {
		log.Fatalf("Invalid bot parameters: %v", err)
	}

}

var (
	collection = db.Connect()
	repo       = &commands.Repository{Collection: collection}

	commandHandlers = map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
		"split-it":   repo.HandleSplit,
		"index-hunt": repo.HandleIndexSessionEntry,
		//"register-team": repo.RegisterHunt,
	}
)

func init() {
	s.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		if h, ok := commandHandlers[i.ApplicationCommandData().Name]; ok {
			h(s, i)
		}
	})
}

func main() {
	s.AddHandler(func(s *discordgo.Session, r *discordgo.Ready) {
		log.Printf("Logged in as: %v#%v", s.State.User.Username, s.State.User.Discriminator)
	})

	err := s.Open()
	if err != nil {
		log.Fatalf("Cannot open the session: %v", err)
	}

	log.Println("Adding commands...")
	registeredCommands := make([]*discordgo.ApplicationCommand, len(commands.Commands))

	for i, v := range commands.Commands {
		fmt.Println(i, v)
		cmd, err := s.ApplicationCommandCreate(s.State.User.ID, "", v)
		if err != nil {
			log.Panicf("Cannot create '%v' command: %v", v.Name, err)
		}
		registeredCommands[i] = cmd
	}

	defer func(s *discordgo.Session) {
		err := s.Close()
		if err != nil {
			fmt.Printf("something bad happened: %v\n", err)
		}
	}(s)

	//guilds := s.State.Guilds
	//fmt.Println(guilds[0].Name)
	//fmt.Println(s.GuildMembers(guilds[0].ID, "", 1000))

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	log.Println("Press Ctrl+C to exit")
	<-stop

}
