package apis

import (
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/go-chi/jwtauth"
	"github.com/nicksrandall/restful-starter-kit/errors"
	"github.com/nicksrandall/restful-starter-kit/models"
	"github.com/nicksrandall/restful-starter-kit/utils"
)

type Credential struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type TokenPayload struct {
	Token string `json:"token"`
}

func Auth(tokenAuth *jwtauth.JWTAuth) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var credential Credential
		if err := json.NewDecoder(r.Body).Decode(&credential); err != nil {
			utils.HandleError(w, errors.Unauthorized(err.Error()))
			return
		}

		identity := authenticate(credential)
		if identity == nil {
			utils.HandleError(w, errors.Unauthorized("invalid credential"))
			return
		}

		_, tokenString, err := tokenAuth.Encode(jwt.MapClaims{
			"id":   identity.GetID(),
			"name": identity.GetName(),
			"exp":  time.Now().Add(time.Hour * 72).Unix(),
		})

		utils.Write(w, &TokenPayload{tokenString}, err)
	}
}

func authenticate(c Credential) models.Identity {
	if c.Username == "demo" && c.Password == "pass" {
		return &models.User{ID: "100", Name: "demo"}
	}
	return nil
}
