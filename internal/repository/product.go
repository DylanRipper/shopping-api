package repository

import (
	"shopping-api/internal/dto"
	"shopping-api/internal/model"

	"gorm.io/gorm"
)

type Product interface {
	CreateProduct(product *model.Products) (*model.Products, error)
	GetAllProducts() (interface{}, error)
	GetProductByID(id uint) (*dto.GetProduct, error)
	GetProductsBySubcategoryID(id int) (interface{}, error)
	GetProductsByUserID(id int) (interface{}, error)
	DeleteProduct(id int)
	DeleteProductByID(id int) (interface{}, error)
	GetProductOwner(id int) (int, error)
	GetProductsByName(product string) (interface{}, error)
}

type product struct {
	DB *gorm.DB
}

func NewProductRepository(db *gorm.DB) *product {
	return &product{
		DB: db,
	}
}

// Fungsi untuk membuat menyewakan produk baru
func (p *product) CreateProduct(product *model.Products) (*model.Products, error) {
	if err := p.DB.Create(&product).Error; err != nil {
		return nil, err
	}
	return product, nil
}

// Fungsi untuk mendapatkan seluruh product
func (p *product) GetAllProducts() (interface{}, error) {
	var results []dto.GetAllProduct
	tx := p.DB.Table("products").Select("products.id, products.users_id, products.name, subcategories.subcategory_name, products.subcategory_id, products.city_id, cities.city_name, products.price, products.description, products.stock, photos.url").Group("products.id").Joins(
		"join subcategories on subcategories.id = products.subcategory_id").Joins(
		"join photos on photos.products_id = products.id").Joins(
		"join cities on products.city_id = cities.id").Where("products.deleted_at IS NULL").Find(&results)
	if tx.Error != nil || tx.RowsAffected < 1 {
		return nil, tx.Error
	}
	return results, nil
}

// Fungsi untuk mendapatkan product berdasarkan id product
func (p *product) GetProductByID(id uint) (*dto.GetProduct, error) {
	var result dto.GetProduct
	tx := p.DB.Table("products").Select("products.id, products.users_id, users.created_at, users.nama, users.phone_number, products.name, products.subcategory_id, subcategories.subcategory_name, products.city_id, cities.city_name, products.price, products.description, products.stock, products.latitude, products.longitude").Group("products.id").Joins(
		"join subcategories on subcategories.id = products.subcategory_id").Joins(
		"join photos on photos.products_id = products.id").Joins(
		"join users on products.users_id = users.id").Joins(
		"join cities on products.city_id = cities.id").Where("products.id = ?", id).Find(&result)
	if tx.Error != nil || tx.RowsAffected < 1 {
		return nil, tx.Error
	}
	return &result, nil
}

// Fungsi untuk mendapatkan product berdasarkan subcategory id
func (p *product) GetProductsBySubcategoryID(id int) (interface{}, error) {
	var results []dto.GetAllProduct
	tx := p.DB.Table("products").Select("products.id, products.users_id, products.name, subcategories.subcategory_name, products.subcategory_id, products.city_id, cities.city_name, products.price, products.description, products.stock, photos.url").Group("products.id").Joins(
		"join subcategories on subcategories.id = products.subcategory_id").Joins(
		"join photos on photos.products_id = products.id").Joins(
		"join cities on products.city_id = cities.id").Where("products.deleted_at IS NULL AND products.subcategory_id = ?", id).Find(&results)
	if tx.Error != nil || tx.RowsAffected < 1 {
		return nil, tx.Error
	}
	return results, nil
}

// Fungsi untuk mendapatkan product berdasarkan users_id
func (p *product) GetProductsByUserID(id int) (interface{}, error) {
	var results []dto.GetAllProduct
	tx := p.DB.Table("products").Select("products.id, products.users_id, products.name, subcategories.subcategory_name, products.subcategory_id, products.city_id, cities.city_name, products.price, products.description, products.stock, photos.url").Group("products.id").Joins(
		"join subcategories on subcategories.id = products.subcategory_id").Joins(
		"join photos on photos.products_id = products.id").Joins(
		"join cities on products.city_id = cities.id").Where("products.deleted_at IS NULL AND products.users_id = ?", id).Find(&results)
	if tx.Error != nil || tx.RowsAffected < 1 {
		return nil, tx.Error
	}
	return results, nil
}

// function database untuk menghapus product  by id
func (p *product) DeleteProduct(id int) {
	p.DB.Exec("DELETE from products WHERE id = ?", id)
}

// Fungsi untuk product by id
func (p *product) DeleteProductByID(id int) (interface{}, error) {
	var product model.Products
	if err := p.DB.Where("id = ?", id).Delete(&product).Error; err != nil {
		return nil, err
	}
	return "deleted", nil
}

func (p *product) GetProductOwner(id int) (int, error) {
	var ownerProduct int
	tx := p.DB.Raw("SELECT users_id FROM products WHERE id = ?", id).Scan(&ownerProduct)
	if tx.Error != nil {
		return 0, tx.Error
	}
	return ownerProduct, nil
}

// Fungsi untuk mendapatkan product berdasarkan search name
func (p *product) GetProductsByName(product string) (interface{}, error) {
	var results []dto.GetAllProduct
	var search = "%" + product + "%"
	tx := p.DB.Table("products").Select("products.id, products.users_id, products.name, subcategories.subcategory_name, products.subcategory_id, products.city_id, cities.city_name, products.price, products.description, products.stock, photos.url").Group("products.id").Joins(
		"join subcategories on subcategories.id = products.subcategory_id").Joins(
		"join photos on photos.products_id = products.id").Joins(
		"join cities on products.city_id = cities.id").Where("products.deleted_at IS NULL AND products.name LIKE ?", search).Find(&results)
	if tx.Error != nil || tx.RowsAffected < 1 {
		return nil, tx.Error
	}
	return results, nil
}
