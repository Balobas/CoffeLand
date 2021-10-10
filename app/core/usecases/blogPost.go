package usecases

import (
	"CoffeLand/app/core/data"
	"github.com/google/uuid"
	"github.com/pkg/errors"
)

type BlogPostUsecases struct {
	*data.Core
}

func NewBlogPostUsecases(core *data.Core) *BlogPostUsecases {
	return &BlogPostUsecases{core}
}

func(b *BlogPostUsecases) GetByTitleLike(name string) ([]data.BlogPost, error) {
	return b.BlogPostRepo.GetByTitleLike(name)
}

func(b *BlogPostUsecases) GetAll() ([]data.BlogPost, error) {
	return b.BlogPostRepo.GetAll()
}

func(b *BlogPostUsecases) put(post data.BlogPost) (string, error) {
	if err := post.Validate(); err != nil {
		return "", errors.WithStack(err)
	}
	if len(post.ID) == 0 {
		post.ID = uuid.New().String()
	}

	if err := b.BlogPostRepo.Store(post); err != nil {
		return "",  errors.WithStack(err)
	}

	return post.ID, nil
}
