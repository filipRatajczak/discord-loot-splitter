package db

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"loot-summary/db/model"
)

type Repository struct {
	Collection *mongo.Collection
}

func (r *Repository) SaveSession(sessionEntry model.SessionEntry) error {

	_, err := r.Collection.InsertOne(context.TODO(), sessionEntry)
	if err != nil {
		log.Println("Error inserting into the database", sessionEntry, err)
		return err
	}

	return nil
}

func (r *Repository) FindAllByUsername(username string) []model.SessionEntry {

	var res []model.SessionEntry

	filter := bson.D{{"username", username}}

	find, err := r.Collection.Find(context.TODO(), filter)

	if err != nil {
		log.Println("Error finding all hunts", err)
	}

	for find.Next(context.TODO()) {
		each := model.SessionEntry{}
		if err := find.Decode(&each); err != nil {
			log.Fatal(err)
		}
		res = append(res, each)
	}

	return res
}
