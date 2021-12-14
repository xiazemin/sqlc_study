// Code generated by sqlc. DO NOT EDIT.

package gorse

import (
	"encoding/json"
	"time"
)

type Feedback struct {
	FeedbackType string    `json:"feedback_type"`
	UserID       string    `json:"user_id"`
	ItemID       string    `json:"item_id"`
	TimeStamp    time.Time `json:"time_stamp"`
	Comment      string    `json:"comment"`
}

type Item struct {
	ItemID    string          `json:"item_id"`
	TimeStamp time.Time       `json:"time_stamp"`
	Labels    json.RawMessage `json:"labels"`
	Comment   string          `json:"comment"`
}

type Measurement struct {
	Name      string    `json:"name"`
	TimeStamp time.Time `json:"time_stamp"`
	Value     float64   `json:"value"`
	Comment   string    `json:"comment"`
}

type User struct {
	UserID    string          `json:"user_id"`
	Labels    json.RawMessage `json:"labels"`
	Comment   string          `json:"comment"`
	Subscribe json.RawMessage `json:"subscribe"`
}