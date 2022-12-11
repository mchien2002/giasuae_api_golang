package repositories

import (
	"giasuaeapi/src/entities"

	"github.com/mashingan/smapping"
	"gorm.io/gorm"
)

type NewClassRepository interface {
	InsertNewClass(nc *entities.NewClassesReq) error
	UpdateNewClass(nc *entities.NewClassesReq) error
	DeleteNewClass(nc *entities.NewclassesDetail) error
	FindAllNewClass() []entities.NewclasssesSet
	FindByID(id int) entities.NewclassesDetail
}

type newClassConnection struct {
	connection *gorm.DB
}

// DeleteNewClass implements NewClassRepository
func (db *newClassConnection) DeleteNewClass(nc *entities.NewclassesDetail) error {
	panic("unimplemented")
}

// FindAllNewClass implements NewClassRepository
func (db *newClassConnection) FindAllNewClass() []entities.NewclasssesSet {
	var newclasses []entities.NewclasssesSet
	db.connection.Table("newclasses").Select(`newclasses.id,
	newclasses.address,
	newclasses.district,
	newclasses.sobuoi,
	newclasses.time,
	newclasses.salary,
	newclasses.require,
	newclasses.status,
	newclasses.contact,
	newclasses.created_at,
	(SELECT 
		GROUP_CONCAT((SELECT c.name
			FROM
				classes c
			WHERE
				(c.id = ncc.id_class))
			SEPARATOR ', ')
		FROM
			classes_of_newclasses ncc
		WHERE
			(ncc.id_newclass = newclasses.id)) AS classes,
	(SELECT 
		GROUP_CONCAT((SELECT s.name
			FROM
				subjects s
			WHERE
				(s.id = ncs.id_subject))
			SEPARATOR ', ')
		FROM
			subjects_of_newclasses ncs
		WHERE
			(ncs.id_newclass = newclasses.id)) AS subjects,
	(SELECT 
		GROUP_CONCAT((SELECT ctg.name
			FROM
				categories ctg
			WHERE
				(ctg.id = ncctg.id_category))
			SEPARATOR ', ')
		FROM
			categories_of_newclasses ncctg
		WHERE
			(ncctg.id_newclass = newclasses.id)) AS categories`).Group("newclasses.id").Find(&newclasses)
	return newclasses
}

// FindByID implements NewClassRepository
func (db *newClassConnection) FindByID(id int) entities.NewclassesDetail {
	var newclasses entities.NewclassesDetail
	db.connection.Limit(1).Table("newclasses").Where("id = ?", id).Scan(&newclasses)
	newclasses.Subjects = getListSubjectOfNC(db, id)
	newclasses.Classes = getListClassOfNC(db, id)
	newclasses.Categories = getListCategoryOfNC(db, id)
	return newclasses
}

// InsertNewClass implements NewClassRepository
func (db *newClassConnection) InsertNewClass(nc *entities.NewClassesReq) error {
	var subOfNC []entities.SubjectsOfNewclasses
	var classOfNC []entities.ClassesOfNewclasses
	var ctgOfNC []entities.CategoriesOfNewclasses
	db.connection.Table("newclasses").Create(&nc)
	for _, value := range nc.Subjects {
		subOfNC = append(subOfNC, entities.SubjectsOfNewclasses{
			ID_newclass: nc.ID,
			ID_subject:  value,
		})
	}
	for _, value := range nc.Classes {
		classOfNC = append(classOfNC, entities.ClassesOfNewclasses{
			ID_newclass: nc.ID,
			ID_class:    value,
		})
	}
	for _, value := range nc.Categories {
		ctgOfNC = append(ctgOfNC, entities.CategoriesOfNewclasses{
			ID_newclass: nc.ID,
			ID_category: value,
		})
	}
	db.connection.Create(&subOfNC)
	db.connection.Create(&classOfNC)
	db.connection.Create(&ctgOfNC)
	return nil
}

// UpdateNewClass implements NewClassRepository
func (db *newClassConnection) UpdateNewClass(nc *entities.NewClassesReq) error {
	delListSubjectOfNC(db, nc.ID)
	delListClassOfNC(db, nc.ID)
	delListCategoryOfNC(db, nc.ID)
	var newClassSet entities.NewclasssesSet
	var subOfNC []entities.SubjectsOfNewclasses
	var clasOfNC []entities.ClassesOfNewclasses
	var ctgOfNC []entities.CategoriesOfNewclasses
	smapping.FillStruct(&newClassSet, smapping.MapFields(&nc))
	err := db.connection.Table("newclasses").Save(&newClassSet)
	if err.Error != nil {
		return err.Error
	}
	for _, value := range nc.Subjects {
		subOfNC = append(subOfNC, entities.SubjectsOfNewclasses{
			ID_newclass: nc.ID,
			ID_subject:  value,
		})
	}
	for _, value := range nc.Classes {
		clasOfNC = append(clasOfNC, entities.ClassesOfNewclasses{
			ID_newclass: nc.ID,
			ID_class:    value,
		})
	}
	for _, value := range nc.Categories {
		ctgOfNC = append(ctgOfNC, entities.CategoriesOfNewclasses{
			ID_newclass: nc.ID,
			ID_category: value,
		})
	}

	err1 := db.connection.Table("subjects_of_newclasses").Save(&subOfNC)
	err2 := db.connection.Table("subjects_of_newclasses").Save(&clasOfNC)
	err3 := db.connection.Table("subjects_of_newclasses").Save(&ctgOfNC)
	if err1.Error != nil {
		return err1.Error
	}
	if err2.Error != nil {
		return err1.Error
	}
	if err3.Error != nil {
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
func getListClassOfNC(db *newClassConnection, ncId int) []entities.Class {
	var classes []entities.Class
	db.connection.Table("classes").Joins("inner join classes_of_newclasses on id_newclass = ?", ncId).Where("id = id_class").Scan(&classes)
	return classes
}
func getListCategoryOfNC(db *newClassConnection, ncId int) []entities.Category {
	var categories []entities.Category
	db.connection.Table("categories").Joins("inner join categories_of_newclasses on id_newclass = ?", ncId).Where("id = id_category").Scan(&categories)
	return categories
}

func delListSubjectOfNC(db *newClassConnection, ncId int) {
	db.connection.Table("subjects_of_newclasses").Where("id_newclass = ?", ncId).Delete(&entities.SubjectsOfNewclasses{})
}
func delListClassOfNC(db *newClassConnection, ncId int) {
	db.connection.Table("classes_of_newclasses").Where("id_newclass = ?", ncId).Delete(&entities.ClassesOfNewclasses{})
}
func delListCategoryOfNC(db *newClassConnection, ncId int) {
	db.connection.Table("categories_of_newclasses").Where("id_newclass = ?", ncId).Delete(&entities.CategoriesOfNewclasses{})
}