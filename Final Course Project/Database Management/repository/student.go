package repository

import (
	"a21hc3NpZ25tZW50/model"

	"gorm.io/gorm"
)

type StudentRepository interface {
	FetchAll() ([]model.Student, error)
	FetchByID(id int) (*model.Student, error)
	Store(s *model.Student) error
	Update(id int, s *model.Student) error
	Delete(id int) error
	FetchWithClass() (*[]model.StudentClass, error)
}

type studentRepoImpl struct {
	db *gorm.DB
}

func NewStudentRepo(db *gorm.DB) *studentRepoImpl {
	return &studentRepoImpl{db}
}

func (s *studentRepoImpl) FetchAll() ([]model.Student, error) {
	
	result := []model.Student{}
	rows, err := s.db.Table("students").Select("*").Rows()
	
	defer rows.Close()
	for rows.Next() { 
	  s.db.ScanRows(rows, &result)
	}

	return result, err
}

func (s *studentRepoImpl) Store(student *model.Student) error {
	
	s.db.Create(&student)
	
	return nil 
}

func (s *studentRepoImpl) Update(id int, student *model.Student) error {

	s.db.Model(&model.Student{}).Where("id = ?", id).Updates(student)

	return nil 
}

func (s *studentRepoImpl) Delete(id int) error {
	
	s.db.Where("id = ?", id).Delete(&model.Student{})

	return nil 
}

func (s *studentRepoImpl) FetchByID(id int) (*model.Student, error) {

	var result model.Student
	s.db.Raw("SELECT * FROM students WHERE id = ?", id).Scan(&result)

	return &result, nil
}

func (s *studentRepoImpl) FetchWithClass() (*[]model.StudentClass, error) {
	
	finalResult := []model.StudentClass{}
	result := s.db.Table("students").Select("students.name as name, students.address, classes.name as class_name, classes.professor, classes.room_number").Joins("left join classes on classes.id = students.class_id").Scan(&finalResult)

	return &finalResult, result.Error
}
