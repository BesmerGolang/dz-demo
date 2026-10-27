package req

type OrderReq struct {
	ProductId uint   `json:"product_id" binding:"required"`
	UserName  string `json:"name" binding:"required,min=2"`
	Email     string `json:"email" binding:"required,email"`
}
type ProductReq struct {
	Name        string   `json:"name" binding:"required,min=3"`
	Description string   `json:"description" binding:"required"`
	Images      []string `json:"images" binding:"required"`
}
