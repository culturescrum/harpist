package auth

import (
  "encoding/json"
  "net/http"
  "strings"
  "fmt"
  "time"

  "github.com/culturescrum/harpist/models"

  "github.com/gorilla/context"
  "github.com/dgrijalva/jwt-go"
)

type UserClaim struct {
  *jwt.StandardClaims
  TokenType string
  models.User
}

// CreateToken implements token creation endpoint
func CreateToken(w http.ResponseWriter, req *http.Request) {
    var user models.User
    // TODO: retrieve user
    _ = json.NewDecoder(req.Body).Decode(&user)
    token := jwt.New(jwt.SigningMethodHS256)
    token.Claims = &UserClaim{
      &jwt.StandardClaims{
        // TODO: Use config value for timeouts
        ExpiresAt: time.Now().Add(time.Minute*30).Unix(),
      },
      "level1",
      user,
    }
    // TODO: Use configuration value for secret
    tokenString, error := token.SignedString([]byte("secret"))
    if error != nil {
      fmt.Println(error)
    }
    json.NewEncoder(w).Encode(models.JwtToken{Token: tokenString})
}

// ValidateUserMiddleware is a wrapper function for http handlers that
// processes JWT tokens
func ValidateUserMiddleware(next http.HandlerFunc) http.HandlerFunc {
    return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
        authorizationHeader := req.Header.Get("authorization")
        if authorizationHeader != "" {
            bearerToken := strings.Split(authorizationHeader, " ")
            if len(bearerToken) == 2 {
                token, error := jwt.ParseWithClaims(bearerToken[1], &UserClaim{}, func(token *jwt.Token) (interface{}, error) {
                  // TODO: use configuration secret
                  return []byte("secret"), nil
                })
                if error != nil {
                    json.NewEncoder(w).Encode(Exception{Message: error.Error()})
                    return
                }
                if token.Valid {
                    claims := token.Claims.(*UserClaim)
                    context.Set(req, "user_id", claims.User.ID)
                    context.Set(req, "username", claims.User.Username)
                    context.Set(req, "user", claims.User)
                    next(w, req)
                } else {
                    json.NewEncoder(w).Encode(Exception{Message: "Invalid authorization token"})
                }
            }
        } else {
            json.NewEncoder(w).Encode(Exception{Message: "An authorization header is required"})
        }
    })
}
