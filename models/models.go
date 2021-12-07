package models

import (
	"time"
)

type User struct {
	UserID    int
	FirstName string
	LastName  string
	Email     string
}

type CurrentEvents struct {
	Response []struct {
		Id          int `json:"id"`
		Competition struct {
			Id       int    `json:"id"`
			Name     string `json:"name"`
			Location struct {
				Country string `json:"country"`
				City    string `json:"city"`
			} `json:"location"`
		} `json:"competition"`
		Circuit struct {
			Id    int    `json:"id"`
			Name  string `json:"name"`
			Image string `json:"image"`
		} `json:"circuit"`
		Season int    `json:"season"`
		Type   string `json:"type"`
		Laps   struct {
			Current interface{} `json:"current"`
			Total   int         `json:"total"`
		} `json:"laps"`
		Distance string      `json:"distance"`
		Timezone string      `json:"timezone"`
		Date     time.Time   `json:"date"`
		Weather  interface{} `json:"weather"`
		Status   string      `json:"status"`
	} `json:"response"`
}
