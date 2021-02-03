package utils

import "gorm.io/gorm"

func Defer(tx *gorm.DB, code *int) {
	if *code == 0 {
		tx.Commit()
	} else {
		tx.Rollback()
	}
}
