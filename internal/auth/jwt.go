package auth

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

// UserClaim is for web interactions with the API and authentication
type UserClaim struct {
	*jwt.StandardClaims
	TokenType string
	Username  string
	Name      string
	UserID    uint
}
