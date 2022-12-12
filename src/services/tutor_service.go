package services

import (
	"giasuaeapi/src/entities"
	"giasuaeapi/src/repositories"
)

type TutorService interface {
	InsertTutor(tutor *entities.TutorReq) error
	UpdateTutor(tutor *entities.TutorReq) error
	DeleteTutor(tutor *entities.TutorReq) error
	FindAllTutor() []entities.TutorSet
	FindByID(id int) entities.TutorDetail
}
type tutorService struct {
	TutorRepository repositories.TutorRepository
}

// DeleteTutor implements TutorService
func (*tutorService) DeleteTutor(tutor *entities.TutorReq) error {
	panic("unimplemented")
}

// UpdateTutor implements TutorService
func (*tutorService) UpdateTutor(tutor *entities.TutorReq) error {
	panic("unimplemented")
}

// FindAllTutor implements TutorService
func (svc *tutorService) FindAllTutor() []entities.TutorSet {
	return svc.TutorRepository.FindAllTutor()
}

// FindByID implements TutorService
func (svc *tutorService) FindByID(id int) entities.TutorDetail {
	return svc.TutorRepository.FindByID(id)
}

// InsertTutor implements TutorService
func (svc *tutorService) InsertTutor(tutor *entities.TutorReq) error {
	return svc.TutorRepository.InsertTutor(tutor)
}

func NewTutorService(repo repositories.TutorRepository) TutorService {
	return &tutorService{
		TutorRepository: repo,
	}
}
