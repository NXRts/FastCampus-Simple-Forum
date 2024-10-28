package memberships

import (
	"context"

	config "github.com/NXRts/fsatcampus/internal/configs"
	"github.com/NXRts/fsatcampus/internal/model/memberships"
)

type membershipsRepository interface {
	GetUser(ctx context.Context, email, username string) (*memberships.UserModel, error)
	CreateUser(ctx context.Context, model memberships.UserModel) error
}

type service struct {
	cfg             *config.Config
	membershipsRepo membershipsRepository
}

func NewService(cfg *config.Config, membershipsRepo membershipsRepository) *service {
	return &service{
		cfg:             cfg,
		membershipsRepo: membershipsRepo,
	}
}
