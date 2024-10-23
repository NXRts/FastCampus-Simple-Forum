package memberships

import (
	"context"

	"github.com/NXRts/fsatcampus/internal/model/memberships"
)

type membershipsRepository interface {
	GetUser(ctx context.Context, email, username string) (*memberships.UserModel, error)
	CreateUser(ctx context.Context, model memberships.UserModel) error
}

type service struct {
	membershipsRepo membershipsRepository
}

func NewService(membershipsRepo membershipsRepository) *service {
	return &service{
		membershipsRepo: membershipsRepo,
	}
}
