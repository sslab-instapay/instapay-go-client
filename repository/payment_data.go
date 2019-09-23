package repository

import (
	"github.com/sslab-instapay/instapay-go-client/model"
	"github.com/sslab-instapay/instapay-go-client/db"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"log"
)

func InsertPaymentData(paymentData model.PaymentData) (model.PaymentData, error) {
	database, err := db.GetDatabase()
	if err != nil {
		return model.PaymentData{}, err
	}

	collection := database.Collection("payments")

	insertResult, err := collection.InsertOne(context.TODO(), paymentData)
	if err != nil {
		return model.PaymentData{}, err
	}

	fmt.Println(insertResult.InsertedID)

	return paymentData, nil
}

func GetPaymentDatasByPaymentId(paymentId int64) ([]model.PaymentData, error) {

	database, err := db.GetDatabase()
	if err != nil {
		return nil, err
	}

	filter := bson.M{
		"paymentId": paymentId,
	}

	collection := database.Collection("payments")

	cur, err := collection.Find(context.TODO(), filter)
	var paymentDatas []model.PaymentData

	defer cur.Close(context.Background())
	for cur.Next(context.Background()) {
		var paymentData model.PaymentData
		err := cur.Decode(&paymentData)
		if err != nil {
			log.Println("Decode Error")
		}
		paymentDatas = append(paymentDatas, paymentData)
	}

	return paymentDatas, nil

}