package internal

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-oauth2/oauth2/v4/errors"
	"github.com/go-oauth2/oauth2/v4/manage"
	"github.com/go-oauth2/oauth2/v4/server"
	account "github.com/netivitton/go-rest-api/internal/account"
	middleware "github.com/netivitton/go-rest-api/internal/middleware"
	"github.com/netivitton/go-rest-api/utils"

	// "github.com/netivitton/go-rest-api/user"

	"context"

	ginglog "github.com/szuecs/gin-glog"
	ginoauth2 "github.com/zalando/gin-oauth2"

	"github.com/jackc/pgx/v4"
	pg "github.com/vgarvardt/go-oauth2-pg/v4"
	"github.com/vgarvardt/go-pg-adapter/pgx4adapter"
)

// InitRouter initialize routing information
func InitRouter() *gin.Engine {
	config, err := utils.LoadConfig()
	if err != nil {
		log.Fatal("cannot connect DB", err)
	}
	account.InitAccountModel()
	pgxConn, _ := pgx.Connect(context.TODO(), config.DB_HOST_ALL)

	manager := manage.NewDefaultManager()

	// use PostgreSQL token store with pgx.Connection adapter
	adapter := pgx4adapter.NewConn(pgxConn)
	tokenStore, _ := pg.NewTokenStore(adapter, pg.WithTokenStoreGCInterval(time.Minute))
	defer tokenStore.Close()

	clientStore, _ := pg.NewClientStore(adapter)

	manager.MapTokenStorage(tokenStore)
	manager.MapClientStorage(clientStore)

	srv := server.NewDefaultServer(manager)
	srv.SetAllowGetAccessRequest(true)
	srv.SetClientInfoHandler(server.ClientFormHandler)

	srv.SetInternalErrorHandler(func(err error) (re *errors.Response) {
		log.Println("Internal Error:", err.Error())
		return
	})

	srv.SetResponseErrorHandler(func(re *errors.Response) {
		log.Println("Response Error:", re.Error.Error())
	})

	router := gin.New()
	router.Use(ginglog.Logger(3 * time.Second))
	router.Use(ginoauth2.RequestLogger([]string{"uid"}, "data"))
	router.Use(gin.Recovery())

	public := router.Group("/api")
	public.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Hello to public world"})
	})
	public.GET("/login", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Hello to public world"})
	})
	auth := public.Group("/oauth2")
	{
		auth.GET("/token", func(c *gin.Context) {
			srv.HandleTokenRequest(c.Writer, c.Request)
		})
	}

	private := router.Group("/api/private")
	privateGroup := router.Group("/api/privateGroup")
	privateUser := router.Group("/api/privateUser")
	privateService := router.Group("/api/privateService")
	//glog.Infof("Register allowed users: %+v and groups: %+v and services: %+v", USERS, TEAMS, SERVICES)

	private.Use(middleware.HandleTokenVerify(srv))
	// privateGroup.Use(ginoauth2.Auth(zalando.GroupCheck(TEAMS), zalando.OAuth2Endpoint))
	// privateUser.Use(ginoauth2.Auth(zalando.UidCheck(USERS), zalando.OAuth2Endpoint))
	// //privateService.Use(ginoauth2.Auth(zalando.UidCheck(SERVICES), zalando.OAuth2Endpoint))
	// privateService.Use(ginoauth2.Auth(zalando.ScopeAndCheck("uidcheck", "uid", "bar"), zalando.OAuth2Endpoint))

	private.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Hello from private for groups and users"})
	})
	privateGroup.GET("/", func(c *gin.Context) {
		// uid, okUid := c.Get("uid")
		// if team, ok := c.Get("team"); ok && okUid {
		// 	c.JSON(200, gin.H{"message": fmt.Sprintf("Hello from private for groups to %s member of %s", uid, team)})
		// } else {
		// 	c.JSON(200, gin.H{"message": "Hello from private for groups without uid and team"})
		// }
	})
	privateUser.GET("/", func(c *gin.Context) {
		// if v, ok := c.Get("cn"); ok {
		// 	c.JSON(200, gin.H{"message": fmt.Sprintf("Hello from private for users to %s", v)})
		// } else {
		// 	c.JSON(200, gin.H{"message": "Hello from private for users without cn"})
		// }
	})
	privateService.GET("/", func(c *gin.Context) {
		// if v, ok := c.Get("cn"); ok {
		// 	c.JSON(200, gin.H{"message": fmt.Sprintf("Hello from private for services to %s", v)})
		// } else {
		// 	c.JSON(200, gin.H{"message": "Hello from private for services without cn"})
		// }
	})
	// r.StaticFS("/export", http.Dir(export.GetExcelFullPath()))
	// r.StaticFS("/upload/images", http.Dir(upload.GetImageFullPath()))
	// r.StaticFS("/qrcode", http.Dir(qrcode.GetQrCodeFullPath()))

	// r.POST("/auth", api.GetAuth)
	// r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	// r.POST("/upload", api.UploadImage)

	// apiv1 := r.Group("/user")
	// // apiv1.Use(jwt.JWT())
	// {
	// 	//获取标签列表
	// 	apiv1.GET("/tags", user.GetTags)
	// 	//新建标签
	// }

	return router
}
