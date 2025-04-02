package search

import "github.com/gin-gonic/gin"

type SearchUC interface {
	SearchWithPagination(c *gin.Context)
}
