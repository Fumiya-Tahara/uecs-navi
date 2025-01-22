package domain

type Edge struct {
	ID           int
	WorkflowID   int
	SourceNodeID string
	TargetNodeID string
}
