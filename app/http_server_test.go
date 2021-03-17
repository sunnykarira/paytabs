package app

import (
	"errors"
	"fmt"
	"math/rand"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestNewHttpServer(t *testing.T) {
	app := App{
		externals: map[string]interface{}{
			"rand": rand.Int(),
		},
	}
	server := NewHttpServer(app)
	newApp := server.App
	if app.externals["rand"].(int) != newApp.externals["rand"].(int) {
		t.Error("Invalid app found in http server")
	}

	if server.router == nil {
		t.Error("No router present in server")
	}
}

func TestRouteHandler(t *testing.T) {
	app := App{}
	server := NewHttpServer(app)
	successBody := fmt.Sprintf("Success_%d", rand.Int())
	successHandler := func(*Context, *http.Request, HttpParams) (HttpResponseBody, error) {
		return []byte(successBody), nil
	}

	errorCode := fmt.Sprintf("CODE_%d", rand.Int())
	errorMessage := fmt.Sprintf("MESSAGE_%d", rand.Int())
	errorHandler := func(*Context, *http.Request, HttpParams) (HttpResponseBody, error) {
		return nil, BadError(errorCode, errorMessage)
	}

	internalErrorMessage := fmt.Sprintf("ERROR_%d", rand.Int())
	internalErrorHandler := func(*Context, *http.Request, HttpParams) (HttpResponseBody, error) {
		return nil, errors.New(internalErrorMessage)
	}

	tests := []struct {
		method       string
		serverMethod func(string, HttpHandler)
	}{
		{
			method:       "POST",
			serverMethod: server.POST,
		},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("%s Success", tt.method), func(t *testing.T) {
			path := fmt.Sprintf("/test-%s", tt.method)
			tt.serverMethod(path, successHandler)
			req, _ := http.NewRequest(tt.method, path, nil)
			rr := httptest.NewRecorder()
			server.router.ServeHTTP(rr, req)

			if status := rr.Code; status != http.StatusOK {
				t.Errorf("Wrong status")
			}

			if rr.Body.String() != successBody {
				t.Errorf("Wrong Body")
			}
		})

		t.Run(fmt.Sprintf("%s App Error", tt.method), func(t *testing.T) {
			path := fmt.Sprintf("/test-error-%s", tt.method)
			tt.serverMethod(path, errorHandler)
			req, _ := http.NewRequest(tt.method, path, nil)
			rr := httptest.NewRecorder()
			server.router.ServeHTTP(rr, req)

			if status := rr.Code; status != http.StatusBadRequest {
				t.Errorf("Wrong status for app method")
			}
		})

		t.Run(fmt.Sprintf("%s internal Error", tt.method), func(t *testing.T) {
			path := fmt.Sprintf("/test-internal-error-%s", tt.method)
			tt.serverMethod(path, internalErrorHandler)
			req, _ := http.NewRequest(tt.method, path, nil)
			rr := httptest.NewRecorder()
			server.router.ServeHTTP(rr, req)

			if status := rr.Code; status != http.StatusInternalServerError {
				t.Errorf("Wrong status for internal error")
			}
		})
	}

}


func TestServeHTTP(t *testing.T) {
	app := App{}
	server := NewHttpServer(app)

	successBody := fmt.Sprintf("Success_%d", rand.Int())
	successHandler := func(ctx *Context, r *http.Request, p HttpParams) (HttpResponseBody, error) {
		return []byte(successBody), nil
	}

	server.POST("/test", successHandler)

	req, _ := http.NewRequest("POST", "/test", nil)
	rr := httptest.NewRecorder()
	server.ServeHTTP(rr, req)

	if rr.Body.String() != successBody {
		t.Errorf("Serve http does not get right body")
	}

}
