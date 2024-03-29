package model

import (
	"apiserver/pkg/auth"
	"apiserver/pkg/constvar"
	"fmt"
	"gopkg.in/go-playground/validator.v9"
)

type UserModel struct {
	BaseModel
	Username string `json:"username" gorm:"column:username;not null" binding:"required" validate:"min=1,max=32"`
	Password string `json:"password" gorm:"column:password;not null" binding:"required" validate:"min=5,max=128"`
}

func (u *UserModel) TableName() string {
	return "tb_users"
}

func (u *UserModel) Create() error {
	return DB.Self.Create(&u).Error
}

func (u *UserModel) Update() error {
	return DB.Self.Save(u).Error
}

func (u *UserModel) Delete() error {
	return DB.Self.Delete(&u).Error
}

func ListUser(username string, offset, limit int) ([]*UserModel, uint64, error) {
	if limit == 0 {
		limit = constvar.DefaultLimit
	}
	users := make([]*UserModel, 0)
	var count uint64
	var where string
	if len(username) > 0 {
		where = fmt.Sprintf("username like '%%%s%%'", username)
	}
	if err := DB.Self.Model(&UserModel{}).Where(where).Count(&count).Error; err != nil {
		return users, count, err
	}

	if err := DB.Self.Where(where).Limit(limit).Offset(offset).Find(&users).Error; err != nil {
		return users, count, err
	}

	return users, count, nil

}

func GetUser(username string) (*UserModel, error) {
	u := &UserModel{}
	d := DB.Self.Where("username=?", username).First(&u)
	return u, d.Error
}

func (u *UserModel) Encrypt() (err error) {
	u.Password, err = auth.Encrypt(u.Password)
	return
}

func (u *UserModel) Validate() error {
	validate := validator.New()
	return validate.Struct(u)
}
