package auth

import (
	"fmt"
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
	var user table.User 
	qx 		:= repo.db
	qx.Find(&user)
	if qx.Error != nil {
		log.Error(ctx, qx.Error, "FindAllByCustomer: ")
	}
	log.Info(ctx, "Repo : ",user)
	helper.AdjustStructToStruct(user,&users)
	return
}

func (repo *Repository) FindByUsername(ctx context.Context, username string) (user table.User, err error) {
	log.Info(ctx, "repo u : "+username)
	qx := repo.db.Where("username = ?", username).First(&user)
	err = qx.Error
	return
}



func (repo *Repository) Register(ctx context.Context, p PayloadUserRegister) (err error) {
	var table table.User
	helper.AdjustStructToStruct(p,&table)
	user := repo.db.Save(&table)
	err = user.Error
	if err != nil { return }
	log.Info(ctx, fmt.Sprintf("Created : %s",table.ID))

	return
}