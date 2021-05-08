package api

import "github.com/gin-gonic/gin"

func RunApi() *gin.Engine {
	r := gin.Default()
	RunApiOnRouter(r)
	return r
}

func RunApiOnRouter(r *gin.Engine) {
	apiGroup := r.Group("/api/simple")
	{
		apiGroup.GET("get-time", GetTime)
		apiGroup.POST("echo", Echo)
	}
}
