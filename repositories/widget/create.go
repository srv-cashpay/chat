package widget

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
	"time"
	"unicode/utf8"

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

	err := ValidateWidgetInput(req.FullName, req.Email, req.Whatsapp, req.BusinessName)
	if err != nil {
		return dto.WidgetResponse{}, fmt.Errorf("validasi gagal: %w", err)
	}

	if err := r.DB.Create(&chat).Error; err != nil {
		if strings.Contains(err.Error(), "duplicate") || strings.Contains(err.Error(), "UNIQUE") {
			return dto.WidgetResponse{}, fmt.Errorf("the schedule has been made", chat.Email)
		}
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

func ValidateWidgetInput(name, email, whatsapp, business string) error {
	// ----- Name: max 20 chars, only letters -----
	if utf8.RuneCountInString(name) > 20 {
		return errors.New("name maksimal 20 karakter")
	}
	if matched, _ := regexp.MatchString(`^[A-Za-z\s]+$`, name); !matched {
		return errors.New("name hanya boleh huruf")
	}

	// ----- Email: max 25 chars, valid email format -----
	if utf8.RuneCountInString(email) > 25 {
		return errors.New("email maksimal 25 karakter")
	}
	if matched, _ := regexp.MatchString(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`, email); !matched {
		return errors.New("format email tidak valid")
	}

	// ----- Whatsapp: max 13 chars, only numbers -----
	if utf8.RuneCountInString(whatsapp) > 13 {
		return errors.New("whatsapp maksimal 13 karakter")
	}
	if matched, _ := regexp.MatchString(`^\d+$`, whatsapp); !matched {
		return errors.New("whatsapp hanya boleh angka")
	}

	// ----- Business Name: max 30 chars, only letters -----
	if utf8.RuneCountInString(business) > 30 {
		return errors.New("nama bisnis maksimal 30 karakter")
	}
	if matched, _ := regexp.MatchString(`^[A-Za-z\s]+$`, business); !matched {
		return errors.New("nama bisnis hanya boleh huruf")
	}

	return nil
}
