package config

import (
	"errors"
	"fmt"
	"os"
	"sftp/util"
	"strconv"
	"time"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const (
	dbConnectRetryWaitTime = 3 * time.Second
)

var (
	logger = util.NewLogger()
)

type DatabaseConfig struct {
	DBHost            string
	DBPort            int
	DBName            string
	DBUser            string
	DBPassword        string
	DBIdleConnection  int
	DBMaxConnection   int
	DBConnMaxLifeTime time.Duration
	DBTimeZone        string
}

// ConnectDatabaseWithGorm is connect database with gorm.
// And returns gorm.DB instance.
func (d DatabaseConfig) ConnectDatabaseWithGorm(retry int) (*gorm.DB, error) {
	if retry <= 0 {
		return nil, errors.New("connect retry over")
	}
	var url string
	if len(d.DBPassword) != 0 {
		url = fmt.Sprintf("%v:%v@tcp(%v:%v)/%v", d.DBUser, d.DBPassword, d.DBHost, d.DBPort, d.DBName)
	} else {
		url = fmt.Sprintf("%v@tcp(%v:%v)/%v", d.DBUser, d.DBHost, d.DBPort, d.DBName)
	}
	url += fmt.Sprintf("?charset=utf8&parseTime=True&loc=%s", d.DBTimeZone)

	db, gormErr := gorm.Open(mysql.Open(url), &gorm.Config{})
	if gormErr != nil {
		time.Sleep(dbConnectRetryWaitTime)
		logger.Warnf("failed to connect database, error: %v connect retry...", gormErr)
		return d.ConnectDatabaseWithGorm(retry - 1)
	}

	sqlDB, sqlErr := db.DB()
	if sqlErr != nil {
		logger.Errorf("failed to get sqldb, error: %v", sqlErr)
		return nil, sqlErr
	}
	sqlDB.SetMaxOpenConns(d.DBMaxConnection)
	sqlDB.SetMaxIdleConns(d.DBIdleConnection)
	sqlDB.SetConnMaxLifetime(d.DBConnMaxLifeTime)
	return db, nil
}

// GetDatabaseConfig is get DatabaseConfig from env.
func GetDatabaseConfig() DatabaseConfig {

	err := godotenv.Load(".env")

	if err != nil {
		logger.Errorf("Error loading environment: %v", err)
	}

	port, _ := strconv.Atoi(os.Getenv("DBPort"))

	config := DatabaseConfig{
		DBHost:            os.Getenv("DBHost"),
		DBPort:            port,
		DBName:            os.Getenv("DBName"),
		DBUser:            os.Getenv("DBUser"),
		DBPassword:        os.Getenv("DBPassword"),
		DBIdleConnection:  100,
		DBMaxConnection:   100,
		DBConnMaxLifeTime: time.Duration(100) * time.Second,
		DBTimeZone:        os.Getenv("DBTimeZone"),
	}

	return config
}
