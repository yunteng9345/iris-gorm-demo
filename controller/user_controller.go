package controller

import (
	"github.com/kataras/iris"
	"iris-gorm-demo/models"
	"iris-gorm-demo/service"
	"log"
)

type UserController struct {
	Ctx     iris.Context
	Service service.UserService
}

func NewUserController() *UserController {
	return &UserController{Service: service.NewUserService()}
}

func (g *UserController) PostLogin() models.Result {
	var m map[string]string
	err := g.Ctx.ReadJSON(&m)
	if err != nil {
		log.Println("ReadJSON Error:", err)
	}
	result := g.Service.Login(m)
	return result
}

func (g *UserController) PostSave() (result models.Result) {
	var user models.User
	if err := g.Ctx.ReadJSON(&user); err != nil {
		log.Println(err)
		result.Msg = "数据错误"
		return
	}

	return g.Service.Save(user)
}
