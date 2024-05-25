package models

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
)

type JSONMap []string

func (c JSONMap) Value() (driver.Value, error) {
	return json.Marshal(c)
}

func (c *JSONMap) Scan(v interface{}) error {
	switch tv := v.(type) {
	case []uint8:
		return json.Unmarshal([]byte(tv), &c)
	}
	return errors.New("unsupported type")
}

type Product struct {
	ProductID     int     `gorm:"primaryKey" json:"ProductID"`
	Name          string  `json:"name"`
	Description   string  `json:"description"`
	Price         float64 `json:"price"`
	Region        string  `json:"region"`
	Weight        float64 `json:"weight"`
	FlavorProfile JSONMap `json:"FlavorProfile"`
	GrindOption   JSONMap `json:"GrindOption"`
	RoastLevel    int     `json:"RoastLevel"`
	ImageURL      string  `json:"ImageURL"`
	Stock         int     `json:"stock"`
}
