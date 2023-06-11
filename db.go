package flow

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"gorm.io/gorm"
)

var db *gorm.DB

func SetDB(d *gorm.DB) {
	db = d
}

type Array[T any] []T

func (a *Array[T]) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New(fmt.Sprint("Failed to scan Array value:", value))
	}
	if len(bytes) > 0 {
		return json.Unmarshal(bytes, &a)
	}
	//a = make([]T, 0)
	return nil
}

func (a Array[T]) Value() (driver.Value, error) {
	if a == nil {
		return "[]", nil
	}
	s, err := json.Marshal(a)
	fmt.Println(string(s))
	fmt.Println(err)
	return json.Marshal(a)
}
