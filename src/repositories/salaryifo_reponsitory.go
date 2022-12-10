package repositories

import (
	"giasuaeapi/src/entities"

	"gorm.io/gorm"
)

type SalaryinfoRepository interface {
	InsertSalaryinfo(sal *entities.Salaryinfo) error
	UpdateSalaryinfo(sal *entities.Salaryinfo) error
	DeleteSalaryinfo(id int) error
	FindAllSalaryinfo() []entities.Salaryinfo
	FindByID(id int) entities.Salaryinfo
	FindByType(type_teacher int) []entities.Salaryinfo
}

type salaryinfoConnection struct {
	connection *gorm.DB
}

// FindByType implements SalaryinfoRepository
func (*salaryinfoConnection) FindByType(type_teacher int) []entities.Salaryinfo {
	panic("unimplemented")
}

// DeleteSalaryinfo implements SalaryinfoRepository
func (db *salaryinfoConnection) DeleteSalaryinfo(id int) error {
	err := db.connection.Delete(&entities.Salaryinfo{}, id)
	if err.Error != nil {
		return err.Error
	}
	return nil
}

// FindAllSalaryinfo implements SalaryinfoRepository
func (db *salaryinfoConnection) FindAllSalaryinfo() []entities.Salaryinfo {
	var sals []entities.Salaryinfo
	db.connection.Table("salaryinfos").Select("sal.id, sal.type_teacher, sal.two_sessions, sal.three_sessions, sal.four_sessions, sal.five_sessions, sal.created_at, categories.name AS id_category ").Joins("sal left join categories on categories.id = sal.id_category").Find(&sals)
	return sals
}

// FindByID implements SalaryinfoRepository
func (db *salaryinfoConnection) FindByID(id int) entities.Salaryinfo {
	var sal entities.Salaryinfo
	db.connection.First(&sal, id)
	return sal
}

// InsertSalaryinfo implements SalaryinfoRepository
func (db *salaryinfoConnection) InsertSalaryinfo(sal *entities.Salaryinfo) error {
	err := db.connection.Save(&sal)
	if err.Error != nil {
		return err.Error
	}
	return nil
}

// UpdateSalaryinfo implements SalaryinfoRepository
func (db *salaryinfoConnection) UpdateSalaryinfo(sal *entities.Salaryinfo) error {
	err := db.connection.Save(&sal)
	if err.Error != nil {
		return err.Error
	}
	return nil
}

func NewSalaryinfoRepository(dbConn *gorm.DB) SalaryinfoRepository {
	return &salaryinfoConnection{
		connection: dbConn,
	}
}
