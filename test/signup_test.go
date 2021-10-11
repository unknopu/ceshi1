package test

import (
	"bytes"
	"ceshi1/account/model"
	"ceshi1/account/model/apperrors"
	"ceshi1/account/model/mocks"
	"ceshi1/account/handler"

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

	// ####################################################################################
	// Separator	Separator	Separator	Separator	Separator	Separator	Separator
	// Separator	Separator	Separator	Separator	Separator	Separator	Separator
	// ####################################################################################
	t.Run("Email and Password Required", func(t *testing.T) {
		// we just wang this to show that it is not called in this case
		mockUserService := new(mocks.MockUserService)
		mockUserService.On("Signup", mock.AnythingOfType("*gin.Context"), mock.AnythingOfType("*model.User")).Return(nil)

		// a response recorder for getting writen http response
		rr := httptest.NewRecorder()

		// we do not need a middleware as we dont yet have authorized user
		router := gin.Default()

		handler.NewHandler(&handler.Config{
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

	// ####################################################################################
	// Separator	Separator	Separator	Separator	Separator	Separator	Separator
	// Separator	Separator	Separator	Separator	Separator	Separator	Separator
	// ####################################################################################
	t.Run("Invalid email", func(t *testing.T) {
		// we just wang this to show that it is not called in this case
		mockUserService := new(mocks.MockUserService)
		mockUserService.On("Signup", mock.AnythingOfType("*gin.Context"),
			mock.AnythingOfType("*model.User")).Return(nil)

		// a response recorder for getting writen http response
		rr := httptest.NewRecorder()

		// we do not need a middleware as we dont yet have authorized user
		router := gin.Default()

		handler.NewHandler(&handler.Config{
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

	// ####################################################################################
	// Separator	Separator	Separator	Separator	Separator	Separator	Separator
	// Separator	Separator	Separator	Separator	Separator	Separator	Separator
	// ####################################################################################
	t.Run("Password too short", func(t *testing.T) {
		// we just wang this to show that it is not called in this case
		mockUserService := new(mocks.MockUserService)
		mockUserService.On("Signup", mock.AnythingOfType("*gin.Context"),
			mock.AnythingOfType("*model.User")).Return(nil)

		// a response recorder for getting writen http response
		rr := httptest.NewRecorder()

		// we do not need a middleware as we dont yet have authorized user
		router := gin.Default()

		handler.NewHandler(&handler.Config{
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

	// ####################################################################################
	// Separator	Separator	Separator	Separator	Separator	Separator	Separator
	// Separator	Separator	Separator	Separator	Separator	Separator	Separator
	// ####################################################################################
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

		handler.NewHandler(&handler.Config{
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

	// ####################################################################################
	// Separator	Separator	Separator	Separator	Separator	Separator	Separator
	// Separator	Separator	Separator	Separator	Separator	Separator	Separator
	// ####################################################################################
	t.Run("Succussful Token Creation", func(t *testing.T) {
		u := &model.User{
			Email:    "alice@bob.com",
			Password: "supersafepassword",
		}

		mockTokenResp := &model.TokenPair{
			IDToken:      "idToken",
			RefreshToken: "refreshToken",
		}

		mockUserService := new(mocks.MockUserService)
		mockTokenService := new(mocks.MockTokenService)

		mockUserService.On("Signup", mock.AnythingOfType("*gin.Context"), u).
			Return(nil)
		mockTokenService.On("NewPairFromUser", mock.AnythingOfType("*gin.Context"), u, "").
			Return(mockTokenResp, nil)

		// a response recorder for getting written http response
		rr := httptest.NewRecorder()

		// we do not need a middleware as we do not yet have authorized user
		router := gin.Default()

		handler.NewHandler(&handler.Config{
			R:            router,
			UserService:  mockUserService,
			TokenService: mockTokenService,
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

		request.Header.Set("Content-Type", "application/json")

		router.ServeHTTP(rr, request)

		respBody, err := json.Marshal(gin.H{
			"tokens": mockTokenResp,
		})
		assert.NoError(t, err)

		assert.Equal(t, http.StatusCreated, rr.Code)
		assert.Equal(t, respBody, rr.Body.Bytes())

		mockUserService.AssertExpectations(t)
		mockTokenService.AssertExpectations(t)

	})

	// ####################################################################################
	// Separator	Separator	Separator	Separator	Separator	Separator	Separator
	// Separator	Separator	Separator	Separator	Separator	Separator	Separator
	// ####################################################################################
	t.Run("Failed Token Creation", func(t *testing.T) {
		u := &model.User{
			Email:    "bob@bob.com",
			Password: "avalidpassword",
		}

		mockErrorResponse := apperrors.NewInternal()
		mockUserService := new(mocks.MockUserService)
		mockTokenService := new(mocks.MockTokenService)

		mockUserService.On("Signup", mock.AnythingOfType("*gin.Context"), u).
	Return(nil)
		mockTokenService.On("NewPairFromUser", mock.AnythingOfType("*gin.Context"), u, "").
	Return(nil, mockErrorResponse)

		// a response recorder for getting written http response
		rr := httptest.NewRecorder()

		// don't need a middleware as we don't yet have authorized user
		router := gin.Default()

		handler.NewHandler(&handler.Config{
			R:            router,
			UserService:  mockUserService,
			TokenService: mockTokenService,
		})

		// create a request body with empty email and password
		reqBody, err := json.Marshal(gin.H{
			"email":    u.Email,
			"password": u.Password,
		})
		assert.NoError(t, err)

		// use bytes.NewBuffer to create a reader
		request, err := http.NewRequest(http.MethodPost, "/signup", bytes.NewBuffer(reqBody))
		assert.NoError(t, err)

		request.Header.Set("Content-Type", "application/json")

		router.ServeHTTP(rr, request)

		respBody, err := json.Marshal(gin.H{
			"error": mockErrorResponse,
		})
		assert.NoError(t, err)

		assert.Equal(t, mockErrorResponse.Status(), rr.Code)
		assert.Equal(t, respBody, rr.Body.Bytes())

		mockUserService.AssertExpectations(t)
		mockTokenService.AssertExpectations(t)
	})
}
