package devops

import (
	"context"

	devopsDao "github.com/flipped-aurora/gin-vue-admin/server/dao/devops"
	devopsModel "github.com/flipped-aurora/gin-vue-admin/server/model/devops"
)

type ToolTypeService struct{}

func (t *ToolTypeService) GetToolTypeByID(ctx context.Context, id uint) (toolType devopsModel.ToolType, err error) {
	dao := devopsDao.NewToolTypeDao()
	toolType, err = dao.GetToolTypeByID(ctx, id)
	return
}

func (t *ToolTypeService) GetToolType(ctx context.Context) (toolType []devopsModel.ToolType, err error) {
	dao := devopsDao.NewToolTypeDao()
	toolType, err = dao.GetToolType(ctx)
	return
}

func (t *ToolTypeService) CreateToolType(ctx context.Context, toolType *devopsModel.ToolType) (err error) {
	dao := devopsDao.NewToolTypeDao()
	err = dao.Create(ctx, toolType)
	return
}

func (t *ToolTypeService) UpdateToolType(ctx context.Context, toolType devopsModel.ToolType) (err error) {
	dao := devopsDao.NewToolTypeDao()
	err = dao.Update(ctx, toolType)
	return
}

func (t *ToolTypeService) DeleteToolType(ctx context.Context, toolType devopsModel.ToolType) (err error) {
	dao := devopsDao.NewToolTypeDao()
	err = dao.Delete(ctx, toolType)
	return
}
