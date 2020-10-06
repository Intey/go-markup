// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	"go-markup/domain"
	"go-markup/models"
	"go-markup/restapi/worker"
)

var marks = make(map[int64]domain.Mark)
var lastID int64 = 1

//go:generate swagger generate server --target ../../go-markup --name Worker --spec ../swagger/worker.yaml --api-package worker --principal interface{}

func configureFlags(api *worker.WorkerAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *worker.WorkerAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	api.UseSwaggerUI()
	// To continue using redoc as your UI, uncomment the following line
	// api.UseRedoc()

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	api.GetHealthHandler = worker.GetHealthHandlerFunc(func(params worker.GetHealthParams) middleware.Responder {
		return middleware.NotImplemented("operation worker.GetHealth has not yet been implemented")
	})

	api.PostMarksHandler = worker.PostMarksHandlerFunc(func(params worker.PostMarksParams) middleware.Responder {
		lastID++
		marks[lastID] = domain.Mark{
			ID:       lastID,
			Position: params.Mark.Position,
			Entity:   params.Mark.Entity}

		return worker.NewPostMarksCreated().WithPayload(&models.MarkOut{ID: lastID})
	})
	api.GetMarksHandler = worker.GetMarksHandlerFunc(func(params worker.GetMarksParams) middleware.Responder {
		payload := []*models.MarkOut{}

		for _, m := range marks {
			payload = append(payload, &models.MarkOut{ID: m.ID})
		}
		return worker.NewGetMarksOK().WithPayload(payload)
	})

	api.PreServerShutdown = func() {}

	api.ServerShutdown = func() {}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix"
func configureServer(s *http.Server, scheme, addr string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return handler
}
