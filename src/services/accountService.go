package services

import (
	"giasuaeapi/src/entities"
	"giasuaeapi/src/repositories"
)

type AccountService interface {
	InsertAccount(acc *entities.Account)
	UpdateAccount(acc *entities.Account)
	DeleteAccount(acc *entities.Account)
	FindAllAccount() []entities.Account
	FindByID(id uint64) entities.Account
}
type accountService struct {
	AccountReponsitory repositories.AccountReponsitory
}

// DeleteAccount implements AccountService
func (accsv*accountService) DeleteAccount(acc *entities.Account) {
	accsv.AccountReponsitory.DeleteAccount(acc)
}

// FindAllAccount implements AccountService
func (accsv*accountService) FindAllAccount() []entities.Account {
	return accsv.AccountReponsitory.FindAllAccount()
}

// FindByID implements AccountService
func (accsv*accountService) FindByID(id uint64) entities.Account {
	return accsv.AccountReponsitory.FindByID(id)
}

// InsertAccount implements AccountService
func (accsv*accountService) InsertAccount(acc *entities.Account) {
	accsv.AccountReponsitory.InsertAccount(acc)
}

// UpdateAccount implements AccountService
func (accsv*accountService) UpdateAccount(acc *entities.Account) {
	accsv.AccountReponsitory.UpdateAccount(acc)
}

func NewAccountService(accRepo repositories.AccountReponsitory) AccountService {
	return &accountService{
		AccountReponsitory: accRepo,
	}
}
