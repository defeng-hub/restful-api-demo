package host

import (
	"github.com/go-playground/validator/v10"
	"time"
)

type Vendor int

var (
	validate = validator.New()
)

const (
	PrivateIDC Vendor = iota
	Tencent
	AliYun
	HuaWei
)

type QueryHostRequest struct {
	PageSize   uint64 `json:"page_size,omitempty"`
	PageNumber uint64 `json:"page_number,omitempty"`
	Keywords   string `json:"kws"`
}

func NewQueryHostRequest(pageSize uint64, pageNumber uint64, keywords string) *QueryHostRequest {
	return &QueryHostRequest{PageSize: pageSize, PageNumber: pageNumber, Keywords: keywords}
}

func (q *QueryHostRequest) OffSet() int64 {
	return int64((q.PageNumber - 1) * q.PageSize)
}

type HostSet struct {
	Total int     `json:"total"`
	Items []*Host `json:"items"`
}

func NewHostSet() *HostSet {
	return &HostSet{}
}

type Host struct {
	*Resource
	*Describe
}

func NewHost() *Host {
	return &Host{
		&Resource{},
		&Describe{},
	}
}

// Validate  Host 校验
func (h *Host) Validate() error {
	return validate.Struct(h)
}

// InjectDefault 注入default 默认值
func (h *Host) InjectDefault() {
	if h.CreateAt == 0 {
		h.CreateAt = time.Now().UnixMilli()
	}
}

type Resource struct {
	Id          string `json:"id"  validate:"required"`     // 全局唯一Id
	Vendor      Vendor `json:"vendor"`                      // 厂商
	Region      string `json:"region"  validate:"required"` // 地域
	CreateAt    int64  `json:"create_at"`                   // 创建时间
	ExpireAt    int64  `json:"expire_at"`                   // 过期时间
	Type        string `json:"type"  validate:"required"`   // 规格
	Name        string `json:"name"  validate:"required"`   // 名称
	Description string `json:"description"`                 // 描述
	Status      string `json:"status"`                      // 服务商中的状态
	UpdateAt    int64  `json:"update_at"`                   // 更新时间
	SyncAt      int64  `json:"sync_at"`                     // 同步时间
	Account     string `json:"account"`                     // 资源的所属账号
	PublicIP    string `json:"public_ip"`                   // 公网IP
	PrivateIP   string `json:"private_ip"`                  // 内网IP
}

type Describe struct {
	CPU          int    `json:"cpu" validate:"required"`    // 核数
	Memory       int    `json:"memory" validate:"required"` // 内存
	GPUAmount    int    `json:"gpu_amount"`                 // GPU数量
	GPUSpec      string `json:"gpu_spec"`                   // GPU类型
	OSType       string `json:"os_type"`                    // 操作系统类型，分为Windows和Linux
	OSName       string `json:"os_name"`                    // 操作系统名称
	SerialNumber string `json:"serial_number"`              // 序列号
}
