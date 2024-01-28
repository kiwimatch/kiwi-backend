package middlewares

import (
	"fmt"
	"os"

	"github.com/1liale/maze-backend/models"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func PropDBEnv(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Set("db", db)
		ctx.Next()
	}
}

func InitDB() *gorm.DB {
	user := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	db_name := os.Getenv("POSTGRES_DB")
	host := os.Getenv("POSTGRES_HOST")
	port := os.Getenv("POSTGRES_PORT")

	dsn := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s", user, password, host, port, db_name)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		logrus.Warn("Cannot connect to: ", dsn)
		logrus.Info("Proceeding without db connection!")
	} else {
		logrus.Info("Successfully connected to: ", dsn)
		db.AutoMigrate(&models.User{}, &models.MazeRecord{})
	}

	return db
}
