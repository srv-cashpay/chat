package widget

import (
	"github.com/srv-cashpay/chat/dto"
	"gorm.io/gorm"
)

type DomainRepository interface {
	Create(req dto.WidgetRequest) (dto.WidgetResponse, error)
}

type widgetRepository struct {
	DB *gorm.DB
}

func NewChatWidgetRepository(DB *gorm.DB) DomainRepository {
	return &widgetRepository{
		DB: DB,
	}
}
