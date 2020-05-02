package middleware

import (
	"github.com/iris-contrib/middleware/cors"
	"github.com/kataras/iris/v12/context"
)

// CORS CORS Handler
var CORS context.Handler

func initCORS() {
	CORS = cors.New(cors.Options{
		AllowedOrigins:   []string{"*"}, // Allows everything, use that to change the hosts.
		AllowCredentials: true,
		AllowedHeaders:   []string{"*"}, // If do not set this will not be able to customize the header.
	})
}