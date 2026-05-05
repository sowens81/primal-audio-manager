package models

type Rating struct {
	Count   int     `json:"count"`
	Average float64 `json:"average"`
}
