package commands

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"loot-summary/db"
	"loot-summary/db/model"
	"loot-summary/mapper"
	"loot-summary/splitter"
)

func (r *Repository) HandleIndexSessionEntry(s *discordgo.Session, i *discordgo.InteractionCreate) {
	data := i.ApplicationCommandData()

	session := mapper.MapPlainStringToSession(data.Options[0].StringValue())

	repo := &db.Repository{Collection: r.Collection}

	sessionEntry := model.SessionEntry{
		Username: i.Member.User.Username,
		Session:  session,
	}

	err := repo.SaveSession(sessionEntry)

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

	repo := &db.Repository{Collection: r.Collection}
	res := repo.FindAllByUsername(i.Member.User.Username)

	loot := splitter.AggregateTransfers(res)

	err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: fmt.Sprintf(BuildResponseMessage(loot)),
		},
	})
	if err != nil {
		fmt.Println("Error", err)
	}
}
