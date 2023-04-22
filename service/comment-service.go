package service

import (
	"h8-assignment-final-project/models"
	"h8-assignment-final-project/repository"
)

type CommentService interface {
	FindAll() []models.Comment
	FindByID(id uint) (models.Comment, error)
	InsertComment(comment models.Comment) (models.Comment, error)
	UpdateComment(comment models.Comment) (models.Comment, error)
	DeleteComment(id uint) error
}

type commentService struct {
	commentRepository repository.CommentRepository
}

func NewCommentService(commentRepo repository.CommentRepository) CommentService {
	return &commentService{
		commentRepository: commentRepo,
	}
}

func (s *commentService) FindAll() []models.Comment {
	comments, _ := s.commentRepository.FindAll()
	return comments
}

func (s *commentService) FindByID(id uint) (models.Comment, error) {
	comment, err := s.commentRepository.FindByID(id)
	if err != nil {
		return models.Comment{}, err
	}
	return comment, nil
}

func (s *commentService) InsertComment(comment models.Comment) (models.Comment, error) {
	newComment, err := s.commentRepository.InsertComment(comment)
	if err != nil {
		return models.Comment{}, err
	}
	return newComment, nil
}

func (s *commentService) UpdateComment(comment models.Comment) (models.Comment, error) {
	updatedComment, err := s.commentRepository.UpdateComment(comment)
	if err != nil {
		return models.Comment{}, err
	}
	return updatedComment, nil
}

func (s *commentService) DeleteComment(id uint) error {
	err := s.commentRepository.DeleteComment(id)
	if err != nil {
		return err
	}
	return nil
}
