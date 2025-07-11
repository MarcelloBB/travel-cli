package model

type Request struct {
	IdRequest    uint   `gorm:"primaryKey" json:"id_request"`
	IdCollection uint   `json:"id_collection"`
	Title        string `json:"title"`
	Description  string `json:"description"`
	Method       string `json:"method"`
	Body         string `json:"body"`
	Params       string `json:"params"`
}
