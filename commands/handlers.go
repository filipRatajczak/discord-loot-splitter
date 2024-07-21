package commands

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"loot-summary/db"
	"loot-summary/mapper"
	"loot-summary/splitter"
	"strconv"
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

	repo := &db.Repository{Collection: r.Collection}
	res := repo.FindAll()

	loot := splitter.AggregateTransfers(res)

	err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: fmt.Sprintf(buildResponseMessage(loot)),
		},
	})
	if err != nil {
		panic(err)
	}
}

func buildResponseMessage(transfers map[splitter.Transfer]int) string {
	output := ""

	for k, v := range transfers {
		output += "``" + k.From + "`` should transfer **" + strconv.Itoa(v) + "** to ``" + k.To + "``. [transfer " + strconv.Itoa(v) + " to " + k.To + " ]\n"
	}

	output += "\n``Loot split used Loot-Splitter discord bot ❤️``"

	return output
}
