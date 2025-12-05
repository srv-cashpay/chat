package widget

import (
	"github.com/srv-cashpay/chat/dto"
)

func (r *widgetRepository) Create(req dto.WidgetRequest) (dto.WidgetResponse, error) {
	if err := r.DB.Create(&req).Error; err != nil {
		return dto.WidgetResponse{}, err
	}
	return dto.WidgetResponse{
		ID:           req.ID,
		FullName:     req.FullName,
		Email:        req.Email,
		Whatsapp:     req.Whatsapp,
		BusinessName: req.BusinessName,
	}, nil
}
