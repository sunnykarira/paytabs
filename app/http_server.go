package app

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type HttpParams interface {
	ByName(string) string
}

type HttpResponseBody []byte

type HttpHandler func(*Context, *http.Request, HttpParams) (HttpResponseBody, error)


type HttpServer struct {
	router      *httprouter.Router
	App         App
}

func routerHandler(path string, httpHandler HttpHandler, server *HttpServer) httprouter.Handle {
	routeHandler := func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		ctx := NewContext(server.App)
		ctx.AddHandlerName(path)
		responseBody, err := httpHandler(ctx, r, p)

		// Set response header
		for k := range ctx.ResponseHeaders {
			w.Header().Set(k, ctx.ResponseHeaders[k])
		}
		w.Header().Set("Content-type", "application/json")

		if err != nil {
			errorBody, httpCode := ErrorHandler(err)
			http.Error(w, string(errorBody), httpCode)
			return
		}
		w.Write(responseBody)
	}
	return routeHandler
}

func NewHttpServer(app App) *HttpServer {
	router := httprouter.New()
	return &HttpServer{
		router: router,
		App:    app,
	}
}

func (s *HttpServer) POST(path string, httpHandler HttpHandler) {
	routeHandler := routerHandler(path, httpHandler, s)
	s.router.POST(path, routeHandler)
}

func (s *HttpServer) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	s.router.ServeHTTP(w, req)
}

/////////////////////////App Middlewares//////////////////////////////////////

func ErrorHandler(err error) (HttpResponseBody, int) {

	appErr, ok := err.(*Error)

	if !ok {
		appErr = InternalServerError(err.Error(),
			"Something went wrong, please try after sometime")
	}

	responseBody := "{\"code\":\"" + appErr.CodeErr() + "\",\"message\":\"" +
		appErr.Error() + "\"}"
	return []byte(responseBody), appErr.HttpCodeErr()
}

/////////////////////////////////////////////////////////////////////////////
