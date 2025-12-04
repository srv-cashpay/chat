package widget

import (
	"github.com/labstack/echo/v4"
	"github.com/srv-cashpay/chat/dto"
	res "github.com/srv-cashpay/util/s/response"
)

func (h *domainHandler) Create(c echo.Context) error {
	var req dto.WidgetRequest
	var resp dto.WidgetResponse

	err := c.Bind(&req)
	if err != nil {
		return res.ErrorBuilder(&res.ErrorConstant.BadRequest, err).Send(c)
	}

	resp, err = h.serviceWidget.Create(req)
	if err != nil {
		return res.ErrorResponse(err).Send(c)
	}

	return res.SuccessResponse(resp).Send(c)

}
