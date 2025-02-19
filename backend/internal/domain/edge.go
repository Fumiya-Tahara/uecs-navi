package domain

type Edge struct {
	ID           int
	WorkflowID   int
	SourceNodeID string
	TargetNodeID string
}

func NewEdgeWithID(id, workflowID int, sourceNodeID, targetNodeID string) *Edge {
	return &Edge{
		ID:           id,
		WorkflowID:   workflowID,
		SourceNodeID: sourceNodeID,
		TargetNodeID: targetNodeID,
	}
}

func NewEdge(workflowID int, sourceNodeID, targetNodeID string) *Edge {
	return &Edge{
		WorkflowID:   workflowID,
		SourceNodeID: sourceNodeID,
		TargetNodeID: targetNodeID,
	}
}
