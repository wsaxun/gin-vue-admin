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

type ToolContentApi struct{}

func (t *ToolContentApi) FindToolContentByID(c *gin.Context) {
	typeID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil || utils.VerifyIsZero(typeID) {
		response.FailWithMessage("参数校验失败", c)
		return
	}
	service := devopsService.ToolContentService{}
	info, err := service.GetToolContentByID(c, uint(typeID))
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

func (t *ToolContentApi) FindToolContentByName(c *gin.Context) {
	typeID, err := strconv.ParseUint(c.Query("toolTypeID"), 10, 64)
	if err != nil || utils.VerifyIsZero(typeID) {
		response.FailWithMessage("参数校验失败", c)
		return
	}
	service := devopsService.ToolContentService{}
	info, err := service.GetToolContentByName(c, uint(typeID), c.Query("name"))
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

func (t *ToolContentApi) FindContentByType(c *gin.Context) {
	var toolContent devopsModel.ToolContent
	// ShouldBindQuery 函数根据struct定义的form tag进行绑定
	// ShouldBindJSON 函数根据struct定义的json tag进行绑定
	err := c.ShouldBindQuery(&toolContent)
	if err != nil {
		response.FailWithMessage("查询失败", c)
		return
	}
	service := devopsService.ToolContentService{}
	info, err := service.GetToolContent(c, toolContent)
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

func (t *ToolContentApi) CreateToolContent(c *gin.Context) {
	var toolContent devopsModel.ToolContent
	if err := c.ShouldBindJSON(&toolContent); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if utils.VerifyIsZero(
		strings.TrimSpace(toolContent.Name),
		strings.TrimSpace(toolContent.Code),
		strings.TrimSpace(toolContent.Language),
		toolContent.ToolTypeID,
	) {
		response.FailWithMessage("参数校验失败", c)
		return
	}

	toolContent.CreatedBy = utils.GetUserID(c)
	service := devopsService.ToolContentService{}
	if err := service.CreateToolContent(c, &toolContent); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithDetailed(map[string]interface{}{"id": toolContent.ID}, "操作成功", c)
}

func (t *ToolContentApi) UpdateToolContentByID(c *gin.Context) {
	var toolContent devopsModel.ToolContent
	if err := c.ShouldBindJSON(&toolContent); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	typeID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if utils.VerifyIsZero(
		strings.TrimSpace(toolContent.Name),
		strings.TrimSpace(toolContent.Code),
		strings.TrimSpace(toolContent.Language),
	) || err != nil || utils.VerifyIsZero(typeID) {
		response.FailWithMessage("参数校验失败", c)
		return
	}
	toolContent.ID = uint(typeID)
	toolContent.UpdatedBy = utils.GetUserID(c)
	service := devopsService.ToolContentService{}
	if err := service.UpdateToolContent(c, toolContent); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithDetailed(map[string]interface{}{"id": toolContent.ID}, "操作成功", c)
}

func (t *ToolContentApi) DeleteToolContentByID(c *gin.Context) {
	typeID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil || utils.VerifyIsZero(typeID) {
		response.FailWithMessage("参数校验失败", c)
		return
	}
	service := devopsService.ToolContentService{}
	toolContent := devopsModel.ToolContent{
		ID: uint(typeID),
	}
	if err := service.DeleteToolContent(c, toolContent); err != nil {
		response.FailWithMessage("操作失败", c)
	} else {
		response.OkWithMessage("操作成功", c)
	}
}
