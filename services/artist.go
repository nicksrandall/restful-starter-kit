package services

import (
	"context"
	"database/sql"

	"github.com/nicksrandall/restful-starter-kit/app"
	"github.com/nicksrandall/restful-starter-kit/models"
)

// artistDAO specifies the interface of the artist DAO needed by ArtistService.
type artistDAO interface {
	// Get returns the artist with the specified artist ID.
	Get(ctx context.Context, id int) (*models.Artist, error)
	// Count returns the number of artists.
	Count(ctx context.Context) (int, error)
	// Query returns the list of artists with the given offset and limit.
	Query(ctx context.Context, offset, limit int) ([]models.Artist, error)
	// Create saves a new artist in the storage.
	Create(ctx context.Context, artist *models.Artist) error
	// Update updates the artist with given ID in the storage.
	Update(ctx context.Context, id int, artist *models.Artist) error
	// Delete removes the artist with given ID from the storage.
	Delete(ctx context.Context, id int) error
}

// ArtistService provides services related with artists.
type ArtistService struct {
}

// Get returns the artist with the specified the artist ID.
func (s *ArtistService) Get(ctx context.Context, id int) (*models.Artist, error) {
	// TODO implement
	return nil, nil
}

// Create creates a new artist.
func (s *ArtistService) Create(ctx context.Context, model *models.Artist) (*models.Artist, error) {
	if err := model.Validate(); err != nil {
		return nil, err
	}
	return nil, nil
}

// Update updates the artist with the specified ID.
func (s *ArtistService) Update(ctx context.Context, id int, model *models.Artist) (*models.Artist, error) {
	if err := model.Validate(); err != nil {
		return nil, err
	}
	return nil, nil
}

// Delete deletes the artist with the specified ID.
func (s *ArtistService) Delete(ctx context.Context, id int) (*models.Artist, error) {
	return nil, nil
}

// Count returns the number of artists.
func (s *ArtistService) Count(ctx context.Context) (int, error) {
	return 0, nil
}

// Query returns the artists with the specified offset and limit.
func (s *ArtistService) Query(ctx context.Context, offset, limit int) (list []models.Artist, err error) {
	dbx := app.DBFromContext(ctx)
	err = dbx.SelectContext(ctx, &list, "select id, name from artist")
	if err == sql.ErrNoRows {
		return make([]models.Artist, 0), nil
	}
	return
}
