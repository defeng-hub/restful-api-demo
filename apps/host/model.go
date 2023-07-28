package host

import (
	"fmt"
	"time"

	"dario.cat/mergo"
	"github.com/go-playground/validator/v10"
)

type Vendor int

var (
	// validate 必须先new一个实例，然后在实例上使用 v.Struct( obj )，  obj 就是 某个结构体的实例
	validate = validator.New()
)

const (
	// 枚举默认值0
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

type UPDATE_MODE string

const (
	UPDATE_MODE_PUT   UPDATE_MODE = "put"   // 全量更新
	UPDATE_MODE_PATCH UPDATE_MODE = "patch" // 局部更新
)

type UpdateHostRequest struct {
	UpdateMode UPDATE_MODE `json:"update_mode"`
	*Host
}
type DeleteHostRequest struct {
	ID string
}

type Host struct {
	//资源公共属性部分
	*Resource
	//资源独有属性
	*Describe
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
	SerialNumber string `json:"serial_number"`              // 主板序列号
}
type DescribeHostRequest struct {
	Id string
}

func NewDescribeHostRequestWithId(id string) *DescribeHostRequest {
	return &DescribeHostRequest{
		Id: id,
	}
}
func NewQueryHostRequest(pageSize uint64, pageNumber uint64, keywords string) *QueryHostRequest {
	return &QueryHostRequest{PageSize: pageSize, PageNumber: pageNumber, Keywords: keywords}
}
func NewPutUpdateHostRequest(id string) *UpdateHostRequest {
	h := NewHost()
	h.Id = id
	return &UpdateHostRequest{
		UpdateMode: UPDATE_MODE_PUT,
		Host:       h,
	}
}
func NewPatchUpdateHostRequest(id string) *UpdateHostRequest {
	h := NewHost()
	h.Id = id
	return &UpdateHostRequest{
		UpdateMode: UPDATE_MODE_PATCH,
		Host:       h,
	}
}

// 对象全量更新
func (h *Host) Put(obj *Host) error {
	if obj.Id != h.Id {
		return fmt.Errorf("id not equal")
	}

	*h.Resource = *obj.Resource
	*h.Describe = *obj.Describe
	return nil
}

// 对象的局部更新
func (h *Host) Patch(obj *Host) error {
	// if obj.Name != "" {
	// 	h.Name = obj.Name
	// }
	// if obj.CPU != 0 {
	// 	h.CPU = obj.CPU
	// }
	// 比如 obj.A  obj.B  只想修改obj.B该属性
	//return mergo.MergeWithOverwrite(h, obj)
	return mergo.Merge(h, *obj, mergo.WithOverride)
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
