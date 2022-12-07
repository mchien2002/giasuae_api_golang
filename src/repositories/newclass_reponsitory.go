package repositories

import (
	"giasuaeapi/src/dto"
	"giasuaeapi/src/entities"

	"github.com/mashingan/smapping"
	"gorm.io/gorm"
)

type NewClassRepository interface {
	InsertNewClass(nc *dto.NewClassesReq) error
	UpdateNewClass(nc *dto.NewClassesReq) error
	DeleteNewClass(nc *entities.Newclasses) error
	FindAllNewClass() []entities.Newclasses
	FindByID(id int) entities.Newclasses
}

type newClassConnection struct {
	connection *gorm.DB
}

// DeleteNewClass implements NewClassRepository
func (db *newClassConnection) DeleteNewClass(nc *entities.Newclasses) error {
	panic("unimplemented")
}

// FindAllNewClass implements NewClassRepository
func (db *newClassConnection) FindAllNewClass() []entities.Newclasses {
	var newclasses []entities.Newclasses
	db.connection.Table("newclasses").Scan(&newclasses)
	for index := range newclasses {
		newclasses[index].Subjects = getListSubjectOfNC(db, newclasses[index].ID)
	}
	return newclasses
}

// FindByID implements NewClassRepository
func (db *newClassConnection) FindByID(id int) entities.Newclasses {
	var newclass entities.Newclasses
	db.connection.Limit(1).Table("newclasses").Where("id = ?", id).Scan(&newclass)
	newclass.Subjects = getListSubjectOfNC(db, id)
	return newclass
}

// InsertNewClass implements NewClassRepository
func (db *newClassConnection) InsertNewClass(nc *dto.NewClassesReq) error {
	var subOfNC []entities.SubjectsOfNewclasses
	db.connection.Table("newclasses").Create(&nc)
	for _, value := range nc.Subjects {
		subOfNC = append(subOfNC, entities.SubjectsOfNewclasses{
			ID_newclass: nc.ID,
			ID_subject:  value,
		})
	}
	db.connection.Create(&subOfNC)
	return nil
}

// UpdateNewClass implements NewClassRepository
func (db *newClassConnection) UpdateNewClass(nc *dto.NewClassesReq) error {
	delListSubjectOfNC(db, nc.ID)
	var newClassSet dto.NewclasssesSet
	var subOfNC []entities.SubjectsOfNewclasses
	smapping.FillStruct(&newClassSet, smapping.MapFields(&nc))
	err := db.connection.Table("newclasses").Save(&newClassSet)
	if err.Error !=  nil{
		return err.Error
	}
	for _, value := range nc.Subjects {
		subOfNC = append(subOfNC, entities.SubjectsOfNewclasses{
			ID_newclass: nc.ID,
			ID_subject:  value,
		})
	}

	err1 := db.connection.Table("subjects_of_newclasses").Save(&subOfNC)
	if err1.Error != nil {
		return err1.Error
	}
	return nil
}

func NewNewClassRepository(dbConn *gorm.DB) NewClassRepository {
	return &newClassConnection{
		connection: dbConn,
	}
}

func getListSubjectOfNC(db *newClassConnection, ncId int) []entities.Subject {
	var subjects []entities.Subject
	db.connection.Table("subjects").Joins("inner join subjects_of_newclasses on id_newclass = ?", ncId).Where("id = id_subject").Scan(&subjects)
	return subjects
}

func delListSubjectOfNC(db *newClassConnection, ncId int) {
	db.connection.Table("subjects_of_newclasses").Where("id_newclass = ?", ncId).Delete(&entities.SubjectsOfNewclasses{})
}
