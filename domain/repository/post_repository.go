package repository

import (
	"github.com/enylvia/boilerplate-go/domain/model"
	"gorm.io/gorm"
)

type RepositoryStruct struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *RepositoryStruct {
	return &RepositoryStruct{db}
}

type Repository interface {
	// Add methods here
	Create(post *model.Post) error
	Get() ([]model.Post, error)
}

func (r *RepositoryStruct) Create(post *model.Post) error {
	err := r.db.Create(post).Error
	if err != nil {
		return err
	}
	return nil
}
func (r *RepositoryStruct) Get() ([]model.Post, error) {
	var posts []model.Post
	err := r.db.Find(&posts).Error
	if err != nil {
		return nil, err
	}
	return posts, nil
}
