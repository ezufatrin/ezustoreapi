package infrastructure

import (
	"ezustore/src/domain"
	"ezustore/src/infrastructure/database"
)

type productRepo struct{}

func NewProductRepo(db interface{}) domain.ProductRepository { return &productRepo{} }

func (r *productRepo) Create(p *domain.Product) error { return database.DB.Create(p).Error }
func (r *productRepo) Update(p *domain.Product) error { return database.DB.Save(p).Error }
func (r *productRepo) Delete(id uint) error           { return database.DB.Delete(&domain.Product{}, id).Error }
func (r *productRepo) GetByID(id uint) (*domain.Product, error) {
	var p domain.Product
	err := database.DB.First(&p, id).Error
	return &p, err
}
func (r *productRepo) GetAll() ([]domain.Product, error) {
	var ps []domain.Product
	err := database.DB.Find(&ps).Error
	return ps, err
}
