// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package wire

import (
	"github.com/harisaginting/tech-test-kredivo/api/v1/auth"
	"github.com/harisaginting/tech-test-kredivo/api/v1/report"
	"github.com/harisaginting/tech-test-kredivo/api/v1/user"
	"github.com/harisaginting/tech-test-kredivo/api/v1/watchlist"
	"gorm.io/gorm"
)

// Injectors from wire.go:

func ApiUser(db *gorm.DB) user.Controller {
	repository := user.ProviderRepository(db)
	service := user.ProviderService(repository)
	controller := user.ProviderController(service)
	return controller
}

func ApiAuth(db *gorm.DB) auth.Controller {
	repository := auth.ProviderRepository(db)
	service := auth.ProviderService(repository)
	controller := auth.ProviderController(service)
	return controller
}

func ApiWatchlist(db *gorm.DB) watchlist.Controller {
	repository := watchlist.ProviderRepository(db)
	service := watchlist.ProviderService(repository)
	controller := watchlist.ProviderController(service)
	return controller
}

func ApiReport(db *gorm.DB) report.Controller {
	repository := report.ProviderRepository(db)
	service := report.ProviderService(repository)
	controller := report.ProviderController(service)
	return controller
}