package api

import (
	"github.com/harisaginting/tech-test-kredivo/pkg/middleware"	
	"github.com/harisaginting/tech-test-kredivo/pkg/wire"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RestV1(r *gin.RouterGroup, db *gorm.DB) {
	// Dependency injection
	apiAuth := wire.ApiAuth(db)
	apiUser := wire.ApiUser(db)

	member := middleware.Start(1)
	// group rest
	rest := r.Group("rest")
	{
		// group v1
		v1 := rest.Group("v1")
		{
			// auth
			apiAuthGroup := v1.Group("auth")
			{
				apiAuthGroup.POST("/register", apiAuth.Register)
				apiAuthGroup.POST("/login", apiAuth.Login)
				apiAuthGroup.GET("/me", member.MustMember(), apiAuth.Me)
			}
			// user
			apiUserGroup := v1.Group("user")
			{
				apiUserGroup.GET("/", apiUser.List)
			}
		}
	}

	return
}