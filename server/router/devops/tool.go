package devops

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/gin-gonic/gin"
)

type ToolRouter struct {
}

// InitToolRouter 初始化 脚本分类 路由信息
func (s *ToolRouter) InitToolRouter(Router *gin.RouterGroup) {
	devopsWithoutRecord := Router.Group("/devops")
	var toolTypeApi = v1.ApiGroupApp.DevopsApiGroup.ToolTypeApi
	{
		devopsWithoutRecord.GET("/tool/toolType/:id", toolTypeApi.FindToolTypeByID)      // 根据ID获取脚本分类
		devopsWithoutRecord.GET("/tool/toolTypes", toolTypeApi.FindToolType)             // 获取脚本分类列表
		devopsWithoutRecord.POST("/tool/toolType", toolTypeApi.CreateToolType)           // 创建分类
		devopsWithoutRecord.PUT("/tool/toolType/:id", toolTypeApi.UpdateToolTypeByID)    // 更新分类
		devopsWithoutRecord.DELETE("/tool/toolType/:id", toolTypeApi.DeleteToolTypeByID) // 删除分类
	}

	var toolContentApi = v1.ApiGroupApp.DevopsApiGroup.ToolContentApi
	{
		devopsWithoutRecord.GET("/tool/toolContent/:id", toolContentApi.FindToolContentByID)      // 根据ID获取脚本
		devopsWithoutRecord.GET("/tool/toolContent", toolContentApi.FindContentByType)            // 根据脚本分类获取脚本
		devopsWithoutRecord.GET("/tool/toolContent/fuzzy", toolContentApi.FindToolContentByName)  // 根据脚本分类&&脚本名 模糊查询
		devopsWithoutRecord.POST("/tool/toolContent", toolContentApi.CreateToolContent)           // 创建脚本
		devopsWithoutRecord.PUT("/tool/toolContent/:id", toolContentApi.UpdateToolContentByID)    // 更新脚本
		devopsWithoutRecord.DELETE("/tool/toolContent/:id", toolContentApi.DeleteToolContentByID) // 删除脚本
	}

	var cmdbHostApi = v1.ApiGroupApp.DevopsApiGroup.CmdbHostApi
	{
		devopsWithoutRecord.GET("/cmdb/host/:id", cmdbHostApi.FindCmdbHostByID)      // 根据ID获取主机
		devopsWithoutRecord.GET("/cmdb/host", cmdbHostApi.FindCmdbHost)              // 获取主机
		devopsWithoutRecord.POST("/cmdb/host", cmdbHostApi.CreateCmdbHost)           // 创建主机
		devopsWithoutRecord.PUT("/cmdb/host/:id", cmdbHostApi.UpdateCmdbHostByID)    // 更新主机
		devopsWithoutRecord.DELETE("/cmdb/host/:id", cmdbHostApi.DeleteCmdbHostByID) // 删除主机
	}

}
