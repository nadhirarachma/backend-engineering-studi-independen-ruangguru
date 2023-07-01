package repository

import (
	"a21hc3NpZ25tZW50/model"

	"gorm.io/gorm"
)

type SessionsRepository interface {
	AddSessions(session model.Session) error
	DeleteSession(token string) error
	UpdateSessions(session model.Session) error
	SessionAvailName(name string) error
	SessionAvailToken(token string) (model.Session, error)
}

type sessionsRepoImpl struct {
	db *gorm.DB
}

func NewSessionRepo(db *gorm.DB) *sessionsRepoImpl {
	return &sessionsRepoImpl{db}
}

func (s *sessionsRepoImpl) AddSessions(session model.Session) error {
	
	s.db.Create(&session)
	
	return nil
}

func (s *sessionsRepoImpl) DeleteSession(token string) error {
	
	s.db.Where("token = ?", token).Delete(&model.Session{})

	return nil
}

func (s *sessionsRepoImpl) UpdateSessions(session model.Session) error {
	
	s.db.Model(&model.Session{}).Where("username = ?", session.Username).Updates(session)

	return nil
}

func (s *sessionsRepoImpl) SessionAvailName(name string) error {

	var result model.Session
	response := s.db.Model(&model.Session{}).Where("username = ?", name).Take(&result)

	return response.Error
}

func (s *sessionsRepoImpl) SessionAvailToken(token string) (model.Session, error) {
	
	var result model.Session
	response := s.db.Model(&model.Session{}).Where("token = ?", token).Take(&result)

	return result, response.Error
}
