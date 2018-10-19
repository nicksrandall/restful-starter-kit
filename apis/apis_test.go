package apis

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi"
	"github.com/nicksrandall/restful-starter-kit/app"
	"github.com/nicksrandall/restful-starter-kit/testdata"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

type apiTestCase struct {
	tag      string
	method   string
	url      string
	body     string
	status   int
	response string
}

func newRouter() chi.Router {
	logger := zap.NewNop()

	router := chi.NewRouter()
	router.Use(
		app.LoggerMiddleware(logger),
		app.DBMiddleware(testdata.DB),
	)

	return router
}

func testAPI(router chi.Router, method, URL, body string) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(method, URL, bytes.NewBufferString(body))
	req.Header.Set("Content-Type", "application/json")
	res := httptest.NewRecorder()
	router.ServeHTTP(res, req)
	return res
}

func runAPITests(t *testing.T, router chi.Router, tests []apiTestCase) {
	for _, test := range tests {
		res := testAPI(router, test.method, test.url, test.body)
		assert.Equal(t, test.status, res.Code, test.tag)
		if test.response != "" {
			assert.JSONEq(t, test.response, res.Body.String(), test.tag)
		}
	}
}
