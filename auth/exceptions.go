package auth

// Exception implements a structure for passing exceptions via json
type Exception struct {
    Message string `json:"message"`
}
