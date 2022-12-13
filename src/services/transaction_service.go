package services

import (
	"giasuaeapi/src/entities"
	"giasuaeapi/src/repositories"
)

type TransService interface {
	InsertTrans(trans *entities.TransactionhistoriesReq) error
	FindAllTrans() []entities.Transactionhistories
	FindByIDAcc(id int) entities.Transactionhistories
	FilterTrans(key interface{}) []entities.Transactionhistories
}
type transService struct {
	TransRepository repositories.TransRepository
}

// FilterTrans implements TransService
func (svc *transService) FilterTrans(key interface{}) []entities.Transactionhistories {
	return svc.TransRepository.FilterTrans(key)
}

// FindAllTrans implements TransService
func (svc *transService) FindAllTrans() []entities.Transactionhistories {
	return svc.TransRepository.FindAllTrans()
}

// FindByIDAcc implements TransService
func (svc *transService) FindByIDAcc(id int) entities.Transactionhistories {
	return svc.TransRepository.FindByIDAcc(id)
}

// InsertTrans implements TransService
func (svc *transService) InsertTrans(trans *entities.TransactionhistoriesReq) error {
	return svc.TransRepository.InsertTrans(trans)
}

func NewTransService(Repo repositories.TransRepository) TransService {
	return &transService{
		TransRepository: Repo,
	}
}
