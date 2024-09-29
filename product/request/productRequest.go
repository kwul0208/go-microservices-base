package request

type ProductRequest struct {
	Id          int64  `json:"id"`
	ProductName string `json:"name" binding:"required,min=3"`
	Description string `json:"description" binding:"required,min=3"`
}
