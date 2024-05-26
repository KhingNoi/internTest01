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
	Id            int     `gorm:"primaryKey" json:"id"`
	Name          string  `json:"name"`
	Description   string  `json:"description"`
	Price         float64 `json:"price"`
	Region        string  `json:"region"`
	Weight        float64 `json:"weight"`
	FlavorProfile JSONMap `json:"flavor_profile" gorm:"type:json"`
	GrindOption   JSONMap `json:"grind_option" gorm:"type:json"`
	RoastLevel    *int    `json:"roast_level"`
	ImageURL      string  `json:"image_url"`
	Stock         int     `json:"stock"`
}

type ProductWithPaginate struct {
	Page  *int `json:"page,omitempty" query:"page"`
	Size  *int `json:"size,omitempty" query:"size"`
	Total *int `json:"total,omitempty" query:"total"`
	Data  []Product
}

type SearchTermPayload struct {
	Name string `json:"name,omitempty"`
}

type StockData struct {
	Name  string `json:"name"`
	Stock int    `json:"stock"`
}
