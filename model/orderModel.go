package model

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

const (
	Ordered          string = "ordered"
	OrderRejected    string = "orderRejected"
	Cooking          string = "cooking"
	Delivering       string = "delivering"
	DeliverCompleted string = "deliverCompleted"
)

type MenuInfo struct {
	Name     string `bson: "name"`
	Quantity int64  `bson: "quantity"`
	Review   string `bson: "review"`
	Grade    int64  `bson: "grade"`
}
type User struct {
	Pnum    string `bson: "pnum"`
	Address string `bson: "address"`
}

type Order struct {
	MenuInfo []MenuInfo `bson: "menuinfo"`
	User     User       `bson:	"user"`
	State    string     `bson: "state"`
	Time     string     `bson: "time"`
	Number   int64      `bson: "number"`
}

func (u User) UserMap() bson.M {
	return bson.M{
		"user": bson.M{
			"pnum":    u.Pnum,
			"address": u.Address,
		},
	}
}

func (p *Model) InsertOrder(order Order) string {
	newData := Order{
		MenuInfo: append([]MenuInfo{}, order.MenuInfo...),
		User:     order.User,
		State:    order.State,
		Time:     order.Time,
		Number:   order.Number,
	}
	_, err := p.colOrder.InsertOne(context.TODO(), newData)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Order inserted with Name: %s\n", newData.MenuInfo)
	return "Order Index : " + order.Time + strconv.Itoa(int(order.Number))
}

func (p *Model) GetOrderByTime(time string) int64 {
	filter := bson.D{{Key: "time", Value: time}}
	cursor, err := p.colOrder.CountDocuments(context.TODO(), filter)
	if err != nil {
		panic(err)
	}
	return cursor
}

func (p *Model) UpdateOrderState(order Order) bool {
	filter := bson.D{{Key: "time", Value: order.Time}, {Key: "number", Value: order.Number}}
	field := bson.M{
		"$set": bson.M{
			"state": order.State,
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
	filter := user
	fmt.Println(filter.Pnum)
	cursor, err := p.colOrder.Find(context.TODO(), filter.UserMap())
	if err != nil {
		return nil
	}

	var orders []Order
	if err = cursor.All(context.TODO(), &orders); err != nil {
		return nil
	}
	fmt.Println(&orders)
	return orders
}

func (p *Model) AddOrderMenu(order Order) bool {
	orders := p.GetOrderByUser(order.User)
	for _, value := range orders {
		fmt.Println("state : ", value.State)
		//1 또는 2의 상태여야 Order가 추가로 가능한 것 같네요.
		//주문 상태를 더 쉽게 알아보기 위해 개선할 수 있을까요?
		if value.State == Ordered || value.State == Cooking {
			filter := bson.D{{"time", value.Time}, {"number", value.Number}}
			field := bson.M{
				"$set": bson.M{
					"menuinfo": append(value.MenuInfo, order.MenuInfo...),
				},
			}
			_, err := p.colOrder.UpdateOne(context.TODO(), filter, field)
			if err != nil {
				panic(err)
			}
			return true
		}
	}
	t := time.Now()
	time := strconv.Itoa(t.Year()) + "-" + strconv.Itoa(int(t.Month())) + "-" + strconv.Itoa(t.Day())
	number := p.GetOrderByTime(time) + 1
	//1, 2가 아닌 주문 상태에서는 새 주문으로 받는 것 같은데, menu 정보가 들어가지 않을 것 같습니다.
	p.InsertOrder(Order{
		MenuInfo: order.MenuInfo,
		User:     order.User,
		State:    1,
		Time:     time,
		Number:   number,
	})
	return false
}
