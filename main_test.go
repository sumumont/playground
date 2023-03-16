package main

import (
	"fmt"
	"testing"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func TestGORM(t *testing.T) {
	//user := User{Name: "jinzhu"}
	//
	//DB.Create(&user)
	//
	//var result User
	//if err := DB.First(&result, user.ID).Error; err != nil {
	//	t.Errorf("Failed, got error: %v", err)
	//}

	result := []map[string]interface{}{}
	tx := DB.Model(&Student{}).Select("students.*,b.*").Joins("left join schools as b on students.school_id = b.id")

	rows, err := tx.Rows()
	if err != nil {
		panic(err)
	}

	cl, err := rows.Columns()
	fmt.Println("columns:", cl)
	tx = tx.Find(&result)
	fmt.Println("result:", result)
	err = tx.Error
	if err != nil {
		t.Errorf("Failed, got error: %v", err)
	}
}
