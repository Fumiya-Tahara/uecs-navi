package pages

import (
	"log"

	"github.com/Fumiya-Tahara/uecs-navi.git/internal/domain"
	"github.com/Fumiya-Tahara/uecs-navi.git/internal/usecase/interfaces"
)

type WorkflowPageService struct {
	workflowRepository          interfaces.WorkflowRepositoryInterface
	workflowOperationRepository interfaces.WorkflowOperationRepositoryInterface
	nodeRepository              interfaces.NodeRepositoryInterface
	edgeRepository              interfaces.EdgeRepositoryInterface
}

func NewWorkflowPageService(wr interfaces.WorkflowRepositoryInterface, wor interfaces.WorkflowOperationRepositoryInterface, nr interfaces.NodeRepositoryInterface, er interfaces.EdgeRepositoryInterface) *WorkflowPageService {
	return &WorkflowPageService{
		workflowRepository: wr,
		nodeRepository:     nr,
		edgeRepository:     er,
	}
}

func (wps WorkflowPageService) GetWorkflowsWithUI(m304ID int) (*[]domain.Workflow, error) {
	workflows, err := wps.workflowRepository.GetWorkflowsFromM304(m304ID)
	if err != nil {
		return nil, err
	}

	for i, v := range *workflows {
		nodes, err := wps.nodeRepository.GetNodesFromWorkflow(v.ID)
		if err != nil {
			log.Printf("Error getting nodes: %v", err)
			continue
		}

		edges, err := wps.edgeRepository.GetEdgesFromWorkflow(v.ID)
		if err != nil {
			log.Printf("Error getting edges: %v", err)
			continue
		}

		(*workflows)[i].Node = *nodes
		(*workflows)[i].Edge = *edges
	}

	return workflows, nil
}

func (wps WorkflowPageService) CreateWorkflowWithUI(workflow *domain.Workflow) error {
	workflowID, err := wps.workflowRepository.CreateWorkflow(*workflow)
	if err != nil {
		log.Printf("Error creating workflow: %v", err)
		return err
	}

	workflowOperation := workflow.Operations
	workflowOperation.WorkflowID = workflowID
	_, err = wps.workflowOperationRepository.CreateWorkflowOperation(workflowOperation)
	if err != nil {
		log.Printf("Error creating workflow operation: %v", err)
		return err
	}

	nodes := workflow.Node
	for _, node := range nodes {
		node.WorkflowID = workflowID
		_, err := wps.nodeRepository.CreateNode(node)
		if err != nil {
			log.Printf("Error creating node: %v", err)
		}
	}

	edges := workflow.Edge
	for _, edge := range edges {
		edge.WorkflowID = workflowID
		_, err := wps.edgeRepository.CreateEdge(edge)
		if err != nil {
			log.Printf("Error creating node: %v", err)
		}
	}

	return nil
}
