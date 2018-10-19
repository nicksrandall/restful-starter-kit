package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/jwtauth"
	"github.com/jmoiron/sqlx"
	"github.com/nicksrandall/restful-starter-kit/apis"
	"github.com/nicksrandall/restful-starter-kit/app"
	"github.com/nicksrandall/restful-starter-kit/errors"
	"github.com/nicksrandall/restful-starter-kit/services"
)

func main() {
	// load application configurations
	if err := app.LoadConfig("./config"); err != nil {
		panic(fmt.Errorf("Invalid application configuration: %s", err))
	}

	// load error messages
	if err := errors.LoadMessages(app.Config.ErrorFile); err != nil {
		panic(fmt.Errorf("Failed to read the error message file: %s", err))
	}

	dbx, err := app.InitDB()
	if err != nil {
		panic(err)
	}

	// wire up API routing
	http.Handle("/", buildRouter(dbx))

	// start the server
	address := fmt.Sprintf(":%v", app.Config.ServerPort)
	app.Logger(nil).Sugar().Infof("server %v is started at %v\n", app.Version, address)
	panic(http.ListenAndServe(address, nil))
}

func buildRouter(db *sqlx.DB) chi.Router {
	router := chi.NewRouter()

	var corsConfig = cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Origin", "Accept", "Authorization", "Content-Type", "content-type", "X-Amz-Date", "X-Api-Key", "X-Amz-Security-Token"},
		AllowCredentials: true,
		MaxAge:           300,
	})

	tokenAuth := jwtauth.New(app.Config.JWTSigningMethod, []byte(app.Config.JWTSigningKey), nil)
	_ = tokenAuth

	router.Use(middleware.Heartbeat("/ping"))
	router.Post("/auth", apis.Auth(tokenAuth))

	router.Route("/v1", func(r chi.Router) {
		r.Use(
			middleware.RequestID,
			middleware.RealIP,
			middleware.Recoverer,
			corsConfig.Handler,
			app.LoggerMiddleware(app.Logger(nil)),
			app.DBMiddleware(db),
			middleware.DefaultCompress,
			jwtauth.Verifier(tokenAuth),
			jwtauth.Authenticator,
		)

		apis.ServeArtistResource(r, &services.ArtistService{})
		// wire up more resource APIs here
	})

	return router
}
