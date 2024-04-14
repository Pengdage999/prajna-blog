package middleware

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"time"
)

// 跨域 CORS是一个W3C标准，全称是"跨域资源共享"（Cross-origin resource sharing）。

func Cors() gin.HandlerFunc {
	//return func(c *gin.Context) {
	//	cors.New(cors.Config{
	//		AllowOrigins:  []string{"*"},
	//		AllowMethods:  []string{"*"},
	//		AllowHeaders:  []string{"Origin"},
	//		ExposeHeaders: []string{"Content-Length", "Authorization"},
	//		//AllowCredentials: true,
	//		//AllowOriginFunc: func(origin string) bool {
	//		//	return origin == "https://github.com"
	//		//},
	//		MaxAge: 12 * time.Hour,
	//	})
	//}

	return cors.New(
		cors.Config{
			AllowAllOrigins:  true,
			AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
			AllowHeaders:     []string{"*"},
			ExposeHeaders:    []string{"Content-Length", "Authorization", "Content-Type"},
			AllowCredentials: true,
			MaxAge:           12 * time.Hour,
		},
	)

	//return cors.Default()
}
