package commands

import "github.com/bwmarrin/discordgo"

var Commands = []*discordgo.ApplicationCommand{

	{
		Name:        "index-hunt",
		Description: "Indexes a session from Party Hunt Analyzer",
		Type:        discordgo.ChatApplicationCommand,
		Options: []*discordgo.ApplicationCommandOption{
			{
				Type:        discordgo.ApplicationCommandOptionString,
				Name:        "string-option",
				Description: "String option",
				Required:    true,
			},
		},
	},
	{
		Name:        "split-it",
		Description: "asdfasdf",
		Type:        discordgo.ChatApplicationCommand,
	},
}
