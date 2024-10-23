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
