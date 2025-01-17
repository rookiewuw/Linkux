package routes

import (
	"Linkux/controllers"
	"Linkux/logger"
	"Linkux/middleware"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func Setup(mode string) *gin.Engine {
	if mode == gin.ReleaseMode {
		gin.SetMode(gin.ReleaseMode)
	}

	//初始化gin框架内置的校验器使用的翻译器
	controllers.InitTrans("zh")

	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true), middleware.RateLimitMiddleware(2*time.Second, 1)) // 令牌桶容量为1，每两秒钟填充1个

	r.LoadHTMLGlob("./template/*")
	r.Static("/static", "./static")

	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"msg": "404 NOT FOUND",
		})
	})
	return r
}
