package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/srv-cashpay/chat/configs"
	h_widget "github.com/srv-cashpay/chat/handlers/widget"
	r_widget "github.com/srv-cashpay/chat/repository/widget"
	s_widget "github.com/srv-cashpay/chat/services/widget"
	"github.com/srv-cashpay/middlewares/middlewares"
)

var (
	DB      = configs.InitDB()
	JWT     = middlewares.NewJWTService()
	widgetR = r_widget.NewChatWidgetRepository(DB)
	widgetS = s_widget.NewChatWidgetService(widgetR, JWT)
	widgetH = h_widget.NewChatWidgetHandler(widgetS)
)

func New() *echo.Echo {
	e := echo.New()

	widget := e.Group("/api/widget", middlewares.AuthorizeJWT(JWT))
	{
		widget.POST("/create/paid", widgetH.Paid)
		widget.POST("/create/unpaid", widgetH.Unpaid)
		widget.PUT("/update/:id", widgetH.Update)
		widget.GET("/:id", widgetH.GetById)
		widget.GET("/requirement", widgetH.Requirement)
	}

	return e
}
