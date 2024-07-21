package splitter

import (
	"loot-summary/model"
	"math"
)

type Transfer struct {
	From string
	To   string
}

func AggregateTransfers(session []model.Session) map[Transfer]int {
	allSessions := make(map[Transfer]int)
	for _, a := range session {
		oneSession := CreateTransfersFromSession(&a)
		for k, v := range oneSession {
			if _, ok := allSessions[k]; ok {
				allSessions[k] += v
			} else {
				allSessions[k] = v
			}
		}

	}
	return allSessions
}

func CreateTransfersFromSession(session *model.Session) map[Transfer]int {

	balancePerPlayer := session.Balance / len(session.Players)

	m := make([]model.Player, len(session.Players))

	transfers := make(map[Transfer]int)

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
