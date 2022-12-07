package services

import (
	"giasuaeapi/src/dto"
	"giasuaeapi/src/entities"
	"giasuaeapi/src/repositories"
)

type NewClassService interface {
	InsertNewClass(nc *dto.NewClassesReq) error
	UpdateNewClass(nc *dto.NewClassesReq) error
	DeleteNewClass(nc *entities.Newclasses) error
	FindAllNewClass() []entities.Newclasses
	FindByID(id int) entities.Newclasses
}
type newClassService struct {
	NewClassRepository repositories.NewClassRepository
}

// DeleteNewClass implements NewClassService
func (*newClassService) DeleteNewClass(nc *entities.Newclasses) error {
	panic("unimplemented")
}

// FindAllNewClass implements NewClassService
func (svc *newClassService) FindAllNewClass() []entities.Newclasses {
	return svc.NewClassRepository.FindAllNewClass()
}

// FindByID implements NewClassService
func (svc *newClassService) FindByID(id int) entities.Newclasses {
	return svc.NewClassRepository.FindByID(id)
}

// InsertNewClass implements NewClassService
func (svc *newClassService) InsertNewClass(nc *dto.NewClassesReq) error {
	return svc.NewClassRepository.InsertNewClass(nc)
}

// UpdateNewClass implements NewClassService
func (svc *newClassService) UpdateNewClass(nc *dto.NewClassesReq) error {
	return svc.NewClassRepository.UpdateNewClass(nc)
}

func NewNewClassService(newRepo repositories.NewClassRepository) NewClassService {
	return &newClassService{
		NewClassRepository: newRepo,
	}
}
