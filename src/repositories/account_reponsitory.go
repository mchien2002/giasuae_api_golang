package repositories

import (
	"giasuaeapi/src/entities"
	"log"

	"github.com/mashingan/smapping"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AccountReponsitory interface {
	InsertAccount(acc *entities.Account)
	UpdateAccount(acc *entities.Account)
	DeleteAccount(acc *entities.Account)
	FindAllAccount() []entities.Account
	VerifyCredential(username string) interface{}
	FindByID(id uint64) entities.Account
}

type accountReponsitory struct {
	connection *gorm.DB
}

// VerifyCredential implements AccountReponsitory
func (db *accountReponsitory) VerifyCredential(username string) interface{} {
	var acc entities.Account
	accToken := entities.AccountWithToken{}
	res := db.connection.Where("username = ?", username).Take(&acc)
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
func (db *accountReponsitory) FindByID(id uint64) entities.Account {
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
func (db *accountReponsitory) InsertAccount(acc *entities.Account) {
	acc.Password = hashPass([]byte(acc.Password))
	db.connection.Save(&acc)
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
	hash, err := bcrypt.GenerateFromPassword(pss, 13)
	if err != nil {
		log.Println(err)
		panic("Failed to hash password")
	}
	return string(hash)
}
