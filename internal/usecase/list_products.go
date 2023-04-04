package usecase

import "github.com/elysantos/go-api-messages/internal/entity"

type ListProductsResponse struct {
	ID    string
	Name  string
	Price float64
}

type ListProductUseCase struct {
	ProductRepository entity.ProductRepository
}

func NewListProductUseCase(productRepository entity.ProductRepository) *ListProductUseCase {
	return &ListProductUseCase{ProductRepository: productRepository}
}

func (u *ListProductUseCase) Execute() ([]*ListProductsResponse, error) {
	products, err := u.ProductRepository.FindAll()
	if err != nil {
		return nil, err
	}
	var productsResponse []*ListProductsResponse
	for _, product := range products {
		productsResponse = append(productsResponse, &ListProductsResponse{
			ID:    product.ID,
			Name:  product.Name,
			Price: product.Price,
		})
	}
	return productsResponse, nil
}
