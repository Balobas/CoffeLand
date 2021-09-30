package data

import "errors"

type BlogPost struct {
	ID string
	Title string
	Text string
	Date string
}

func(bp *BlogPost) Validate() error {
	if len(bp.Title) == 0 {
		return errors.New("empty blog post title")
	}
	if len(bp.Text) == 0 {
		return errors.New("empty blog post text")
	}
	return nil
}

type BlogPostRepository interface {
	GetByID(id string) (BlogPost, error)
	GetByTitleLike(name string) ([]BlogPost, error)
	GetAll() ([]BlogPost, error)
	Store(post BlogPost) error
}
