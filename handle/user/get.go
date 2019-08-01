package user

import (
	"apiserver/handle"
	"apiserver/model"
	"apiserver/pkg/errno"
	"apiserver/util"
	"github.com/gin-gonic/gin"
	"github.com/lexkong/log"
	"github.com/lexkong/log/lager"
)

func Get(c *gin.Context) {
	log.Info("User List function called.", lager.Data{"X-Request-Id:": util.GetReqID(c)})

	username := c.Param("username")

	if username == "" {
		handle.SendResponse(c, errno.ErrValidation, nil)
		return
	}

	user, err := model.GetUser(username)

	if err != nil {
		handle.SendResponse(c, err, nil)
		return
	}

	handle.SendResponse(c, nil, user)

}
