package repository

import (
	"github.com/kwul0208/product/pkg/db"
	models "github.com/kwul0208/product/pkg/model"
)

type ProductRepository interface {
	GetAll() ([]models.Product, error)
	GetOne(Id int) (models.Product, error)
	CreateProduct(product models.Product) (models.Product, error)
}

type productRepository struct {
	db db.Handler
}

func NewProductRepository(db db.Handler) *productRepository {
	return &productRepository{db}
}

func (pr *productRepository) GetAll() ([]models.Product, error) {
	var products []models.Product

	err := pr.db.DB.Find(&products).Error

	return products, err
}

func (pr *productRepository) GetOne(Id int) (models.Product, error) {
	var product models.Product

	err := pr.db.DB.Find(&product, Id).Error

	return product, err

}

func (pr *productRepository) CreateProduct(product models.Product) (models.Product, error) {
	err := pr.db.DB.Create(&product).Error

	return product, err
}
