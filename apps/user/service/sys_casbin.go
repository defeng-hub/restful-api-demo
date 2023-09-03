package service

import (
	"errors"
	"github.com/casbin/casbin/v2"
	"gorm.io/gorm"
	"restful-api-demo/apps/user"
	"restful-api-demo/apps/user/model"
	"restful-api-demo/apps/user/model/request"
	"restful-api-demo/conf"
	"sync"

	gormadapter "github.com/casbin/gorm-adapter/v3"
	_ "github.com/go-sql-driver/mysql"
)

type CasbinService struct {
	db *gorm.DB
}

func (s *CasbinService) Name() string {
	return user.AppName + ImplMap["sys_casbin"]
}
func (s *CasbinService) Config() {
	s.db, _ = conf.C().MySQL.GetGormDB()
}

// UpdateCasbin
//@description: 更新casbin权限
//@param: authorityId string, casbinInfos []request.CasbinInfo
//@return: error
func (casbinService *CasbinService) UpdateCasbin(authorityId string, casbinInfos []request.CasbinInfo) error {
	casbinService.ClearCasbin(0, authorityId)
	var rules [][]string
	for _, v := range casbinInfos {
		cm := model.CasbinModel{
			Ptype:       "p",
			AuthorityId: authorityId,
			Path:        v.Path,
			Method:      v.Method,
		}
		rules = append(rules, []string{cm.AuthorityId, cm.Path, cm.Method})
	}
	e := casbinService.Casbin()
	success, _ := e.AddPolicies(rules)
	if !success {
		return errors.New("存在相同api,添加失败,请联系管理员")
	}
	return nil
}

// UpdateCasbinApi
//@description: API更新随动
//@param: oldPath string, newPath string, oldMethod string, newMethod string
//@return: error
func (casbinService *CasbinService) UpdateCasbinApi(oldPath string, newPath string, oldMethod string, newMethod string) error {
	err := casbinService.db.Table("casbin_rule").Model(&model.CasbinModel{}).Where("v1 = ? AND v2 = ?", oldPath, oldMethod).Updates(map[string]interface{}{
		"v1": newPath,
		"v2": newMethod,
	}).Error
	return err
}

// GetPolicyPathByAuthorityId
//@description: 获取权限列表
//@param: authorityId string
//@return: pathMaps []request.CasbinInfo
func (casbinService *CasbinService) GetPolicyPathByAuthorityId(authorityId string) (pathMaps []request.CasbinInfo) {
	e := casbinService.Casbin()
	list := e.GetFilteredPolicy(0, authorityId)
	for _, v := range list {
		pathMaps = append(pathMaps, request.CasbinInfo{
			Path:   v[1],
			Method: v[2],
		})
	}
	return pathMaps
}

// ClearCasbin
//@description: 清除匹配的权限
//@param: v int, p ...string
//@return: bool
func (casbinService *CasbinService) ClearCasbin(v int, p ...string) bool {
	e := casbinService.Casbin()
	success, _ := e.RemoveFilteredPolicy(v, p...)
	return success
}

var (
	syncedEnforcer *casbin.SyncedEnforcer
	once           sync.Once
)

// Casbin
//@description: 持久化到数据库  引入自定义规则
//@return: *casbin.Enforcer
func (casbinService *CasbinService) Casbin() *casbin.SyncedEnforcer {

	once.Do(func() {
		casbinService.db, _ = conf.C().MySQL.GetGormDB()
		a, err := gormadapter.NewAdapterByDB(casbinService.db)
		if err != nil {
			conf.L().Named("Init Casbin").Error(err)
		}
		syncedEnforcer, err = casbin.NewSyncedEnforcer("./etc/rbac_model.conf", a)
		if err != nil {
			conf.L().Named("Init Casbin").Error(err)
		}
	})

	_ = syncedEnforcer.LoadPolicy()
	return syncedEnforcer
}
