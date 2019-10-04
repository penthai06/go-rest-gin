package rest

import (
	"github.com/penthai06/go-rest-gin/backend/src/dblayer"
	"github.com/penthai06/go-rest-gin/backend/src/models"
	"github.com/gin-gonic/gin"
)

type HandlerInterface interface {
  GetProducts(c *gin.Context)
  GetPromos(c *gin.Context)
  AddUser(c *gin.Context)
  SignIn(c *gin.Context)
  SignOut(c *gin.Context)
  GetOrders(c *gin.Context)
  Charge(c *gin.Context)
}

type Handler struct{
	db dblayer.DBLayer
}

func NewHandler() (*Handler, error) {
  //This creates a new pointer to the Handler object
	return new(Handler), nil
}