package api

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"math/rand"
	"restful-api-demo/apps/user/model"
	"restful-api-demo/apps/user/service"
	"restful-api-demo/common/logger"
	"restful-api-demo/conf"
	"strconv"

	utils "restful-api-demo/apps/user/common"
	"restful-api-demo/apps/user/common/request"
	"restful-api-demo/apps/user/common/response"
	modelReq "restful-api-demo/apps/user/model/request"
	modelResp "restful-api-demo/apps/user/model/response"
)

type UserApi struct {
	Srv *service.UserService
	L   logger.Logger
}

//生产n位随机字符
func generateRandomString(length int) string {
	charset := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	result := make([]byte, length)

	for i := 0; i < length; i++ {
		result[i] = charset[rand.Intn(len(charset))]
	}

	return string(result)
}

// @Summary 用户登录
func (b *UserApi) Login(c *gin.Context) {
	var loginform struct {
		Username string `json:"username"` // 用户名
		Password string `json:"password"` // 密码
	}

	err := c.ShouldBindJSON(&loginform)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	rule := utils.Rules{"Username": {"required"}, "Password": {"required"}}
	if err := utils.Verify(loginform, rule); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	u := &model.SysUser{Username: loginform.Username, Password: loginform.Password}
	if err, user := b.Srv.Login(u); err != nil {
		b.L.Error("登陆失败! 用户名不存在或者密码错误, error:", err)
		response.FailWithMessage("用户名不存在或者密码错误", c)
	} else {
		b.tokenNext(c, *user)
	}
}

// @Summary 登录以后签发jwt
func (b *UserApi) tokenNext(c *gin.Context, user model.SysUser) {
	j := &utils.JWT{SigningKey: []byte(conf.C().Jwt.SigningKey)} // 唯一签名
	claims := j.CreateClaims(modelReq.BaseClaims{
		UUID:        user.UUID,
		ID:          user.ID,
		NickName:    user.NickName,
		Username:    user.Username,
		AuthorityId: user.AuthorityId,
	})
	token, err := j.CreateToken(claims)
	if err != nil {
		b.L.Error("获取token失败!", zap.Error(err))
		response.FailWithMessage("获取token失败", c)
		return
	}

	// token入库redis
	//var jwt service.JwtService
	//if err := jwt.SetRedisJWT(token, user.Username); err != nil {
	//	b.L.Error("Redis写入token失败!", zap.Error(err))
	//	response.FailWithMessage("获取token失败", c)
	//	return
	//}

	//if !b.Srv.VerifyAllowLogin() {
	//	b.L.Error("阻止登录!", zap.Error(err))
	//	response.FailWithMessage("获取token失败", c)
	//	return
	//}

	response.OkWithDetailed(modelResp.LoginResponse{
		User:      user,
		Token:     token,
		ExpiresAt: claims.StandardClaims.ExpiresAt * 1000,
	}, "登录成功", c)
	return
}

// @Summary 用户注册账号
func (b *UserApi) Register(c *gin.Context) {
	var r modelReq.Register
	_ = c.ShouldBindJSON(&r)

	rule := utils.Rules{"Username": {"required"}, "NickName": {"required"},
		"Password": {"required"}, "AuthorityId": {"required"}}

	if err := utils.Verify(r, rule); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	var authorities []model.SysAuthority
	for _, v := range r.AuthorityIds {
		authorities = append(authorities, model.SysAuthority{
			AuthorityId: v,
		})
	}
	user := model.SysUser{Username: r.Username, NickName: r.NickName, Password: r.Password, HeaderImg: r.HeaderImg,
		AuthorityId: r.AuthorityId, Authorities: authorities}
	err, userReturn := b.Srv.Register(user)
	if err != nil {
		b.L.Error("注册失败!", zap.Error(err))
		response.FailWithDetailed(modelResp.SysUserResponse{User: nil}, "注册失败", c)
	} else {
		response.OkWithDetailed(modelResp.SysUserResponse{User: userReturn}, "注册成功", c)
	}
}

