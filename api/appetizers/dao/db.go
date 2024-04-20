package dao

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"alan/restaurant/appetizers/shared/domain"
)

var (
	dsn = fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)
)

// IAppetizerStore is interface of appetizer store
type IAppetizerStore interface {
	Create(c *gin.Context, appetizer domain.Appetizer) (*domain.Appetizer, error)
	ListAll(c *gin.Context) ([]domain.Appetizer, error)
	GetById(c *gin.Context, id uint) (*domain.Appetizer, error)
	PatchById(c *gin.Context, id uint, appetizer domain.Appetizer) (*domain.Appetizer, error)
	DeleteById(c *gin.Context, id uint) error
}

// NewAppetizerStore initiates db connection
func NewAppetizerStore(log *logrus.Entry) IAppetizerStore {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: dbLogger(),
	})
	if err != nil {
		log.Panicf("Got error when connect to database, the error is '%v'\n", err)
	}

	return &AppetizerStore{
		DB:  db,
		log: log,
	}
}

func dbLogger() logger.Interface {
	return logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second,   // Slow SQL threshold
			LogLevel:                  logger.Silent, // Log level
			IgnoreRecordNotFoundError: false,         // Ignore ErrRecordNotFound error for logger
			Colorful:                  false,         // Disable color
		},
	)
}
