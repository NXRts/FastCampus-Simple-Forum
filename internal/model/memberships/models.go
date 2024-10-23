package memberships

import "time"

type (
	SingUpRequest struct {
		Email     string `json:"email"`
		Usernamem string `json:"username"`
		Password  string `json:"password"`
	}
)

type (
	UserModel struct {
		Id        int64     `db:"id"`
		Email     string    `db:"email"`
		Username  string    `db:"username"`
		Password  string    `db:"password"`
		CreatedAt time.Time `db:"created_at"`
		UpdateAt  time.Time `db:"update_at"`
		CreatedBy string    `db:"created_by"`
		UpdateBy  string    `db:"upadte_by"`
	}
)
