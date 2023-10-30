package handler

import (
    "net/http"
    "log"
	"github.com/gin-gonic/gin"
	"github.com/jj-attaq/api-todo/account/model/apperrors"
	"github.com/jj-attaq/api-todo/account/model"
)

// Get's user's details
func (h *Handler) Me(c *gin.Context) {
    // log.Printf("This is the '/me' handler: '%v'\n", c.Request.URL)
    user, exists := c.Get("user")
    log.Printf("This is 'user': '%v'\n", exists)

    if !exists {
        log.Printf("Unable to extract user from request context for unknown reason: %v\n", c)
        err := apperrors.NewInternal()
        c.JSON(err.Status(), gin.H{
            "error": err,
        })

        return
    }

    uid := user.(*model.User).UID
    // context
    ctx := c.Request.Context()

    u, err := h.UserService.Get(ctx, uid)
    if err != nil {
        log.Printf("Unable to find user: %v\n%v", uid, err)
        e := apperrors.NewNotFound("user", uid.String())

        c.JSON(e.Status(), gin.H{
            "error": e,
        })
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "user": u,
    })
}
