package auth

import (
	"errors"
	"os"
	"time"

	"github.com/aclel/sense/store"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

type JWTAuth struct {
	SecretKey []byte
}

var jwtAuth *JWTAuth = nil

// Initialise a JWTAuth object which has a secretKey
func InitJWTAuth() (*JWTAuth, error) {
	if jwtAuth == nil {
		secretKey, err := getSecretKey()
		if err != nil {
			return nil, err
		}

		jwtAuth = &JWTAuth{
			SecretKey: secretKey,
		}
	}
	return jwtAuth, nil
}

// GenerateToken generates a JWT for the given user. The token is signed with a private key.
// The user's role is included in the token for authorization.
func (jwtAuth *JWTAuth) GenerateToken(user *store.User) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims["exp"] = time.Now().Add(time.Hour * time.Duration(24)).Unix() // 24 hour expiry
	token.Claims["iat"] = time.Now().Unix()
	token.Claims["sub"] = user.Email
	// token.Claims["role"] = user.Role
	tokenString, err := token.SignedString(jwtAuth.SecretKey)
	if err != nil {
		panic(err)
		return "", err
	}

	return tokenString, nil
}

// Authenticate authenticates the given user. It compares the emails and passwords of the two given Users.
// Returns true if authenticated.
func (jwtAuth *JWTAuth) Authenticate(dbUser *store.User, user *store.User) bool {

	// Check that the emails and password hashes are the same
	return user.Email == dbUser.Email && bcrypt.CompareHashAndPassword([]byte(dbUser.Password), []byte(user.Password)) == nil
}

func getSecretKey() ([]byte, error) {
	key := os.Getenv("FMS_SECRET_KEY")
	if key == "" {
		return []byte(""), errors.New("No environment variable named FMS_SECRET_KEY")
	}

	return []byte(key), nil
}
