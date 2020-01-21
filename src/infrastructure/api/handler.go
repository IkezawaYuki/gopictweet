package api

import "github.com/jinzhu/gorm"

type Client struct {
	db *gorm.DB
}

func NewClient(db *gorm.DB) *Client {
	return &Client{db: db}
}
