package widget

import (
	"time"

	"github.com/srv-cashpay/chat/dto"
	"github.com/srv-cashpay/chat/entity"
)

func (r *widgetRepository) Create(req dto.WidgetRequest) (dto.WidgetResponse, error) {
	chat := entity.ChatWidget{
		ID:           req.ID,
		FullName:     req.FullName,
		Email:        req.Email,
		Whatsapp:     req.Whatsapp,
		BusinessName: req.BusinessName,
		CreatedAt:    time.Now(),
	}

	if err := r.DB.Create(&chat).Error; err != nil {
		return dto.WidgetResponse{}, err
	}

	return dto.WidgetResponse{
		ID:           chat.ID,
		FullName:     chat.FullName,
		Email:        chat.Email,
		Whatsapp:     chat.Whatsapp,
		BusinessName: chat.BusinessName,
	}, nil
}
