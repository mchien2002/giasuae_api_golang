package repositories

import (
	"giasuaeapi/src/entities"

	"gorm.io/gorm"
)

type TransRepository interface {
	InsertTrans(trans *entities.TransactionhistoriesReq) error
	FindAllTrans() []entities.Transactionhistories
	FindByIDAcc(id int) entities.Transactionhistories
	FilterTrans(key interface{}) []entities.Transactionhistories
}

type transConnection struct {
	connection *gorm.DB
}

// FilterTrans implements TransRepository
func (db *transConnection) FilterTrans(key interface{}) []entities.Transactionhistories {
	var transs []entities.Transactionhistories
	db.connection.Table("transactionhistories").Select("trans.id, trans.amount, trans.content, trans.status, trans.created_at, accounts.username AS id_account ").Joins("trans left join accounts on accounts.id = trans.id_account").Where("trans.id LIKE ? OR trans.amount LIKE ? OR trans.content LIKE ? OR trans.status LIKE ? OR trans.created_at LIKE ? OR accounts.username LIKE ? ", key, key, key, key, key, key).Find(&transs)
	return transs
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
func (db *transConnection) InsertTrans(trans *entities.TransactionhistoriesReq) error {
	err := db.connection.Table("transactionhistories").Save(&trans)
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
