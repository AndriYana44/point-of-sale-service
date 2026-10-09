package services

import (
	"errors"

	"github.com/AndriYana44/msc-pos-product/models"
	"github.com/AndriYana44/msc-pos-product/repositories"
	"gorm.io/gorm"
)

type ProductService struct {
	repository *repositories.ProductRepository
}

func NewProductService(
	repository *repositories.ProductRepository,
) *ProductService {
	return &ProductService{
		repository: repository,
	}
}

func (s *ProductService) Create(product *models.Product) error {
	return s.repository.Create(product)
}

func (s *ProductService) FindAll() ([]models.Product, error) {
	return s.repository.FindAll()
}

func (s *ProductService) FindByID(id uint) (*models.Product, error) {
	product, err := s.repository.FindByID(id)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("product not found")
		}

		return nil, err
	}

	return product, nil
}

func (s *ProductService) Update(
	id uint,
	data *models.Product,
) (*models.Product, error) {

	product, err := s.repository.FindByID(id)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("product not found")
		}

		return nil, err
	}

	product.SKU = data.SKU
	product.Barcode = data.Barcode
	product.Name = data.Name
	product.CategoryID = data.CategoryID
	product.Price = data.Price
	product.IsActive = data.IsActive

	if err := s.repository.Update(product); err != nil {
		return nil, err
	}

	return product, nil
}

func (s *ProductService) Delete(id uint) error {
	_, err := s.FindByID(id)

	if err != nil {
		return err
	}

	return s.repository.Delete(id)
}
