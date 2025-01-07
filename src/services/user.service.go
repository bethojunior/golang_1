package services

import (
	"errors"
)

type UserService struct {
	users map[string]User
}

type User struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type UserInput struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func NewUserService() *UserService {
	return &UserService{
		users: make(map[string]User),
	}
}

func (us *UserService) CreateUser(input UserInput) (User, error) {
	id := generateID() // Substitua por sua lógica de geração de ID

	if input.Name == "" || input.Email == "" {
		return User{}, errors.New("name and email are required")
	}

	user := User{
		ID:    id,
		Name:  input.Name,
		Email: input.Email,
	}

	us.users[id] = user
	return user, nil
}

func (us *UserService) GetUser(id string) (User, error) {
	user, exists := us.users[id]
	if !exists {
		return User{}, errors.New("user not found")
	}

	return user, nil
}

func (us *UserService) UpdateUser(id string, input UserInput) (User, error) {
	user, exists := us.users[id]
	if !exists {
		return User{}, errors.New("user not found")
	}

	if input.Name != "" {
		user.Name = input.Name
	}

	if input.Email != "" {
		user.Email = input.Email
	}

	us.users[id] = user
	return user, nil
}

func (us *UserService) DeleteUser(id string) error {
	_, exists := us.users[id]
	if !exists {
		return errors.New("user not found")
	}

	delete(us.users, id)
	return nil
}

// Função auxiliar para gerar um ID único
func generateID() string {
	// Substitua pela lógica de geração de IDs (como UUID ou similar)
	return "some-unique-id"
}
