package router

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/GenkiHirano/gqlgen-tutorial/graph"
	"github.com/GenkiHirano/gqlgen-tutorial/graph/generated"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Setup(engine *gin.Engine) {
	engine.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"GET", "OST", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders: []string{"Origin", "Content-Type", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization"},
		MaxAge:       12 * 60 * 60,
	}))

	engine.POST("/query", gqlHandler())
}

func gqlHandler() gin.HandlerFunc {
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))
	return func(c *gin.Context) {
		srv.ServeHTTP(c.Writer, c.Request)
	}
}
