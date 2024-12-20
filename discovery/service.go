package discovery

import (
	"context"

	"github.com/BeReal-Candidates/backend-go-test/util"
)

type Service interface {
	AddPost(context.Context, *AddPost) (*Post, error)
	GetPosts(ctx context.Context, cursor string) (Page[*Post], error)
	Close()
}

type service struct {
	posts []*Post
	pageSize uint
}

func NewService(pageSize uint) Service {
	return &service{
		posts: make([]*Post, 0, 10),
		pageSize: pageSize,
	}
}

func (s *service) AddPost(ctx context.Context, p *AddPost) (*Post, error) {
	id := util.MustGetNewID()
	post := &Post{
		ID:          id,
		OwnerID:     p.OwnerID,
		FrontPicUrl: p.FrontPicUrl,
		BackPicUrl:  p.BackPicUrl,
	}
	s.posts = append(s.posts, post)
	return post, nil
}

// TODO: implement it!
func (s *service) GetPosts(ctx context.Context, cursor string) (Page[*Post], error) {
	return Page[*Post]{
		Data:   s.posts,
		Cursor: s.posts[len(s.posts)-1].ID,
	}, nil
}

func (s *service) Close() {}
