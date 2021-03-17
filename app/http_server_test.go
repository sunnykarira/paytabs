package app

import (
	"errors"
	"fmt"
	"math/rand"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
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
			method:       "GET",
			serverMethod: server.GET,
		},
		{
			method:       "POST",
			serverMethod: server.POST,
		},
		{
			method:       "PUT",
			serverMethod: server.PUT,
		},
		{
			method:       "OPTIONS",
			serverMethod: server.OPTIONS,
		},
		{
			method:       "PATCH",
			serverMethod: server.PATCH,
		},
		{
			method:       "HEAD",
			serverMethod: server.HEAD,
		},
		{
			method:       "DELETE",
			serverMethod: server.DELETE,
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

func TestRouteHandlerForResponseHeaders(t *testing.T) {
	app := App{}
	server := NewHttpServer(app)

	key := "Test"
	value := fmt.Sprintf("VALUE_%d", rand.Int())
	successBody := fmt.Sprintf("Success_%d", rand.Int())
	successHandler := func(ctx *Context, r *http.Request, p HttpParams) (HttpResponseBody, error) {
		ctx.ResponseHeaders[key] = value
		return []byte(successBody), nil
	}

	server.GET("/test", successHandler)

	req, _ := http.NewRequest("GET", "/test", nil)
	rr := httptest.NewRecorder()
	server.router.ServeHTTP(rr, req)

	if rr.Header().Get(key) != value {
		t.Errorf("Header not set for http server")
	}

}

func TestRouteHandlerForMiddlewares(t *testing.T) {
	app := App{}
	server := NewHttpServer(app)

	key := "Test"
	value := fmt.Sprintf("VALUE_%d", rand.Int())

	server.AddMiddleware(
		func(m []HttpMiddleware, ctx *Context, r *http.Request, p HttpParams) (HttpResponseBody, error) {
			ctx.ResponseHeaders[key] = value
			return m[0](m[1:], ctx, r, p)
		},
	)

	successBody := fmt.Sprintf("Success_%d", rand.Int())
	var handlerName string
	successHandler := func(ctx *Context, r *http.Request, p HttpParams) (HttpResponseBody, error) {
		handlerName = ctx.GetHandlerName()
		return []byte(successBody), nil
	}

	server.GET("/test", successHandler)

	req, _ := http.NewRequest("GET", "/test", nil)
	rr := httptest.NewRecorder()
	server.router.ServeHTTP(rr, req)

	assert.NotNil(t, handlerName)

	if rr.Header().Get(key) != value {
		t.Errorf("Header not set for http server; means middleware is not called")
	}

}

func TestServeHTTP(t *testing.T) {
	app := App{}
	server := NewHttpServer(app)

	successBody := fmt.Sprintf("Success_%d", rand.Int())
	successHandler := func(ctx *Context, r *http.Request, p HttpParams) (HttpResponseBody, error) {
		return []byte(successBody), nil
	}

	server.GET("/test", successHandler)

	req, _ := http.NewRequest("GET", "/test", nil)
	rr := httptest.NewRecorder()
	server.ServeHTTP(rr, req)

	if rr.Body.String() != successBody {
		t.Errorf("Serve http does not get right body")
	}

}
