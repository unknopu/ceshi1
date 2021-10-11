package handler

import (
	"ceshi1/account/model"
	"ceshi1/account/model/apperrors"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// signup request is not exported, hence the lowercase name
// it is used for validation and json marshalling
type signupReq struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,gte=6,lte=30"`
}

// signup handler
func (h *Handler) Signup(c *gin.Context) {

	// // define a variable to which we will bind incoming
	// // json body , including email, password
	var req signupReq

	// bind incoming json to struct and check for validation errors
	if ok := bindData(c, &req); !ok {
		return
	}

	u := &model.User{
		Email:    req.Email,
		Password: req.Password,
	}

	err := h.UserService.Signup(c, u)
	if err != nil {
		log.Printf("Failed to sign up user: %v\n", err.Error())
		c.JSON(apperrors.Status(err), gin.H{
			"error": err,
		})
		return
	}

	// create token pair as strings
	tokens, err := h.TokenService.NewPairFromUser(c, u, "")
	if err != nil {
		log.Printf("Failed to create tokens for user: %v", err.Error())
		// may eventually implement rollback logic here
		// that, if we fail to create tokens after creating a user,
		// we make sure to clear/remove the created user in the database
		c.JSON(apperrors.Status(err), gin.H{
			"error": err,
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"tokens": tokens,
	})
}
