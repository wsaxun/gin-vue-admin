package devops

import (
	"context"

	devopsDao "github.com/flipped-aurora/gin-vue-admin/server/dao/devops"
	devopsModel "github.com/flipped-aurora/gin-vue-admin/server/model/devops"
)

type ToolContentService struct{}

func (t *ToolContentService) GetToolContentByID(ctx context.Context, id uint) (toolContent devopsModel.ToolContent, err error) {
	dao := devopsDao.NewToolContentDao()
	toolContent, err = dao.GetToolContentByID(ctx, id)
	return
}

func (t *ToolContentService) GetToolContentByName(ctx context.Context, id uint, name string) (toolContents []devopsModel.ToolContent, err error) {
	dao := devopsDao.NewToolContentDao()
	toolContents, err = dao.GetToolContentByName(ctx, id, name)
	return
}

func (t *ToolContentService) GetToolContent(ctx context.Context, toolContent devopsModel.ToolContent) (toolContents []devopsModel.ToolContent, err error) {
	dao := devopsDao.NewToolContentDao()
	toolContents, err = dao.GetToolContentByType(ctx, toolContent)
	return
}

func (t *ToolContentService) CreateToolContent(ctx context.Context, toolContent *devopsModel.ToolContent) (err error) {
	dao := devopsDao.NewToolContentDao()
	err = dao.Create(ctx, toolContent)
	return
}

func (t *ToolContentService) UpdateToolContent(ctx context.Context, toolContent devopsModel.ToolContent) (err error) {
	dao := devopsDao.NewToolContentDao()
	err = dao.Update(ctx, toolContent)
	return
}

func (t *ToolContentService) DeleteToolContent(ctx context.Context, toolContent devopsModel.ToolContent) (err error) {
	dao := devopsDao.NewToolContentDao()
	err = dao.Delete(ctx, toolContent)
	return
}
