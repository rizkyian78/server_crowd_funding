package Users

import (
	"errors"
	"fmt"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	RegisterUser(input RegisterUserInput) (User, error)
	LoginUser(input LoginInput) (User, error)
	CheckEmail(input CheckEmailInput) (bool, error)
	SaveAvatar(ID string, fileLocation string) (User, error)
	GetUserByID(ID string) (User, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) RegisterUser(input RegisterUserInput) (User, error) {
	user := User{}
	user.Name = input.Name
	user.Email = input.Email
	user.Occupation = input.Occupation
	user.ID = uuid.NewString()
	fmt.Println(user)
	passHash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)
	if err != nil {
		return user, err
	}
	user.HashedPassword = string(passHash)
	user.Role = "user"
	newUser, err := s.repository.Save(user)
	if err != nil {
		return newUser, err
	}
	return newUser, nil
}

func (s *service) LoginUser(input LoginInput) (User, error) {
	email := input.Email
	password := input.Password
	user, err := s.repository.FindByEmail(email)
	if err != nil {
		return user, err
	}
	if user.ID == "" {
		return user, errors.New("User Not Found")
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.HashedPassword), []byte(password))
	if err != nil {
		return user, errors.New("Wrong Password")
	}
	return user, nil
}

func (s *service) CheckEmail(input CheckEmailInput) (bool, error) {
	email := input.Email
	user, err := s.repository.FindByEmail(email)
	if err != nil {
		return false, err
	}
	if user.ID == "" {
		return true, nil
	}
	return false, nil
}

func (s *service) SaveAvatar(ID string, fileLocation string) (User, error) {
	user, err := s.repository.FindByID(ID)
	if err != nil {
		return user, err
	}
	user.AvatarFileName = fileLocation
	updatedUser, err := s.repository.Update(user)
	if err != nil {
		return updatedUser, err
	}
	return updatedUser, nil
}

func (s *service) GetUserByID(ID string) (User, error) {
	user, err := s.repository.FindByID(ID)
	if err != nil {
		return user, err
	}
	if user.ID == "" {
		return user, errors.New("No User Found")
	}
	return user, nil
}
