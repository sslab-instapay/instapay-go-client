package repository

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"log"
	"github.com/sslab-instapay/instapay-go-client/model"
	"github.com/sslab-instapay/instapay-go-client/db"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

func GetChannelIdList() ([]string, error) {

	database, err := db.GetDatabase()
	if err != nil {
		return nil, err
	}

	collection := database.Collection("channels")

	cur, err := collection.Find(context.TODO(), bson.D{})

	if err != nil {
		return nil, err
	}
	var channelIds []string

	defer cur.Close(context.Background())
	for cur.Next(context.Background()) {
		var channel model.Channel
		err := cur.Decode(&channel)
		if err != nil {
			log.Fatal(err)
		}
		// To get the raw bson bytes use cursor.Current
		channelIds = append(channelIds, channel.ChannelId.String())
	}

	return channelIds, nil
}

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

func GetChannelById(channelId primitive.ObjectID) (model.Channel, error){

	database, err := db.GetDatabase()
	if err != nil {
		return model.Channel{}, err
	}

	filter := bson.M{
		"_id": channelId,
	}

	collection := database.Collection("channels")

	channel := model.Channel{}
	singleRecord := collection.FindOne(context.TODO(), filter)
	if err := singleRecord.Decode(&channel); err != nil{
		log.Fatal(err)
	}
	return channel, nil
}
