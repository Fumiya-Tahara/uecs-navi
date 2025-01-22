package domain

type WorkflowOperation struct {
	ID         int
	WorkflowID int
	Relays     [8]bool
}

type Workflow struct {
	ID         int
	M304ID     int
	Name       string
	Operations WorkflowOperation
	Node       []Node
	Edge       []Edge
}

func NewWorkflowOperationWithID(id, workflowID int, relays [8]bool) *WorkflowOperation {
	return &WorkflowOperation{
		ID:         id,
		WorkflowID: workflowID,
		Relays:     relays,
	}
}

func NewWorkflowOperation(workflowID int, relays [8]bool) *WorkflowOperation {
	return &WorkflowOperation{
		WorkflowID: workflowID,
		Relays:     relays,
	}
}

func NewWorkflowWithID(id, m304ID int, name string, operations WorkflowOperation, node []Node, edge []Edge) *Workflow {
	return &Workflow{
		ID:         id,
		M304ID:     m304ID,
		Name:       name,
		Operations: operations,
		Node:       node,
		Edge:       edge,
	}
}

func NewWorkflow(m304ID int, name string, operations WorkflowOperation, node []Node, edge []Edge) *Workflow {
	return &Workflow{
		M304ID:     m304ID,
		Name:       name,
		Operations: operations,
		Node:       node,
		Edge:       edge,
	}
}
