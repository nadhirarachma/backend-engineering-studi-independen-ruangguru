package repository

import (
	"a21hc3NpZ25tZW50/model"

	"gorm.io/gorm"
)

type ClassRepository interface {
	FetchAll() ([]model.Class, error)
}

type classRepoImpl struct {
	db *gorm.DB
}

func NewClassRepo(db *gorm.DB) *classRepoImpl {
	return &classRepoImpl{db}
}

func (s *classRepoImpl) FetchAll() ([]model.Class, error) {

	result := []model.Class{}
	rows, err := s.db.Table("classes").Select("*").Rows()
	
	defer rows.Close()
	for rows.Next() { 
	  s.db.ScanRows(rows, &result)
	}

	return result, err
}
