package models

type UserData struct {
	UID      int64  `json:"uid string" db:"uid"`
	Name     string `json:"name" db:"name"`
	Email    string `json:"email" db:"email"`
	Password string `json:"-" db:"password"`
	Token    string `json:"token"`
}

type UserSignupParams struct {
	Name       string `json:"name" binding:"required" db:"username"`
	Password   string `json:"password" binding:"required" db:"password"`
	RePassword string `json:"rePassword" binding:"required,eqfield=Password"`
	Email      string `json:"email" binding:"required" db:"email"`
}

type UserLoginParams struct {
	Name     string `json:"name" binding:"required" db:"name"`
	Password string `json:"password" binding:"required" db:"password"`
}

type UserLogoutParams struct {
	UID   int64  `json:"uid string" db:"uid"`
	Name  string `db:"username" json:"name"`
	Email string `db:"email" json:"email"`
	Token string `json:"token"`
}

type UserUpdateParams struct {
	Token string `json:"token" binding:"required"`
	UID   int64  `json:"uid string" db:"uid"`
}
