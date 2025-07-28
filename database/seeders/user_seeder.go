package seeders

import (
	"user-service/constants"
	"user-service/domain/models"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func RunUserSeeder(db *gorm.DB) {
	password, _ := bcrypt.GenerateFromPassword([]byte("admin123"), bcrypt.DefaultCost)
	user := models.User{
		UUID:     uuid.New(),
		Name:     "Administrator",
		Username: "admin",
		Password: string(password),
		Email:    "admin@gmail.com",
		Phone:    "081234567890",
		RoleID:   constants.Admin,
	}

	err := db.FirstOrCreate(&user, models.User{Username: user.Username}).Error
	if err != nil {
		logrus.Errorf("Failed to seed user: %v", err)
		panic(err)
	}
	logrus.Infof("User %s seeded successfully", user.Username)
}
