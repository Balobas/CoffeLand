package repo

import (
	"CoffeLand/app/core/data"
	"CoffeLand/app/interfaces/database"
	"fmt"
	"testing"
)

func getBlogPostRepo(t *testing.T) data.BlogPostRepository {
	db, err := database.MySQLDB()
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	return NewBlogPostRepo(db)
}

func TestBlogPost_GetAll(t *testing.T) {
	bp := getBlogPostRepo(t)
	posts, err := bp.GetAll()
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	fmt.Println(posts)
}

func TestBlogPost_GetByID(t *testing.T) {
	bp := getBlogPostRepo(t)
	post, err := bp.GetByID("ooooooooo")
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	fmt.Println(post)
}

func TestBlogPost_GetByTitleLike(t *testing.T) {
	bp := getBlogPostRepo(t)
	posts, err := bp.GetByTitleLike("tit")
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	fmt.Println(posts)
}

func TestBlogPost_Store(t *testing.T) {
	post := data.BlogPost{
		ID:    "ooooooooo",
		Title: "title",
		Text:  "simple.text.simple.text",
		Date:  "28-09-2021",
	}

	bp := getBlogPostRepo(t)

	if err := bp.Store(post); err != nil {
		t.Error(err)
		t.FailNow()
	}
}