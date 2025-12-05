package widget

import (
	s "github.com/srv-cashpay/chat/services/widget"

	"github.com/labstack/echo/v4"
)

type DomainHandler interface {
	Create(c echo.Context) error
}

type domainHandler struct {
	serviceWidget s.WidgetService
}

func NewChatWidgetHandler(service s.WidgetService) DomainHandler {
	return &domainHandler{
		serviceWidget: service,
	}
}
