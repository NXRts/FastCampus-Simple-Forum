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
		CreateAt     time.Time `db:"create_at"`
		UpdateAt     time.Time `db:"update_at"`
		CreateBy     string    `db:"create_by"`
		UpdateBy     string    `db:"update_by"`
	}
)

// post_title VARCHAR(250) NOT NULL,
// post_content LONGTEXT NOTy NULL,
// post_hashtags LONGTEXT NOT NULL,
// created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
// updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
// created_by LONGTEXT NOT NULL,
// updated_by LONGTEXT NOT NULL
