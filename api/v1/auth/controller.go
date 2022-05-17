package auth

import (
	 "github.com/gin-gonic/gin"
	 "github.com/harisaginting/tech-test-kredivo/pkg/http/response"
	 "github.com/harisaginting/tech-test-kredivo/pkg/tracer"
	 "github.com/harisaginting/tech-test-kredivo/pkg/log"
)

type Controller struct {
	service Service
}

func ProviderController(s Service) Controller {
	return Controller{
		service: s,
	}
}

func (ctrl *Controller) Register(c *gin.Context) {
	ctx  := c.Request.Context()
	span := tracer.Span(ctx, "ListUserController")
	defer span.End()
	log.Info(ctx,"Controller User")

	var resData ResponseList
	ctrl.service.List(ctx, &resData)
	
	// return
	response.Json(c,resData)
	return
}


func (ctrl *Controller) List(c *gin.Context) {
	ctx  := c.Request.Context()
	span := tracer.Span(ctx, "ListUserController")
	defer span.End()
	log.Info(ctx,"Controller User")

	var resData ResponseList
	ctrl.service.List(ctx, &resData)
	
	// return
	response.Json(c,resData)
	return
}