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
				Name:        "session",
				Description: "Session from Party Hunt Analyzer",
				Required:    true,
			},
		},
	},
	//{
	//	Name:        "register-team",
	//	Description: "Register a team",
	//	Type:        discordgo.ChatApplicationCommand,
	//	Options: []*discordgo.ApplicationCommandOption{
	//		{
	//			Type:        discordgo.ApplicationCommandOptionString,
	//			Name:        "session",
	//			Description: "Session from Party Hunt Analyzer",
	//			Required:    true,
	//		},
	//	},
	//},
	//{
	//	Name:        "register-team",
	//	Description: "Register a team",
	//	Type:        discordgo.ChatApplicationCommand,
	//	Options: []*discordgo.ApplicationCommandOption{
	//		{
	//			Type:        discordgo.ApplicationCommandOptionString,
	//			Name:        "session",
	//			Description: "Session from Party Hunt Analyzer",
	//			Required:    true,
	//		},
	//	},
	//},
	{
		Name:        "split-it",
		Description: "Splits all sessions",
		Type:        discordgo.ChatApplicationCommand,
		Options: []*discordgo.ApplicationCommandOption{
			{
				Type:        discordgo.ApplicationCommandOptionString,
				Name:        "date",
				Description: "Session from Party Hunt Analyzer",
				Required:    true,
			},
		},
	}, {
		Name:        "split-it",
		Description: "Splits all sessions",
		Type:        discordgo.ChatApplicationCommand,
		Options: []*discordgo.ApplicationCommandOption{
			{
				Type:        discordgo.ApplicationCommandOptionString,
				Name:        "session",
				Description: "Session from Party Hunt Analyzer",
				Required:    true,
			},
		},
	},
}
