package user

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	RegisterUser(input RegisterUserInput) (User, error)
	Login(input LoginInput) (User, error)
	IsEmailAvailable(input CheckEmailInput) (bool, error)
	SaveAvatar(id int, fileLocation string) (User, error)
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

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)
	if err != nil {
		return user, err
	}
	user.PasswordHash = string(passwordHash)

	user.Role = "user"

	newUser, err := s.repository.Save(user)
	if err != nil {
		return newUser, err
	}

	return newUser, nil
}

func (s *service) Login(input LoginInput) (User, error) {
	email := input.Email
	password := input.Password

	userExist, err := s.repository.FindByEmail(email)
	if err != nil {
		return userExist, err
	}

	if userExist.ID == 0 {
		return userExist, errors.New("User not found on that email")
	}

	err = bcrypt.CompareHashAndPassword([]byte(userExist.PasswordHash), []byte(password))

	if err != nil {
		return userExist, err
	}

	return userExist, nil
}

func (s *service) IsEmailAvailable(input CheckEmailInput) (bool, error) {
	email := input.Email

	s.repository.FindByEmail(email)

	userExist, err := s.repository.FindByEmail(email)
	if err != nil {
		return false, err
	}

	if userExist.ID == 0 {
		return true, nil
	}

	return false, nil
}

func (s *service) SaveAvatar(id int, fileLocation string) (User, error) {

	// get user by id
	userExist, err := s.repository.FindById(id)
	if err != nil {
		return userExist, err
	}

	userExist.AvatarFileName = fileLocation

	result, err := s.repository.Update(userExist)
	if err != nil {
		return result, err
	}

	return result, nil

}
