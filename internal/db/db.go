package db

import (
	"fmt"
	"hoxt/data"
	"hoxt/internal/modules"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func InitDataBase() {
	defer fmt.Println("[DATABASE: RELOADED]")
	if data.Configs.DBFilename == "" {
		data.Configs.DBFilename = "data.db"
	}
	var err error
	DB, err = gorm.Open(sqlite.Open(data.Configs.DBFilename), &gorm.Config{
		TranslateError: true, // Позволяет GORM возвращать понятные ошибки
		Logger:         logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		panic("failed to connect database")
	}

	DB.Exec("PRAGMA foreign_keys = ON")
	err = DB.AutoMigrate(&modules.Paste{})
	if err != nil {
		panic(err.Error())
	}

	cfg, err := data.LoadConfig("./data/config.json")
	if err != nil {
		log.Fatalln(err)
	}

	if len(cfg.Pastes) > 0 {
		for _, p := range cfg.Pastes {
			var paste modules.Paste

			result := DB.Where("title = ?", p.Title).First(&paste)

			if result.Error != nil {
				DB.Create(&modules.Paste{
					Title:    p.Title,
					Content:  p.Content,
					Author:   "WEBSITE SYSTEM",
					IsTitled: p.IsTitled,
				})
			} else {
				paste.Content = p.Content
				DB.Save(&paste)
			}
		}
	}

}
