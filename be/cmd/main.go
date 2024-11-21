package main

import (
	gql "github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	"github.com/samber/lo"
	. "github.com/seprich/keycape/internal/config"
	"github.com/seprich/keycape/internal/graph"
	"github.com/seprich/keycape/internal/logger"
	"github.com/seprich/keycape/internal/routes"
)

func main() {
	r := gin.New()
	r.Use(gzip.Gzip(gzip.DefaultCompression))
	r.Use(logger.GinLogger)
	r.Use(logger.GinRecovery)

	gqlServer := gql.New(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}}))
	gqlServer.AddTransport(transport.POST{})
	r.POST("/graphql", gin.WrapH(gqlServer))

	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"result": "ok",
		})
	})

	routes.RegisterAPIRoutes(r.Group("/api"))
	lo.Must0(r.Run(":" + Config.String("server.port")))
}
