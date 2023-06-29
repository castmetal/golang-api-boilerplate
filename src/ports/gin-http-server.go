package ports

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/castmetal/golang-api-boilerplate/src/config"
	rest_controllers "github.com/castmetal/golang-api-boilerplate/src/controllers/rest"
	"github.com/castmetal/golang-api-boilerplate/src/domains/common/logger"
	"github.com/castmetal/golang-api-boilerplate/src/infra/redis"
	"github.com/jackc/pgx/v4/pgxpool"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type GinServer struct {
	config    config.EnvStruct
	pgConn    *pgxpool.Pool
	redisConn redis.IRedisClient
}

func NewGinServer(config config.EnvStruct, pgConn *pgxpool.Pool, redisConn redis.IRedisClient) ServerInterface {
	return &GinServer{
		config:    config,
		pgConn:    pgConn,
		redisConn: redisConn,
	}
}

func (s *GinServer) CreateServer(ctx context.Context) {
	logger.Setup(s.config.Logger.Env)

	isDev := IsDev(s.config)
	if isDev {
		ServerLog(s.config)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.Default()
	cfg := cors.DefaultConfig()
	cfg.AllowOrigins = []string{"*"}

	_ = router.SetTrustedProxies([]string{})
	router.RemoteIPHeaders = []string{"X-Forwarded-For"}

	router.GET("/healthCheck", func(c *gin.Context) {
		c.IndentedJSON(http.StatusOK, gin.H{
			"alive": true,
		})

		logger.Info(ctx, "server_up")
	})

	controllers := rest_controllers.NewRestControllers(s.config, router, s.pgConn, s.redisConn)
	controllers.SetRestControllers()

	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.Use(cors.New(cfg))
	srvr := &http.Server{
		Addr:           ":" + s.config.Server.Port,
		Handler:        router,
		ReadTimeout:    15 * time.Second,
		WriteTimeout:   15 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	log.SetFlags(0)

	logger.Info(ctx, fmt.Sprintf("Server listen at port :%s", s.config.Server.Port))

	log.Fatal(srvr.ListenAndServe())
}
