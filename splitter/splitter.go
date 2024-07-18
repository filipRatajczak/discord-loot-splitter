package splitter

import (
	"fmt"
	"loot-summary/model"
	"math"
)

type Transfer struct {
	From string
	To   string
}

func SplitLoot(session []model.Session) string {

	var allSessions []Transfer

	for _, a := range session {

		oneSession := Split(&a)

		for k, v := range oneSession {

			val, ok := allSessions["foo"]
			if ok {
				fmt.Println("bob lives in", val)
			} else {
				fmt.Println("bob is not in the map")
			}

			allSessions = append(allSessions, oneSession...)

		}

	}
	return ""
}

func Split(session *model.Session) map[Transfer]int {

	balancePerPlayer := session.Balance / len(session.Players)

	m := make([]model.Player, len(session.Players))

	var transfers map[Transfer]int

	for i, player := range session.Players {
		profit := balancePerPlayer + player.Balance*(-1)
		m[i] = model.Player{
			Name:    player.Name,
			Balance: profit,
		}
	}

	for i := range m {
		if m[i].Balance > 0 {
			continue
		} else {
			for j := range m {
				if math.Abs(float64(m[i].Balance)) >= float64(m[j].Balance) && m[j].Balance > 0 {
					m[i].Balance += m[j].Balance
					transfers[Transfer{
						From: m[i].Name,
						To:   m[j].Name,
					}] = m[j].Balance
					m[j].Balance = 0
				} else if math.Abs(float64(m[i].Balance)) < float64(m[j].Balance) && math.Abs(float64(m[i].Balance)) > 0 {
					m[j].Balance -= int(math.Abs(float64(m[i].Balance)))
					transfers[Transfer{
						From: m[i].Name,
						To:   m[j].Name,
					}] = int(math.Abs(float64(m[i].Balance)))
					m[i].Balance = 0
					break
				}
			}
		}
	}
	return transfers
}
