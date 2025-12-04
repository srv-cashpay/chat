package dto

type WidgetRequest struct {
	ID           string `json:"id"`
	FullName     string `json:"full_name"`
	Email        string `json:"email"`
	Whatsapp     string `json:"whatsapp"`
	BusinessName string `json:"business_name"`
}

type WidgetResponse struct {
	ID           string `json:"id"`
	FullName     string `json:"full_name"`
	Email        string `json:"email"`
	Whatsapp     string `json:"whatsapp"`
	BusinessName string `json:"business_name"`
}
