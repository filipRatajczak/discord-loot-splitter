package commands

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"loot-summary/db"
	"loot-summary/mapper"
	"loot-summary/splitter"
)

func (r *Repository) HandleRegisterHuntSession(s *discordgo.Session, i *discordgo.InteractionCreate) {
	data := i.ApplicationCommandData()

	session := mapper.MapPlainStringToSession(data.Options[0].StringValue())

	repo := &db.Repository{Collection: r.Collection}
	err := repo.SaveSession(session)

	if err != nil {
		return
	}

	err = s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: fmt.Sprintf(
				"Session successfully saved!",
			),
		},
	})
	if err != nil {
		panic(err)
	}

}

func (r *Repository) HandleSplitIt(s *discordgo.Session, i *discordgo.InteractionCreate) {
	//data := i.ApplicationCommandData()

	repo := &db.Repository{Collection: r.Collection}
	res := repo.FindAll()

	splitter.Split(&res[0])

	err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: fmt.Sprintf(
				"You picked asdf autocompletion"),
		},
	})
	if err != nil {
		panic(err)
	}
}
