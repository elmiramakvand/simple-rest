package api

import "github.com/gin-gonic/gin"

func RunApi() *gin.Engine {
	r := gin.Default()
	RunApiOnRouter(r)
	return r
}

func RunApiOnRouter(r *gin.Engine) {
	cacheQueryGroup := r.Group("/api/simple")
	{
		cacheQueryGroup.GET("get-time", GetTime)
		cacheQueryGroup.POST("echo", Echo)
	}
}
