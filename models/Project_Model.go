package model



type ProjectModel struct {
	Id string `json:"id"`
	Name string `json:"name"`
	Description string `json:"description"`
	Image string `json:"image"`
	Stack []string `json:"stack"`
	Repo *string `json:"repo"`
	Demo *string `json:"demo"` 
}