package devops

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	devopsModel "github.com/flipped-aurora/gin-vue-admin/server/model/devops"
	systemReq "github.com/flipped-aurora/gin-vue-admin/server/model/devops/request"
	devopsService "github.com/flipped-aurora/gin-vue-admin/server/service/devops"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strconv"
	"strings"
)

type CmdbHostApi struct{}

func (c *CmdbHostApi) FindCmdbHostByID(ctx *gin.Context) {
	id := ctx.Param("id")
	hostID, err := strconv.ParseUint(id, 10, 64)

	if err != nil || utils.VerifyIsZero(hostID) {
		response.FailWithMessage("参数校验失败", ctx)
		return
	}
	service := devopsService.CmdbHostService{}
	info, err := service.GetCmdbHostByID(ctx, uint(hostID))
	if err != nil {
		response.FailWithMessage("查询失败", ctx)
		return
	}
	fields := ctx.Query("fields")
	if utils.VerifyIsZero(fields) {
		response.OkWithData(info, ctx)
		return
	}
	response.OkWithData(utils.FilterFields(info, strings.Split(fields, ",")...), ctx)

}

func (c *CmdbHostApi) FindCmdbHost(ctx *gin.Context) {
	service := devopsService.CmdbHostService{}
	var pageInfo systemReq.SearchHostParams
	err := ctx.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), ctx)
		return
	}
	err = utils.Verify(pageInfo.PageInfo, utils.PageInfoVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), ctx)
		return
	}
	info, total, err := service.GetCmdbHosts(ctx, &pageInfo.CMDBHost, &pageInfo.PageInfo, pageInfo.OrderKey, pageInfo.Desc)
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", ctx)
		return
	}
	fields := ctx.Query("fields")
	if utils.VerifyIsZero(fields) {
		response.OkWithDetailed(response.PageResult{
			List:     info,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "查询成功", ctx)
		return
	}
	response.OkWithDetailed(response.PageResult{
		List:     utils.FilterFields(info, strings.Split(fields, ",")...),
		Total:    total,
		Page:     pageInfo.Page,
		PageSize: pageInfo.PageSize,
	}, "查询成功", ctx)

}

func (c *CmdbHostApi) CreateCmdbHost(ctx *gin.Context) {
	var cmdbHost devopsModel.CMDBHost
	if err := ctx.ShouldBindJSON(&cmdbHost); err != nil {
		response.FailWithMessage(err.Error(), ctx)
		return
	}

	if utils.VerifyIsZero(strings.TrimSpace(cmdbHost.Hostname)) {
		response.FailWithMessage("主机名不能为空", ctx)
		return
	}
	cmdbHost.CreatedBy = utils.GetUserID(ctx)
	service := devopsService.CmdbHostService{}
	if err := service.CreateCmdbHost(ctx, &cmdbHost); err != nil {
		response.FailWithMessage(err.Error(), ctx)
		return
	}
	data := map[string]interface{}{"id": cmdbHost.ID}
	response.OkWithDetailed(data, "操作成功", ctx)
}

func (c *CmdbHostApi) UpdateCmdbHostByID(ctx *gin.Context) {
	var cmdbHost devopsModel.CMDBHost
	if err := ctx.ShouldBindJSON(&cmdbHost); err != nil {
		response.FailWithMessage(err.Error(), ctx)
		return
	}

	id := ctx.Param("id")
	hostID, err := strconv.ParseUint(id, 10, 64)
	if err != nil || utils.VerifyIsZero(hostID) || utils.VerifyIsZero(strings.TrimSpace(cmdbHost.Hostname)) {
		response.FailWithMessage("参数校验失败", ctx)
		return
	}

	cmdbHost.ID = uint(hostID)
	cmdbHost.UpdatedBy = utils.GetUserID(ctx)
	service := devopsService.CmdbHostService{}
	if err := service.UpdateCmdbHost(ctx, &cmdbHost); err != nil {
		response.FailWithMessage(err.Error(), ctx)
		return
	}
	data := map[string]interface{}{"id": cmdbHost.ID, "hostname": cmdbHost.Hostname}
	response.OkWithDetailed(data, "操作成功", ctx)
}

func (c *CmdbHostApi) DeleteCmdbHostByID(ctx *gin.Context) {
	id := ctx.Param("id")
	hostID, err := strconv.ParseUint(id, 10, 64)
	if err != nil || utils.VerifyIsZero(hostID) {
		response.FailWithMessage("id必须为数值", ctx)
		return
	}
	service := devopsService.CmdbHostService{}
	cmdbHost := devopsModel.CMDBHost{
		ID: uint(hostID),
	}
	if err := service.DeleteCmdbHost(ctx, &cmdbHost); err != nil {
		response.FailWithMessage(err.Error(), ctx)
	} else {
		response.OkWithMessage("操作成功", ctx)
	}
}
