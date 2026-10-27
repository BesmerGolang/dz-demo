package req

import "net/http"

type OrderReq struct {
	ProductId uint   `json:"product_id" validate:"required"`
	UserName  string `json:"name" validate:"required"`
	Email     string `json:"email" validate:"required,email"`
}
type ProductReq struct {
	Name        string   `json:"name" validate:"required"`
	Description string   `json:"description" validate:"required"`
	Images      []string `json:"images" validate:"required,email"`
}

func HandleBody[T any](w *http.ResponseWriter, r *http.Request) (*T, error) {
	body, err := Decode[T](r.Body)
	if err != nil {
		return nil, err
	}
	err = IsValid(body)
	if err != nil {
		return nil, err
	}
	return &body, nil
}
