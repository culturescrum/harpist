package auth

import (
  //"encoding/json"
  "net/http"
  //"strings"
  //"fmt"

  //"github.com/culturescrum/harpist/models"

  //"github.com/gorilla/context"
  //"github.com/dgrijalva/jwt-go"
)

// ValidateGroupOwnerMiddleware is a wrapper function for http handlers that
// processes JWT tokens
func ValidateGroupOwnerMiddleware(next http.HandlerFunc) http.HandlerFunc {
    // TODO: handle db calls to retrieve tokenized user from
    // database
    // User{
    //  ID: 'user_id'
    //  Username: 'username'
    // }
    return next
}
