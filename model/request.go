package model

type Request struct {
	IdRequest    uint    `gorm:"primaryKey" json:"id_request"`
	IdCollection uint    `json:"id_collection"`
	Method       string  `json:"method"`
	Url          string  `json:"url"`
	Title        string  `json:"title"`
	Headers      *string `json:"headers"`
	Body         *string `json:"body"`
	Params       *string `json:"params"`
}
