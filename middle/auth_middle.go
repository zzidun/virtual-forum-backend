package middle

/*
func Auth_Middle() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// 获取authrization
		authorization := ctx.GetHeader("Authorization")

		// 验证格式
		if authorization == "" || !strings.HasPrefix(authorization, "Bearer") {
			response.Response_Make(ctx, http.StatusUnauthorized, http.StatusUnauthorized, nil, "认证失败")
			ctx.Abort()
			return
		}

		authorization = authorization[7:]
		token, claims, err := util.Token_Parse(authorization)
		if err != nil || !token.Valid {
			response.Response_Make(ctx, http.StatusUnauthorized, http.StatusUnauthorized, nil, "认证失败")
			ctx.Abort()
			return
		}

		// 获取userid
		user_id := claims.UserId
		db := model.DB_Get()
		var user model.User
		db.First(&user, user_id)

		// 用户不存在
		if user.ID == 0 {
			response.Response_Make(ctx, http.StatusUnauthorized, http.StatusUnauthorized, nil, "认证失败")
			ctx.Abort()
			return
		}

		// 认证成功
		ctx.Set("user", user)
		ctx.Next()
	}
}
*/
