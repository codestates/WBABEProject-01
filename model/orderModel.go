package model

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

type MenuInfo struct {
	Name     string `bson: "name"`
	Quantity int64  `bson: "quantity"`
}
type User struct {
	Pnum    string `bson: "pnum"`
	Address string `bson: "address"`
}

type Order struct {
	MenuInfo []MenuInfo `bson: "menuinfo"`
	User     User       `bson:	"user"`
	State    int64      `bson: "state"`
	Score    int64      `bson: "score"`
	Review   string     `bson: "review"`
	Time     string     `bson: "time"`
	Number   int64      `bson: "number"`
}

func (u User) Map() bson.M {
	return bson.M{
		"user": bson.M{
			"pnum":    u.Pnum,
			"address": u.Address,
		},
	}
}

func (p *Model) InsertOrder(menuInfo MenuInfo, user User) string {
	t := time.Now()
	time := strconv.Itoa(t.Year()) + "-" + strconv.Itoa(int(t.Month())) + "-" + strconv.Itoa(t.Day())
	number := p.GetOrderByTime(time) + 1
	newData := Order{
		MenuInfo: append([]MenuInfo{}, menuInfo),
		User:     user,
		State:    1,
		Score:    0,
		Review:   "",
		Time:     time,
		Number:   number,
	}
	_, err := p.colOrder.InsertOne(context.TODO(), newData)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Order inserted with Name: %s\n", newData.MenuInfo)
	return "Order Index : " + time + strconv.Itoa(int(number))
}

func (p *Model) GetOrderByTime(time string) int64 {
	filter := bson.D{{Key: "time", Value: time}}
	cursor, err := p.colOrder.CountDocuments(context.TODO(), filter)
	if err != nil {
		panic(err)
	}
	return cursor
}

func (p *Model) UpdateOrderState(time string, number int64, state int64) bool {
	filter := bson.D{{Key: "time", Value: time}, {Key: "number", Value: number}}
	field := bson.M{
		"$set": bson.M{
			"state": state,
		},
	}
	cursor, err := p.colOrder.UpdateOne(context.TODO(), filter, field)
	if err != nil {
		panic(err)
	}
	if cursor.MatchedCount == 1 {
		return true
	}
	return false
}

func (p *Model) GetOrderByUser(user User) []Order {
	filter := User{
		Pnum:    user.Pnum,
		Address: user.Address,
	}
	cursor, err := p.colOrder.Find(context.TODO(), filter.Map())
	if err != nil {
		return nil
	}
	defer cursor.Close(context.TODO())

	var orders []Order
	if err = cursor.All(context.TODO(), &orders); err != nil {
		return nil
	}
	fmt.Println(&orders)
	return orders
}

func (p *Model) AddOrderMenu(user User, menuInfo MenuInfo) bool {
	orders := p.GetOrderByUser(user)
	for _, value := range orders {
		fmt.Println("state : ", value.State)
		if value.State == 3 {
			fmt.Println(p.InsertOrder(menuInfo, user))
			return false
		} else if value.State == 1 || value.State == 2 {
			filter := bson.D{{"time", value.Time}, {"number", value.Number}}
			field := bson.M{
				"$set": bson.M{
					"menuinfo": append(value.MenuInfo, menuInfo),
				},
			}
			_, err := p.colOrder.UpdateOne(context.TODO(), filter, field)
			if err != nil {
				panic(err)
			}
			return true
		}
	}
	return false
}

// func (p *Model)
