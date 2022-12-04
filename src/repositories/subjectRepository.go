package repositories

import (
	"giasuaeapi/src/entities"

	"gorm.io/gorm"
)

type SubjectRepository interface {
	InsertSubject(s *entities.Subject) error
	UpdateSubject(s *entities.Subject) error
	DeleteSubject(s *entities.Subject) error
	FindAllSubject() []entities.Subject
	FindByID(id uint64) entities.Subject
}

type subjectConnection struct {
	connection *gorm.DB
}

// DeleteSubject implements SubjectRepository
func (db *subjectConnection) DeleteSubject(s *entities.Subject) error {
	panic("unimplemented")
}

// FindAllSubject implements SubjectRepository
func (db *subjectConnection) FindAllSubject() []entities.Subject {
	var subjects []entities.Subject
	db.connection.Find(&subjects)
	return subjects
}

// FindByID implements SubjectRepository
func (db *subjectConnection) FindByID(id uint64) entities.Subject {
	var sub entities.Subject
	db.connection.First(&sub, id)
	return sub
}

// InsertSubject implements SubjectRepository
func (db *subjectConnection) InsertSubject(s *entities.Subject) error {
	err := db.connection.Save(&s)
	if err != nil {
		return err.Error
	}
	return nil
}

// UpdateSubject implements SubjectRepository
func (db *subjectConnection) UpdateSubject(s *entities.Subject) error {
	err := db.connection.Save(&s)
	if err != nil {
		return err.Error
	}
	return nil
}

func NewSubjectRepository(dbConn *gorm.DB) SubjectRepository {
	return &subjectConnection{
		connection: dbConn,
	}
}
