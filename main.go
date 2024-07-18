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
	s, err = discordgo.New("Bot " + os.Getenv("TOKEN"))
	if err != nil {
		log.Fatalf("Invalid bot parameters: %v", err)
	}

}

var (
	collection = db.Connect()
	repo       = &commands.Repository{Collection: collection}

	commandHandlers = map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
		"split-it":   repo.HandleSplitIt,
		"index-hunt": repo.HandleRegisterHuntSession,
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
	//
	//collection := db.Connect()
	//
	//repository := db.Repository{Collection: collection}
	//
	//session := mapper.MapPlainStringToSession("Session data: From 2024-04-19, 14:48:52 to 2024-04-19, 17:03:44\nSession: 02:14h\nLoot Type: Leader\nLoot: 4,198,105\nSupplies: 1,029,065\nBalance: 3,169,040\nCiapa Ciapa (Leader)\n    Loot: 4,160,617\n    Supplies: 518,856\n    Balance: 3,641,761\n    Damage: 5,042,586\n    Healing: 1,378,299\nHailey Honeyy\n    Loot: 37,488\n    Supplies: 510,209\n    Balance: -472,721\n    Damage: 6,977,651\n    Healing: 3,630,966")
	//
	//repository.SaveSession(session)
	//
	//repository.FindAll()

	defer s.Close()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	log.Println("Press Ctrl+C to exit")
	<-stop

	log.Println("Gracefully shutting down.")
}
