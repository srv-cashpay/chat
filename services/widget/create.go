package widget

import (
	"fmt"

	"github.com/srv-cashpay/chat/dto"
	util "github.com/srv-cashpay/util/s"
	res "github.com/srv-cashpay/util/s/response"
)

func (s *widgetService) Create(req dto.WidgetRequest) (dto.WidgetResponse, error) {
	req.ID = util.GenerateRandomString()

	if !util.IsEmailExists(req.Email) {
		return dto.WidgetResponse{}, res.ErrorBuilder(&res.ErrorConstant.RegisterMailNotExists, nil)
	}

	created, err := s.Repo.Create(req)
	if err != nil {
		return dto.WidgetResponse{}, fmt.Errorf("gagal menyimpan %w", err)
	}
	return created, nil
}
