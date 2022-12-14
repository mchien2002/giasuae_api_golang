package repositories

import (
	"fmt"
	"giasuaeapi/src/entities"
	"log"

	"github.com/mashingan/smapping"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AccountReponsitory interface {
	InsertAccount(acc *entities.Account) error
	UpdateAccount(acc *entities.Account)
	DeleteAccount(acc *entities.Account)
	FindAllAccount() []entities.Account
	VerifyCredential(username string) interface{}
	FindByID(id int) entities.Account
	FilterAccount(username string) []entities.Account
	UpdatePassword(pass string, id int) error
}

type accountReponsitory struct {
	connection *gorm.DB
}

// UpdatePassword implements AccountReponsitory
func (db *accountReponsitory) UpdatePassword(pass string, id int) error {
	pass = hashPass([]byte(pass))
	err := db.connection.Table("accounts").Where("id = ?", id).Update("password", pass)
	if err.Error != nil {
		return err.Error
	}
	return nil
}

// FilterAccount implements AccountReponsitory
func (db *accountReponsitory) FilterAccount(username string) []entities.Account {
	var account []entities.Account
	db.connection.Table("accounts").Where("username LIKE ?", username).Scan(&account)
	return account
}

// VerifyCredential implements AccountReponsitory
func (db *accountReponsitory) VerifyCredential(username string) interface{} {
	var acc entities.Account
	accToken := entities.AccountWithToken{}
	res := db.connection.Where("username = ?", username).Take(&acc)
	fmt.Println(acc)
	if res.Error == nil {
		error := smapping.FillStruct(&accToken, smapping.MapFields(&acc))
		if error != nil {
			log.Fatalf("Failed map %v: ", error)
		}
		return accToken
	}
	return false
}

// FilterAccount implements AccountReponsitory
func (db *accountReponsitory) FindByID(id int) entities.Account {
	var acc entities.Account
	db.connection.First(&acc, id)
	return acc
}

// DeleteAccount implements AccountReponsitory
func (db *accountReponsitory) DeleteAccount(acc *entities.Account) {
	db.connection.Delete(&acc)
}

// FindAllAccount implements AccountReponsitory
func (db *accountReponsitory) FindAllAccount() []entities.Account {
	var accs []entities.Account
	db.connection.Find(&accs)
	return accs
}

// InsertAccount implements AccountReponsitory
func (db *accountReponsitory) InsertAccount(acc *entities.Account) error {
	acc.Password = hashPass([]byte(acc.Password))
	err := db.connection.Save(&acc)
	if err.Error != nil {
		return err.Error
	}
	return nil
}

// UpdateAccount implements AccountReponsitory
func (db *accountReponsitory) UpdateAccount(acc *entities.Account) {
	db.connection.Save(&acc)
}

func NewAccountReponsitory(dbConn *gorm.DB) AccountReponsitory {
	return &accountReponsitory{
		connection: dbConn,
	}
}

func hashPass(pss []byte) string {
	hash, err := bcrypt.GenerateFromPassword(pss, 15)
	if err != nil {
		log.Println(err)
		panic("Failed to hash password")
	}
	return string(hash)
}
