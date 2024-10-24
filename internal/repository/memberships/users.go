package memberships

import (
	"context"
	"database/sql"

	"github.com/NXRts/fsatcampus/internal/model/memberships"
)

func (r *repository) GetUser(ctx context.Context, email, username string) (*memberships.UserModel, error) {
	query := `SELECT id, email, password, username, create_at, update_at, create_by, update_by FROM users WHERE email = ? OR username = ?`
	row := r.db.QueryRowContext(ctx, query, email, username)

	var response memberships.UserModel
	err := row.Scan(&response.Id, &response.Email, &response.Password, &response.Username, &response.CreateAt, &response.UpdateAt, &response.CreateBy, &response.UpdateBy)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &response, nil
}

func (r *repository) CreateUser(ctx context.Context, model memberships.UserModel) error {
	query := `INSERT INTO users (email, password, username, create_at, create_by, update_at, update_by) VALUES (?, ?, ?, ?, ?, ?, ?)`
	_, err := r.db.ExecContext(ctx, query, model.Email, model.Password, model.Username, model.CreateAt, model.CreateBy, model.UpdateAt, model.UpdateBy)
	if err != nil {
		return err
	}
	return nil
}
