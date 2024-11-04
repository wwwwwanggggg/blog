package router

import (
	"blog/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitRouter(r *gin.Engine) {
	r.Use(middleware.Error)
	r.Use(middleware.GinLogger(), middleware.GinRecovery(true))
	r.LoadHTMLGlob("./static/*")
	apiRouter := r.Group("/api")
	{
		apiRouter.POST("/", ctr.Admin.Login) // 管理员登录
		apiRouter.Use(middleware.CheckRole(2, "admin"))
		apiRouter.DELETE("/", ctr.Admin.Logout) // 管理员登出
		fileRouter := apiRouter.Group("/file")
		{
			fileRouter.POST("/file", ctr.FileManagement.UploadFile) // 上传文件
			fileRouter.GET("/file", ctr.GetFile)                    // 获取文件
			fileRouter.POST("/dir", ctr.FileManagement.NewDir)      // 创建新目录
			fileRouter.DELETE("/", ctr.FileManagement.DeleteDir)    // 删除目录或者文件
		}
	}
	r.GET("/", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "index.html", nil)
	})

}
