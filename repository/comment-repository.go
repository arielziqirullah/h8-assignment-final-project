package repository

import (
	"h8-assignment-final-project/models"

	"gorm.io/gorm"
)

type CommentRepository interface {
	FindAll() ([]models.Comment, error)
	FindByID(id uint) (models.Comment, error)
	InsertComment(comment models.Comment) (models.Comment, error)
	UpdateComment(comment models.Comment) (models.Comment, error)
	DeleteComment(id uint) error
}

type commentRepository struct {
	db *gorm.DB
}

func NewCommentRepository(db *gorm.DB) CommentRepository {
	return &commentRepository{
		db: db,
	}
}

func (db *commentRepository) FindAll() ([]models.Comment, error) {
	var comments []models.Comment
	err := db.db.Find(&comments).Error
	if err != nil {
		return []models.Comment{}, err
	}
	return comments, nil
}

func (db *commentRepository) FindByID(id uint) (models.Comment, error) {
	var comment models.Comment
	err := db.db.Where("id = ?", id).Take(&comment).Error
	if err != nil {
		return models.Comment{}, err
	}
	return comment, nil
}

func (db *commentRepository) InsertComment(comment models.Comment) (models.Comment, error) {
	err := db.db.Create(&comment).Error
	if err != nil {
		return models.Comment{}, err
	}
	return comment, nil
}

func (db *commentRepository) UpdateComment(comment models.Comment) (models.Comment, error) {
	err := db.db.Save(&comment).Error
	if err != nil {
		return models.Comment{}, err
	}
	return comment, nil
}

func (db *commentRepository) DeleteComment(id uint) error {
	var comment models.Comment
	err := db.db.Where("id = ?", id).Delete(&comment).Error
	if err != nil {
		return err
	}

	return nil
}
