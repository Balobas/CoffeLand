package repo

import (
	"CoffeLand/app/core/data"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type BlogPost struct {
	db gorm.DB
}

func NewBlogPostRepo(db gorm.DB) data.BlogPostRepository {
	return &BlogPost{db:db}
}

func (b BlogPost) GetByID(id string) (data.BlogPost, error) {
	var post data.BlogPost
	result := b.db.Where("id = ?", id).Find(&post)
	if result.Error != nil {
		return data.BlogPost{}, result.Error
	}
	if result.RowsAffected == 0 {
		return data.BlogPost{}, gorm.ErrRecordNotFound
	}

	return post, nil
}

func (b BlogPost) GetByTitleLike(name string) ([]data.BlogPost, error) {
	return handleErrorsAndReturnBlogPosts(b.db.Where("title LIKE ?", name + "%"))
}

func (b BlogPost) GetAll() ([]data.BlogPost, error) {
	return handleErrorsAndReturnBlogPosts(&b.db)
}

func (b BlogPost) Store(post data.BlogPost) error {
	_, err := b.GetByID(post.ID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return b.db.Create(&post).Error
		}
		return errors.WithStack(err)
	}

	return b.db.Save(post).Error
}


func handleErrorsAndReturnBlogPosts(where *gorm.DB) ([]data.BlogPost, error) {
	var posts []data.BlogPost
	result := where.Find(&posts)
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return []data.BlogPost{}, gorm.ErrRecordNotFound
	}
	return posts, nil
}
