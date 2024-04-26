package server

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"html/template"
	"io"
	"robot-system-server/docs"
	"robot-system-server/internal/handler"
	"robot-system-server/internal/middleware"
	"robot-system-server/pkg/jwt"
	"robot-system-server/pkg/log"
	"robot-system-server/pkg/server/http"
	"strings"
)

func NewHTTPServer(
	logger *log.Logger,
	conf *viper.Viper,
	jwt *jwt.JWT,
	userHandler *handler.UserHandler,
	beasenHandler *handler.BeasenHandler,
	fortuneHandler *handler.FortuneHandler,
	signInHandler *handler.SignInHandler,
	spiritSignHandler *handler.SpiritSignHandler,
	fileHandler *handler.FileHandler,
) *http.Server {
	if conf.GetString("env") == "local" {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}
	s := http.NewServer(
		gin.Default(),
		logger,
		http.WithServerHost(conf.GetString("http.host")),
		http.WithServerPort(conf.GetInt("http.port")),
	)
	s.Static("/css", "./web/static/css")
	s.StaticFile("favicon.ico", "./web/static/favicon.ico")
	// 加载模板引擎，并指定自定义函数

	t, err := loadTemplate()
	if err != nil {
		panic(err)
	}
	s.SetHTMLTemplate(t)

	//s.SetHTMLTemplate(tmpl)
	if conf.GetString("env") == "local" {
		// swagger doc
		docs.SwaggerInfo.BasePath = "/v1"
		s.GET("/swagger/*any", ginSwagger.WrapHandler(
			swaggerfiles.Handler,
			//ginSwagger.URL(fmt.Sprintf("http://localhost:%d/swagger/doc.json", conf.GetInt("app.http.port"))),
			ginSwagger.DefaultModelsExpandDepth(-1),
		))
	}

	s.Use(
		middleware.CORSMiddleware(),
		middleware.ResponseLogMiddleware(logger),
		middleware.RequestLogMiddleware(logger),
		//middleware.SignMiddleware(log),
	)
	s.GET("/", func(ctx *gin.Context) {
		ctx.HTML(200, "index.html", nil)
	})

	api := s.Group("/api")
	{
		api.POST("/getFortuneOfToday", fortuneHandler.GetFortuneOfToday)
		api.GET("/get/:qq", fortuneHandler.GetFortuneOfTodayOld)
		api.POST("/signIn", signInHandler.SignIn)
		api.POST("/querySignInData", signInHandler.QuerySignInData)
		api.POST("/addSignInPoints", signInHandler.AddSignInPoints)
		api.GET("/getSen", beasenHandler.RandResult)
	}
	file := s.Group("/file")
	{
		file.POST("/uploadFile", fileHandler.UploadFile)
		file.GET("/img/:bucket/:file", fileHandler.Img)
	}
	lq := s.Group("/lq")
	{
		lq.POST("/spiritSign", spiritSignHandler.OneSignPerDay)
		lq.GET("/v/cs/:qq", spiritSignHandler.ViewCs)
		lq.GET("/v/yl/:qq", spiritSignHandler.ViewYl)
		lq.GET("/v/gy/:qq", spiritSignHandler.ViewGy)
	}

	v1 := s.Group("/v1")
	{
		// No route group has permission
		noAuthRouter := v1.Group("/")
		{
			noAuthRouter.POST("/register", userHandler.Register)
			noAuthRouter.POST("/login", userHandler.Login)
		}
		// Non-strict permission routing group
		noStrictAuthRouter := v1.Group("/").Use(middleware.NoStrictAuth(jwt, logger))
		{
			noStrictAuthRouter.GET("/user", userHandler.GetProfile)
		}

		// Strict permission routing group
		strictAuthRouter := v1.Group("/").Use(middleware.StrictAuth(jwt, logger))
		{
			strictAuthRouter.PUT("/user", userHandler.UpdateProfile)
		}
	}

	return s
}

// loadTemplate 加载由 go-assets-builder 嵌入的模板
// go-assets-builder web/templates -p server -o internal/server/assets.go
func loadTemplate() (*template.Template, error) {
	t := template.New("")
	for name, file := range Assets.Files {
		if file.IsDir() || !strings.HasSuffix(name, ".html") {
			continue
		}
		h, err := io.ReadAll(file)
		if err != nil {
			return nil, err
		}
		t, err = t.New(name).Funcs(template.FuncMap{
			"replace": func(s, old, new string) template.HTML {
				return template.HTML(strings.ReplaceAll(s, old, new))
			},
			"split": strings.Split,
		}).Parse(string(h))
		if err != nil {
			return nil, err
		}
	}
	return t, nil
}
