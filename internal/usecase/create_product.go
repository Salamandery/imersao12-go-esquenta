package usecase

import "github.com/salamandery/imersao12-go-esquenta/internal/entity"

type CreateProductInputDto struct {
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

type CreateProductOutputDto struct {
	ID    string
	Name  string
	Price float64
}

type CreateProductUseCase struct {
	ProductRepository entity.ProductRepository
}

func NewCreateProductUseCase(productRepository entity.ProductRepository) *CreateProductUseCase {
	return &CreateProductUseCase{
		ProductRepository: productRepository,
	}
}

func (u *CreateProductUseCase) Execute(Input CreateProductInputDto) (*CreateProductOutputDto, error) {
	product := entity.NewProduct(Input.Name, Input.Price)
	err := u.ProductRepository.Create(product)
	if err != nil {
		return nil, err
	}
	return &CreateProductOutputDto{
		ID:    product.ID,
		Name:  product.Name,
		Price: product.Price,
	}, nil
}
