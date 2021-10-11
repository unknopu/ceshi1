package service

import (
	"ceshi1/account/handler"
	"ceshi1/account/model/repository"
	"fmt"
	"io/ioutil"

	"log"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// will initialize a hanler starting from data sources
// which inject into repository layer
// which inject into service layer
// which inject into handler layer
func Inject(d *gorm.DB) (*gin.Engine, error) {
	log.Println("Injecting data sources")

	// repository layer
	userRepository := repository.NewUserRepository(d)
	userService := NewUserService(&USConfig{
		UserRepository: userRepository,
	})
	
	// load RSA keys
	privKeyFile := "./rsa_private.pem"
	priv, err := ioutil.ReadFile(privKeyFile)
	if err != nil {
		return nil, fmt.Errorf("could not read private key pem file: %v", err)
	}
	privKey, err := jwt.ParseRSAPrivateKeyFromPEM(priv)
	if err != nil {
		return nil, fmt.Errorf("could not parse private key: %v", err)
	}

	pubKeyFile := "./rsa_public.pem"
	pub, err := ioutil.ReadFile(pubKeyFile)
	if err != nil {
		return nil, fmt.Errorf("could not read public key pem file: %v", err)
	}
	pubKey, err := jwt.ParseRSAPublicKeyFromPEM(pub)
	if err != nil {
		return nil, fmt.Errorf("could not parse private key: %v", err)
	}

	// load refresh token secret from env variable
	refreshSecret := "1qaz2wsx3edc4rfv"

	tokenService := NewTokenService(&TSConfig{
		PrivKey: privKey,
		PubKey: pubKey,
		RefreshSecret: refreshSecret,
	})

	// initialize gin.Engine
	router := gin.Default()
	handler.NewHandler(&handler.Config{
		R: router,
		UserService: userService,
		TokenService: tokenService,
	})

	return router, nil
}

