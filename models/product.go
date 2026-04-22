package models

type Product struct {
	ID    int    `json:"id"`
	Nomi  string `json:"nomi"`
	Narxi int    `json:"narxi"`
}

var Ombor = make(map[int]Product)
var OxirgiID = 0
