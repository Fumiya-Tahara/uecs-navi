package domain

type WorkflowOperation struct {
	ID         int
	WorkflowID int
	Relay1     *bool
	Relay2     *bool
	Relay3     *bool
	Relay4     *bool
	Relay5     *bool
	Relay6     *bool
	Relay7     *bool
	Relay8     *bool
}

type Workflow struct {
	ID         int
	M304ID     int
	Name       string
	Operations WorkflowOperation
	Node       []Node
	Edge       []Edge
}

func NewWorkflowWithBasicInfo(id, m304ID int, name string) *Workflow {
	return &Workflow{
		ID:     id,
		M304ID: m304ID,
		Name:   name,
	}
}

func NewWorkflowOperationWithID(id, workflowID int, relay1, relay2, relay3, relay4, relay5, relay6, relay7, relay8 *bool) *WorkflowOperation {
	return &WorkflowOperation{
		ID:         id,
		WorkflowID: workflowID,
		Relay1:     relay1,
		Relay2:     relay2,
		Relay3:     relay3,
		Relay4:     relay4,
		Relay5:     relay5,
		Relay6:     relay6,
		Relay7:     relay7,
		Relay8:     relay8,
	}
}

func NewWorkflowOperation(workflowID int, relay1, relay2, relay3, relay4, relay5, relay6, relay7, relay8 *bool) *WorkflowOperation {
	return &WorkflowOperation{
		WorkflowID: workflowID,
		Relay1:     relay1,
		Relay2:     relay2,
		Relay3:     relay3,
		Relay4:     relay4,
		Relay5:     relay5,
		Relay6:     relay6,
		Relay7:     relay7,
		Relay8:     relay8,
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
