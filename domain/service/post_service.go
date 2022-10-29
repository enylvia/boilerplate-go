package service

import (
	"github.com/enylvia/boilerplate-go/domain/model"
	"github.com/enylvia/boilerplate-go/domain/repository"
)

type ServiceStruct struct {
	repository repository.Repository
}

type Service interface {
	// Add methods here
	CreatePost(post *model.Post) error
	GetPost() ([]model.Post, error)
}

func NewService(repository repository.Repository) *ServiceStruct {
	return &ServiceStruct{repository}
}

func (s *ServiceStruct) CreatePost(post *model.Post) error {
	err := s.repository.Create(post)
	if err != nil {
		return err
	}
	return nil
}

func (s *ServiceStruct) GetPost() ([]model.Post, error) {
	posts, err := s.repository.Get()
	if err != nil {
		return nil, err
	}
	return posts, nil
}
