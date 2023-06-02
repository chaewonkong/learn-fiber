package database

import (
	"log"

	"github.com/chaewonkong/learn-fiber/entity"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type (
	// Database 데이터베이스 객체
	Database interface {
		Migrate()
	}
	database struct {
		*gorm.DB
	}
)

// NewDatabase 데이터베이스 생성자
func NewDatabase() Database {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	return &database{db}
}

// Migrate entity를 추가해 자동으로 마이그레이션 진행
// entity가 추가될 때마다 아래에 코드 수정해 migration 관리
func (db *database) Migrate() {
	err := db.AutoMigrate(&entity.User{})
	if err != nil {
		log.Fatal(err)
	}
}
