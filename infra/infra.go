package infra

import (
	"finance/config"
	"finance/controller"
	"finance/database"
	"finance/helper"
	"finance/repository"
	"finance/service"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type IntegrationContext struct {
	Cfg   *config.Config
	DB    *gorm.DB
	Log   *zap.Logger
	Ctl   *controller.AllController
	Cache *database.Cache
	// Middleware middleware.AllHandler
}

func NewIntegrateContext() (*IntegrationContext, error) {

	errorHandler := func(err error) (*IntegrationContext, error) {
		return nil, err
	}

	config, err := config.SetConfig()
	if err != nil {
		return errorHandler(err)
	}

	log, err := helper.InitLog(config)
	if err != nil {
		return errorHandler(err)
	}

	jwt := helper.NewJwt(config, log)

	db, err := database.SetDatabase(config)
	if err != nil {
		return errorHandler(err)
	}

	rdb := database.NewCache(config, 0)

	repo := repository.NewAllRepo(db, log)

	service := service.NewAllService(repo, log, jwt)

	// middleware := middleware.NewMiddleware(service, log)

	handler := controller.NewAllController(service, log, rdb)

	return &IntegrationContext{
		Cfg:   config,
		DB:    db,
		Log:   log,
		Ctl:   handler,
		Cache: rdb,
		// Middleware: *middleware,
	}, nil
}
