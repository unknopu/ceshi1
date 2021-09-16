package handler

import (
	"bytes"
	"ceshi1/account/model"
	"ceshi1/account/model/apperrors"
	"ceshi1/account/model/mocks"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestSignup(t *testing.T) {
	// setup
	gin.SetMode(gin.TestMode)

	t.Run("Email and Password Required", func(t *testing.T) {
		// we just wang this to show that it is not called in this case
		mockUserService := new(mocks.MockUserService)
		mockUserService.On("Signup", mock.AnythingOfType("*gin.Context"), mock.AnythingOfType("*model.User")).Return(nil)

		// a response recorder for getting writen http response
		rr := httptest.NewRecorder()

		// we do not need a middleware as we dont yet have authorized user
		router := gin.Default()

		NewHandler(&Config{
			R:           router,
			UserService: mockUserService,
		})

		// create a request body with empty email and password
		reqBody, err := json.Marshal(gin.H{
			"email": "",
		})
		assert.NoError(t, err)

		// user bytes.NewBuffer to create a reader
		request, err := http.NewRequest(http.MethodPost, "/signup", bytes.NewBuffer(reqBody))
		assert.NoError(t, err)

		request.Header.Set("Content-Type", "application/json")

		router.ServeHTTP(rr, request)

		assert.Equal(t, 400, rr.Code)
		mockUserService.AssertNotCalled(t, "Signup")
	})

	t.Run("Invalid email", func(t *testing.T) {
		// we just wang this to show that it is not called in this case
		mockUserService := new(mocks.MockUserService)
		mockUserService.On("Signup", mock.AnythingOfType("*gin.Context"),
			mock.AnythingOfType("*model.User")).Return(nil)

		// a response recorder for getting writen http response
		rr := httptest.NewRecorder()

		// we do not need a middleware as we dont yet have authorized user
		router := gin.Default()

		NewHandler(&Config{
			R:           router,
			UserService: mockUserService,
		})

		// create a request body with empty email and password
		reqBody, err := json.Marshal(gin.H{
			"email":    "alice@bob",
			"password": "superpassword",
		})
		assert.NoError(t, err)

		// user bytes.NewBuffer to create a reader
		request, err := http.NewRequest(http.MethodPost, "/signup", bytes.NewBuffer(reqBody))
		assert.NoError(t, err)

		request.Header.Set("content-type", "application/json")

		router.ServeHTTP(rr, request)
		assert.Equal(t, 400, rr.Code)
		mockUserService.AssertNotCalled(t, "Signup")
	})

	t.Run("Password too short", func(t *testing.T) {
		// we just wang this to show that it is not called in this case
		mockUserService := new(mocks.MockUserService)
		mockUserService.On("Signup", mock.AnythingOfType("*gin.Context"),
			mock.AnythingOfType("*model.User")).Return(nil)

		// a response recorder for getting writen http response
		rr := httptest.NewRecorder()

		// we do not need a middleware as we dont yet have authorized user
		router := gin.Default()

		NewHandler(&Config{
			R:           router,
			UserService: mockUserService,
		})

		// create a request body with empty email and password
		reqBody, err := json.Marshal(gin.H{
			"email":    "alice@bob.com",
			"password": "super",
		})
		assert.NoError(t, err)

		// user bytes.NewBuffer to create a reader
		request, err := http.NewRequest(http.MethodPost, "/signup", bytes.NewBuffer(reqBody))
		assert.NoError(t, err)

		request.Header.Set("content-type", "application/json")

		router.ServeHTTP(rr, request)
		assert.Equal(t, 400, rr.Code)
		mockUserService.AssertNotCalled(t, "Signup")
	})

	t.Run("Error returned from UserService", func(t *testing.T) {
		u := &model.User{
			Email:    "alice@bob.com",
			Password: "supersafepassword",
		}

		// we just wang this to show that it is not called in this case
		mockUserService := new(mocks.MockUserService)
		mockUserService.On("Signup", mock.AnythingOfType("*gin.Context"),
			mock.AnythingOfType("*model.User")).Return(apperrors.NewConflict("User Already Exists", u.Email))

		// a response recorder for getting writen http response
		rr := httptest.NewRecorder()

		// we do not need a middleware as we dont yet have authorized user
		router := gin.Default()

		NewHandler(&Config{
			R:           router,
			UserService: mockUserService,
		})

		// create a request body with empty email and password
		reqBody, err := json.Marshal(gin.H{
			"email":    u.Email,
			"password": u.Password,
		})
		assert.NoError(t, err)

		// user bytes.NewBuffer to create a reader
		request, err := http.NewRequest(http.MethodPost, "/signup", bytes.NewBuffer(reqBody))
		assert.NoError(t, err)

		request.Header.Set("content-type", "application/json")

		router.ServeHTTP(rr, request)
		assert.Equal(t, 409, rr.Code)
		mockUserService.AssertExpectations(t)
	})
}
