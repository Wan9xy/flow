package gorm

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	g "gorm.io/gorm"
)

type gorm struct {
	g.DB
}

func NewGorm(db g.DB) *gorm {
	return &gorm{db}
}

type Array[T any] []T

// 实现 sql.Scanner 接口，Scan 将 value 扫描至 Jsonb
func (a *Array[T]) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New(fmt.Sprint("Failed to scan Array value:", value))
	}
	if len(bytes) > 0 {
		return json.Unmarshal(bytes, a)
	}
	*a = make([]T, 0)
	return nil
}

// 实现 driver.Valuer 接口，Value 返回 json value
func (a Array[T]) Value() (driver.Value, error) {
	if a == nil {
		return "[]", nil
	}
	return convertor.ToString(a), nil
}
