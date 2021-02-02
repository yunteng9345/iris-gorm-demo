package service

import (
	"print-chn/models"
	"print-chn/repo"
	"print-chn/utils"
)

type UserService interface {
	Login(m map[string]string) (result models.Result)
	Save(user models.User) (result models.Result)
}

type userService struct {

}

func NewUserService() UserService  {
	return &userService{}
}

var userRespo = repo.NewUserRepository()

// login
func (u userService)Login(m map[string]string) (result models.Result)  {
	if m["username"] == "" {
		result.Code = -1
		result.Msg = "请输入用户名"
		return
	}
	if m["password"] == "" {
		result.Code = -1
		result.Msg = "请输入密码"
		return
	}
	user := userRespo.GetUserByUserNameAndPwd(m["username"],utils.GetMD5String(m["password"]))
	if user.ID == 0 {
		result.Code = -1
		result.Msg = "用户名或密码错误!"
	}
	//user.Token =
	result.Code = 0
	result.Data = user
	result.Msg = "登录成功"
	return
}

// save
func (u userService)Save(user models.User) (result models.Result) {
	if user.ID == 0 {
		agen := userRespo.GetUserByUsername(user.Username)
		if agen.ID != 0 {
			result.Msg = "登录名重复，保存失败"
		}
	}
	code, p := userRespo.Save(user)
	if code == -1 {
		result.Code = -1
		result.Msg = "保存失败"
		return
	}
	result.Code = 0
	result.Data = p
	return
}
