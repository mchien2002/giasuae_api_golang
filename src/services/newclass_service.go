package services

import (
	"giasuaeapi/src/entities"
	"giasuaeapi/src/repositories"
)

type NewClassService interface {
	InsertNewClass(nc *entities.NewClassesReq) error
	UpdateNewClass(nc *entities.NewClassesReq) error
	DeleteNewClass(id int) error
	FindAllNewClass() []entities.NewclasssesSet
	FindByID(id int) entities.NewclassesDetail
	FilterNewClass(value ...interface{}) []entities.NewclasssesSet
}
type newClassService struct {
	NewClassRepository repositories.NewClassRepository
}

// FilterNewClass implements NewClassService
func (svc *newClassService) FilterNewClass(value ...interface{}) []entities.NewclasssesSet {
	return svc.NewClassRepository.FilterNewClass(value)
}

// DeleteNewClass implements NewClassService
func (svc *newClassService) DeleteNewClass(id int) error {
	return svc.NewClassRepository.DeleteNewClass(id)
}

// FindAllNewClass implements NewClassService
func (svc *newClassService) FindAllNewClass() []entities.NewclasssesSet {
	return svc.NewClassRepository.FindAllNewClass()
}

// FindByID implements NewClassService
func (svc *newClassService) FindByID(id int) entities.NewclassesDetail {
	return svc.NewClassRepository.FindByID(id)
}

// InsertNewClass implements NewClassService
func (svc *newClassService) InsertNewClass(nc *entities.NewClassesReq) error {
	return svc.NewClassRepository.InsertNewClass(nc)
}

// UpdateNewClass implements NewClassService
func (svc *newClassService) UpdateNewClass(nc *entities.NewClassesReq) error {
	return svc.NewClassRepository.UpdateNewClass(nc)
}

func NewNewClassService(newRepo repositories.NewClassRepository) NewClassService {
	return &newClassService{
		NewClassRepository: newRepo,
	}
}
