package user

import (
	"apiserver/handle"
	"apiserver/model"
	"apiserver/pkg/errno"
	"apiserver/util"
	"github.com/gin-gonic/gin"
	"github.com/lexkong/log"
	"github.com/lexkong/log/lager"
	"strconv"
)

func Update(c *gin.Context) {
	log.Info("User Update function called.", lager.Data{"X-Request-Id:": util.GetReqID(c)})
	userId, _ := strconv.Atoi(c.Param("id"))

	var u model.UserModel

	if err := c.Bind(&u); err != nil {
		handle.SendResponse(c, errno.ErrBind, nil)
		return
	}

	u.Id = uint64(userId)

	if err := u.Validate(); err != nil {
		handle.SendResponse(c, errno.ErrValidation, nil)
		return
	}

	if err := u.Encrypt(); err != nil {
		handle.SendResponse(c, errno.ErrEncrypt, nil)
		return
	}

	if err := u.Update(); err != nil {
		handle.SendResponse(c, errno.ErrDatabase, nil)
		return
	}

	handle.SendResponse(c, nil, nil)

}
