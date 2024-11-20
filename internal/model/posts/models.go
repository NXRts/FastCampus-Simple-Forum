package posts

import "time"

type (
	CreatePostRequest struct {
		PostTitle    string   `json:"postTitle"`
		PostContent  string   `json:"postContent"`
		PostHashtags []string `json:"postHashtags"`
	}
)

type (
	PostModel struct {
		Id           int64     `db:"id"`
		UserId       int64     `db:"user_id"`
		PostTitle    string    `db:"post_title"`
		PostContent  string    `db:"post_content"`
		PostHashtags string    `db:"post_hashtags"`
		CreateAt     time.Time `db:"created_at"`
		UpdateAt     time.Time `db:"updated_at"`
		CreateBy     string    `db:"created_by"`
		UpdateBy     string    `db:"updated_by"`
	}
)
