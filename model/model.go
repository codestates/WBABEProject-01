package model

import (
	"context"
	"encoding/json"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Model struct {
	client  *mongo.Client
	colMenu *mongo.Collection
}
type Menu struct {
	Name      string `bson:"name"`
	IsOrder   bool   `bson:"isOrder"`
	Quantity  int64  `bson:"quantity"`
	Price     int64  `bson:"price"`
	Origin    string `bson:"origin"`
	Spicy     int64  `bson:"spicy"`
	IsVisible bool   `bson:"isvisible"`
}

func NewModel() (*Model, error) {
	r := &Model{}

	var err error
	mgUrl := "mongodb://127.0.0.1:27017"
	if r.client, err = mongo.Connect(context.Background(), options.Client().ApplyURI(mgUrl)); err != nil {
		return nil, err
	} else if err := r.client.Ping(context.Background(), nil); err != nil {
		return nil, err
	} else {
		db := r.client.Database("go-final")
		r.colMenu = db.Collection("tMenu")
	}

	return r, nil
}

// func (p *Model) GetMenuByName(name string) []Menu {
// 	filter := bson.D{{"name", name}}
// 	cursor, err := p.colMenu.Find(context.TODO(), filter)
// 	if err != nil {
// 		panic(err)
// 	}

// 	var pers []Menu

// 	if err = cursor.All(context.TODO(), &pers); err != nil {
// 		panic(err)
// 	}

// 	for _, result := range pers {
// 		cursor.Decode(&result)
// 		output, err := json.MarshalIndent(result, "", "    ")
// 		if err != nil {
// 			panic(err)
// 		}
// 		fmt.Printf("%s\n", output)
// 	}
// 	return pers
// }

func (p *Model) InsertMenu(name string, quantity, price int64, origin string, spicy int64, isorder bool) Menu {
	newData := Menu{
		Name:      name,
		IsOrder:   isorder,
		Quantity:  quantity,
		Price:     price,
		Origin:    origin,
		Spicy:     spicy,
		IsVisible: true,
	}
	_, err := p.colMenu.InsertOne(context.TODO(), newData)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Menu inserted with Name: %s\n", newData.Name)
	return newData
}

func (p *Model) DeleteMenu(menuName string) int64 {
	filter := bson.D{{Key: "name", Value: menuName}}
	field := bson.M{
		"$set": bson.M{
			"isVisible": false,
		},
	}
	cursor, err := p.colMenu.UpdateOne(context.TODO(), filter, field)
	if err != nil {
		panic(err)
	}
	return cursor.MatchedCount
}

func (p *Model) UpdateMenu(name string, quantity, price int64, origin string, spicy int64, isorder bool) int64 {
	filter := bson.D{{Key: "name", Value: name}}
	field := bson.M{
		"$set": bson.M{
			"name":     name,
			"quantity": quantity,
			"price":    price,
			"origin":   origin,
			"spicy":    spicy,
			"isOrder":  isorder,
		},
	}
	cursor, err := p.colMenu.UpdateOne(context.TODO(), filter, field)
	if err != nil {
		panic(err)
	}
	return cursor.MatchedCount
}

func (p *Model) GetMenu() []Menu {
	filter := bson.D{}
	cursor, err := p.colMenu.Find(context.TODO(), filter)
	if err != nil {
		panic(err)
	}

	var pers []Menu

	if err = cursor.All(context.TODO(), &pers); err != nil {
		panic(err)
	}

	for _, result := range pers {
		cursor.Decode(&result)
		output, err := json.MarshalIndent(result, "", "    ")
		if err != nil {
			panic(err)
		}
		fmt.Printf("%s\n", output)
	}
	return pers
}
