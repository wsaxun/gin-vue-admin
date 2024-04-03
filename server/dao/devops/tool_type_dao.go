package devops

import (
	"context"
	"errors"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	devopsModel "github.com/flipped-aurora/gin-vue-admin/server/model/devops"
	"gorm.io/gorm"
)

type ToolTypeDao struct{}

func NewToolTypeDao() ToolTypeDao {
	return ToolTypeDao{}
}

func (t *ToolTypeDao) GetToolTypeByID(ctx context.Context, id uint) (toolType devopsModel.ToolType, err error) {
	err = global.GVA_DB.WithContext(ctx).Where("id = ?", id).First(&toolType).Error
	return
}

func (t *ToolTypeDao) GetToolType(ctx context.Context) (toolType []devopsModel.ToolType, err error) {
	err = global.GVA_DB.WithContext(ctx).Find(&toolType).Error
	return
}

func (t *ToolTypeDao) Create(ctx context.Context, toolType *devopsModel.ToolType) (err error) {
	// 创建之前查询是否出现相同分类名
	err = global.GVA_DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		var count int64
		if tx.Model(&devopsModel.ToolType{}).Where("name = ?", toolType.Name).Count(&count); count != 0 {
			return errors.New("存在重复的脚本分类")
		}
		return tx.Create(toolType).Error
	})
	return
}

func (t *ToolTypeDao) Update(ctx context.Context, toolType devopsModel.ToolType) (err error) {
	// Save 方法会保存所有字段，包括零值
	// gin c.ShouldBindJSON()方法对于json没有的字段，会初始化为零值
	// err = global.GVA_DB.Save(toolType).Error
	var tool devopsModel.ToolType
	err = global.GVA_DB.WithContext(ctx).Model(&tool).Where("id = ?", toolType.ID).Updates(devopsModel.ToolType{
		Name:      toolType.Name,
		UpdatedBy: toolType.UpdatedBy,
	}).Error
	return
}

func (t *ToolTypeDao) Delete(ctx context.Context, toolType devopsModel.ToolType) (err error) {
	////手工事务
	//tx := global.GVA_DB.WithContext(ctx).Begin()
	//defer func() {
	//	if r := recover(); r != nil {
	//		tx.Rollback()
	//	}
	//}()
	//
	//if err = tx.Error; err != nil {
	//	tx.Rollback()
	//	return
	//}
	//
	//if err = tx.Unscoped().Delete(toolType).Error; err != nil {
	//	tx.Rollback()
	//	return
	//}
	//// 删除脚本分类时候,把脚本关联的分类id置为初值0
	//var toolContent devopsModel.ToolContent
	//if err = tx.Model(&toolContent).Where("tool_type_id = ?", toolType.ID).Updates(devopsModel.ToolContent{ToolTypeID: 0}).Error; err != nil {
	//	tx.Rollback()
	//	return
	//}
	//return tx.Commit().Error

	// 自动事务
	err = global.GVA_DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		var count int64
		if tx.Model(&devopsModel.ToolContent{}).Where("tool_type_id = ?", toolType.ID).Count(&count); count != 0 {
			return errors.New("关联的脚本不为空")
		}

		if err = tx.Unscoped().Delete(toolType).Error; err != nil {
			return err
		}
		return nil
	})
	return
}
