package posts

import "time"

type (
	CreateCommentRequest struct {
		CommentContent string `json:"commentContent"`
	}
)

type (
	CommentModel struct {
		ID             int64     `db:"id"`
		PostID         int64     `db:"post_id"`
		UserID         int64     `db:"user_id"`
		CommentContent string    `db:"comment_content"`
		CreateAt       time.Time `db:"created_at"`
		UpdateAt       time.Time `db:"updated_at"`
		CreateBy       string    `db:"created_by"`
		UpdateBy       string    `db:"updated_by"`
	}
)
