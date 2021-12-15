package controller

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"

	"zzidun.tech/vforum0/dto"
	"zzidun.tech/vforum0/model"
	"zzidun.tech/vforum0/response"
	"zzidun.tech/vforum0/util"
)

func User_Register(ctx *gin.Context) {
	db := util.DB_Get()

	// 获取参数
	var request_user = model.User{}
	ctx.Bind(&request_user)
	name := request_user.Name
	email := request_user.Email
	password := request_user.Password

	// 数据验证
	if len(name) == 0 {
		name = util.String_Random(10)
	}
	if len(email) == 0 {
		response.Response_Fail_Make(ctx, nil, "邮箱格式错误")
		return
	}
	if len(password) < 6 {
		response.Response_Fail_Make(ctx, nil, "密码必须不少于6位")
		return
	}
	if util.Email_Exist(db, email) {
		response.Response_Fail_Make(ctx, nil, "邮箱已被注册")
		return
	}
	// 新建用户
	password_mixed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		response.Response_Make(ctx, http.StatusInternalServerError,
			http.StatusInternalServerError, nil, "加密错误")
		return
	}
	user := model.User{
		Name:     name,
		Email:    email,
		Password: string(password_mixed),
	}
	db.Create(&user)

	// 发放token
	token, err := util.Token_Release(&user)
	if err != nil {
		response.Response_Make(ctx, http.StatusInternalServerError,
			http.StatusInternalServerError, nil, "系统异常")
		log.Printf("token generate error : %v", err)
		return
	}

	// 返回结果
	response.Response_Success_Make(ctx, gin.H{"token": token}, "登陆成功")
	return
}

func User_Login(ctx *gin.Context) {
	db := util.DB_Get()

	// 获取参数
	var request_user = model.User{}
	ctx.Bind(&request_user)
	email := request_user.Email
	password := request_user.Password

	// 数据验证
	if len(email) == 0 {
		response.Response_Fail_Make(ctx, nil, "邮箱格式错误")
		return
	}
	if len(password) < 6 {
		response.Response_Fail_Make(ctx, nil, "密码必须不少于6位")
		return
	}

	// 判断邮箱是否存在
	var user model.User
	db.Where("email = ?", email).First(&user)
	if user.ID == 0 {
		response.Response_Fail_Make(ctx, nil, "邮箱未注册")
		return
	}

	// 判断密码
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		response.Response_Fail_Make(ctx, nil, "密码错误")
		return
	}

	// 发放token
	token, err := util.Token_Release(&user)
	if err != nil {
		response.Response_Make(ctx, http.StatusInternalServerError,
			http.StatusInternalServerError, nil, "系统异常")
		log.Printf("token generate error : %v", err)
		return
	}

	// 返回结果
	response.Response_Success_Make(ctx, gin.H{"token": token}, "登陆成功")
	return
}

func User_Info(ctx *gin.Context) {
	user, _ := ctx.Get("user")

	response.Response_Success_Make(ctx, gin.H{"user": dto.UserDto_Make(user.(model.User))}, "")
	return
}
