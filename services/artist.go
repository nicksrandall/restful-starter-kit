package services

import (
	"context"
	"database/sql"

	"github.com/gofrs/uuid"

	"github.com/nicksrandall/restful-starter-kit/app"
	"github.com/nicksrandall/restful-starter-kit/models"
)

// ArtistService provides services related with artists.
type ArtistService struct{}

// Get returns the artist with the specified the artist ID.
func (s *ArtistService) Get(ctx context.Context, id uuid.UUID) (*models.Artist, error) {
	dbx := app.DBFromContext(ctx)
	var model models.Artist
	err := dbx.QueryRowxContext(ctx, "select id, name FROM artist where id = $1 LIMIT 1", id).StructScan(&model)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	return &model, err
}

// Create creates a new artist.
func (s *ArtistService) Create(ctx context.Context, model *models.Artist) (*models.Artist, error) {
	if err := model.Validate(); err != nil {
		return nil, err
	}
	model.ID = uuid.Must(uuid.NewV4())
	dbx := app.DBFromContext(ctx)
	_, err := dbx.NamedExecContext(ctx, `
		INSERT INTO artist (id, name) VALUES
		(:id, :name)
	`, model)
	return model, err
}

// Update updates the artist with the specified ID.
func (s *ArtistService) Update(ctx context.Context, id uuid.UUID, model *models.Artist) (*models.Artist, error) {
	if err := model.Validate(); err != nil {
		return nil, err
	}
	dbx := app.DBFromContext(ctx)
	model.ID = id
	_, err := dbx.NamedExecContext(ctx, "UPDATE artist SET name = :name WHERE id = :id;", model)
	return model, err
}

// Delete deletes the artist with the specified ID.
func (s *ArtistService) Delete(ctx context.Context, id uuid.UUID) (bool, error) {
	dbx := app.DBFromContext(ctx)
	_, err := dbx.ExecContext(ctx, "DELETE FROM artist where id = $1;", id)
	if err == sql.ErrNoRows {
		return false, nil
	}
	return true, err
}

// Count returns the number of artists.
func (s *ArtistService) Count(ctx context.Context) (count int, err error) {
	dbx := app.DBFromContext(ctx)
	err = dbx.GetContext(ctx, &count, "SELECT count(id) FROM artist;")
	return count, err
}

// Query returns the artists with the specified offset and limit.
func (s *ArtistService) Query(ctx context.Context, offset, limit int) (list []models.Artist, err error) {
	dbx := app.DBFromContext(ctx)
	err = dbx.SelectContext(ctx, &list, "select id, name FROM artist;")
	if err == sql.ErrNoRows {
		return make([]models.Artist, 0), nil
	}
	return list, err
}
