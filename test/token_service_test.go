package test

import (
	"ceshi1/account/model"
	"ceshi1/account/service"

	"context"
	"fmt"
	"io/ioutil"
	"testing"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)
func TestNewPairFromUser(t *testing.T) {
	priv, _ := ioutil.ReadFile("../rsa_private.pem")
	privKey, _ := jwt.ParseRSAPrivateKeyFromPEM(priv)
	pub, _ := ioutil.ReadFile("../rsa_public.pem")
	pubKey, _ := jwt.ParseRSAPublicKeyFromPEM(pub)
	secret := "anotsorandomtestsecret"

	// instantiate a common token service to be used by all tests
	tokenService := service.NewTokenService(&service.TSConfig{
		PrivKey:               privKey,
		PubKey:                pubKey,
		RefreshSecret:         secret,
	})

	// include password to make sure it is not serialized
	// since json tag is "-"
	uid, _ := uuid.NewRandom()
	u := &model.User{
		UID:      uid,
		Email:    "bob@bob.com",
		Password: "blarghedymcblarghface",
	}

	t.Run("Returns a token pair with proper values", func(t *testing.T) {
		ctx := context.TODO()                                    // updated from context.TODO()
		tokenPair, err := tokenService.NewPairFromUser(ctx, u, "") // replaced "" with prevID from setup
		assert.NoError(t, err)
	
		var s string
		assert.IsType(t, s, tokenPair.IDToken)

		// decode the Base64URL encoded string
		// simpler to use jwt library which is already imported
		idTokenClaims := &service.IDTokenCustomClaims{}
		

		_, err = jwt.ParseWithClaims(tokenPair.IDToken, idTokenClaims, func(token *jwt.Token) (interface{}, error) {
			return pubKey, nil
		})
		assert.NoError(t, err)
		fmt.Println(idTokenClaims.User)
		fmt.Println(idTokenClaims.User)
		fmt.Println(idTokenClaims.User)
		fmt.Println(idTokenClaims.User)
		

		// assert claims on idToken
		expectedClaims := []interface{}{
			u.UID,
			u.Email,
			u.Name,
			u.ImageURL,
			u.Website,
		}
		actualIDClaims := []interface{}{
			idTokenClaims.User.UID,
			idTokenClaims.User.Email,
			idTokenClaims.User.Name,
			idTokenClaims.User.ImageURL,
			idTokenClaims.User.Website,
		}
		

		assert.ElementsMatch(t, expectedClaims, actualIDClaims)
		assert.Empty(t, idTokenClaims.User.Password) // password should never be encoded to json

		expiresAt := time.Unix(idTokenClaims.StandardClaims.ExpiresAt, 0)
		expectedExpiresAt := time.Now().Add(15 * time.Minute)
		assert.WithinDuration(t, expectedExpiresAt, expiresAt, 5*time.Second)

		refreshTokenClaims := &service.RefreshTokenCustomClaims{}
		_, err = jwt.ParseWithClaims(tokenPair.RefreshToken, refreshTokenClaims, func(token *jwt.Token) (interface{}, error) {
			return []byte(secret), nil
		})
		assert.IsType(t, s, tokenPair.RefreshToken)

		// assert claims on refresh token
		assert.NoError(t, err)
		assert.Equal(t, u.UID, refreshTokenClaims.UID)

		expiresAt = time.Unix(refreshTokenClaims.StandardClaims.ExpiresAt, 0)
		expectedExpiresAt = time.Now().Add(3 * 24 *time.Hour)
		assert.WithinDuration(t, expectedExpiresAt, expiresAt, 5*time.Second)
	})
}