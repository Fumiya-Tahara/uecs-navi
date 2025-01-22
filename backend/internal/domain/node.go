package domain

type Position struct {
	X float64
	Y float64
}

type Node struct {
	ID         int
	WorkflowID int
	NodeID     string
	Type       string
	Data       map[string]interface{}
	Position   Position
}
