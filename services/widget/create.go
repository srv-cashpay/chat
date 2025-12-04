package widget

import (
	"fmt"

	"github.com/srv-cashpay/chat/dto"
	util "github.com/srv-cashpay/util/s"
)

func (s *widgetService) Create(req dto.WidgetRequest) (dto.WidgetResponse, error) {
	chat := dto.WidgetRequest{
		ID:           util.GenerateRandomString(),
		FullName:     req.FullName,
		Email:        req.Email,
		Whatsapp:     req.Whatsapp,
		BusinessName: req.BusinessName,
	}

	// Simpan ke database
	created, err := s.Repo.Create(chat)
	if err != nil {
		return dto.WidgetResponse{}, fmt.Errorf("gagal menyimpan transaksi: %w", err)
	}

	// Return response yang lebih lengkap
	return dto.WidgetResponse{
		ID:           created.ID,
		FullName:     created.FullName,
		Email:        created.Email,
		Whatsapp:     created.Whatsapp,
		BusinessName: created.BusinessName,
	}, nil
}
