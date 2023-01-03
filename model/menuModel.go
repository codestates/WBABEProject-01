package model

import (
	"context"
	"encoding/json"
	"fmt"
	"lecture/go-final/config"
	"sort"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Model struct {
	client    *mongo.Client
	colMenu   *mongo.Collection
	colOrder  *mongo.Collection
	colReview *mongo.Collection
}
type Menu struct {
	Name      string `bson:"name" `
	IsOrder   bool   `bson:"isorder" `
	Quantity  int64  `bson:"quantity"`
	Price     int64  `bson:"price"`
	Origin    string `bson:"origin"`
	Spicy     int64  `bson:"spicy"`
	IsVisible bool   `bson:"isvisible"`
}

func NewModel(config *config.Config) (*Model, error) {
	r := &Model{}

	var err error
	mgUrl := fmt.Sprintf("%v", config.DB["admin"]["host"])
	if r.client, err = mongo.Connect(context.Background(), options.Client().ApplyURI(mgUrl)); err != nil {
		return nil, err
	} else if err := r.client.Ping(context.Background(), nil); err != nil {
		return nil, err
	} else {
		db := r.client.Database("go-final")
		r.colMenu = db.Collection("tMenu")
		r.colOrder = db.Collection("tOrder")
		r.colReview = db.Collection("tReview")
	}

	return r, nil
}

func (p *Model) InsertMenu(menu Menu) Menu {
	newData := Menu{
		Name:      menu.Name,
		IsOrder:   menu.IsOrder,
		Quantity:  menu.Quantity,
		Price:     menu.Price,
		Origin:    menu.Origin,
		Spicy:     menu.Spicy,
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
			"isvisible": false,
		},
	}
	cursor, err := p.colMenu.UpdateOne(context.TODO(), filter, field)
	if err != nil {
		panic(err)
	}
	return cursor.MatchedCount
}

func (p *Model) UpdateMenu(menu Menu) int64 {
	filter := bson.D{{Key: "name", Value: menu.Name}}
	field := bson.M{
		"$set": bson.M{
			"name":      menu.Name,
			"quantity":  menu.Quantity,
			"price":     menu.Price,
			"origin":    menu.Origin,
			"spicy":     menu.Spicy,
			"isorder":   menu.IsOrder,
			"isvisible": menu.IsVisible,
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

	var menus []Menu

	if err = cursor.All(context.TODO(), &menus); err != nil {
		panic(err)
	}

	for _, result := range menus {
		cursor.Decode(&result)
		output, err := json.MarshalIndent(result, "", "    ")
		if err != nil {
			panic(err)
		}
		fmt.Printf("%s\n", output)
	}
	return menus
}

func (p *Model) SortMenu() []interface{} {
	menus := p.GetMenu()
	array := []interface{}{}
	for _, menu := range menus {
		array = append(array, []interface{}{p.GetReviewWithMenu(menu.Name).Grade, menu.Name})
	}
	sort.Slice(array, func(i, j int) bool {
		//Grade의 type이 int64이므로 type assertion을 통해 값에 접근할 때는
		//int가 아닌 int64로 변경해주어야 합니다.
		//그렇지 않으면 firstElementI와 firstElementJ는 int의 zero value가 출력됩니다.
		firstElementI, _ := array[i].([]interface{})[0].(int64)
		firstElementJ, _ := array[j].([]interface{})[0].(int64)
		return firstElementI > firstElementJ
	})
	return array
}
