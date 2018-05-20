package harpist

import (
	// "encoding/json"
	// "net/http"
	// "strings"
	// "fmt"
	// "time"
	// "context"

	// "github.com/culturescrum/harpist/models"

	"github.com/dgrijalva/jwt-go"
)

// Exception implements a structure for passing exceptions via json
type Exception struct {
	Message string `json:"message"`
}

// UserClaim is for web interactions with the API and authentication
type UserClaim struct {
	*jwt.StandardClaims
	TokenType string
	Username  string
	Name      string
	UserID    uint
	// TODO: implement group pinning here to reduce membership look-ups
}
