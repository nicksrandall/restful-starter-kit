package apis

import (
	"net/http"
	"testing"

	"github.com/go-chi/jwtauth"
	"github.com/nicksrandall/restful-starter-kit/app"
)

func TestAuth(t *testing.T) {
	router := newRouter()
	tokenAuth := jwtauth.New(app.Config.JWTSigningMethod, []byte("secret"), nil)
	router.Post("/auth", Auth(tokenAuth))
	runAPITests(t, router, []apiTestCase{
		{"t1 - successful login", "POST", "/auth", `{"username":"demo", "password":"pass"}`, http.StatusOK, ""},
		{"t2 - unsuccessful login", "POST", "/auth", `{"username":"demo", "password":"bad"}`, http.StatusUnauthorized, ""},
	})
}
