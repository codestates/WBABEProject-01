package controller

import (
	"lecture/go-final/model"
	"strconv"

	"github.com/gin-gonic/gin"
)

// 주문자 주문
func (p *Controller) InsertOrder(c *gin.Context) {
	menuName := c.PostForm("menuName")
	menuQuantity := c.PostForm("menuQuantity")
	orderUserPnum := c.PostForm("orderUserPnum")
	orderUserAddress := c.PostForm("orderUserAddress")
	strQuantity, _ := strconv.ParseInt(menuQuantity, 10, 64)

	c.JSON(200, p.md.InsertOrder(model.MenuInfo{menuName, strQuantity}, model.User{orderUserPnum, orderUserAddress}))
}

// 피주문자 주문 상태 변경
func (p *Controller) UpdateOrderState(c *gin.Context) {
	orderTime := c.PostForm("orderTime")
	orderCount := c.PostForm("orderCount")
	orderState := c.PostForm("orderState")
	strCount, _ := strconv.ParseInt(orderCount, 10, 64)
	strState, _ := strconv.ParseInt(orderState, 10, 64)
	c.JSON(200, p.md.UpdateOrderState(orderTime, strCount, strState))
}

// 유저 주문내역 조회
func (p *Controller) GetOrderByUser(c *gin.Context) {
	orderUserPnum := c.Query("orderUserPnum")
	orderUserAddress := c.Query("orderUserAddress")
	//To do : 최신순 정렬
	c.JSON(200, p.md.GetOrderByUser(model.User{orderUserPnum, orderUserAddress}))
}

// 유저 추가주문 및 주문 변경
func (p *Controller) AddOrderMenu(c *gin.Context) {
	orderUserPnum := c.PostForm("orderUserPnum")
	orderUserAddress := c.PostForm("orderUserAddress")
	menuName := c.PostForm("menuName")
	menuQuantity := c.PostForm("menuQuantity")

	strQuantity, _ := strconv.ParseInt(menuQuantity, 10, 64)
	c.JSON(200, p.md.AddOrderMenu(model.User{orderUserPnum, orderUserAddress}, model.MenuInfo{menuName, strQuantity}))
}
