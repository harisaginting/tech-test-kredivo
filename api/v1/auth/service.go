package auth

import (
	"errors"
	"context"
	"encoding/json"
	"github.com/google/uuid"
	"github.com/harisaginting/tech-test-kredivo/pkg/tracer"
	"github.com/harisaginting/tech-test-kredivo/pkg/cache"
 	"github.com/harisaginting/tech-test-kredivo/pkg/utils/helper"
)

type Service struct {
	repo Repository
}

func ProviderService(r Repository) Service {
	return Service{
		repo: r,
	}
}

func (service *Service) Register(ctx context.Context, p PayloadUserRegister)(err error, res ResponseUserRegister) {
	trace := tracer.Span(ctx,"Register")
	defer trace.End()

	check, _ := service.repo.FindByUsername(ctx, p.Username)
	if check.ID != 0 {
		err  = errors.New("username already used")
		return
	}

	p.Password, err =  helper.HashPassword(p.Password)
	if err != nil { return }

	err = service.repo.Register(ctx, p)
	if err != nil { return }


	token := uuid.NewString()
	cacheKey 	 	:= cache.CreateCacheKey("auth:"+token)
	cacheData, err  := json.Marshal(p.Username)
	if err != nil { return }
	err = cache.SetKeyWithExpired(cacheKey, cacheData,"120m")
	if err != nil { return }
	res.Token = token
	tracer.SetAttributeInt(trace,"register User",p.Username)
	return
}

func (service *Service) Login(ctx context.Context, p PayloadUserLogin)(err error, res ResponseUserLogin) {
	trace := tracer.Span(ctx,"Login")
	defer trace.End()

	check, _ := service.repo.FindByUsername(ctx, p.Username)
	if check.ID == 0 {
		err  = errors.New("invalid Username")
		return
	}

	valid :=  helper.CheckPasswordHash(p.Password, check.Password)
	if !valid {
		err  = errors.New("invalid Password")
		return
	}

	helper.AdjustStructToStruct(check,&res)
	token 			:= uuid.NewString()
	cacheKey 	 	:= cache.CreateCacheKey("auth:"+token)
	cacheData, err  := json.Marshal(p.Username)
	if err != nil { return }
	err = cache.SetKeyWithExpired(cacheKey, cacheData,"120m")
	if err != nil { return }
	res.Token = token
	tracer.SetAttributeInt(trace,"Login",p.Username)
	return
}

func (service *Service) GetByUsername(ctx context.Context, p string)(err error, res ResponseMe) {
	trace := tracer.Span(ctx,"FindByUsername")
	defer trace.End()

	user, err := service.repo.FindByUsername(ctx, p)
	if err != nil { return }
	if user.ID == 0 {
		err  = errors.New("username not found")
		return
	}
	helper.AdjustStructToStruct(user,&res)
	tracer.SetAttributeInt(trace,"Find User",p)
	return
}

func (service *Service) List(ctx context.Context, res *ResponseList) {
	trace := tracer.Span(ctx,"ListUser")
	defer trace.End()
	users := service.repo.FindAll(ctx)
	res.Items = users
	res.Total = len(users)

	tracer.SetAttributeInt(trace,"total User",res.Total)
	return
}