// @Summary 用户修改密码
func (b *UserApi) ChangePassword(c *gin.Context) {
	var user modelReq.ChangePasswordStruct
	_ = c.ShouldBindJSON(&user)
	if err := utils.Verify(user, utils.ChangePasswordVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	u := &model.SysUser{Username: user.Username, Password: user.Password}
	if err, _ := b.Srv.ChangePassword(u, user.NewPassword); err != nil {
		b.L.Error("修改失败!", zap.Error(err))
		response.FailWithMessage("修改失败，原密码与当前账户不符", c)
	} else {
		response.OkWithMessage("修改成功", c)
	}
}

// @Summary 分页获取用户列表
func (b *UserApi) GetUserList(c *gin.Context) {
	var pageInfo request.PageInfo
	_ = c.ShouldBindJSON(&pageInfo)
	if err := utils.Verify(pageInfo, utils.PageInfoVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err, list, total := b.Srv.GetUserInfoList(pageInfo); err != nil {
		b.L.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取成功", c)
	}
}

// @Summary 更改用户权限
func (b *UserApi) SetUserAuthority(c *gin.Context) {
	var sua modelReq.SetUserAuth
	_ = c.ShouldBindJSON(&sua)
	if UserVerifyErr := utils.Verify(sua, utils.SetUserAuthorityVerify); UserVerifyErr != nil {
		response.FailWithMessage(UserVerifyErr.Error(), c)
		return
	}
	userID := utils.GetUserID(c)
	uuid := utils.GetUserUuid(c)
	if err := b.Srv.SetUserAuthority(userID, uuid, sua.AuthorityId); err != nil {
		b.L.Error("修改失败!", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
	} else {
		claims := utils.GetUserInfo(c)
		j := &utils.JWT{SigningKey: []byte(conf.C().Jwt.SigningKey)} // 唯一签名
		claims.AuthorityId = sua.AuthorityId
		if token, err := j.CreateToken(*claims); err != nil {
			b.L.Error("修改失败!", zap.Error(err))
			response.FailWithMessage(err.Error(), c)
		} else {
			c.Header("new-token", token)
			c.Header("new-expires-at", strconv.FormatInt(claims.ExpiresAt, 10))
			response.OkWithMessage("修改成功", c)
		}

	}
}

// @Summary 设置用户权限
func (b *UserApi) SetUserAuthorities(c *gin.Context) {
	var sua modelReq.SetUserAuthorities
	err := c.ShouldBindJSON(&sua)
	if err != nil {
		response.FailWithMessage("数据校验失败", c)
	}
	if err := b.Srv.SetUserAuthorities(sua.ID, sua.AuthorityIds); err != nil {
		b.L.Error("修改失败!", zap.Error(err))
		response.FailWithMessage("修改失败", c)
	} else {
		response.OkWithMessage("修改成功", c)
	}
}

// @Summary 删除用户
func (b *UserApi) DeleteUser(c *gin.Context) {
	var reqId request.GetById
	_ = c.ShouldBindJSON(&reqId)
	if err := utils.Verify(reqId, utils.IdVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	jwtId := utils.GetUserID(c)
	if jwtId == uint(reqId.ID) {
		response.FailWithMessage("删除失败, 自杀失败", c)
		return
	}
	if err := b.Srv.DeleteUser(reqId.ID); err != nil {
		b.L.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Summary 设置用户信息
func (b *UserApi) SetUserInfo(c *gin.Context) {
	var user model.SysUser
	_ = c.ShouldBindJSON(&user)
	if err := utils.Verify(user, utils.IdVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err, ReqUser := b.Srv.SetUserInfo(user); err != nil {
		b.L.Error("设置失败!", zap.Error(err))
		response.FailWithMessage("设置失败", c)
	} else {
		response.OkWithDetailed(gin.H{"userInfo": ReqUser}, "设置成功", c)
	}
}

// @Summary 设置用户信息
func (b *UserApi) SetSelfInfo(c *gin.Context) {
	var user model.SysUser
	err := c.ShouldBindJSON(&user)
	if err != nil {
		response.FailWithDetailed(err, "设置失败", c)
	}

	user.ID = utils.GetUserID(c)
	if err, ReqUser := b.Srv.SetUserInfo(user); err != nil {
		b.L.Error("设置失败!", zap.Error(err))
		response.FailWithMessage("设置失败", c)
	} else {
		response.OkWithDetailed(gin.H{"userInfo": ReqUser}, "设置成功", c)
	}
}

// @Summary 获取用户信息
func (b *UserApi) GetUserInfo(c *gin.Context) {
	uuid := utils.GetUserUuid(c)
	if err, ReqUser := b.Srv.GetUserInfo(uuid); err != nil {
		b.L.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(gin.H{"userInfo": ReqUser}, "获取成功", c)
	}
}

// @Summary 用户修改密码
func (b *UserApi) ResetPassword(c *gin.Context) {
	var user model.SysUser
	_ = c.ShouldBindJSON(&user)
	if err := b.Srv.ResetPassword(user.ID); err != nil {
		b.L.Error("重置失败!", zap.Error(err))
		response.FailWithMessage("重置失败"+err.Error(), c)
	} else {
		response.OkWithMessage("重置成功", c)
	}
}
