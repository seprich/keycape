package main

import (
	gql "github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	"github.com/samber/lo"
	"github.com/seprich/go-future/async"
	. "github.com/seprich/keycape/internal/config"
	"github.com/seprich/keycape/internal/db_schema"
	"github.com/seprich/keycape/internal/graph"
	"github.com/seprich/keycape/internal/logger"
	"github.com/seprich/keycape/internal/routes"
	"net/http"
)

// Tracks readiness

func main() {
	migrationsFuture := async.NewFuture(db_schema.RunMigrations)

	r := gin.New()
	r.Use(gzip.Gzip(gzip.DefaultCompression))
	r.Use(logger.GinLogger)
	r.Use(logger.GinRecovery)

	// GraphQL
	gqlServer := gql.New(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}}))
	gqlServer.AddTransport(transport.POST{})
	r.POST("/graphql", gin.WrapH(gqlServer))

	// Public endpoints
	r.GET("/health", func(c *gin.Context) {
		if migrationsFuture.HasResult() {
			c.JSON(http.StatusOK, gin.H{"status": "READY"})
		} else {
			c.JSON(http.StatusOK, gin.H{"status": "UP"})
		}
	})
	r.GET("/health/liveness", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "UP"})
	})
	r.GET("/health/readiness", func(c *gin.Context) {
		if migrationsFuture.HasResult() {
			c.JSON(http.StatusOK, gin.H{"status": "READY"})
		} else {
			c.JSON(http.StatusServiceUnavailable, gin.H{"status": "NOT READY"})
		}
	})

	routes.RegisterAPIRoutes(r.Group("/api"))

	lo.Must0(r.Run(":" + Config.String("server.port")))
}
