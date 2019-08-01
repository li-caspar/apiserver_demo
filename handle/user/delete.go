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

func Delete(c *gin.Context) {
	log.Info("User Delete function called.", lager.Data{"X-Request-Id:": util.GetReqID(c)})
	userId, _ := strconv.Atoi(c.Param("id"))

	var u model.UserModel

	u.Id = uint64(userId)

	if err := u.Delete(); err != nil {
		handle.SendResponse(c, errno.ErrDatabase, nil)
		return
	}

	handle.SendResponse(c, nil, nil)

}
