package memberships

import (
	"context"
	"errors"
	"time"

	"github.com/NXRts/fsatcampus/internal/model/memberships"
	"golang.org/x/crypto/bcrypt"
)

func (s *service) SingUp(ctx context.Context, req memberships.SingUpRequest) error {
	user, err := s.membershipsRepo.GetUser(ctx, req.Email, req.Usernamem)
	if err != nil {
		return err
	}

	if user != nil {
		return errors.New("username or email alredy exist")
	}

	pass, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	now := time.Now()
	model := memberships.UserModel{
		Email:     req.Email,
		Username:  req.Usernamem,
		Password:  string(pass),
		CreatedAt: now,
		UpdateAt:  now,
		CreatedBy: req.Email,
		UpdateBy:  req.Email,
	}
	return s.membershipsRepo.CreateUser(ctx, model)

}
