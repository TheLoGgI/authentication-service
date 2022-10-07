package providers

import (
	"net/http"
)

// type Token struct {
// 	user string
// 	type string
// 	exp int
// 	iat int
// }

// {
// 	"sub": "612f5efd4c1cedb60f46baf1",
// 	"tid": "2835540a-cd00-420d-9a27-f361ced1e2fa",
// 	"type": "refresh",
// 	"exp": 1672959099,
// 	"iat": 1665183098
//   }

func Authorize(w http.ResponseWriter, r *http.Request) {

	// Generate JWT Token

	w.Header().Set("Content-Type", "application/json")
	// w.Header().Set("Set-Cookie", )

	w.Write([]byte("Hello wold"))
}
