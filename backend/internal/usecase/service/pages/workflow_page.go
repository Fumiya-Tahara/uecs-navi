package pages

import (
	"log"

	"github.com/Fumiya-Tahara/uecs-navi.git/internal/domain"
	"github.com/Fumiya-Tahara/uecs-navi.git/internal/usecase/interfaces"
)

type WorkflowPageService struct {
	workflowRepository interfaces.WorkflowRepositoryInterface
	nodeRepository     interfaces.NodeRepositoryInterface
	edgeRepository     interfaces.EdgeRepositoryInterface
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

func (wps WorkflowPageService) CreateWorkflowsWithUI(workflow *domain.Workflow) error {
	nodes := workflow.Node
	for _, v := range nodes {
		_, err := wps.nodeRepository.CreateNode(v)
		if err != nil {
			log.Printf("Error creating node: %v", err)
		}
	}

	edges := workflow.Edge
	for _, v := range edges {
		_, err := wps.edgeRepository.CreateEdge(v)
		if err != nil {
			log.Printf("Error creating node: %v", err)
		}
	}

	_, err := wps.workflowRepository.CreateWorkflow(*workflow)
	if err != nil {
		log.Printf("Error creating workflow: %v", err)
	}

	// operationの追加を考える

	return nil
}
