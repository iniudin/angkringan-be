package response

import (
	"angkringan/pkg/entity"
	"time"
)

type ProductResponse struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Price       float64   `json:"price"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func ToProductResponse(product *entity.Product) *ProductResponse {
	return &ProductResponse{
		ID:          product.ID,
		Name:        product.Name,
		Description: product.Description,
		Price:       product.Price,
		CreatedAt:   product.CreatedAt,
		UpdatedAt:   product.UpdatedAt,
	}
}

func ToProductResponses(products []entity.Product) *[]ProductResponse {
	var productResponses []ProductResponse
	for _, product := range products {
		productResponses = append(
			productResponses,
			ProductResponse{
				ID:          product.ID,
				Name:        product.Name,
				Description: product.Description,
				Price:       product.Price,
				CreatedAt:   product.CreatedAt,
				UpdatedAt:   product.UpdatedAt,
			},
		)
	}
	return &productResponses
}
