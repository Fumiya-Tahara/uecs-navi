package dto

type Workflow struct {
	ID     int    `json:"id"`
	M304ID int    `json:"m304_id"`
	Name   string `json:"name"`
}

func NewWorkflow(id, m304ID int, name string) *Workflow {
	return &Workflow{
		ID:     id,
		Name:   name,
		M304ID: m304ID,
	}
}
