package controller

import (
	"lecture/go-final/model"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type MenuInfo struct {
	Name     string `bson: "name"`
	Quantity int64  `bson: "quantity"`
}
type InputUser struct {
	Pnum    string `bson: "pnum"`
	Address string `bson: "address"`
}

type InputOrder struct {
	MenuInfo []MenuInfo `bson: "menuinfo"`
	User     InputUser  `bson:	"user"`
}

type InputOrderState struct {
	Time   string `bson: "time"`
	Number int64  `bson:"number"`
	State  int64  `bson: "state"`
}

// InsertOrder godoc
// @Summary call InsertOrder, return ok by json.
// @Description DB에 주문 추가
// @name InsertOrder
// @Accept  json
// @Produce  json
// @Param order body InputOrder true "InputOrder"
// @Router /order/insertOrder [post]
// @Success 200 {object} Controller
func (p *Controller) InsertOrder(c *gin.Context) {
	var form model.Order

	t := time.Now()
	time := strconv.Itoa(t.Year()) + "-" + strconv.Itoa(int(t.Month())) + "-" + strconv.Itoa(t.Day())
	number := p.md.GetOrderByTime(time) + 1
	form.State = 1
	form.Time = time
	form.Number = number
	if err := c.ShouldBind(&form); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	c.JSON(200, p.md.InsertOrder(form))
}

// UpdateOrderState godoc
// @Summary call UpdateOrderState, return ok by json.
// @Description 주문 상태를 변경
// @name UpdateOrderState
// @Accept  json
// @Produce  json
// @Param orderstate body InputOrderState true "InputOrderState"
// @Router /order/updateOrderState [put]
// @Success 200 {object} Controller
func (p *Controller) UpdateOrderState(c *gin.Context) {
	var form model.Order
	if err := c.ShouldBind(&form); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	c.JSON(200, p.md.UpdateOrderState(form))
}

// GetOrderByUser godoc
// @Summary call GetOrderByUser, return ok by json.
// @Description 유저별 주문 내역 조회
// @name GetOrderByUser
// @Accept  json
// @Produce  json
// @Param user query InputUser true "InputUser"
// @Router /order/getOrderByUser [get]
// @Success 200 {object} Controller
func (p *Controller) GetOrderByUser(c *gin.Context) {
	var form model.User
	if err := c.ShouldBind(&form); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	c.JSON(200, p.md.GetOrderByUser(form))
	//To do : 최신순 정렬
}

// AddOrderMenu godoc
// @Summary call AddOrderMenu, return ok by json.
// @Description 유저 추가 주문
// @name AddOrderMenu
// @Accept  json
// @Produce  json
// @Param user body InputOrder true "InputOrder"
// @Router /order/addOrderMenu [put]
// @Success 200 {object} Controller
func (p *Controller) AddOrderMenu(c *gin.Context) {
	var form model.Order
	if err := c.ShouldBind(&form); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	c.JSON(200, p.md.AddOrderMenu(form))
}
