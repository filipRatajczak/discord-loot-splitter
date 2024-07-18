package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Session struct {
	ID        *primitive.ObjectID `bson:"_id,omitempty"`
	Date      string
	StartTime string
	EndTime   string
	Duration  string
	LootType  string
	Loot      int
	Supplies  int
	Balance   int
	Players   []Player
}

type Player struct {
	Name     string
	Loot     int
	Supplies int
	Balance  int
	Damage   int
	Healing  int
}
type Summary struct {
	Message []string
}
