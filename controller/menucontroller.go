package controller

import (
	"lecture/go-final/model"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type MenuInput struct {
	Name     string `bson:"name"`
	IsOrder  bool   `bson:"isorder"`
	Quantity int64  `bson:"quantity"`
	Price    int64  `bson:"price"`
	Origin   string `bson:"origin"`
	Spicy    int64  `bson:"spicy"`
}

type Controller struct {
	md *model.Model
}

func NewCTL(rep *model.Model) (*Controller, error) {
	r := &Controller{md: rep}
	return r, nil
}

// InsertMenu godoc
// @Summary call InsertMenu, return ok by json.
// @Description DB에 menu 추가 []
// @name InsertMenu
// @Accept  json
// @Produce  json
// @Param menu body MenuInput true "menuInput"
// @Router /menu/ [post]
// @Success 200 {object} Controller
func (p *Controller) InsertMenu(c *gin.Context) {
	var form model.Menu
	if err := c.ShouldBind(&form); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	c.JSON(200, p.md.InsertMenu(form))
}

// DeleteMenu godoc
// @Summary call DeleteMenu, return ok by json.
// @Description DB에 menu 추가
// @name InsertMenu
// @Accept  json
// @Produce  json
// @Param menuName body string true "menuName"
// @Router /menu/ [delete]
// @Success 200 {object} Controller
func (p *Controller) DeleteMenu(c *gin.Context) {
	var form string
	if err := c.ShouldBind(&form); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	c.JSON(200, gin.H{"DeletedCount : ": p.md.DeleteMenu(form)})
}

// UpdateMenu godoc
// @Summary call UpdateMenu, return ok by json.
// @MenuName으로 조회 후, 다른 필드들 값을 업데이트
// @name UpdateMenu
// @Accept  json
// @Produce  json
// @Param menu body MenuInput true "menuInput"
// @Router /menu/ [put]
// @Success 200 {object} Controller
func (p *Controller) UpdateMenu(c *gin.Context) {
	var form model.Menu
	if err := c.ShouldBind(&form); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	c.JSON(200, gin.H{"UpdatedCount : ": p.md.UpdateMenu(form)})
}

// GetMenu godoc
// @Summary call GetMenu, return ok by json.
// @MenuName으로 조회 후, 다른 필드들 값을 업데이트
// @name GetMenu
// @Accept  json
// @Produce  json
// @Router /menu/ [get]
// @Success 200 {object} Controller
func (p *Controller) GetMenu(c *gin.Context) {
	c.JSON(200, p.md.GetMenu())
}

// GetMenu godoc
// @Summary call SortMenu, return ok by json.
// @MenuName으로 조회 후, 다른 필드들 값을 업데이트
// @name SortMenu
// @Accept  json
// @Produce  json
// @Router /menu/grade [get]
// @Success 200 {object} Controller
func (p *Controller) SortMenu(c *gin.Context) {
	c.JSON(200, p.md.SortMenu())
}

// GetMenu godoc
// @Summary call RecommendMenu, return ok by json.
// @추천 메뉴 설정
// @name RecommendMenu
// @Accept  json
// @Produce  json
// @Param name query string true "name"
// @Param recommend query bool true "recommend"
// @Router /menu/admin [put]
// @Success 200 {object} Controller
func (p *Controller) RecommendMenu(c *gin.Context) {
	name := c.Query("name")
	recommend, _ := strconv.ParseBool(c.Query("recommend"))
	c.JSON(200, p.md.RecommendMenu(name, recommend))
}
