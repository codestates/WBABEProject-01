package controller

import (
	"lecture/go-final/model"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	md *model.Model
}

func NewCTL(rep *model.Model) (*Controller, error) {
	r := &Controller{md: rep}
	return r, nil
}

// InsertMenu godoc
// @Summary call InsertMenu, return ok by json.
// @Description DB에 menu 추가
// @name InsertMenu
// @Accept  json
// @Produce  json
// @Param menuName formData string true "menuName"
// @Param menuIsOrder formData int true "menuIsorder"
// @Param menuQuantity formData string true "menuQuantity"
// @Param menuPrice formData string true "menuPrice"
// @Param menuOrigin formData int true "menuOrigin"
// @Param menuSpicy formData string true "menuSpicy"
// @Router /menu/insertMenu [post]
// @Success 200 {object} Controller
func (p *Controller) InsertMenu(c *gin.Context) {
	menuName := c.PostForm("menuName")         //string
	menuIsOrder := c.PostForm("menuIsOrder")   //bool
	menuQuantity := c.PostForm("menuQuantity") //int
	menuPrice := c.PostForm("menuPrice")       //int
	menuOrigin := c.PostForm("menuOrigin")     //string
	menuSpicy := c.PostForm("menuSpicy")       //int
	str_menuSpicy, _ := strconv.ParseInt(menuSpicy, 10, 64)
	str_menuPrice, _ := strconv.ParseInt(menuPrice, 10, 64)
	str_menuIsOrder, _ := strconv.ParseBool(menuIsOrder)
	str_menuQuantity, _ := strconv.ParseInt(menuQuantity, 10, 64)
	c.JSON(200, p.md.InsertMenu(menuName, str_menuQuantity, str_menuPrice, menuOrigin, str_menuSpicy, str_menuIsOrder))
}

// InsertMenu godoc
// @Summary call DeleteMenu, return ok by json.
// @Description DB에 menu 추가
// @name InsertMenu
// @Accept  json
// @Produce  json
// @Param menuName formData string true "menuName"
// @Router /menu/deletetMenu [put]
// @Success 200 {object} Controller
func (p *Controller) DeleteMenu(c *gin.Context) {
	menuName := c.PostForm("menuName")
	c.JSON(200, gin.H{"DeletedCount : ": p.md.DeleteMenu(menuName)})
}

// UpdateMenu godoc
// @Summary call UpdateMenu, return ok by json.
// @MenuName으로 조회 후, 다른 필드들 값을 업데이트
// @name UpdateMenu
// @Accept  json
// @Produce  json
// @Param menuName formData string true "menuName"
// @Param menuQuantity formData int true "menuQuantity"
// @Param menuPrice formData string true "menuPrice"
// @Param menuOrigin formData int true "menuOrigin"
// @Param menuSpicy formData string true "menuSpicy"
// @Param menuIsOrder formData int true "menuIsOrder"
// @Router /menu/UpdateMenu [put]
// @Success 200 {object} Controller
func (p *Controller) UpdateMenu(c *gin.Context) {
	menuName := c.PostForm("menuName")         //string
	menuIsOrder := c.PostForm("menuIsOrder")   //bool
	menuQuantity := c.PostForm("menuQuantity") //int
	menuPrice := c.PostForm("menuPrice")       //int
	menuOrigin := c.PostForm("menuOrigin")     //string
	menuSpicy := c.PostForm("menuSpicy")       //int

	str_menuSpicy, _ := strconv.ParseInt(menuSpicy, 10, 64)
	str_menuPrice, _ := strconv.ParseInt(menuPrice, 10, 64)
	str_menuIsOrder, _ := strconv.ParseBool(menuIsOrder)
	str_menuQuantity, _ := strconv.ParseInt(menuQuantity, 10, 64)
	c.JSON(200, gin.H{"UpdatedCount : ": p.md.UpdateMenu(menuName, str_menuQuantity, str_menuPrice, menuOrigin, str_menuSpicy, str_menuIsOrder)})
}
