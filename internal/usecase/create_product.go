package usecase

import "github.com/elysantos/go-api-messages/internal/entity"

type CreateProductRequest struct {
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

type CreateProductResponse struct {
	ID    string
	Name  string
	Price float64
}

type CreateProductUseCase struct {
	ProductRepository entity.ProductRepository
}

func NewCreateProductUseCase(productRepository entity.ProductRepository) *CreateProductUseCase {
	return &CreateProductUseCase{ProductRepository: productRepository}
}

func (c *CreateProductUseCase) Execute(input CreateProductRequest) (*CreateProductResponse, error) {
	product := entity.NewProduct(input.Name, input.Price)
	err := c.ProductRepository.Create(product)
	if err != nil {
		return nil, err
	}

	return &CreateProductResponse{
		ID:    product.ID,
		Name:  product.Name,
		Price: product.Price,
	}, nil
}
