package comment

import "github.com/jinzhu/gorm"

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
	GetAllComment() ([]Comment, error)
}

func NewService(db *gorm.DB) *Service {
	return &Service{
		DB: db,
	}
}
