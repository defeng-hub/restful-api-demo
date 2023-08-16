package init

import (
	"github.com/infraboard/mcube/logger/zap"
	"gorm.io/gorm"
	"os"
	"restful-api-demo/apps/user/model"
)

//sys_data_authority_id
//sys_users

func RegisterTables(db *gorm.DB) {
	err := db.AutoMigrate(
		model.SysBaseMenu{},
		model.SysAuthority{},
		model.SysUser{},
	)
	if err != nil {
		zap.L().Named("User Init").Error("register table failed", zap.Error(err))
		os.Exit(0)
	}
	zap.L().Named("User Init").Info("register table success")
}