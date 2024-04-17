package dto

type CreateTaskInput struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Priority    int    `json:"priority"`
}
