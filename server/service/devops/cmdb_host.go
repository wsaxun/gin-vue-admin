package devops

import (
	"context"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"

	devopsDao "github.com/flipped-aurora/gin-vue-admin/server/dao/devops"
	devopsModel "github.com/flipped-aurora/gin-vue-admin/server/model/devops"
)

type CmdbHostService struct{}

func (c *CmdbHostService) GetCmdbHostByID(ctx context.Context, id uint) (cmdbHost devopsModel.CMDBHost, err error) {
	dao := devopsDao.NewCmdbHostDao()
	cmdbHost, err = dao.GetCmdbHostByID(ctx, id)
	return
}

func (c *CmdbHostService) GetCmdbHosts(ctx context.Context, cmdbHost *devopsModel.CMDBHost, info *request.PageInfo, order string, desc bool) (cmdbHosts []devopsModel.CMDBHost, total int64, err error) {
	dao := devopsDao.NewCmdbHostDao()
	cmdbHosts, total, err = dao.GetCmdbHosts(ctx, cmdbHost, info, order, desc)
	return
}

func (c *CmdbHostService) CreateCmdbHost(ctx context.Context, cmdbHost *devopsModel.CMDBHost) (err error) {
	dao := devopsDao.NewCmdbHostDao()
	err = dao.Create(ctx, cmdbHost)
	return
}

func (c *CmdbHostService) UpdateCmdbHost(ctx context.Context, cmdbHost *devopsModel.CMDBHost) (err error) {
	dao := devopsDao.NewCmdbHostDao()
	err = dao.Update(ctx, cmdbHost)
	return
}

func (c *CmdbHostService) DeleteCmdbHost(ctx context.Context, cmdbHost *devopsModel.CMDBHost) (err error) {
	dao := devopsDao.NewCmdbHostDao()
	err = dao.Delete(ctx, cmdbHost)
	return
}
