package models

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// NewEvent - required structure of a new Event object
type NewEvent struct {
	Name string   `json:"name" form:"name"`
	Tags []string `json:"tags" form:"tags"`
}

// EventDocument - any user or server event
type EventDocument struct {
	ID           primitive.ObjectID `bson:"_id,omitempty"`
	Name         string             `bson:"name,omitempty" validate:"required,min=2,max=30"`
	Tags         []string           `bson:"tags,omitempty"`
	CreatedDate  string             `bson:"createdDate,omitempty"`
	ModifiedDate string             `bson:"modifiedDate,omitempty"`
}

// AddEvent adds a new Event to the events collection
func (db *DB) AddEvent(event *NewEvent) error {
	ed := EventDocument{
		Name:         event.Name,
		Tags:         event.Tags,
		CreatedDate:  time.Now().UTC().String(),
		ModifiedDate: time.Now().UTC().String(),
	}
	ctx := context.Background()
	eventsCollection := db.Database("lifehubdb").Collection("events")

	res, err := eventsCollection.InsertOne(ctx, ed)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(res)

	return nil
}

// AllEvents returns all events from the events collection
func (db *DB) AllEvents() ([]EventDocument, error) {
	ctx := context.Background()
	eventsCollection := db.Database("lifehubdb").Collection("events")
	cur, err := eventsCollection.Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}
	defer cur.Close(ctx)

	var events []EventDocument
	if err = cur.All(ctx, &events); err != nil {
		log.Fatal(err)
	}

	fmt.Println("events:")
	fmt.Printf("%v\n", events)

	return events, nil
}
