package controller

import (
	"log"
	"net/http"

	"github.com/Fumiya-Tahara/uecs-navi.git/internal/controller/converter"
	"github.com/Fumiya-Tahara/uecs-navi.git/internal/controller/dto"
	"github.com/Fumiya-Tahara/uecs-navi.git/internal/usecase/service/pages"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	workflowPageService     pages.WorkflowPageService
	timeSchedulePageService pages.TimeSchedulePageService
}

func NewHandler(wps pages.WorkflowPageService, tsps pages.TimeSchedulePageService) *Handler {
	return &Handler{
		workflowPageService:     wps,
		timeSchedulePageService: tsps,
	}
}

func (h Handler) GetClimateData(c *gin.Context) {

}

func (h Handler) GetHouses(c *gin.Context) {

}

func (h Handler) CreateHouse(c *gin.Context) {

}

func (h Handler) GetDevice(c *gin.Context, houseId int) {

}

func (h Handler) CreateDevice(c *gin.Context, houseId int) {

}

func (h Handler) GetTimeSchedules(c *gin.Context) {
	var req dto.M304IDRequest
	if err := c.BindJSON(&req); err != nil {
		log.Printf("Error binding json: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"status": "BadRequest"})
		return
	}

	timeSchedule, err := h.timeSchedulePageService.GetTimeSchedule(int(req))
	if err != nil {
		log.Printf("Error getting time schedule: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"status": "InternalServerError"})
		return
	}

	workflows, err := h.timeSchedulePageService.GetWorkflows(int(req))
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

func (h Handler) GetWorkflowsWithUI(c *gin.Context) {
	var m304ID int
	if err := c.BindJSON(&m304ID); err != nil {
		log.Printf("Error binding JSON: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"status": "BadRequest", "error": "Invalid request body"})
		return
	}

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
