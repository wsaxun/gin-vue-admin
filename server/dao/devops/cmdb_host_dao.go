package devops

import (
	"context"
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	devopsModel "github.com/flipped-aurora/gin-vue-admin/server/model/devops"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
)

type CmdbHostDao struct{}

func NewCmdbHostDao() CmdbHostDao {
	return CmdbHostDao{}
}

func (c *CmdbHostDao) GetCmdbHostByID(ctx context.Context, id uint) (host devopsModel.CMDBHost, err error) {
	err = global.GVA_DB.WithContext(ctx).Where("id = ?", id).First(&host).Error
	return
}

func (c *CmdbHostDao) GetCmdbHosts(ctx context.Context, cmdbHost *devopsModel.CMDBHost, info *request.PageInfo, order string, desc bool) (hosts []devopsModel.CMDBHost, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	query := global.GVA_DB.WithContext(ctx).Model(cmdbHost)
	err = query.Where(cmdbHost).Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	if order != "" {
		if !utils.VerifyJsonFiled(cmdbHost, order) {
			return nil, 0, fmt.Errorf("非法的排序字段: %v", order)
		}
		if desc {
			order += " desc"
		}
		query = query.Order(order)
	} else {
		query = query.Order("id")
	}

	err = query.Limit(limit).Offset(offset).Find(&hosts).Error
	return hosts, total, err
}

func (c *CmdbHostDao) Create(ctx context.Context, host *devopsModel.CMDBHost) (err error) {
	err = global.GVA_DB.WithContext(ctx).Create(host).Error
	return
}

func (c *CmdbHostDao) Update(ctx context.Context, host *devopsModel.CMDBHost) (err error) {
	// Save 方法会保存所有字段，包括零值
	// gin c.ShouldBindJSON()方法对于json没有的字段，会初始化为零值
	// err = global.GVA_DB.Save(toolType).Error
	var hosts devopsModel.CMDBHost
	err = global.GVA_DB.WithContext(ctx).Model(&hosts).Where("id = ?", host.ID).Updates(host).Error
	return
}

func (c *CmdbHostDao) Delete(ctx context.Context, host *devopsModel.CMDBHost) (err error) {
	err = global.GVA_DB.WithContext(ctx).Unscoped().Delete(host).Error
	return err
}
