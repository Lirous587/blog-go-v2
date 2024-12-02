package repository

import (
	"blog/models"
	"fmt"
	"github.com/jmoiron/sqlx"
)

const (
	userValiteFailed = "用户身份验证失败"
)

type UserRepo interface {
	//Signup(data *models.UserSignupParams) error
	//Login(data *models.UserLoginParams) error
	//Logout(data *models.UserLoginParams) error
	//Update(data *models.UserUpdateParams) error
	CheckExist(data *models.UserSignupParams) (bool, error)
	Save(data *models.UserData) error
	Validate(data *models.UserLoginParams) (*models.UserData, error)
}

type UserRepoMySQL struct {
	db *sqlx.DB
}

func NewUserRepoMySQL(db *sqlx.DB) *UserRepoMySQL {
	return &UserRepoMySQL{
		db: db,
	}
}

func (r *UserRepoMySQL) CheckExist(data *models.UserSignupParams) (bool, error) {
	//用户名
	sqlStr := `SELECT count(*) FROM user where name = ? OR email = ?`
	var count int8
	if err := db.Get(&count, sqlStr, data.Name, data.Email); err != nil {
		return false, err
	}
	return count > 0, nil
}

func (r *UserRepoMySQL) Save(data *models.UserData) error {
	sqlStr := `INSERT INTO user (name,password,email,uid) values(:name,:password,:email,:uid)`
	_, err := db.NamedExec(sqlStr, data)
	return err
}

func (r *UserRepoMySQL) Validate(data *models.UserLoginParams) (*models.UserData, error) {
	sqlStr := `SELECT uid, name, password, email FROM user WHERE name = ?`
	waitValidatePwd := data.Password
	ret := new(models.UserData)
	if err := db.Get(ret, sqlStr, data.Name); err != nil {
		return nil, err
	}
	if waitValidatePwd != ret.Password {
		return nil, fmt.Errorf(userValiteFailed)
	}
	return ret, nil
}

//func Login(u *models.User) (err error) {
//	oldPassword := *(&u.Password)
//	sqlStr := `select user_id,username,password from users where username = ?`
//	err = db.Get(u, sqlStr, u.Username)
//	if errors.Is(err, sql.ErrNoRows) {
//		return err
//	} else if err != nil {
//		return err
//	}
//	encryptedPassword := encryptPassword(oldPassword)
//
//	if encryptedPassword != u.Password {
//		return errors.New(loginFailed)
//	}
//	return nil
//}
//
//func UpdateUserMsg(p *models.UserParams, id int64) (err error) {
//	if err = CheckUserExist2(p.Username, p.Email); err != nil {
//		return err
//	}
//	sqlStr := `UPDATE users SET username = ?,password = ?,email = ? where user_id = ?`
//	_, err = db.Exec(sqlStr, p.Username, p.Password, p.Email, id)
//	return err
//}
//
//func GetUserMsg(p *models.UserParams, id int64) error {
//	sqlStr := `SELECT username,email FROM users where user_id = ?`
//	return db.Get(p, sqlStr, id)
//}
