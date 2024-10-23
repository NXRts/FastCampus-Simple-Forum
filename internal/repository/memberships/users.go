package memberships

import (
	"context"

	"github.com/NXRts/fsatcampus/internal/model/memberships"
)

func (r *repository) GetUser(ctx context.Context, email, username string) (*memberships.UserModel, error) {
	query := `SELECT id, email, password, username, created_at, created_by, upadted_at, updated_by FROM users WHERE email = $1 OR username = $`
	row := r.db.QueryRowContext(ctx, query, email, username)

	var response memberships.UserModel
	err := row.Scan(&response.Id, &response.Email, &response.Password, &response.Username, &response.CreatedAt, &response.UpdateAt, &response.CreatedBy, &response.UpdateBy)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

//	func (r *repository) CreateUser(ctx context.Context, model memberships.UserModel) error {
//		query := `INSERT INTO users (email, password, username, created_at, created_by, upadted_at, updated_by) VALUES ($1, $2, $3, $4, $5, $6, $7)`
//		_, err := r.db.ExecContext(ctx, query, model.Email, model.Password, model.CreatedAt, model.CreatedBy, model.UpdateAt, model.UpdateBy)
//		if err != nil {
//			return err
//		}
//	}
func (r *repository) CreateUser(ctx context.Context, model memberships.UserModel) error {
	query := `INSERT INTO users (email, password, username, created_at, created_by, updated_at, updated_by) VALUES (?, ?, ?, ?, ?, ?, ?)`
	_, err := r.db.ExecContext(ctx, query, model.Email, model.Password, model.Username, model.CreatedAt, model.CreatedBy, model.UpdateAt, model.UpdateBy)
	if err != nil {
		return err
	}
	return nil
}
