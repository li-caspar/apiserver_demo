package user

import (
	"apiserver/handle"
	"apiserver/model"
	"apiserver/pkg/auth"
	"apiserver/pkg/errno"
	"apiserver/pkg/token"
	"fmt"
	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var u model.UserModel

	if err := c.Bind(&u); err != nil {
		handle.SendResponse(c, errno.ErrBind, nil)
		return
	}

	d, err := model.GetUser(u.Username)
	if err != nil {
		handle.SendResponse(c, errno.ErrUserNotFound, nil)
		return
	}

	if err := auth.Compare(d.Password, u.Password); err != nil {
		handle.SendResponse(c, errno.ErrPasswordIncorrect, nil)
		return
	}

	t, err := token.Sign(c, token.Context{ID: d.Id, Username: d.Username}, "")
	fmt.Println(err)
	if err != nil {
		handle.SendResponse(c, errno.ErrToken, nil)
		return
	}

	handle.SendResponse(c, nil, model.Token{Token: t})

}
