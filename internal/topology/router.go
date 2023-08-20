package topology

import (
	"ClusterWatcher/internal/pkg/core"
	"ClusterWatcher/internal/pkg/errno"
	"ClusterWatcher/internal/pkg/log"
	"ClusterWatcher/internal/topology/controller/v1/record"
	"ClusterWatcher/internal/topology/store"
	"github.com/gin-gonic/gin"
)

func installRouters(g *gin.Engine) error {
	// 注册 404 Handler.
	g.NoRoute(func(c *gin.Context) {
		core.WriteResponse(c, errno.ErrPageNotFound, nil)
	})

	//注册 /healthz handler.
	g.GET("/healthz", func(c *gin.Context) {
		log.Infow("Healthz function called")

		core.WriteResponse(c, nil, map[string]string{"status": "ok"})
	})

	rc := record.New(store.S)
	//uc := user.New(store.S)
	/*
		/topology/record/pod
		/topology/record/container
		/topology/record/host

		/topology-by-group

		/topology-by-layer

	*/

	// 创建 v1 路由分组
	v1 := g.Group("/v1")
	{
		recordv1 := v1.Group("/topology")
		{
			recordv1.GET("record/:name", rc.List)
			recordv1.GET("render-by-group", rc.Group)
			recordv1.GET("render-by-layer", rc.Layer)
		}

		//userv1 := v1.Group("/users")
		//{
		//	userv1.POST("", uc.Create)                             // 创建用户
		//	userv1.PUT(":name/change-password", uc.ChangePassword) // 修改用户密码
		//	userv1.Use(mw.Authn(), mw.Authz(authz))
		//	userv1.GET(":name", uc.Get)       // 获取用户详情
		//	userv1.PUT(":name", uc.Update)    // 更新用户
		//	userv1.GET("", uc.List)           // 列出用户列表，只有 root 用户才能访问
		//	userv1.DELETE(":name", uc.Delete) // 删除用户
		//}
	}
	return nil
}
