package req

import (
	"dz4/pkg/res"
	"net/http"
)

type OrderReq struct {
	ProductId uint   `json:"product_id" validate:"required"`
	UserName  string `json:"name" validate:"required"`
	Email     string `json:"email" validate:"required,email"`
}
type ProductReq struct {
	Name        string   `json:"name" validate:"required"`
	Description string   `json:"description" validate:"required"`
	Images      []string `json:"images" validate:"required,dive,url"`
}

func HandleBody[T any](w *http.ResponseWriter, r *http.Request) (*T, error) {
	body, err := Decode[T](r.Body)
	if err != nil {
		res.Json(*w, err.Error(), 402)
		return nil, err
	}
	err = IsValid(body)
	if err != nil {
		res.Json(*w, err.Error(), 402)
		return nil, err
	}
	return &body, nil
}
