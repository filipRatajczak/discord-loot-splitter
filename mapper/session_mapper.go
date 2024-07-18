package mapper

import (
	"fmt"
	"loot-summary/model"
	"math"
	"regexp"
	"strconv"
	"strings"
)

func MapPlainStringToSession(sessionRaw string) model.Session {
	var session model.Session

	reSession := regexp.MustCompile(`Session data: From ([^,]+), ([^ ]+) to ([^,]+), ([^ ]+) Session: ([^ ]+) Loot Type: ([^ ]+) Loot: ([\d,]+) Supplies: ([\d,]+) Balance: (-?[\d,]+)`)
	rePlayer := regexp.MustCompile(`(\w+ \w+(?: \(Leader\))?) Loot: ([\d,]+) Supplies: ([\d,]+) Balance: (-?[\d,]+) Damage: ([\d,]+) Healing: ([\d,]+)`)

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
		if len(match) < 7 {
			continue
		}

		player := model.Player{
			Name:     removeLeaderSuffix(match[1]),
			Loot:     toInt(match[2]),
			Supplies: toInt(match[3]),
			Balance:  toInt(match[4]),
			Damage:   toInt(match[5]),
			Healing:  toInt(match[6]),
		}
		session.Players = append(session.Players, player)
	}

	fmt.Println(session)

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

func removeLeaderSuffix(s string) string {

	value := strings.ReplaceAll(s, " (Leader)", "")

	return value
}
