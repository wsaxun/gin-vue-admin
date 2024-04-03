package devops

import (
	"gorm.io/gorm"
	"time"
)

type ToolContent struct {
	ID          uint           `gorm:"primarykey" json:"id"`   // 主键ID
	CreatedAt   time.Time      `json:"createAt"`               // 创建时间
	UpdatedAt   time.Time      `json:"updateAt"`               // 更新时间
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"deletedAt"` // 删除时间
	Language    string         `json:"language" form:"language" gorm:"column:language;not null;comment:语言"`
	Name        string         `json:"name" form:"name" gorm:"column:name;not null;comment:脚本名"`
	Timeout     uint           `json:"timeout" form:"timeout" gorm:"column:timeout;comment:超时"`
	Description string         `json:"description" form:"description" gorm:"column:description;comment:描述"`
	Code        string         `json:"code" form:"code" gorm:"column:code;type:text;not null;comment:脚本内容"`
	ToolTypeID  uint           `json:"toolTypeID" form:"toolTypeID" gorm:"column:tool_type_id;comment:脚本分类"`
	CreatedBy   uint           `json:"createdBy" form:"createdBy" gorm:"column:created_by;comment:创建者"`
	UpdatedBy   uint           `json:"updatedBy" form:"updatedBy" gorm:"column:updated_by;comment:更新者"`
	DeletedBy   uint           `json:"deletedBy" form:"deletedBy" gorm:"column:deleted_by;comment:删除者"`
}

func (ToolContent) TableName() string {
	return "devops_tool_content"
}
