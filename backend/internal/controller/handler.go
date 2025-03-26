package controller

import (
	"log"
	"net/http"

	"github.com/Fumiya-Tahara/uecs-navi.git/internal/controller/converter"
	"github.com/Fumiya-Tahara/uecs-navi.git/internal/controller/dto"
	"github.com/Fumiya-Tahara/uecs-navi.git/internal/controller/generated"
	"github.com/Fumiya-Tahara/uecs-navi.git/internal/usecase/service/pages"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	workflowPageService     pages.WorkflowPageService
	timeSchedulePageService pages.TimeSchedulePageService
	pageUtilitiesService    pages.PageUtilitiesService
}

func NewHandler(wps pages.WorkflowPageService, tsps pages.TimeSchedulePageService, pus pages.PageUtilitiesService) *Handler {
	return &Handler{
		workflowPageService:     wps,
		timeSchedulePageService: tsps,
		pageUtilitiesService:    pus,
	}
}

func (h Handler) GetClimateData(c *gin.Context) {
	climateData, err := h.pageUtilitiesService.GetClimateData()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "InternalServerError"})
		return
	}

	c.JSON(http.StatusOK, climateData)
}

func (h Handler) GetHouses(c *gin.Context) {

}

func (h Handler) CreateHouse(c *gin.Context) {

}

func (h Handler) GetDevice(c *gin.Context, houseId int) {

}

func (h Handler) CreateDevice(c *gin.Context, houseId int) {

}

func (h Handler) GetM304s(c *gin.Context) {
	m304IDs, err := h.pageUtilitiesService.GetM304IDs()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "InternalServerError"})
		return
	}

	c.JSON(http.StatusOK, m304IDs)
}

func (h Handler) GetTimeSchedule(c *gin.Context, m304ID generated.M304ID) {
	timeSchedule, err := h.timeSchedulePageService.GetTimeSchedule(int(m304ID))
	if err != nil {
		log.Printf("Error getting time schedule: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"status": "InternalServerError"})
		return
	}

	workflows, err := h.timeSchedulePageService.GetWorkflows(int(m304ID))
	if err != nil {
		log.Printf("Error getting workflows: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"status": "InternalServerError"})
		return
	}

	timeScheduleRes := converter.ToDtoTimeScheduleResponse(*timeSchedule, *workflows)

	c.JSON(http.StatusOK, timeScheduleRes)
}

func (h Handler) CreateAndBuildTimeSchedule(c *gin.Context) {
	var req dto.TimeScheduleRequest
	if err := c.BindJSON(&req); err != nil {
		log.Printf("Error binding json: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"status": "BadRequest"})
		return
	}

	newTimeSchedule := converter.ToDomainTimeSchedule(req)
	err := h.timeSchedulePageService.CreateAndBuildTimeSchedule(*newTimeSchedule)
	if err != nil {
		log.Printf("Error creating workflow: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"status": "InternalServerError"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Workflow created successfully"})
}

func (h Handler) UpdateAndBuildTimeSchedule(c *gin.Context) {
	var req dto.TimeScheduleRequest
	if err := c.BindJSON(&req); err != nil {
		log.Printf("Error binding json: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"status": "BadRequest"})
		return
	}

	newTimeSchedule := converter.ToDomainTimeSchedule(req)
	err := h.timeSchedulePageService.UpdateAndBuildTimeSchedule(*newTimeSchedule)
	if err != nil {
		log.Printf("Error creating workflow: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"status": "InternalServerError"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Workflow created successfully"})
}

func (h Handler) GetWorkflowsWithUI(c *gin.Context, m304ID generated.M304ID) {
	workflows, err := h.workflowPageService.GetWorkflowsWithUI(m304ID)
	if err != nil {
		log.Printf("Error getting workflows with UI: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"status": "BadRequest"})
		return
	}

	workflowsRes := converter.ToDtoWorkflowWithUI(*workflows)

	c.JSON(http.StatusOK, workflowsRes)
}

func (h Handler) CreateWorkflowWithUI(c *gin.Context) {
	var req dto.WorkflowWithUIRequest
	if err := c.BindJSON(&req); err != nil {
		log.Printf("Error binding json: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"status": "BadRequest"})
		return
	}

	newWorkflow := converter.ToDomainWorkflow(req)
	err := h.workflowPageService.CreateWorkflowWithUI(newWorkflow)
	if err != nil {
		log.Printf("Error creating workflow: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"status": "InternalServerError"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Workflow created successfully"})
}

func (h Handler) UpdateWorkflowWithUI(c *gin.Context) {
	var req dto.WorkflowWithUIRequest
	if err := c.BindJSON(&req); err != nil {
		log.Printf("Error binding json: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"status": "BadRequest"})
		return
	}

	newWorkflow := converter.ToDomainWorkflow(req)
	err := h.workflowPageService.UpdateWorkflowWithUI(newWorkflow)
	if err != nil {
		log.Printf("Error creating workflow: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"status": "InternalServerError"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Workflow created successfully"})
}

func (h Handler) DeleteWorkflowWithUI(c *gin.Context, workflowId generated.WorkflowId) {
	if err := h.workflowPageService.DeleteWorkflowWithUI(workflowId); err != nil {
		log.Printf("Error deleting workflow: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"status": "InternalServerError"})

		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Workflow deleted successfully"})
}
