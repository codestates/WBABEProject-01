package controller

import (
	"lecture/go-final/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

type InputGrade struct {
	OrderTime   string     `bson: "time"`
	OrderNumber int64      `bson: "number"`
	MenuName    string     `bson: "name"`
	User        model.User `bson: "user"`
	Review      string     `bson: "review"`
	Grade       int64      `bson: "grade"`
}

// UpdateMenuGrade godoc
// @Summary call UpdateMenuGrade, return ok by json.
// @Description 메뉴의 평점 및 리뷰를 남김
// @name UpdateMenuGrade
// @Accept  json
// @Produce  json
// @Param review body InputGrade true "InputGrade"
// @Router /review/reviewer [put]
// @Success 200 {object} Controller
func (p *Controller) UpdateMenuGrade(c *gin.Context) {
	var form model.Review
	if err := c.ShouldBind(&form); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	c.JSON(200, p.md.UpdateMenuGrade(form))
}

// GetReviewWithMenu godoc
// @Summary call GetReviewWithMenu, return ok by json.
// @Description 해당 메뉴의 평균 평점과 리뷰를 조회
// @name GetReviewWithMenu
// @Accept  json
// @Produce  json
// @Param menuname query string true "menuName"
// @Router /review/reviewer [get]
// @Success 200 {object} Controller
func (p *Controller) GetReviewWithMenu(c *gin.Context) {
	form := c.Query("menuname")
	if err := c.ShouldBind(&form); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	c.JSON(200, p.md.GetReviewWithMenu(form))
}
