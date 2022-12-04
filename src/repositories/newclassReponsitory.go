package repositories

import (
	"giasuaeapi/src/dto"
	"giasuaeapi/src/entities"

	"gorm.io/gorm"
)

type NewClassRepository interface {
	InsertNewClass(nc *dto.Newclasses) error
	UpdateNewClass(nc *entities.NewClass) error
	DeleteNewClass(nc *entities.NewClass) error
	FindAllNewClass() []entities.NewClass
	FindByID(id int) entities.NewClass
}

type newClassConnection struct {
	connection *gorm.DB
}

// DeleteNewClass implements NewClassRepository
func (db *newClassConnection) DeleteNewClass(nc *entities.NewClass) error {
	panic("unimplemented")
}

// FindAllNewClass implements NewClassRepository
func (db *newClassConnection) FindAllNewClass() []entities.NewClass {
	panic("unimplemented")

}

// FindByID implements NewClassRepository
func (db *newClassConnection) FindByID(id int) entities.NewClass {
	panic("unimplemented")

}

// InsertNewClass implements NewClassRepository
func (db *newClassConnection) InsertNewClass(nc *dto.Newclasses) error {
	var subOfNC entities.SubjectsOfNewClass
	err1 := db.connection.Save(&nc)
	if err1 != nil {
		return err1.Error
	}
	for id := range nc.Subjects {
		err2 := db.connection.Select("id_newclass = ? id_subject = ?", nc.ID, id).Save(&subOfNC)
		if err2 != nil {
			return err2.Error
		}
	}
	return nil

}

// UpdateNewClass implements NewClassRepository
func (db *newClassConnection) UpdateNewClass(nc *entities.NewClass) error {
	panic("unimplemented")

}

func NewNewClassRepository(dbConn *gorm.DB) NewClassRepository {
	return &newClassConnection{
		connection: dbConn,
	}
}
