package apis

import (
	"context"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/nicksrandall/restful-starter-kit/models"
	"github.com/nicksrandall/restful-starter-kit/utils"
)

type (
	// artistService specifies the interface for the artist service needed by artistResource.
	artistService interface {
		Get(ctx context.Context, id int) (*models.Artist, error)
		Query(ctx context.Context, offset, limit int) ([]models.Artist, error)
		Count(ctx context.Context) (int, error)
		Create(ctx context.Context, model *models.Artist) (*models.Artist, error)
		Update(ctx context.Context, id int, model *models.Artist) (*models.Artist, error)
		Delete(ctx context.Context, id int) (*models.Artist, error)
	}

	// artistResource defines the handlers for the CRUD APIs.
	artistResource struct {
		service artistService
	}
)

// ServeArtist sets up the routing of artist endpoints and the corresponding handlers.
func ServeArtistResource(router chi.Router, service artistService) {
	r := &artistResource{service}
	router.Get("/artists/{id}", http.HandlerFunc(r.get))
	router.Get("/artists", http.HandlerFunc(r.query))
	router.Post("/artists", http.HandlerFunc(r.create))
	router.Put("/artists/{id}", http.HandlerFunc(r.update))
	router.Delete("/artists/{id}", http.HandlerFunc(r.delete))
}

func (r *artistResource) get(res http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	id, _ := strconv.Atoi(chi.URLParam(req, "id"))
	model, err := r.service.Get(ctx, id)
	utils.Write(res, model, err)
}

func (r *artistResource) query(res http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	models, err := r.service.Query(ctx, 0, 10)
	utils.Write(res, models, err)
}

func (r *artistResource) create(res http.ResponseWriter, req *http.Request) {
}

func (r *artistResource) update(res http.ResponseWriter, req *http.Request) {
}

func (r *artistResource) delete(res http.ResponseWriter, req *http.Request) {
}
