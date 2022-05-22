package middleware

import (
	"github.com/harisaginting/tech-test-kredivo/pkg/cache"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	// "time"
	// "github.com/harisaginting/tech-test-kredivo/pkg/utils/helper"
	"github.com/harisaginting/tech-test-kredivo/pkg/log"
)

type middleware struct {
	env int
}

func Start(env int) middleware { 
	return middleware{env: env}
}

func (middleware *middleware) MustMember() gin.HandlerFunc {
	return func(context *gin.Context) {
		// RETURN IF ENV KEYCLOAK NOT 1 or true
		if middleware.env != 1 {
			context.Next()
			return
		}

		// GET TOKEN
		s := strings.SplitN(context.Request.Header.Get("Authorization"), " ", 2)
		if len(s) != 2 {
			msg := "Authorization token is not found"
			middleware.abort(http.StatusUnauthorized, context, msg)
			return
		}		
		var token = s[1]

		cacheKey 	 		:= cache.CreateCacheKey("auth:"+token)
		username, err		:= cache.GetKey(cacheKey)
		if err != nil || username == "" {
			log.Error(context, err)
			msg := "Unauthorized token"
			middleware.abort(http.StatusUnauthorized, context, msg)
			return
		}else{
			log.Info(context, username)
		}
		username = strings.ReplaceAll(username, "\"", "")
		context.Set("username", username)
		context.Next()
	}
}

func (middleware *middleware) abort(status int, context *gin.Context, message interface{}) {
	context.AbortWithStatusJSON(status, gin.H{
		"status":        status,
		"error_message": message,
		"data":          nil,
	})
}