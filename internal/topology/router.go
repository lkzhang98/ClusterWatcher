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
			recordv1.GET("render-by-name", rc.GetName)
			recordv1.GET("render-by-ns", rc.GetNs)
			recordv1.GET("render-by-layer", rc.Layer)
		}

	}
	return nil
}
