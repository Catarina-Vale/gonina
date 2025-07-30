package service

import (
    "yourapp/model"
    "yourapp/repository"
)

type UserService struct {
    Repo *repository.UserRepo
}

func (s *UserService) GetUsers() ([]model.User, error) {
    return s.Repo.GetAll()
}

func (s *UserService) CreateUser(user model.User) error {
    return s.Repo.Create(user)
}

func (s *UserService) UpdateUser(user model.User) error {
    return s.Repo.Update(user)
}

func (s *UserService) DeleteUser(id int) error {
    return s.Repo.Delete(id)
}
