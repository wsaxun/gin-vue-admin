package devops

import (
	"gorm.io/gorm"
	"time"
)

type ToolType struct {
	ID        uint           `gorm:"primarykey" json:"id"`                    // 主键ID
	CreatedAt time.Time      `json:"createAt"`                                // 创建时间
	UpdatedAt time.Time      `json:"updateAt"`                                // 更新时间
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deletedAt" form:"deletedAt"` // 删除时间
	Name      string         `json:"name" form:"name" gorm:"column:name;not null;unique;comment:分类名"`
	CreatedBy uint           `json:"createdBy" form:"createdBy" gorm:"column:created_by;comment:创建者"`
	UpdatedBy uint           `json:"updatedBy" form:"updatedBy" gorm:"column:updated_by;comment:更新者"`
	DeletedBy uint           `json:"deletedBy" form:"deletedBy" gorm:"column:deleted_by;comment:删除者"`
}

func (ToolType) TableName() string {
	return "devops_tool_type"
}
