package mapper

import (
	"fmt"
	"loot-summary/db/model"
	"math"
	"regexp"
	"strconv"
	"strings"
)

func MapPlainStringToSession(sessionRaw string) model.Session {
	var session model.Session

	reSession := regexp.MustCompile(`Session data: From ([^,]+), ([^ ]+) to ([^,]+), ([^ ]+) Session: ([^ ]+) Loot Type: ([^ ]+) Loot: ([\d,]+) Supplies: ([\d,]+) Balance: (-?[\d,]+)`)
	rePlayer := regexp.MustCompile(`([A-z ]+?)( \(Leader\))? Loot: ([\d,]+) Supplies: ([-\d,]+) Balance: ([-\d,]+) Damage: ([-\d,]+) Healing: ([\d,]+)`)

	sessionRaw = strings.ReplaceAll(sessionRaw, "\n", " ")
	sessionRaw = strings.ReplaceAll(sessionRaw, "    ", "")

	sessionMatch := reSession.FindStringSubmatch(sessionRaw)
	if len(sessionMatch) < 8 {
		return session
	}

	session.Date = sessionMatch[1]
	session.StartTime = sessionMatch[2]
	session.EndTime = sessionMatch[4]
	session.Duration = sessionMatch[5]
	session.LootType = sessionMatch[6]
	session.Loot = toInt(sessionMatch[7])
	session.Supplies = toInt(sessionMatch[8])
	session.Balance = toInt(sessionMatch[9])

	playerMatches := rePlayer.FindAllStringSubmatch(sessionRaw, -1)

	for _, match := range playerMatches {
		player := model.Player{
			Name:     trimSpaceLeft(match[1]),
			Loot:     toInt(match[3]),
			Supplies: toInt(match[4]),
			Balance:  toInt(match[5]),
			Damage:   toInt(match[6]),
			Healing:  toInt(match[7]),
		}
		session.Players = append(session.Players, player)
	}

	return session
}

func toInt(s string) int {

	value, err := strconv.ParseInt(strings.ReplaceAll(s, ",", ""), 0, 64)
	if err != nil {
		fmt.Println("error, %v", err)
		return math.MaxInt
	}

	return int(value)
}

func trimSpaceLeft(s string) string {

	value := strings.TrimLeft(s, " ")

	return value
}
