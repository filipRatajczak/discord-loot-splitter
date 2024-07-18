package db

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"loot-summary/model"
)

type Repository struct {
	Collection *mongo.Collection
}

func (r *Repository) SaveSession(session model.Session) error {

	fmt.Println(r.Collection)

	_, err := r.Collection.InsertOne(context.TODO(), session)
	if err != nil {
		log.Println("Error inserting into the database", err)
		return err
	}
	log.Println("Session successfully inserted")

	return nil
}

func (r *Repository) FindAll() []model.Session {

	var res []model.Session

	find, err := r.Collection.Find(context.TODO(), bson.D{{}})
	if err != nil {
		log.Println("Error finding all hunts", err)
	}

	for find.Next(context.TODO()) {
		each := model.Session{}
		if err := find.Decode(&each); err != nil {
			log.Fatal(err)
		}

		res = append(res, each)

		fmt.Printf("%+v\n", each)
	}

	return res
}
