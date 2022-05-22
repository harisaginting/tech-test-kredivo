package auth

import (
	"encoding/json"
	"io/ioutil"
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

	var reqData PayloadUserRegister
	var resData ResponseUserRegister
	request, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		response.BadRequest(c)
		return
	}
	json.Unmarshal(request, &reqData)
	err, resData = ctrl.service.Register(ctx, reqData)
	if err != nil {
		response.Error(c, err.Error())
		return
	}
	response.Json(c,resData)
	return
}

func (ctrl *Controller) Login(c *gin.Context) {
	ctx  := c.Request.Context()
	span := tracer.Span(ctx, "ListUserController")
	defer span.End()
	log.Info(ctx,"Controller User")

	var reqData PayloadUserLogin
	var resData ResponseUserLogin
	request, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		response.BadRequest(c)
		return
	}
	json.Unmarshal(request, &reqData)
	err, resData = ctrl.service.Login(ctx, reqData)
	if err != nil {
		response.Error(c, err.Error())
		return
	}
	response.Json(c,resData)
	return
}

func (ctrl *Controller) Me(c *gin.Context) {
	ctx  := c.Request.Context()
	span := tracer.Span(ctx, "ListUserController")
	defer span.End()
	log.Info(ctx,"Controller User")

	var resData ResponseMe
	username := c.Value("username").(string)
	err, resData := ctrl.service.GetByUsername(ctx, username)
	if err != nil {
		response.Error(c, err.Error())
		return
	}
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