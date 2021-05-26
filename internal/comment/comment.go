package comment

import (
	"github.com/jinzhu/gorm"
)

// Service - the struct for our comment service
type Service struct {
	DB *gorm.DB
}

// Comment - define comment structur
type Comment struct {
	gorm.Model
	Slug   string
	Body   string
	Author string
}

// CommentService - the interface for comment service
type CommentService interface {
	GetComment(ID uint) (Comment, error)
	GetCommentBySlug(slug string) ([]Comment, error)
	PostComment(comment Comment) (Comment, error)
	UpdateComment(ID uint, NewComment Comment) (Comment, error)
	DeleteComment(ID uint) error
	GetAllComments() ([]Comment, error)
}

func NewService(db *gorm.DB) *Service {
	return &Service{
		DB: db,
	}
}

//GetComment - retrieves comments by their ID
func (s *Service) GetComment(ID uint) (Comment, error) {
	var comment Comment

	if result := s.DB.First(&comment, ID); result.Error != nil {
		return Comment{}, result.Error
	}
	return comment, nil
}

//GetCommentBySlug - retrieves comments by slug (path - /article/name)
func (s *Service) GetCommentBySlug(slug string) ([]Comment, error) {
	var comments []Comment
	if result := s.DB.Find(&comments).Where("slug = ?", slug); result.Error != nil {
		return []Comment{}, result.Error
	}
	return comments, nil
}

//PostComment - Adds a new comment to db
func (s *Service) PostComment(comment Comment) (Comment, error) {
	if result := s.DB.Save(&comment); result.Error != nil {
		return Comment{}, result.Error
	}
	return comment, nil
}

//UpdateComment - Update a comment
func (s *Service) UpdateComment(ID uint, newComment Comment) (Comment, error) {
	comment, err := s.GetComment(ID)
	if err != nil {
		return Comment{}, err
	}
	if result := s.DB.Model(&comment).Update(newComment); result.Error != nil {
		return Comment{}, result.Error
	}

	return comment, nil
}

//DeleteComment - Deletes a comment from db by ID
func (s *Service) DeleteComment(ID uint) error {
	if result := s.DB.Delete(&Comment{}, ID); result.Error != nil {
		return result.Error
	}
	return nil
}

//GetAllComments - retrieves all comments from db
func (s *Service) GetAllComments() ([]Comment, error) {
	var comments []Comment
	if result := s.DB.Find(&comments); result.Error != nil {
		return comments, result.Error
	}
	return comments, nil
}
