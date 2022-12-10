package services

import (
	"giasuaeapi/src/entities"
	"giasuaeapi/src/repositories"
)

type SalaryinfoService interface {
	InsertSalaryinfo(sal *entities.Salaryinfo) error
	UpdateSalaryinfo(sal *entities.Salaryinfo) error
	DeleteSalaryinfo(id int) error
	FindAllSalaryinfo() []entities.Salaryinfo
	FindByID(id int) entities.Salaryinfo
	FindByType(type_teacher int) []entities.Salaryinfo
}
type salaryinfoService struct {
	SalaryinfoRepository repositories.SalaryinfoRepository
}

// DeleteSalaryinfo implements SalaryinfoService
func (svc *salaryinfoService) DeleteSalaryinfo(id int) error {
	return svc.SalaryinfoRepository.DeleteSalaryinfo(id)
}

// FindAllSalaryinfo implements SalaryinfoService
func (svc *salaryinfoService) FindAllSalaryinfo() []entities.Salaryinfo {
	return svc.SalaryinfoRepository.FindAllSalaryinfo()
}

// FindByID implements SalaryinfoService
func (svc *salaryinfoService) FindByID(id int) entities.Salaryinfo {
	return svc.SalaryinfoRepository.FindByID(id)
}

// FindByType implements SalaryinfoService
func (svc *salaryinfoService) FindByType(type_teacher int) []entities.Salaryinfo {
	return svc.SalaryinfoRepository.FindByType(type_teacher)
}

// InsertSalaryinfo implements SalaryinfoService
func (svc *salaryinfoService) InsertSalaryinfo(sal *entities.Salaryinfo) error {
	return svc.SalaryinfoRepository.InsertSalaryinfo(sal)
}

// UpdateSalaryinfo implements SalaryinfoService
func (svc *salaryinfoService) UpdateSalaryinfo(sal *entities.Salaryinfo) error {
	return svc.SalaryinfoRepository.UpdateSalaryinfo(sal)
}

func NewSalaryinfoService(subRepo repositories.SalaryinfoRepository) SalaryinfoService {
	return &salaryinfoService{
		SalaryinfoRepository: subRepo,
	}
}
