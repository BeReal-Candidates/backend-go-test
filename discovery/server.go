package discovery

import (
	"context"

	pbdiscovery "github.com/BeReal-Candidates/backend-go-test/protob/discovery/v1"
)

type Server struct {
	service Service
}

func NewServer(service Service) (*Server, error) {
	return &Server{
		service: service,
	}, nil
}

func (s *Server) AddPost(ctx context.Context, req *pbdiscovery.AddPostRequest) (*pbdiscovery.AddPostResponse, error) {
	post, err := s.service.AddPost(ctx, &AddPost{
		OwnerID:     req.Post.Owner,
		FrontPicUrl: req.Post.FrontPicUrl,
		BackPicUrl:  req.Post.FrontPicUrl,
	})
	if err != nil {
		return &pbdiscovery.AddPostResponse{
			Success: false,
		}, err
	}

	return &pbdiscovery.AddPostResponse{
		Success: true,
		Post: &pbdiscovery.Post{
			Id:          post.ID,
			Owner:       post.OwnerID,
			FrontPicUrl: post.FrontPicUrl,
			BackPicUrl:  post.BackPicUrl,
		},
	}, nil
}

func (s *Server) GetPosts(ctx context.Context, req *pbdiscovery.GetPostsRequest) (*pbdiscovery.GetPostsResponse, error) {
	page, err := s.service.GetPosts(ctx, req.Cursor)
	if err != nil {
		return &pbdiscovery.GetPostsResponse{
			Success: false,
		}, err
	}

	data := make([]*pbdiscovery.Post, len(page.Data))
	for i, post := range page.Data {
		data[i] = &pbdiscovery.Post{
			Id:          post.ID,
			Owner:       post.OwnerID,
			FrontPicUrl: post.FrontPicUrl,
			BackPicUrl:  post.BackPicUrl,
		}
	}

	return &pbdiscovery.GetPostsResponse{
		Success: true,
		Cursor:  page.Cursor,
		Data:    data,
	}, nil
}

func (*Server) Close() {}
