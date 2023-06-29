package main

import (
	"context"
	"log"

	_ "github.com/castmetal/golang-api-boilerplate/docs"
	"github.com/castmetal/golang-api-boilerplate/src/config"
	"github.com/castmetal/golang-api-boilerplate/src/infra/redis"
	"github.com/castmetal/golang-api-boilerplate/src/infra/storage"
	"github.com/castmetal/golang-api-boilerplate/src/ports"
)

// @title           Example API
// @version         1.0
// @description     This is a sample server
// @termsOfService  http://swagger.io/terms/

// @contact.name   Michel La Guardia
// @contact.url    https://www.github.com/castmetal
// @contact.email  mlaguardia@gmail.com

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      http://localhost:8088

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/
func main() {
	if err := config.Init(); err != nil {
		log.Fatal("could not initialize environment", err)
		return
	}

	limit := make(chan struct{}, 1)
	ctx := context.Background()

	cfg := config.EnvStruct(config.Env)

	redisConn := redis.NewRedisClient(cfg)
	pgConn, err := storage.NewPostgresConnection(ctx, cfg)
	if err != nil {
		log.Fatal("could not initialize postgres", err)
	}

	for {
		select {
		case <-ctx.Done():
			break
		case limit <- struct{}{}:

			switch config.Env.Server.Type {
			case "http":
				ginServer := ports.NewGinServer(cfg, pgConn, redisConn)
				ginServer.CreateServer(ctx)
				break
			default:
				ginServer := ports.NewGinServer(cfg, pgConn, redisConn)
				ginServer.CreateServer(ctx)
			}
		}
	}

}
