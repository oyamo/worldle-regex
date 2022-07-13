package entity

import "html/template"

type ServerResponse struct {
	Regex        string          `json:"regex"`
	StatusCode   int             `json:"status_code"`
	Results      []string        `json:"results"`
	Pages        int             `json:"pages"`
	Page         int             `json:"page"`
	Total        int             `json:"total"`
	Time         float64         `json:"time"`
	HasNext      bool            `json:"has_next"`
	NextPage     int             `json:"next_page"`
	PreviousPage int             `json:"previous_page"`
	HasPrev      bool            `json:"has_prev"`
	PagesContent []template.HTML `json:"pages_content"`
	PagesActive  []int           `json:"pages_active"`
}
