package router

import (
	"github.com/gogf/gf/v2/net/ghttp"
	"tree-table/internal/controller"
)

// 注册路由
func RegisterRouter(s *ghttp.Server) {
	// 区域管理路由
	s.Group("/api/areas", func(group *ghttp.RouterGroup) {
		//group.Middleware(middleware.Auth) // 假设有认证中间件
		group.Bind(
			controller.AreaController.Create,
			controller.AreaController.Update,
			controller.AreaController.Delete,
			controller.AreaController.GetList,
			controller.AreaController.GetTree,
			controller.AreaController.GetChildren,
			controller.AreaController.GetNodePath,
			controller.AreaController.AdjustSort,
		)
	})
}
