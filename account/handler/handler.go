package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jj-attaq/api-todo/account/model"
)

// holds required services for handler to function
type Handler struct{
    UserService model.UserService
}

// holds services to be injected in handler layer on initialization
type Config struct {
    R *gin.Engine
    UserService model.UserService
}

func NewHandler(c *Config) {
    h := &Handler{
        UserService: c.UserService,
    }

    // g := c.R.Group("localhost:8080/api/account/")
    g := c.R.Group("/api/account")

    g.GET("/me", h.Me)
    g.POST("/signup", h.SignUp)
    g.POST("/signin", h.SignIn)
    g.POST("/signout", h.SignOut)
    g.POST("/tokens", h.Tokens)
}

func (h *Handler) SignUp(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{
        "hello": "it's SignUp",
    })
}
func (h *Handler) SignIn(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{
        "hello": "it's SignIn",
    })
}
func (h *Handler) SignOut(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{
        "hello": "it's SignOut",
    })
}
func (h *Handler) Tokens(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{
        "hello": "it's Tokens",
    })
}
