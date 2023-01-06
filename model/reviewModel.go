package model

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
)

type userReview struct {
	Grade  int64    `bson: "grade"`
	Review []string `bson: "reviews"`
}

type Review struct {
	OrderTime   string `bson: "time"`
	OrderNumber int64  `bson: "number"`
	MenuName    string `bson: "name"`
	User        User   `bson: "user"`
	Review      string `bson: "review"`
	Grade       int64  `bson: "grade"`
}

func (p *Model) UpdateMenuGrade(review Review) bool {
	filter := bson.M{
		"menuinfo.name": review.MenuName,
		"user": bson.M{
			"pnum":    review.User.Pnum,
			"address": review.User.Address,
		},
		"time":   review.OrderTime,
		"number": review.OrderNumber,
		"state":  4,
	}
	cursor, err := p.colOrder.Find(context.TODO(), filter)
	var reviews []Review

	if err = cursor.All(context.TODO(), &reviews); err != nil {
		panic(err)
	}

	if cursor != nil {

		field := bson.M{
			"time":   review.OrderTime,
			"number": review.OrderNumber,
			"name":   review.MenuName,
			"user":   review.User,
			"review": review.Review,
			"grade":  review.Grade,
		}

		_, err := p.colReview.InsertOne(context.TODO(), field)
		if err != nil {
			panic(err)
		}
		return true
	}
	return false
}

func (p *Model) GetReviewWithMenu(menuName string) userReview {
	count := 0
	var totalGrade int64 = 0
	filter := bson.D{{"name", menuName}, {"grade", bson.D{{"$gt", 0}}}}
	cursor, err := p.colReview.Find(context.TODO(), filter)
	var reviews []Review
	if err = cursor.All(context.TODO(), &reviews); err != nil {
		panic(err)
	}
	var userReview userReview
	for _, review := range reviews {
		totalGrade += review.Grade
		userReview.Review = append(userReview.Review, review.Review)
		count += 1
	}
	if count != 0 {
		userReview.Grade = totalGrade / int64(count)
	} else {
		userReview = userReview
		return userReview
	}
	return userReview
}
