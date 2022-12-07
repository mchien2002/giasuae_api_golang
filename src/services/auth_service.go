package services

import (
	"giasuaeapi/src/dto"
	"giasuaeapi/src/entities"
	"giasuaeapi/src/repositories"
	"log"

	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	VerifyCredential(usernam string, password string, role int) interface{}
	CreateUser(user *entities.Account)
	// FindByEmail(email string) entity.User
	// IsDuplicateEmail(email string) bool
}

type authService struct {
	AcountRepository repositories.AccountReponsitory
}

// CreateUser implements AuthService
func (svc *authService) CreateUser(user *entities.Account) {
	svc.AcountRepository.InsertAccount(user)
}

// VerifyCredential implements AuthService
func (svc *authService) VerifyCredential(username string, password string, role int) interface{} {
	res := svc.AcountRepository.VerifyCredential(username)
	if acv, ok := res.(dto.AccountWithToken); ok {
		isComparePass := comparePass(acv.Password, []byte(password))
		if acv.Email == username && isComparePass && acv.Role == role {
			return res
		} else {
			return false
		}
	}
	return false
}

func NewAuthService(accRepo repositories.AccountReponsitory) AuthService {
	return &authService{
		AcountRepository: accRepo,
	}
}

func comparePass(hashedPwd string, plainPassword []byte) bool {
	byteHash := []byte(hashedPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, plainPassword)
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}
