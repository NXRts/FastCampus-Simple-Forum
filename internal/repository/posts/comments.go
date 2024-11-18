package posts

import (
	"context"

	"github.com/NXRts/fsatcampus/internal/model/posts"
)

func (r *repository) CreateComment(ctx context.Context, model posts.CommentModel) error {
	query := `INSERT INTO comments(post_id, user_id, comment_content, created_at, updated_at, created_by, updated_by) VALUES (?, ?, ?, ?, ?, ?, ?)`
	_, err := r.db.ExecContext(ctx, query, model.PostID, model.UserID, model.CommentContent, model.CreateAt, model.UpdateAt, model.CreateBy, model.UpdateBy)
	if err != nil {
		return err
	}
	return nil
}
