package services

import (
	"giasuaeapi/src/entities"
	"giasuaeapi/src/repositories"
)

type NewClassService interface {
	InsertNewClass(nc *entities.NewClassesReq) error
	UpdateNewClass(nc *entities.NewClassesReq) error
	DeleteNewClass(nc *entities.NewclassesDetail) error
	FindAllNewClass() []entities.NewclassesDetail
	FindByID(id int) entities.NewclassesDetail
}
type newClassService struct {
	NewClassRepository repositories.NewClassRepository
}

// DeleteNewClass implements NewClassService
func (*newClassService) DeleteNewClass(nc *entities.NewclassesDetail) error {
	panic("unimplemented")
}

// FindAllNewClass implements NewClassService
func (svc *newClassService) FindAllNewClass() []entities.NewclassesDetail {
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
