package usecase

import (
	"testtesttest/internal/search"
	"testtesttest/pkg/logger"
	"testtesttest/pkg/postgres"

	"github.com/gin-gonic/gin"
)

type Search struct {
	DB     *postgres.Postgres
	Logger logger.Logger
}

func (S *Search) SearchWithPagination(c *gin.Context) {
	S.Logger.Debug("Got Search Request")
}

func NewSearchUC(DB *postgres.Postgres, Logger logger.Logger) search.SearchUC {
	Logger.Info("Search service ready")
	return &Search{DB: DB, Logger: Logger}
}
