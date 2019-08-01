package user

import (
	"apiserver/model"
	"github.com/gin-gonic/gin"
	"strconv"
)

func Test(c *gin.Context) {

	go func() {
		var u model.UserModel
		for i := 1000000; i < 2000000; i++ {
			u = model.UserModel{
				Username: "asdmin" + strconv.Itoa(i),
				Password: "123456",
			}
			u.Create()
		}
	}()

	go func() {
		var u model.UserModel
		for i := 2000000; i < 3000000; i++ {
			u = model.UserModel{
				Username: "asdmin" + strconv.Itoa(i),
				Password: "123456",
			}
			u.Create()
		}
	}()

	go func() {
		var u model.UserModel
		for i := 3000000; i < 4000000; i++ {
			u = model.UserModel{
				Username: "asdmin" + strconv.Itoa(i),
				Password: "123456",
			}
			u.Create()
		}
	}()

	go func() {
		var u model.UserModel
		for i := 4000000; i < 5000000; i++ {
			u = model.UserModel{
				Username: "asdmin" + strconv.Itoa(i),
				Password: "123456",
			}
			u.Create()
		}
	}()

	go func() {
		var u model.UserModel
		for i := 5000000; i < 6000000; i++ {
			u = model.UserModel{
				Username: "asdmin" + strconv.Itoa(i),
				Password: "123456",
			}
			u.Create()
		}
	}()

	go func() {
		var u model.UserModel
		for i := 7000000; i < 8000000; i++ {
			u = model.UserModel{
				Username: "asdmin" + strconv.Itoa(i),
				Password: "123456",
			}
			u.Create()
		}
	}()

}
