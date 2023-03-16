package main

import (
	"fmt"
	"testing"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

type School struct {
	ID   int64  `gorm:"primary_key" json:"id"`
	Name string `json:"name"`
}
type Student struct {
	ID       int64  `gorm:"primary_key" json:"id"`
	Name     string `json:"name"`
	SchoolId int64  `json:"schoolId"`
}

// if column duplicat,rows.Columns() can add suffix such as name,name(1),name(2)?
func TestGORM(t *testing.T) {

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
