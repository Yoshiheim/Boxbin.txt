package helpers

import (
	"hoxt/internal/modules"

	"gorm.io/gorm"
)

func CreatePasteIfTopicExists(db *gorm.DB, paste modules.Paste) (modules.Paste, error) {

	if err := db.Transaction(func(tx *gorm.DB) error {
		// Topic exists, associate it and create the post
		// paste.TopicID = topic.ID
		if err := tx.Create(&paste).Error; err != nil {
			return err
		}

		return nil
	}); err != nil {
		return modules.Paste{}, err
	}
	return paste, nil
}
