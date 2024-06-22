package dto

import "github.com/yatoenough/go-url-shortener/internal/lib/api/response"

type URLRequest struct {
	URL string `json:"url" validate:"required,url"`
}

type URLResponse struct {
	response.Response
	Alias string `json:"alias,omitempty"`
}
