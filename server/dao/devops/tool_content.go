package devops

import (
	"context"
	"errors"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	devopsModel "github.com/flipped-aurora/gin-vue-admin/server/model/devops"
	"gorm.io/gorm"
)

type ToolContentDao struct{}

func NewToolContentDao() ToolContentDao {
	return ToolContentDao{}
}

func (t *ToolContentDao) GetToolContentByID(ctx context.Context, id uint) (toolContent devopsModel.ToolContent, err error) {
	err = global.GVA_DB.WithContext(ctx).Where("id = ?", id).First(&toolContent).Error
	return
}

func (t *ToolContentDao) GetToolContentByName(ctx context.Context, id uint, name string) (toolContents []devopsModel.ToolContent, err error) {
	err = global.GVA_DB.WithContext(ctx).Where("tool_type_id = ? and name LIKE ?", id, "%"+name+"%").Find(&toolContents).Error
	return
}

func (t *ToolContentDao) GetToolContentByType(ctx context.Context, toolContent devopsModel.ToolContent) (toolContents []devopsModel.ToolContent, err error) {
	err = global.GVA_DB.WithContext(ctx).Where(&toolContent).Find(&toolContents).Error
	return
}

func (t *ToolContentDao) Create(ctx context.Context, toolContent *devopsModel.ToolContent) (err error) {
	err = global.GVA_DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		var count int64
		if tx.Model(&devopsModel.ToolContent{}).Where("tool_type_id = ? and name = ?", toolContent.ToolTypeID, toolContent.Name).Count(&count); count != 0 {
			return errors.New("存在重复的脚本名")
		}
		return tx.Create(toolContent).Error
	})
	return
}

func (t *ToolContentDao) Update(ctx context.Context, toolContent devopsModel.ToolContent) (err error) {
	var tool devopsModel.ToolContent
	err = global.GVA_DB.WithContext(ctx).Model(&tool).Where("id = ?", toolContent.ID).Updates(devopsModel.ToolContent{
		Name:        toolContent.Name,
		Language:    toolContent.Language,
		Timeout:     toolContent.Timeout,
		Description: toolContent.Description,
		Code:        toolContent.Code,
		UpdatedBy:   toolContent.UpdatedBy,
		ToolTypeID:  toolContent.ToolTypeID,
	}).Error
	return
}

func (t *ToolContentDao) Delete(ctx context.Context, toolContent devopsModel.ToolContent) (err error) {
	err = global.GVA_DB.WithContext(ctx).Unscoped().Delete(toolContent).Error
	return err
}
