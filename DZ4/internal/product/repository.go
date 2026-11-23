package product

import (
	"dz4/pkg/db"

	"gorm.io/gorm/clause"
)

type ProdRepo struct {
	Db *db.Db
}

func NewProductrRepo(database *db.Db) *ProdRepo {
	return &ProdRepo{
		Db: database,
	}
}
func (repo *ProdRepo) Create(product *Product) (*Product, error) {
	result := repo.Db.DB.Create(product)
	if result.Error != nil {
		return nil, result.Error
	}
	return product, nil
}
func (repo *ProdRepo) Update(product *Product) (*Product, error) {
	result := repo.Db.DB.Clauses(clause.Returning{}).Updates(product)
	if result.Error != nil {
		return nil, result.Error
	}
	return product, nil
}
func (repo *ProdRepo) Delete(prod *Product, id uint) error {
	result := repo.Db.DB.Delete(prod, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
func (repo *ProdRepo) GetById(id uint) (*Product, error) {
	var prod Product
	result := repo.Db.DB.First(&prod, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &prod, nil
}
