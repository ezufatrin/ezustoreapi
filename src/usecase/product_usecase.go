package usecase

import "ezustore/src/domain"

type productUsecase struct{ repo domain.ProductRepository }

func NewProductUsecase(r domain.ProductRepository) domain.ProductUsecase {
	return &productUsecase{repo: r}
}

func (u *productUsecase) Create(name string, price float64, stock int, category string) error {
	p := &domain.Product{Name: name, Price: price, Stock: stock, Category: category}
	return u.repo.Create(p)
}

func (u *productUsecase) Update(id uint, name string, price float64, stock int, category string) error {
	p, err := u.repo.GetByID(id)
	if err != nil {
		return err
	}
	p.Name = name
	p.Price = price
	p.Stock = stock
	p.Category = category
	return u.repo.Update(p)
}

func (u *productUsecase) Delete(id uint) error                     { return u.repo.Delete(id) }
func (u *productUsecase) GetByID(id uint) (*domain.Product, error) { return u.repo.GetByID(id) }
func (u *productUsecase) GetAll() ([]domain.Product, error)        { return u.repo.GetAll() }
