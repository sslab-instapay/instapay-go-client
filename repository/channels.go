package repository

import (
	"github.com/instapay-go-client/db"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"log"
	"github.com/instapay-go-client/model"
)

func GetChannelList() ([]model.Channel, error) {

	database, err := db.GetDatabase()
	if err != nil {
		return nil, err
	}

	collection := database.Collection("channels")

	cur, err := collection.Find(context.TODO(), bson.D{})

	if err != nil {
		return nil, err
	}
	var channels []model.Channel

	defer cur.Close(context.Background())
	for cur.Next(context.Background()) {
		var channel model.Channel
		err := cur.Decode(&channel)
		if err != nil {
			log.Fatal(err)
		}
		// To get the raw bson bytes use cursor.Current
		channels = append(channels, channel)
	}

	return channels, nil
}

//TODO 닫힌채널만 하게 쿼리 수정
func GetClosedChannelList() ([]model.Channel, error) {

	database, err := db.GetDatabase()
	if err != nil {
		return nil, err
	}

	filter := bson.M{"channelStatus": model.CLOSED}
	collection := database.Collection("channels")

	cur, err := collection.Find(context.TODO(), filter)

	if err != nil {
		return nil, err
	}
	var channels []model.Channel

	defer cur.Close(context.Background())
	for cur.Next(context.Background()) {
		var channel model.Channel
		err := cur.Decode(&channel)
		if err != nil {
			log.Fatal(err)
		}
		// To get the raw bson bytes use cursor.Current
		channels = append(channels, channel)
	}

	return channels, nil
}
//TODO 열린채 하게 쿼리 수정
func GetOpenedChannelList() ([]model.Channel, error) {

	database, err := db.GetDatabase()
	if err != nil {
		return nil, err
	}

	filter := bson.M{"channelStatus": bson.M{
		"$not": bson.M{
			"$eq": 3,
		},
	} }
	collection := database.Collection("channels")

	cur, err := collection.Find(context.TODO(), filter)

	if err != nil {
		return nil, err
	}
	var channels []model.Channel

	defer cur.Close(context.Background())
	for cur.Next(context.Background()) {
		var channel model.Channel
		err := cur.Decode(&channel)
		if err != nil {
			log.Fatal(err)
		}
		// To get the raw bson bytes use cursor.Current
		channels = append(channels, channel)
	}

	return channels, nil
}


