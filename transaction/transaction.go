package transaction

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type ITransaction interface {
	Log(Name string, Type string, Detail string, Url string, User string, Idrm string, action string) error
}

type Transaction struct {
	collection *mongo.Collection
}

func NewTransaction(mongoDB *mongo.Database) ITransaction {
	return &Transaction{collection: mongoDB.Collection("logs")}
}

func (t *Transaction) Log(Name string, Type string, Detail string, Url string, User string, Idrm string, action string) error {

	// create log in MongoDB
	doc := bson.D{{Key: "time", Value: time.Now().Format("2-Jan-06 03:04PM")}, {Key: "User", Value: User}, {Key: "id", Value: Idrm}, {Key: "action", Value: action}}

	// เพิ่มฟิลด์ Detail และ Url เฉพาะเมื่อมีการอัปเดต
	if Name != "" {
		doc = append(doc, bson.E{Key: "Name", Value: Name})
	}

	if Type != "" {
		doc = append(doc, bson.E{Key: "Type", Value: Type})
	}

	if Detail != "" {
		doc = append(doc, bson.E{Key: "Detail", Value: Detail})
	}
	if Url != "" {
		doc = append(doc, bson.E{Key: "Url", Value: Url})
	}

	_, err := t.collection.InsertOne(context.TODO(), doc)
	if err != nil {
		log.Println(err.Error())
		return err
	}
	return nil
}
