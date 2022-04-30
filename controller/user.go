package controller

import (
	"database/sql"
	"errors"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"

	"zzidun.tech/vforum0/dao"
	"zzidun.tech/vforum0/dto"
	"zzidun.tech/vforum0/model"
	"zzidun.tech/vforum0/response"
	"zzidun.tech/vforum0/util"
)

func UserRegister(user *model.User) (err error) {
	sqlStr := "select count(user_id) from user where username = ?"
	var count int64
	err = db.Get(&count, sqlStr, user.UserName)
	if err != nil && err != sql.ErrNoRows {
		return err
	}
	if count > 0 {
		// 用户已存在
		return ErrorUserExit
	}
	// 生成user_id
	userID, err := snowflake.GetID()
	if err != nil {
		return ErrorGenIDFailed
	}
	// 生成加密密码
	password := encryptPassword([]byte(user.Password))
	// 把用户插入数据库
	sqlStr = "insert into user(user_id, username, password) values (?,?,?)"
	_, err = db.Exec(sqlStr, userID, user.UserName, password)
	return
}

func SignUpHandler(c *gin.Context) {
	// 1.获取请求参数 2.校验数据有效性
	var fo *models.RegisterForm
	if err := c.ShouldBindJSON(&fo); err != nil {
		// 请求参数有误，直接返回响应
		zap.L().Error("SiginUp with invalid param", zap.Error(err))
		// 判断err是不是 validator.ValidationErrors类型的errors
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			// 非validator.ValidationErrors类型错误直接返回
			ResponseError(c, CodeInvalidParams) // 请求参数错误
			return
		}
		// validator.ValidationErrors类型错误则进行翻译
		ResponseErrorWithMsg(c, CodeInvalidParams, removeTopStruct(errs.Translate(trans)))
		return // 翻译错误
	}

	// 3.业务处理——注册用户
	if err := logic.SignUp(fo); err != nil {
		zap.L().Error("logic.signup failed", zap.Error(err))
		if errors.Is(err, mysql.ErrorUserExit) {
			ResponseError(c, CodeUserExist)
			return
		}
		ResponseError(c, CodeServerBusy)
		return

		if err != nil {
			zap.L().Error("mysql.Register() failed", zap.Error(err))
			ResponseError(c, CodeServerBusy)
			return
		}
	}
	//返回响应
	ResponseSuccess(c, nil)
}

func User_Register(ctx *gin.Context) {
	db := model.DB_Get()

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
	if Email_Exist(db, email) {
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
	db := dao.DatabaseGet()

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

func Email_Exist(db *gorm.DB, email string) bool {
	var user model.User
	db.Where("email = ?", email).First(&user)
	if user.ID != 0 {
		return true
	}
	return false
}

func UserCurrentIDGet(c *gin.Context) (userId uint, err error) {
	_userType, ok := c.Get("userType")
	if !ok {
		err = ErrorUserNotLogin
		return
	}
	userType, ok := _userType.(uint)
	if !ok || userType != 0 {
		err = ErrorUserNotLogin
		return
	}

	_userId, ok := c.Get("userId")
	if !ok {
		err = ErrorUserNotLogin
		return
	}
	userId, ok = _userId.(uint)
	if !ok {
		err = ErrorUserNotLogin
		return
	}
	return
}
