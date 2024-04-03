package devops

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	devopsModel "github.com/flipped-aurora/gin-vue-admin/server/model/devops"
	devopsService "github.com/flipped-aurora/gin-vue-admin/server/service/devops"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"strconv"
	"strings"
)

type ToolTypeApi struct{}

func (t *ToolTypeApi) FindToolTypeByID(c *gin.Context) {
	id := c.Param("id")
	typeID, err := strconv.ParseUint(id, 10, 64)

	if err != nil || utils.VerifyIsZero(typeID) {
		response.FailWithMessage("参数校验失败", c)
		return
	}
	service := devopsService.ToolTypeService{}
	info, err := service.GetToolTypeByID(c, uint(typeID))
	if err != nil {
		response.FailWithMessage("查询失败", c)
		return
	}
	fields := c.Query("fields")
	if utils.VerifyIsZero(fields) {
		response.OkWithData(info, c)
		return
	}
	response.OkWithData(utils.FilterFields(info, strings.Split(fields, ",")...), c)

}

func (t *ToolTypeApi) FindToolType(c *gin.Context) {
	service := devopsService.ToolTypeService{}
	info, err := service.GetToolType(c)
	if err != nil {
		response.FailWithMessage("查询失败", c)
		return
	}
	fields := c.Query("fields")
	if utils.VerifyIsZero(fields) {
		response.OkWithData(info, c)
		return
	}
	response.OkWithData(utils.FilterFields(info, strings.Split(fields, ",")...), c)
}

func (t *ToolTypeApi) CreateToolType(c *gin.Context) {
	var toolType devopsModel.ToolType
	if err := c.ShouldBindJSON(&toolType); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if utils.VerifyIsZero(strings.TrimSpace(toolType.Name)) {
		response.FailWithMessage("分类名不能为空", c)
		return
	}
	toolType.CreatedBy = utils.GetUserID(c)
	service := devopsService.ToolTypeService{}
	if err := service.CreateToolType(c, &toolType); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	data := map[string]interface{}{"id": toolType.ID}
	response.OkWithDetailed(data, "操作成功", c)
}

func (t *ToolTypeApi) UpdateToolTypeByID(c *gin.Context) {
	var toolType devopsModel.ToolType
	if err := c.ShouldBindJSON(&toolType); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	id := c.Param("id")
	typeID, err := strconv.ParseUint(id, 10, 64)
	if err != nil || utils.VerifyIsZero(typeID) || utils.VerifyIsZero(strings.TrimSpace(toolType.Name)) {
		response.FailWithMessage("参数校验失败", c)
		return
	}

	toolType.ID = uint(typeID)
	toolType.UpdatedBy = utils.GetUserID(c)
	service := devopsService.ToolTypeService{}
	if err := service.UpdateToolType(c, toolType); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	data := map[string]interface{}{"id": toolType.ID, "name": toolType.Name}
	response.OkWithDetailed(data, "操作成功", c)
}

func (t *ToolTypeApi) DeleteToolTypeByID(c *gin.Context) {
	id := c.Param("id")
	typeID, err := strconv.ParseUint(id, 10, 64)
	if err != nil || utils.VerifyIsZero(typeID) {
		response.FailWithMessage("id必须为数值", c)
		return
	}
	service := devopsService.ToolTypeService{}
	toolType := devopsModel.ToolType{
		ID: uint(typeID),
	}
	if err := service.DeleteToolType(c, toolType); err != nil {
		response.FailWithMessage(err.Error(), c)
	} else {
		response.OkWithMessage("操作成功", c)
	}
}
