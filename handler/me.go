package handler

import (
	"ceshi1/account/model"
	"ceshi1/account/model/apperrors"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// me handler calls services for getting a user's details
func (h *Handler) Me(c *gin.Context){
	// A *model.User will eventually be added to context in middleware
	user, exists := c.Get("user")

	// this should not happen, as out middleware ought to throw an error.
	// this is an extra safety measure
	// we will extract this logic later as it will be common to allhandler
	// method which require a valid user
	if !exists{
		log.Printf("Unable to extract user from request context for unkown reason: %v", c)
		err := apperrors.NewInternal()
		c.JSON(err.Status(), gin.H{
			"error": err,
		})
		return
	}

	uid := user.(*model.User).UID

	// gin.Context satisfies go's context.Context interface
	u, err := h.UserService.Get(c, uid)
	if err != nil{
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