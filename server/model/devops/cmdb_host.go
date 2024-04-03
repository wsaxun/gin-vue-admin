package devops

import (
	"gorm.io/gorm"
	"time"
)

type CMDBHost struct {
	ID          uint           `gorm:"primarykey" json:"id"`                    // 主键ID
	CreatedAt   time.Time      `json:"createAt"`                                // 创建时间
	UpdatedAt   time.Time      `json:"updateAt"`                                // 更新时间
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"deletedAt" form:"deletedAt"` // 删除时间
	Hostname    string         `json:"hostname" form:"hostname" gorm:"column:hostname;not null;comment:主机名"`
	Env         string         `json:"env" form:"env" gorm:"column:env;comment:环境"`
	Status      string         `json:"status" form:"status" gorm:"column:status;comment:主机状态"`
	OsVersion   string         `json:"osVersion" form:"osVersion" gorm:"column:os_version;comment:操作系统"`
	OsRelease   string         `json:"osRelease" form:"osRelease" gorm:"column:os_release;comment:内核发行版本"`
	Cpu         string         `json:"cpu" form:"cpu" gorm:"column:cpu;comment:cpu"`
	Mem         string         `json:"mem" form:"mem" gorm:"column:mem;comment:内存"`
	Disk        string         `json:"disk" form:"disk" gorm:"column:disk;comment:磁盘"`
	Ip          string         `json:"ip" form:"ip" gorm:"column:ip;comment:ip"`
	Tag         string         `json:"tag" form:"tag" gorm:"column:tag;comment:tag"`
	Description string         `json:"description" form:"description" gorm:"column:description;comment:描述"`
	CreatedBy   uint           `json:"createdBy" form:"createdBy" gorm:"column:created_by;comment:创建者"`
	UpdatedBy   uint           `json:"updatedBy" form:"updatedBy" gorm:"column:updated_by;comment:更新者"`
	DeletedBy   uint           `json:"deletedBy" form:"deletedBy" gorm:"column:deleted_by;comment:删除者"`
}

func (CMDBHost) TableName() string {
	return "devops_cmdb_host"
}
