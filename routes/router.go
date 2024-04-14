package routes

import (
	"fmt"
	"github.com/gin-contrib/multitemplate"
	"github.com/gin-gonic/gin"
	"go-blog/api/v1"
	"go-blog/middleware"
	"go-blog/utils"
	"os"
)

// 前端代码配置变量
var (
	adminIndex  string
	frontIndex  string
	frontStatic string
	adminRoot   string
	frontIco    string
)

// 入口文件 router
// 读取 setting.go 里面的变量
func createMyRender() multitemplate.Renderer {
	p := multitemplate.NewRenderer()
	// 指定环境变量，前端代码的路径
	webRootPath := os.Getenv("GO_BLOG_WEB_ROOT")
	if webRootPath == "" {
		adminIndex = "web/admin/dist/index.html"
		frontIndex = "web/front/dist/index.html"
	} else {
		adminIndex = webRootPath + "/web/admin/dist/index.html"
		frontIndex = webRootPath + "/web/front/dist/index.html"
	}

	p.AddFromFiles("admin", adminIndex)
	p.AddFromFiles("front", frontIndex)
	return p
}

func InitRouter() {
	gin.SetMode(utils.AppMode)
	r := gin.New()
	// 设置信任网络 []string
	// nil 为不计算，避免性能消耗，上线应当设置
	_ = r.SetTrustedProxies(nil)

	r.HTMLRender = createMyRender()
	r.Use(middleware.Logger())
	r.Use(gin.Recovery())
	r.Use(middleware.Cors())

	webRootPath := os.Getenv("GO_BLOG_WEB_ROOT")
	if webRootPath == "" {
		frontStatic = "./web/front/dist/static"
		adminRoot = "./web/admin/dist"
		frontIco = "./web/front/dist/favicon.ico"
	} else {
		frontStatic = webRootPath + "/web/front/dist/static"
		adminRoot = webRootPath + "/web/admin/dist"
		frontIco = webRootPath + "/web/front/dist/favicon.ico"
	}

	r.Static("/static", frontStatic)
	r.Static("/admin", adminRoot)
	r.StaticFile("/favicon.ico", frontIco)

	r.GET("/", func(c *gin.Context) {
		c.HTML(200, "front", nil)
	})

	r.GET("/admin", func(c *gin.Context) {
		c.HTML(200, "admin", nil)
	})

	/*
		后台管理路由接口
	*/
	auth := r.Group("api/v1")
	auth.Use(middleware.JwtToken())
	{
		// 用户模块的路由接口
		auth.GET("admin/users", v1.GetUsers)
		auth.PUT("user/:id", v1.EditUser)
		auth.DELETE("user/:id", v1.DeleteUser)

		//修改密码
		auth.PUT("admin/changepw/:id", v1.ChangeUserPassword)

		// 分类模块的路由接口
		auth.GET("admin/category", v1.GetCategory)
		auth.POST("category/add", v1.AddCategory)
		auth.PUT("category/:id", v1.EditCategory)
		auth.DELETE("category/:id", v1.DeleteCategory)

		// 文章模块的路由接口
		auth.GET("admin/article/info/:id", v1.GetArticleInfo)
		auth.GET("admin/article", v1.GetArticle)
		auth.POST("article/add", v1.AddArticle)
		auth.PUT("article/:id", v1.EditArticle)
		auth.DELETE("article/:id", v1.DeleteArticle)

		// 上传文件
		auth.POST("upload", v1.Upload)

		// 更新个人设置
		auth.GET("admin/profile/:id", v1.GetProfile)
		auth.PUT("profile/:id", v1.UpdateProfile)

		// 评论模块
		auth.GET("comment/list", v1.GetCommentList)
		auth.DELETE("delcomment/:id", v1.DeleteComment)
		auth.PUT("checkcomment/:id", v1.CheckComment)
		auth.PUT("uncheckcomment/:id", v1.UncheckComment)
	}

	/*
		前端展示页面接口
	*/
	router := r.Group("api/v1")
	{
		// 用户信息模块
		router.POST("user/add", v1.AddUser)
		router.GET("user/:id", v1.GetUserInfo)
		router.GET("users", v1.GetUsers)

		// 文章分类信息模块
		router.GET("category", v1.GetCategory)
		router.GET("category/:id", v1.GetCateInfo)

		// 文章模块
		router.GET("article", v1.GetArticle)
		router.GET("article/list/:id", v1.GetCategoryArticle)
		router.GET("article/info/:id", v1.GetArticleInfo)

		// 登录控制模块
		router.POST("login", v1.Login)
		router.POST("loginfront", v1.LoginFront)

		// 获取个人设置信息
		router.GET("profile/:id", v1.GetProfile)

		// 评论模块
		router.POST("addcomment", v1.AddComment)
		router.GET("comment/info/:id", v1.GetComment)
		router.GET("commentfront/:id", v1.GetCommentListFront)
		router.GET("commentcount/:id", v1.GetCommentCount)
	}
	err := r.Run(utils.HttpPort)
	if err != nil {
		fmt.Println("端口启动失败")
		return
	}
}
