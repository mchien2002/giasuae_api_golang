package repositories

import (
	"giasuaeapi/src/entities"

	"gorm.io/gorm"
)

type TutorRepository interface {
	InsertTutor(tutor *entities.TutorReq) error
	UpdateTutor(tutor *entities.TutorReq) error
	DeleteTutor(tutor *entities.TutorReq) error
	FindAllTutor() []entities.TutorSet
	FindByID(id int) entities.TutorDetail
}

type tutorConnection struct {
	connection *gorm.DB
}

// DeleteTutor implements TutorRepository
func (*tutorConnection) DeleteTutor(tutor *entities.TutorReq) error {
	panic("unimplemented")
}

// UpdateTutor implements TutorRepository
func (*tutorConnection) UpdateTutor(tutor *entities.TutorReq) error {
	panic("unimplemented")
}

// FindAllTutor implements TutorRepository
func (db *tutorConnection) FindAllTutor() []entities.TutorSet {
	var tutors []entities.TutorSet
	db.connection.Table("tutors").Select(`tutors.id,
	tutors.name,
	tutors.address,
	tutors.email,
	tutors.phone,
	tutors.school,
	tutors.department,
	tutors.teach_areas,
	tutors.gender,
	tutors.graduate_year,
	tutors.isnow,
	tutors.describe,
	tutors.sobuoi,
	tutors.birth_year,
	tutors.created_at,
	(SELECT 
		GROUP_CONCAT((SELECT c.name
			FROM
				classes c
			WHERE
				(c.id = ncc.id_class))
			SEPARATOR ', ')
		FROM
			classes_of_tutors ncc
		WHERE
			(ncc.id_tutor = tutors.id)) AS classes,
	(SELECT 
		GROUP_CONCAT((SELECT s.name
			FROM
				subjects s
			WHERE
				(s.id = ncs.id_subject))
			SEPARATOR ', ')
		FROM
			subjects_of_tutors ncs
		WHERE
			(ncs.id_tutor = tutors.id)) AS subjects`).Group("tutors.id").Find(&tutors)
	return tutors
}

// FindByID implements TutorRepository
func (db *tutorConnection) FindByID(id int) entities.TutorDetail {
	var tutors entities.TutorDetail
	db.connection.Limit(1).Table("tutors").Where("id = ?", id).Scan(&tutors)
	tutors.Subjects = getListSubjectOfTutor(db, id)
	tutors.Classes = getListClassOfTutor(db, id)
	return tutors
}

// InsertTutor implements TutorRepository
func (db *tutorConnection) InsertTutor(tutor *entities.TutorReq) error {
	var subOfTT []entities.SubjectsOfTutor
	var classOfTT []entities.ClassesOfTutor
	db.connection.Table("tutors").Create(&tutor)
	for _, value := range tutor.Subjects {
		subOfTT = append(subOfTT, entities.SubjectsOfTutor{
			ID_tutor:   tutor.ID,
			ID_subject: value,
		})
	}
	for _, value := range tutor.Classes {
		classOfTT = append(classOfTT, entities.ClassesOfTutor{
			ID_tutor: tutor.ID,
			ID_class: value,
		})
	}

	db.connection.Create(&subOfTT)
	db.connection.Create(&classOfTT)
	return nil
}

func NewTutorRepository(dbConn *gorm.DB) TutorRepository {
	return &tutorConnection{
		connection: dbConn,
	}
}

func getListSubjectOfTutor(db *tutorConnection, id int) []entities.Subject {
	var subjects []entities.Subject
	db.connection.Table("subjects").Joins("inner join subjects_of_tutors on id_tutor = ?", id).Where("id = id_subject").Scan(&subjects)
	return subjects
}
func getListClassOfTutor(db *tutorConnection, id int) []entities.Class {
	var classes []entities.Class
	db.connection.Table("classes").Joins("inner join classes_of_tutors on id_tutor = ?", id).Where("id = id_class").Scan(&classes)
	return classes
}

// func getListCategoryOfNC(db *TutorConnection, ncId int) []entities.Category {
// 	var categories []entities.Category
// 	db.connection.Table("categories").Joins("inner join categories_of_Tutores on id_Tutor = ?", ncId).Where("id = id_category").Scan(&categories)
// 	return categories
// }

// func delListSubjectOfNC(db *TutorConnection, ncId int) {
// 	db.connection.Table("subjects_of_Tutores").Where("id_Tutor = ?", ncId).Delete(&entities.SubjectsOfTutores{})
// }
// func delListClassOfNC(db *TutorConnection, ncId int) {
// 	db.connection.Table("classes_of_Tutores").Where("id_Tutor = ?", ncId).Delete(&entities.ClassesOfTutores{})
// }
// func delListCategoryOfNC(db *TutorConnection, ncId int) {
// 	db.connection.Table("categories_of_Tutores").Where("id_Tutor = ?", ncId).Delete(&entities.CategoriesOfTutores{})
// }
