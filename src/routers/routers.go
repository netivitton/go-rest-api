package routers

import (
	"github.com/gin-gonic/gin"
)

// InitRouter initialize routing information
func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	// r.StaticFS("/export", http.Dir(export.GetExcelFullPath()))
	// r.StaticFS("/upload/images", http.Dir(upload.GetImageFullPath()))
	// r.StaticFS("/qrcode", http.Dir(qrcode.GetQrCodeFullPath()))

	// r.POST("/auth", api.GetAuth)
	// r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	// r.POST("/upload", api.UploadImage)

	apiv1 := r.Group("/user")
	// apiv1.Use(jwt.JWT())
	{
		//获取标签列表
		apiv1.GET("/tags", user.GetTags)
		//新建标签
	}

	return r
}
