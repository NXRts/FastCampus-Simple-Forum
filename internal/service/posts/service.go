package posts

import (
	"context"

	config "github.com/NXRts/fsatcampus/internal/configs"
	"github.com/NXRts/fsatcampus/internal/model/posts"
)

type postRepository interface {
	CreatePost(ctx context.Context, model posts.PostModel) error
}

type service struct {
	cfg      *config.Config
	postRepo postRepository
}

func NewService(cfg *config.Config, postRepo postRepository) *service {
	return &service{
		cfg:      cfg,
		postRepo: postRepo,
	}
}
