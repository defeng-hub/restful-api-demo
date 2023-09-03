package model

type JwtBlacklist struct {
	BASEMODEL
	Jwt string `gorm:"type:text;comment:jwt"`
}

func (f *JwtBlacklist) TableName() string {
	return "jwt_blacklists"
}

// 这个表需要定期清理
