package models

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Event represents any user or server event
type Event struct {
	ID           primitive.ObjectID `bson:"_id,omitempty"`
	Name         string             `bson:"name,omitempty" validate:"required,min=2,max=30"`
	Tags         []string           `bson:"tags,omitempty"`
	CreatedDate  string             `bson:"createdDate,omitempty" validate:"required"`
	ModifiedDate string             `bson:"modifiedDate,omitempty" validate:"required"`
}

// AddEvent adds a new Event to the events collection
func (db *DB) AddEvent(event *Event) error {
	ctx := context.Background()
	eventsCollection := db.Database("lifehubdb").Collection("events")

	res, err := eventsCollection.InsertOne(ctx, event)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(res)

	return nil
}

// AllEvents returns all events from the events collection
func (db *DB) AllEvents() ([]Event, error) {
	ctx := context.Background()
	eventsCollection := db.Database("lifehubdb").Collection("events")
	cur, err := eventsCollection.Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}
	defer cur.Close(ctx)

	var events []Event
	if err = cur.All(ctx, &events); err != nil {
		log.Fatal(err)
	}

	fmt.Println("events:")
	fmt.Printf("%v\n", events)

	return events, nil
}
