package apis

import (
	"testing"

	"github.com/nicksrandall/restful-starter-kit/services"
	"github.com/nicksrandall/restful-starter-kit/testdata"
)

func TestArtist(t *testing.T) {
	testdata.ResetDB()
	router := newRouter()
	ServeArtistResource(router, &services.ArtistService{})

	notFoundError := `{"error_code":"NOT_FOUND", "message":"NOT_FOUND"}`
	nameRequiredError := `{"error_code":"INVALID_DATA","message":"INVALID_DATA","details":[{"field":"name","error":"cannot be blank"}]}`
	_ = notFoundError
	_ = nameRequiredError

	runAPITests(t, router, []apiTestCase{})
}
