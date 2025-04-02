package c_ud

import "github.com/gin-gonic/gin"

// C_UD is for Create, Update, Delete
type C_UDUC interface {
	AddUser(c *gin.Context)
	ChangeUser(c *gin.Context)
	DeleteUser(c *gin.Context)
}
