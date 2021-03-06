// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"log"
	"net/http"

	errors "github.com/go-openapi/errors"
	runtime "github.com/go-openapi/runtime"

    "github.com/dre1080/recover"

	"github.com/gangozero/hackprague2019-be/components/server"
	"github.com/gangozero/hackprague2019-be/generated/restapi/operations"
	"github.com/gangozero/hackprague2019-be/generated/restapi/operations/data"
	"github.com/gangozero/hackprague2019-be/generated/restapi/operations/profile"
)

//go:generate swagger generate server --target ../../generated --name Hackprague2019Be --spec ../../openapi/swagger.yaml

func configureFlags(api *operations.Hackprague2019BeAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.Hackprague2019BeAPI) http.Handler {
	srv := server.NewServer()
	log.Printf("Server started")

	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	api.DataGetGradeListHandler = data.GetGradeListHandlerFunc(srv.GetGradesByID)
	api.DataGetUserGradeListHandler = data.GetUserGradeListHandlerFunc(srv.GetGradesByIDAndUser)
	api.DataPostNewGradeHandler = data.PostNewGradeHandlerFunc(srv.PostGrade)
	api.ProfileGetProfileHandler = profile.GetProfileHandlerFunc(srv.GetProfiles)

	api.ServerShutdown = func() {
		srv.Stop()
	}

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
	recovery := recover.New(&recover.Options{
		Log: log.Print,
	})
	return recovery(handler)
}
