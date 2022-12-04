package services

import (
	"giasuaeapi/src/dto"
	"giasuaeapi/src/entities"
	"giasuaeapi/src/repositories"
)

type NewClassService interface {
	InsertNewClass(nc *dto.Newclasses) error
	UpdateNewClass(nc *entities.NewClass) error
	DeleteNewClass(nc *entities.NewClass) error
	FindAllNewClass() []entities.NewClass
	FindByID(id uint64) entities.NewClass
}
type newClassService struct {
	NewClassRepository repositories.NewClassRepository
}

// DeleteNewClass implements NewClassService
func (*newClassService) DeleteNewClass(nc *entities.NewClass) error {
	panic("unimplemented")
}

// FindAllNewClass implements NewClassService
func (*newClassService) FindAllNewClass() []entities.NewClass {
	panic("unimplemented")
}

// FindByID implements NewClassService
func (*newClassService) FindByID(id uint64) entities.NewClass {
	panic("unimplemented")
}

// InsertNewClass implements NewClassService
func (svc *newClassService) InsertNewClass(nc *dto.Newclasses) error {
	return svc.NewClassRepository.InsertNewClass(nc)
}

// UpdateNewClass implements NewClassService
func (*newClassService) UpdateNewClass(nc *entities.NewClass) error {
	panic("unimplemented")
}

func NewNewClassService(newRepo repositories.NewClassRepository) NewClassService {
	return &newClassService{
		NewClassRepository: newRepo,
	}
}
