package handler

import (
	"ceshi1/account/model"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

// handler struct holds required services for handler to function
type Handler struct {
	UserService model.UserService
}

// config will hold services that will eventually be injected into this
// handler layer on handler initialization
type Config struct{
	R *gin.Engine
	UserService model.UserService
}

// this is initializes the handler with required injected services along with http toutes
// does not return as it deals directly with a reference to the Engine
func NewHandler(c *Config){
	// create ahandler which will later have injected services
	h := &Handler{
		UserService: c.UserService,
	}

	// create an account group
	// ACCOUNT_API_URL = "api/account"
	g := c.R.Group(os.Getenv("ACCOUNT_API_URL"))
	g.GET("/me", h.Me)
	g.POST("/signup", h.Signup)
	g.POST("/signin", h.Signin)
	g.POST("/tokens", h.Tokens)
	g.POST("/image", h.Image)
	g.DELETE("/image", h.DeleteImage)
	g.PUT("/details", h.Details)
}

// sign-in handler
func (h *Handler) Signin(c *gin.Context){
	c.JSON(http.StatusOK, gin.H{
		"hello": "it's signin",
	})
}

func (h *Handler) Signout(c *gin.Context){
	c.JSON(http.StatusOK, gin.H{
		"hello": "it's sign-out",
	})
}

// tokens handler
func (h *Handler) Tokens(c *gin.Context){
	c.JSON(http.StatusOK, gin.H{
		"hello": "it's tokens",
	})
}

// image handler
func (h *Handler) Image(c *gin.Context){
	c.JSON(http.StatusOK, gin.H{
		"hello": "it's image",
	})
}

// delete image handler
func (h *Handler) DeleteImage(c *gin.Context){
	c.JSON(http.StatusOK, gin.H{
		"hello": "it's delete image",
	})
}

// details handler
func (h *Handler) Details(c *gin.Context){
	c.JSON(http.StatusOK, gin.H{
		"hello": "it's details",
	})
}
