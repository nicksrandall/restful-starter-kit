package apis

import (
	"context"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/gofrs/uuid"
	jsoniter "github.com/json-iterator/go"
	"github.com/nicksrandall/restful-starter-kit/models"
	"github.com/nicksrandall/restful-starter-kit/utils"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

type (
	// artistService specifies the interface for the artist service needed by artistResource.
	artistService interface {
		Get(ctx context.Context, id uuid.UUID) (*models.Artist, error)
		Query(ctx context.Context, offset, limit int) ([]models.Artist, error)
		Count(ctx context.Context) (int, error)
		Create(ctx context.Context, model *models.Artist) (*models.Artist, error)
		Update(ctx context.Context, id uuid.UUID, model *models.Artist) (*models.Artist, error)
		Delete(ctx context.Context, id uuid.UUID) (bool, error)
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
	uuid, err := uuid.FromString(chi.URLParam(req, "id"))
	if err != nil {
		http.Error(res, "Invalid ID", http.StatusBadRequest)
		return
	}
	model, err := r.service.Get(ctx, uuid)
	utils.Write(res, model, err)
}

func (r *artistResource) query(res http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	models, err := r.service.Query(ctx, 0, 10)
	utils.Write(res, models, err)
}

func (r *artistResource) create(res http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	var model models.Artist
	if err := json.NewDecoder(req.Body).Decode(&model); err != nil {
		http.Error(res, "Invalid JSON", http.StatusBadRequest)
		return
	}
	m, err := r.service.Create(ctx, &model)
	utils.Write(res, m, err)
}

func (r *artistResource) update(res http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	uuid, err := uuid.FromString(chi.URLParam(req, "id"))
	if err != nil {
		http.Error(res, "Invalid ID", http.StatusBadRequest)
		return
	}
	var model models.Artist
	if err := json.NewDecoder(req.Body).Decode(&model); err != nil {
		http.Error(res, "Invalid JSON", http.StatusBadRequest)
		return
	}
	m, err := r.service.Update(ctx, uuid, &model)
	utils.Write(res, m, err)
}

func (r *artistResource) delete(res http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	uuid, err := uuid.FromString(chi.URLParam(req, "id"))
	if err != nil {
		http.Error(res, "Invalid ID", http.StatusBadRequest)
		return
	}
	ok, err := r.service.Delete(ctx, uuid)
	utils.Write(res, ok, err)
}
