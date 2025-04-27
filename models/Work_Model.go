package model



type WorkModel struct {
	Id string `json:"id"`
	Role string `json:"role"`
	Company string `json:"company"`
	Task []string `json:"task"`
	Stack []string `json:"stack"`
	Image string `json:"image"`
	StartDate *string `json:"start_date"`
	EndDate *string `json:"end_date"`
}


