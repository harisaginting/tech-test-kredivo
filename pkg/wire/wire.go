//go:build wireinject
// +build wireinject

package wire

import (
	googleWire "github.com/google/wire"
	mAuth "github.com/harisaginting/tech-test-kredivo/api/v1/auth"
	mUser "github.com/harisaginting/tech-test-kredivo/api/v1/user"
	mReport "github.com/harisaginting/tech-test-kredivo/api/v1/report"
	mWatchlist "github.com/harisaginting/tech-test-kredivo/api/v1/watchlist"
	"gorm.io/gorm"
)

func ApiUser(db *gorm.DB) mUser.Controller {
	googleWire.Build(
		mUser.ProviderController,
		mUser.ProviderService,
		mUser.ProviderRepository,
	)
	return mUser.Controller{}
}

func ApiAuth(db *gorm.DB) mAuth.Controller {
	googleWire.Build(
		mAuth.ProviderController,
		mAuth.ProviderService,
		mAuth.ProviderRepository,
	)
	return mAuth.Controller{}
}

func ApiWatchlist(db *gorm.DB) mWatchlist.Controller {
	googleWire.Build(
		mWatchlist.ProviderController,
		mWatchlist.ProviderService,
		mWatchlist.ProviderRepository,
	)
	return mWatchlist.Controller{}
}

func ApiReport(db *gorm.DB) mReport.Controller {
	googleWire.Build(
		mReport.ProviderController,
		mReport.ProviderService,
		mReport.ProviderRepository,
	)
	return mReport.Controller{}
}
