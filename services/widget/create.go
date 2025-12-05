package widget

import (
	"fmt"

	"github.com/srv-cashpay/chat/dto"
	util "github.com/srv-cashpay/util/s"
)

func (s *widgetService) Create(req dto.WidgetRequest) (dto.WidgetResponse, error) {
	req.ID = util.GenerateRandomString()
	created, err := s.Repo.Create(req)
	if err != nil {
		return dto.WidgetResponse{}, fmt.Errorf("gagal menyimpan %w", err)
	}
	return created, nil
}
