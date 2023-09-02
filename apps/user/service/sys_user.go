package service

import (
	"errors"
	"fmt"
	"restful-api-demo/apps/user"
	"restful-api-demo/common/logger"
	"restful-api-demo/conf"

	"restful-api-demo/apps/user/common/request"
	"restful-api-demo/apps/user/model"
	"restful-api-demo/utils"

	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

// UserService
type UserService struct {
	db *gorm.DB
	l  logger.Logger
}

func (s *UserService) Config() {
	s.db, _ = conf.C().MySQL.GetGormDB()
	s.l = conf.L().Named(s.Name())
}
func (s *UserService) Name() string {
	return user.AppName + ImplMap["sys_user"]
}

func (userService *UserService) Register(u model.SysUser) (err error, userInter *model.SysUser) {
	var user *model.SysUser
	if !errors.Is(userService.db.Where("username = ?", u.Username).First(&user).Error, gorm.ErrRecordNotFound) { // 判断用户名是否注册
		return errors.New("用户名已注册"), userInter
	}
	// 否则 附加uuid 密码md5简单加密 注册
	u.Password = utils.MD5V([]byte(u.Password))
	u.UUID = uuid.NewV4()
	err = userService.db.Create(&u).Error
	return err, &u
}

// Login
//@description: 用户登录
//@param: u *model.SysUser
//@return: err error, userInter *model.SysUser
func (userService *UserService) Login(u *model.SysUser) (err error, userInter *model.SysUser) {
	if nil == userService.db {
		return fmt.Errorf("db not init"), nil
	}

	var user model.SysUser
	u.Password = utils.MD5V([]byte(u.Password))
	err = userService.db.Where("username = ? AND password = ?", u.Username, u.Password).
		Preload("Authorities").Preload("Authority").First(&user).Error
	//user.Authority.DefaultRouter = "main"
	return err, &user
}

// ChangePassword
//@description: 修改用户密码
//@param: u *model.SysUser, newPassword string
//@return: err error, userInter *model.SysUser
func (userService *UserService) ChangePassword(u *model.SysUser, newPassword string) (err error, userInter *model.SysUser) {
	var user model.SysUser
	u.Password = utils.MD5V([]byte(u.Password))
	err = userService.db.Where("username = ? AND password = ?", u.Username, u.Password).First(&user).Update("password", utils.MD5V([]byte(newPassword))).Error
	return err, u
}

// GetUserInfoList
//@description: 分页获取数据
//@param: info request.PageInfo
//@return: err error, list interface{}, total int64
func (userService *UserService) GetUserInfoList(info request.PageInfo) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := userService.db.Model(&model.SysUser{})
	var userList []model.SysUser
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Preload("Authorities").Preload("Authority").Find(&userList).Error
	return err, userList, total
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: SetUserAuthority
//@description: 设置一个用户的权限
//@param: uuid uuid.UUID, authorityId string
//@return: err error

func (userService *UserService) SetUserAuthority(id uint, uuid uuid.UUID, authorityId string) (err error) {
	assignErr := userService.db.Where("sys_user_id = ? AND sys_authority_authority_id = ?", id, authorityId).First(&model.SysUseAuthority{}).Error
	if errors.Is(assignErr, gorm.ErrRecordNotFound) {
		return errors.New("该用户无此角色")
	}
	err = userService.db.Where("uuid = ?", uuid).First(&model.SysUser{}).Update("authority_id", authorityId).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: SetUserAuthorities
//@description: 设置一个用户的权限
//@param: id uint, authorityIds []string
//@return: err error

func (userService *UserService) SetUserAuthorities(id uint, authorityIds []string) (err error) {
	return userService.db.Transaction(func(tx *gorm.DB) error {
		TxErr := tx.Delete(&[]model.SysUseAuthority{}, "sys_user_id = ?", id).Error
		if TxErr != nil {
			return TxErr
		}
		useAuthority := []model.SysUseAuthority{}
		for _, v := range authorityIds {
			useAuthority = append(useAuthority, model.SysUseAuthority{
				id, v,
			})
		}
		TxErr = tx.Create(&useAuthority).Error
		if TxErr != nil {
			return TxErr
		}
		TxErr = tx.Where("id = ?", id).First(&model.SysUser{}).Update("authority_id", authorityIds[0]).Error
		if TxErr != nil {
			return TxErr
		}
		// 返回 nil 提交事务
		return nil
	})
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteUser
//@description: 删除用户
//@param: id float64
//@return: err error

func (userService *UserService) DeleteUser(id float64) (err error) {
	var user model.SysUser
	err = userService.db.Where("id = ?", id).Delete(&user).Error
	if err != nil {
		return err
	}
	err = userService.db.Delete(&[]model.SysUseAuthority{}, "sys_user_id = ?", id).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: SetUserInfo
//@description: 设置用户信息
//@param: reqUser model.SysUser
//@return: err error, user model.SysUser

func (userService *UserService) SetUserInfo(reqUser model.SysUser) (err error, user model.SysUser) {
	err = userService.db.Updates(&reqUser).Error
	return err, reqUser
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetUserInfo
//@description: 获取用户信息
//@param: uuid uuid.UUID
//@return: err error, user model.SysUser

func (userService *UserService) GetUserInfo(uuid uuid.UUID) (err error, user model.SysUser) {
	var reqUser model.SysUser
	err = userService.db.Preload("Authorities").Preload("Authority").First(&reqUser, "uuid = ?", uuid).Error
	if err != nil {
		return err, reqUser
	}
	var am model.SysMenu
	ferr := userService.db.First(&am, "name = ? AND authority_id = ?", reqUser.Authority.DefaultRouter, reqUser.AuthorityId).Error
	if errors.Is(ferr, gorm.ErrRecordNotFound) {
		reqUser.Authority.DefaultRouter = "404"
	}
	return err, reqUser
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@function: FindUserById
//@description: 通过id获取用户信息
//@param: id int
//@return: err error, user *model.SysUser

func (userService *UserService) FindUserById(id int) (err error, user *model.SysUser) {
	var u model.SysUser
	err = userService.db.Where("`id` = ?", id).First(&u).Error
	return err, &u
}

func (userService *UserService) FindUserByUsername(name string) (err error, user *model.SysUser) {
	var u model.SysUser
	err = userService.db.Where("`username` = ?", name).First(&u).Error
	return err, &u
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@function: FindUserByUuid
//@description: 通过uuid获取用户信息
//@param: uuid string
//@return: err error, user *model.SysUser

func (userService *UserService) FindUserByUuid(uuid string) (err error, user *model.SysUser) {
	var u model.SysUser
	if err = userService.db.Where("`uuid` = ?", uuid).First(&u).Error; err != nil {
		return errors.New("用户不存在"), &u
	}
	return nil, &u
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: resetPassword
//@description: 修改用户密码
//@param: ID uint
//@return: err error

func (userService *UserService) ResetPassword(ID uint) (err error) {
	err = userService.db.Model(&model.SysUser{}).Where("id = ?", ID).Update("password", utils.MD5V([]byte("123456"))).Error
	return err
}
