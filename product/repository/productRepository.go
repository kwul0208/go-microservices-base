package repository

import (
	models "github.com/kwul0208/product/model"
	"gorm.io/gorm"
)

type ProductRepository interface {
	GetAll() ([]models.Product, error)
	GetById(Id int) (models.Product, error)
	Create(product models.Product) (models.Product, error)
	Update(Id int, product models.Product) (models.Product, error)
	Delete(product models.Product) (models.Product, error)
}

type productRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) *productRepository {
	return &productRepository{db}
}

func (pr *productRepository) GetAll() ([]models.Product, error) {
	var products []models.Product

	err := pr.db.Find(&products).Error

	return products, err
}

func (pr *productRepository) GetById(Id int) (models.Product, error) {
	var product models.Product

	err := pr.db.Find(&product, Id).Error

	return product, err

}

func (pr *productRepository) Create(product models.Product) (models.Product, error) {
	err := pr.db.Create(&product).Error

	return product, err
}

func (pr *productRepository) Update(Id int, product models.Product) (models.Product, error) {
	err := pr.db.Where("id = ?", Id).Updates(product).Error

	return product, err
}

func (pr *productRepository) Delete(product models.Product) (models.Product, error) {
	err := pr.db.Delete(&product).Error

	return product, err
}
