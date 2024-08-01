package commands

import (
	"loot-summary/splitter"
	"strconv"
)

func BuildResponseMessage(transfers map[splitter.Transfer]int) string {
	output := ""

	for playerName, amount := range transfers {
		output += "``" + playerName.From + "`` should transfer **" + strconv.Itoa(amount) + "** to ``" + playerName.To + "``. " +
			"[transfer " + strconv.Itoa(amount) + " to " + playerName.To + " ]\n"
	}

	output += "\n``Thank you for splitting loot using Loot-Splitter ❤️``"

	return output
}
