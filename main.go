package main

import (
	"api-messages/config"
	"api-messages/controllers"
	"fmt"
	"net/http"
	"time"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/getsentry/sentry-go"
	sentrygin "github.com/getsentry/sentry-go/gin"
	"github.com/gin-gonic/gin"
)

func main() {

	// To initialize Sentry's handler, you need to initialize Sentry itself beforehand
	if err := sentry.Init(sentry.ClientOptions{
		Dsn:              config.Provider().Sentry.Dns,
		AttachStacktrace: true,
	}); err != nil {
		fmt.Printf("Sentry initialization failed: %v\n", err)
	}

	r := gin.Default()

	r.Use(sentrygin.New(sentrygin.Options{}))
	r.Use(Cors)

	authMiddleware, _ := jwt.New(&jwt.GinJWTMiddleware{
		Realm:       "<realm>",
		Key:         []byte("<private key>"),
		Timeout:     24,
		MaxRefresh:  time.Hour,
		IdentityKey: "currentClientIdentity",
		IdentityHandler: func(c *gin.Context) interface{} {
			claims := jwt.ExtractClaims(c)
			return claims["id"].(string)
		},
		Unauthorized: func(c *gin.Context, code int, message string) {
			c.JSON(code, gin.H{
				"code":    code,
				"message": message,
			})
		},
		TokenLookup:   "header: Authorization",
		TokenHeadName: "Bearer",
		TimeFunc:      time.Now,
	})

	r.GET("/health-check", GetHealthStatus)

	r.Use(authMiddleware.MiddlewareFunc())
	{
		g := r.Group("/api")
		{
			g.GET("/messages", controllers.GetMessages)
			g.POST("/messages", controllers.SaveMessage)
		}
	}

	r.Run(":5009")
}

func GetHealthStatus(c *gin.Context) {
	c.JSON(200, gin.H{
		"status": "Message Service",
	})
}

func Cors(c *gin.Context) {

	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", "*")
	c.Header("Access-Control-Allow-Headers", "*")
	c.Header("Content-Type", "application/json")

	if c.Request.Method != "OPTIONS" {

		c.Next()
	} else {

		// Everytime we receive an OPTIONS request,
		// we just return an HTTP 200 Status Code
		// Like this, Angular can now do the real
		// request using any other method than OPTIONS
		c.AbortWithStatus(http.StatusOK)
	}
}
