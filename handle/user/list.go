package user

import (
	"apiserver/handle"
	"apiserver/service"
	"apiserver/util"
	"github.com/gin-gonic/gin"
	"github.com/lexkong/log"
	"github.com/lexkong/log/lager"
	"strconv"
)

func List(c *gin.Context) {
	log.Info("User List function called.", lager.Data{"X-Request-Id:": util.GetReqID(c)})

	username := c.Query("username")
	offset, _ := strconv.Atoi(c.Query("offset"))
	limit, _ := strconv.Atoi(c.Query("limit"))

	log.Infof("User List funcation params:Username:%s, Offset:%d, Limit:%d", username, offset, limit)

	infos, count, err := service.ListUser(username, offset, limit)

	if err != nil {
		handle.SendResponse(c, err, nil)
		return
	}

	handle.SendResponse(c, nil, ListResponse{
		TotalCount: count,
		UserList:   infos,
	})

}
