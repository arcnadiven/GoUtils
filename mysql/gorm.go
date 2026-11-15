package mysql

import (
	"fmt"

	"github.com/pkg/errors"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type GormParams struct {
	DbUser     string
	DbPassword string
	DbHost     string
	DbPort     string
	DbName     string
	Config     *gorm.Config
}

func NewGormConn(params *GormParams) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		params.DbUser, params.DbPassword, params.DbHost, params.DbPort, params.DbName)
	gormConfig := &gorm.Config{}
	if params.Config != nil {
		gormConfig = params.Config
	}

	conn, err := gorm.Open(mysql.Open(dsn), gormConfig)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return conn, nil
}
