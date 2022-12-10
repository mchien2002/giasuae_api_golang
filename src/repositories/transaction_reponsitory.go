package repositories

import (
	"giasuaeapi/src/entities"

	"gorm.io/gorm"
)

type TransRepository interface {
	InsertTrans(trans *entities.Transactionhistories) error
	FindAllTrans() []entities.Transactionhistories
	FindByIDAcc(id int) entities.Transactionhistories
}

type transConnection struct {
	connection *gorm.DB
}

// FindAllTrans implements TransRepository
func (db *transConnection) FindAllTrans() []entities.Transactionhistories {
	var transs []entities.Transactionhistories
	db.connection.Table("transactionhistories").Select("trans.id, trans.amount, trans.content, trans.status, trans.created_at, accounts.username AS id_account ").Joins("trans left join accounts on accounts.id = trans.id_account").Find(&transs)
	return transs
}

// FindByIDAcc implements TransRepository
func (db *transConnection) FindByIDAcc(id int) entities.Transactionhistories {
	var trans entities.Transactionhistories
	db.connection.Where("id_account = ?", id).First(&trans)
	return trans
}

// InsertTrans implements TransRepository
func (db *transConnection) InsertTrans(trans *entities.Transactionhistories) error {
	err := db.connection.Save(&trans)
	if err.Error != nil {
		return err.Error
	}
	return nil
}

func NewTransRepository(dbConn *gorm.DB) TransRepository {
	return &transConnection{
		connection: dbConn,
	}
}
