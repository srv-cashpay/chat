package widget

import (
	dto "github.com/srv-cashpay/chat/dto"
	m "github.com/srv-cashpay/middlewares/middlewares"

	r "github.com/srv-cashpay/chat/repositories/widget"
)

type WidgetService interface {
	Create(req dto.WidgetRequest) (dto.WidgetResponse, error)
}

type widgetService struct {
	Repo r.DomainRepository
	jwt  m.JWTService
}

func NewChatWidgetService(Repo r.DomainRepository, jwtS m.JWTService) WidgetService {
	return &widgetService{
		Repo: Repo,
		jwt:  jwtS,
	}
}
