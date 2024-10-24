package memberships

import "time"

type (
	SignUpRequest struct {
		Email    string `json:"email"`
		Username string `json:"username"`
		Password string `json:"password"`
	}
)

type (
	UserModel struct {
		Id       int64     `db:"id"`
		Email    string    `db:"email"`
		Username string    `db:"username"`
		Password string    `db:"password"`
		CreateAt time.Time `db:"create_at"`
		UpdateAt time.Time `db:"update_at"`
		CreateBy string    `db:"create_by"`
		UpdateBy string    `db:"update_by"`
	}
)
