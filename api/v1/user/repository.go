package user

import (
	"context"
	"gorm.io/gorm"
	"github.com/harisaginting/tech-test-kredivo/db/table"
	"github.com/harisaginting/tech-test-kredivo/pkg/utils/helper"
	"github.com/harisaginting/tech-test-kredivo/pkg/log"
)


type Repository struct {
	db *gorm.DB
}

func ProviderRepository(db *gorm.DB) Repository {
	return Repository{
		db: db,
	}
}

func (repo *Repository) FindAll(ctx context.Context) (users []User) {
	var table table.User 
	qx 		:= repo.db
	qx.Find(&table)
	if qx.Error != nil {
		log.Error(ctx, qx.Error, "FindAllByCustomer: ")
	}
	log.Info(ctx, "Repo : ",table)
	helper.AdjustStructToStruct(table,&users)
	return
}