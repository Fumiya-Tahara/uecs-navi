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

func NewPosition(x, y float64) *Position {
	return &Position{
		X: x,
		Y: y,
	}
}

func NewNodeWithID(id, workflowID int, nodeID, nodeType string, data map[string]interface{}, position Position) *Node {
	return &Node{
		ID:         id,
		WorkflowID: workflowID,
		NodeID:     nodeID,
		Type:       nodeType,
		Data:       data,
		Position:   position,
	}
}

func NewNode(workflowID int, nodeID, nodeType string, data map[string]interface{}, position Position) *Node {
	return &Node{
		WorkflowID: workflowID,
		NodeID:     nodeID,
		Type:       nodeType,
		Data:       data,
		Position:   position,
	}
}
