package flow

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"testing"
)

func TestDb(t *testing.T) {
	dsn := "root:wjswkwssyr5188@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
	d, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	d.AutoMigrate(&DbNode{}, &Event{})
	if err != nil {
		t.Error(err)
	}
	SetDB(d)
	d.AutoMigrate(&DbNode{}, &Event{}, &Process{})
}
