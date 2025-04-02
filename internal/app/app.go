package app

import (
	"testtesttest/config"
	c_uduc "testtesttest/internal/c_ud/usecase"
	searchuc "testtesttest/internal/search/usecase"
	httpserver "testtesttest/pkg/http_server"
	"testtesttest/pkg/logger"
	"testtesttest/pkg/postgres"
)

type App struct {
	Server httpserver.HttpServer
	Logger logger.Logger
	DB     *postgres.Postgres
}

func NewApp(Logger logger.Logger, Config config.Config) (*App, error) {
	Server := httpserver.NewServer(&Config)
	DB, err := postgres.NewDB(Config)
	if err != nil {
		Logger.Error(err)
		return nil, err
	}
	Logger.Debug("DB connected")
	C_UD := c_uduc.NewAdd(DB, Logger, Config.OpenApi.ApiGender, Config.OpenApi.ApiAge, Config.OpenApi.ApiNation)
	Logger.Debug("Mapping Handlers")
	Search := searchuc.NewSearchUC(DB, Logger)
	Server.MapGet("/GetUser", Search.SearchWithPagination)
	Server.MapPost("/AddUser", C_UD.AddUser)
	Server.MapDelete("/DeleteUser", C_UD.DeleteUser)
	Server.MapPut("/ChangeUser", C_UD.ChangeUser)
	Logger.Debug("Server Ready")
	return &App{Logger: Logger, Server: *Server, DB: DB}, nil
}

func (App *App) Run() {
	App.Logger.Info("Statring Server")
	App.Server.Run()
	defer App.DB.Close()
}
