package repo

import (
	"a21hc3NpZ25tZW50/model"

	"gorm.io/gorm"
)

type TeacherRepo struct {
	db *gorm.DB
}

func NewTeacherRepo(db *gorm.DB) TeacherRepo {
	return TeacherRepo{db}
}

func (t TeacherRepo) Save(data model.Teacher) error {
	
	t.db.Create(&data)

	return nil
}

func (t TeacherRepo) Query() ([]model.Teacher, error) {

	result := []model.Teacher{}
	rows, err := t.db.Table("teachers").Select("*").Rows()

	if err != nil {
		return []model.Teacher{}, err
	}
	
	defer rows.Close()
	for rows.Next() { 
	  t.db.ScanRows(rows, &result)
	}
	return result, nil
}

func (t TeacherRepo) Update(id uint, name string) error {

	t.db.Model(&model.Teacher{}).Where("id = ?", id).Update("name", name)

	return nil 
}

func (t TeacherRepo) Delete(id uint) error {
	
	t.db.Where("id = ?", id).Delete(&model.Teacher{})

	return nil 
}
