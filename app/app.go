package app

import (
	"context"

	"github.com/atadzan/simple-crud/pkg/controller"
	"github.com/atadzan/simple-crud/pkg/repository"
	"github.com/atadzan/simple-crud/third_party/cache"
	"github.com/atadzan/simple-crud/third_party/database"
	"github.com/atadzan/simple-crud/third_party/server"
	"github.com/atadzan/simple-crud/third_party/storage"
)

func Init(configPath string) error {
	// init app config from input path
	appCfg, err := loadConfig(configPath)
	if err != nil {
		return err
	}
	ctx := context.Background()

	// init database pool connection
	dbClient, err := database.New(ctx, database.Config{
		Username: appCfg.Postgres.Username,
		Password: appCfg.Postgres.Password,
		Host:     appCfg.Postgres.Host,
		Port:     appCfg.Postgres.Port,
		DBName:   appCfg.Postgres.DBName,
		SSLMode:  appCfg.Postgres.SSLMode,
	})
	if err != nil {
		return err
	}

	// init cache connection
	cacheClient, err := cache.New(ctx, cache.Params{
		Host:     appCfg.Cache.Host,
		Port:     appCfg.Cache.Port,
		Password: appCfg.Cache.Password,
		DB:       appCfg.Cache.DB,
	})
	if err != nil {
		return err
	}

	// init storage connection
	storageClient, err := storage.New(storage.Params{
		Endpoint:          appCfg.Storage.Endpoint,
		AccessKeyId:       appCfg.Storage.AccessKeyId,
		SecretAccessKeyId: appCfg.Storage.SecretAccessKeyId,
	})
	repo := repository.New(dbClient, storageClient, cacheClient)
	ctl := controller.New(repo)
	app := ctl.InitRoutes()

	server.StartServerWithGracefulShutdown(app, appCfg.HTTP.Port)
	return nil
}
