package router

import (
	"fmt"
	ctl "lecture/go-final/controller"

	"lecture/go-final/docs"

	"github.com/gin-gonic/gin"
	swgFiles "github.com/swaggo/files"
	ginSwg "github.com/swaggo/gin-swagger"
	"golang.org/x/sync/errgroup"
)

type Router struct {
	ct *ctl.Controller
}

func NewRouter(ctl *ctl.Controller) (*Router, error) {
	r := &Router{ct: ctl} //controller 포인터를 ct로 복사, 할당

	return r, nil
}

var (
	g errgroup.Group
)

// cross domain을 위해 사용
func CORS() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		//허용할 header 타입에 대해 열거
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, X-Forwarded-For, Authorization, accept, origin, Cache-Control, X-Requested-With")
		//허용할 method에 대해 열거
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	}
}

// 임의 인증을 위한 함수
func liteAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c == nil {
			c.Abort() // 미들웨어에서 사용, 이후 요청 중지
			return
		}
		//http 헤더내 "Authorization" 폼의 데이터를 조회
		auth := c.GetHeader("Authorization")
		//실제 인증기능이 올수있다. 단순히 출력기능만 처리 현재는 출력예시
		fmt.Println("Authorization-word ", auth)

		c.Next() // 다음 요청 진행
	}
}

// 실제 라우팅
func (p *Router) Idx() *gin.Engine {
	//~생략
	e := gin.Default()
	e.Use(gin.Logger())
	e.Use(gin.Recovery())
	e.Use(CORS())

	e.GET("/health")
	e.GET("/swagger/:any", ginSwg.WrapHandler(swgFiles.Handler))
	docs.SwaggerInfo.Host = "localhost:8080" //swagger 정보 등록

	menu := e.Group("/menu", liteAuth())
	{
		fmt.Println("1")
		// account.GET("/getPersonWithName", p.ct.GetPersonByName) // controller 패키지의 실제 처리 함수
		// account.GET("/getPersonWithPnum", p.ct.GetPersonByPnum)
		menu.POST("/insertMenu", p.ct.InsertMenu)
		menu.PUT("/deleteMenu", p.ct.DeleteMenu)
		menu.PUT("/updateMenu", p.ct.UpdateMenu)
	}
	order := e.Group("/order", liteAuth())
	{
		fmt.Println("2")
		order.POST("/insertOrder", p.ct.InsertOrder)
		order.GET("/getOrderByUser", p.ct.GetOrderByUser)
		order.PUT("/addOrderMenu", p.ct.AddOrderMenu)
		order.PUT("/updateOrderState", p.ct.UpdateOrderState)
	}
	return e
}
