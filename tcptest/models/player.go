package models

type Player struct {
	ID string `json:"id"`
	Color string `json:"color"`
	X float32 `json:"x"`
	Y float32 `json:"y"`
	Z float32 `json:"z"`
	
}

var Players []Player