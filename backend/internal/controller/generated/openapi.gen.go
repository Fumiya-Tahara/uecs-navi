// Package generated provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/oapi-codegen/oapi-codegen/v2 version v2.4.1 DO NOT EDIT.
package generated

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/oapi-codegen/runtime"
)

// ClimateData defines model for ClimateData.
type ClimateData struct {
	Id   *int    `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	Unit *string `json:"unit,omitempty"`
}

// Condition defines model for Condition.
type Condition struct {
	SelectedClimateDataId        *int     `json:"selected_climate_data_id,omitempty"`
	SelectedComparisonOperatorId *int     `json:"selected_comparison_operator_id,omitempty"`
	SetPoint                     *float32 `json:"set_point,omitempty"`
}

// DeviceRequest defines model for DeviceRequest.
type DeviceRequest struct {
	ClimateDataId *int    `json:"climate_data_id,omitempty"`
	M304Id        *int    `json:"m304_id,omitempty"`
	Name          *string `json:"name,omitempty"`
	Rly           *int    `json:"rly,omitempty"`
}

// EdgeRequest defines model for EdgeRequest.
type EdgeRequest struct {
	SourceNodeId *string `json:"source_node_id,omitempty"`
	TargetNodeId *string `json:"target_node_id,omitempty"`
}

// EdgeResponse defines model for EdgeResponse.
type EdgeResponse struct {
	Id           *int    `json:"id,omitempty"`
	SourceNodeId *string `json:"source_node_id,omitempty"`
	TargetNodeId *string `json:"target_node_id,omitempty"`
	WorkflowId   *int    `json:"workflow_id,omitempty"`
}

// HousesResponse defines model for HousesResponse.
type HousesResponse struct {
	Id   *int    `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

// M304ID defines model for M304ID.
type M304ID = int

// M304IDRequest defines model for M304IDRequest.
type M304IDRequest = int

// NodeRequest defines model for NodeRequest.
type NodeRequest struct {
	Data           *map[string]interface{} `json:"data,omitempty"`
	PositionX      *float32                `json:"position_x,omitempty"`
	PositionY      *float32                `json:"position_y,omitempty"`
	Type           *string                 `json:"type,omitempty"`
	WorkflowNodeId *string                 `json:"workflow_node_id,omitempty"`
}

// NodeResponse defines model for NodeResponse.
type NodeResponse struct {
	Data           *map[string]interface{} `json:"data,omitempty"`
	Id             *int                    `json:"id,omitempty"`
	PositionX      *float32                `json:"position_x,omitempty"`
	PositionY      *float32                `json:"position_y,omitempty"`
	Type           *string                 `json:"type,omitempty"`
	WorkflowId     *int                    `json:"workflow_id,omitempty"`
	WorkflowNodeId *string                 `json:"workflow_node_id,omitempty"`
}

// TimeScheduleRequest defines model for TimeScheduleRequest.
type TimeScheduleRequest struct {
	M304Id       *int               `json:"m304_id,omitempty"`
	TimeSchedule *[]TimeScheduleRow `json:"time_schedule,omitempty"`
}

// TimeScheduleResponse defines model for TimeScheduleResponse.
type TimeScheduleResponse struct {
	Id            *int               `json:"id,omitempty"`
	M304Id        *int               `json:"m304_id,omitempty"`
	TimeSchedules *[]TimeScheduleRow `json:"time_schedules,omitempty"`
	Workflows     *[]Workflow        `json:"workflows,omitempty"`
}

// TimeScheduleRow defines model for TimeScheduleRow.
type TimeScheduleRow struct {
	Condition *Condition `json:"condition,omitempty"`

	// EndTime 終了時刻（HH:mm形式）
	EndTime            *string `json:"end_time,omitempty"`
	SelectedWorkflowId *int    `json:"selected_workflow_id,omitempty"`

	// StartTime 開始時刻（HH:mm形式）
	StartTime *string `json:"start_time,omitempty"`
}

// Workflow defines model for Workflow.
type Workflow struct {
	Id     *int    `json:"id,omitempty"`
	M304Id *int    `json:"m304_id,omitempty"`
	Name   *string `json:"name,omitempty"`
}

// WorkflowOperations defines model for WorkflowOperations.
type WorkflowOperations struct {
	Relay1 *bool `json:"relay_1,omitempty"`
	Relay2 *bool `json:"relay_2,omitempty"`
	Relay3 *bool `json:"relay_3,omitempty"`
	Relay4 *bool `json:"relay_4,omitempty"`
	Relay5 *bool `json:"relay_5,omitempty"`
	Relay6 *bool `json:"relay_6,omitempty"`
	Relay7 *bool `json:"relay_7,omitempty"`
	Relay8 *bool `json:"relay_8,omitempty"`
}

// WorkflowUIRequest defines model for WorkflowUIRequest.
type WorkflowUIRequest struct {
	Edges *[]EdgeRequest `json:"edges,omitempty"`
	Nodes *[]NodeRequest `json:"nodes,omitempty"`
}

// WorkflowUIResponse defines model for WorkflowUIResponse.
type WorkflowUIResponse struct {
	Edges *[]EdgeResponse `json:"edges,omitempty"`
	Nodes *[]NodeResponse `json:"nodes,omitempty"`
}

// WorkflowWithUIRequest defines model for WorkflowWithUIRequest.
type WorkflowWithUIRequest struct {
	Relays     *WorkflowOperations `json:"relays,omitempty"`
	Workflow   *Workflow           `json:"workflow,omitempty"`
	WorkflowUI *WorkflowUIRequest  `json:"workflowUI,omitempty"`
}

// WorkflowWithUIResponse defines model for WorkflowWithUIResponse.
type WorkflowWithUIResponse struct {
	Workflow   *Workflow           `json:"workflow,omitempty"`
	WorkflowUI *WorkflowUIResponse `json:"workflowUI,omitempty"`
}

// HouseId defines model for house-id.
type HouseId = int

// M304Id defines model for m304-id.
type M304Id = int

// WorkflowId defines model for workflow-id.
type WorkflowId = int

// CreateDeviceJSONRequestBody defines body for CreateDevice for application/json ContentType.
type CreateDeviceJSONRequestBody = DeviceRequest

// CreateAndBuildTimeScheduleJSONRequestBody defines body for CreateAndBuildTimeSchedule for application/json ContentType.
type CreateAndBuildTimeScheduleJSONRequestBody = TimeScheduleRequest

// UpdateAndBuildTimeScheduleJSONRequestBody defines body for UpdateAndBuildTimeSchedule for application/json ContentType.
type UpdateAndBuildTimeScheduleJSONRequestBody = TimeScheduleRequest

// GetTimeScheduleJSONRequestBody defines body for GetTimeSchedule for application/json ContentType.
type GetTimeScheduleJSONRequestBody = M304IDRequest

// CreateWorkflowWithUIJSONRequestBody defines body for CreateWorkflowWithUI for application/json ContentType.
type CreateWorkflowWithUIJSONRequestBody = WorkflowWithUIRequest

// UpdateWorkflowWithUIJSONRequestBody defines body for UpdateWorkflowWithUI for application/json ContentType.
type UpdateWorkflowWithUIJSONRequestBody = WorkflowWithUIRequest

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// Get climate data list
	// (GET /climate-datas)
	GetClimateData(c *gin.Context)
	// Get houses list
	// (GET /houses)
	GetHouses(c *gin.Context)
	// Create device
	// (POST /houses/{house-id}/devices)
	CreateDevice(c *gin.Context, houseId HouseId)
	// Get M304 ID list
	// (GET /m304s)
	GetM304s(c *gin.Context)
	// Get devices connected to M304
	// (GET /m304s/{m304-id}/devices)
	GetDevices(c *gin.Context, m304Id M304Id)
	// Create and build time schedule
	// (POST /time-schedule)
	CreateAndBuildTimeSchedule(c *gin.Context)
	// Update and build time schedule
	// (PUT /time-schedule)
	UpdateAndBuildTimeSchedule(c *gin.Context)
	// Get time schedule
	// (GET /time-schedule/{m304-id})
	GetTimeSchedule(c *gin.Context, m304Id M304Id)
	// Delete a workflow with UI
	// (DELETE /workflow-with-ui/{workflow-id})
	DeleteWorkflowWithUI(c *gin.Context, workflowId WorkflowId)
	// Create a workflow with UI
	// (POST /workflows-with-ui)
	CreateWorkflowWithUI(c *gin.Context)
	// Update a workflow with UI
	// (PUT /workflows-with-ui)
	UpdateWorkflowWithUI(c *gin.Context)
	// Get workflows with the state of React Flow.
	// (GET /workflows-with-ui/{m304-id})
	GetWorkflowsWithUI(c *gin.Context, m304Id M304Id)
}

// ServerInterfaceWrapper converts contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler            ServerInterface
	HandlerMiddlewares []MiddlewareFunc
	ErrorHandler       func(*gin.Context, error, int)
}

type MiddlewareFunc func(c *gin.Context)

// GetClimateData operation middleware
func (siw *ServerInterfaceWrapper) GetClimateData(c *gin.Context) {

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.GetClimateData(c)
}

// GetHouses operation middleware
func (siw *ServerInterfaceWrapper) GetHouses(c *gin.Context) {

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.GetHouses(c)
}

// CreateDevice operation middleware
func (siw *ServerInterfaceWrapper) CreateDevice(c *gin.Context) {

	var err error

	// ------------- Path parameter "house-id" -------------
	var houseId HouseId

	err = runtime.BindStyledParameterWithOptions("simple", "house-id", c.Param("house-id"), &houseId, runtime.BindStyledParameterOptions{Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter house-id: %w", err), http.StatusBadRequest)
		return
	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.CreateDevice(c, houseId)
}

// GetM304s operation middleware
func (siw *ServerInterfaceWrapper) GetM304s(c *gin.Context) {

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.GetM304s(c)
}

// GetDevices operation middleware
func (siw *ServerInterfaceWrapper) GetDevices(c *gin.Context) {

	var err error

	// ------------- Path parameter "m304-id" -------------
	var m304Id M304Id

	err = runtime.BindStyledParameterWithOptions("simple", "m304-id", c.Param("m304-id"), &m304Id, runtime.BindStyledParameterOptions{Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter m304-id: %w", err), http.StatusBadRequest)
		return
	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.GetDevices(c, m304Id)
}

// CreateAndBuildTimeSchedule operation middleware
func (siw *ServerInterfaceWrapper) CreateAndBuildTimeSchedule(c *gin.Context) {

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.CreateAndBuildTimeSchedule(c)
}

// UpdateAndBuildTimeSchedule operation middleware
func (siw *ServerInterfaceWrapper) UpdateAndBuildTimeSchedule(c *gin.Context) {

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.UpdateAndBuildTimeSchedule(c)
}

// GetTimeSchedule operation middleware
func (siw *ServerInterfaceWrapper) GetTimeSchedule(c *gin.Context) {

	var err error

	// ------------- Path parameter "m304-id" -------------
	var m304Id M304Id

	err = runtime.BindStyledParameterWithOptions("simple", "m304-id", c.Param("m304-id"), &m304Id, runtime.BindStyledParameterOptions{Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter m304-id: %w", err), http.StatusBadRequest)
		return
	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.GetTimeSchedule(c, m304Id)
}

// DeleteWorkflowWithUI operation middleware
func (siw *ServerInterfaceWrapper) DeleteWorkflowWithUI(c *gin.Context) {

	var err error

	// ------------- Path parameter "workflow-id" -------------
	var workflowId WorkflowId

	err = runtime.BindStyledParameterWithOptions("simple", "workflow-id", c.Param("workflow-id"), &workflowId, runtime.BindStyledParameterOptions{Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter workflow-id: %w", err), http.StatusBadRequest)
		return
	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.DeleteWorkflowWithUI(c, workflowId)
}

// CreateWorkflowWithUI operation middleware
func (siw *ServerInterfaceWrapper) CreateWorkflowWithUI(c *gin.Context) {

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.CreateWorkflowWithUI(c)
}

// UpdateWorkflowWithUI operation middleware
func (siw *ServerInterfaceWrapper) UpdateWorkflowWithUI(c *gin.Context) {

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.UpdateWorkflowWithUI(c)
}

// GetWorkflowsWithUI operation middleware
func (siw *ServerInterfaceWrapper) GetWorkflowsWithUI(c *gin.Context) {

	var err error

	// ------------- Path parameter "m304-id" -------------
	var m304Id M304Id

	err = runtime.BindStyledParameterWithOptions("simple", "m304-id", c.Param("m304-id"), &m304Id, runtime.BindStyledParameterOptions{Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter m304-id: %w", err), http.StatusBadRequest)
		return
	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.GetWorkflowsWithUI(c, m304Id)
}

// GinServerOptions provides options for the Gin server.
type GinServerOptions struct {
	BaseURL      string
	Middlewares  []MiddlewareFunc
	ErrorHandler func(*gin.Context, error, int)
}

// RegisterHandlers creates http.Handler with routing matching OpenAPI spec.
func RegisterHandlers(router gin.IRouter, si ServerInterface) {
	RegisterHandlersWithOptions(router, si, GinServerOptions{})
}

// RegisterHandlersWithOptions creates http.Handler with additional options
func RegisterHandlersWithOptions(router gin.IRouter, si ServerInterface, options GinServerOptions) {
	errorHandler := options.ErrorHandler
	if errorHandler == nil {
		errorHandler = func(c *gin.Context, err error, statusCode int) {
			c.JSON(statusCode, gin.H{"msg": err.Error()})
		}
	}

	wrapper := ServerInterfaceWrapper{
		Handler:            si,
		HandlerMiddlewares: options.Middlewares,
		ErrorHandler:       errorHandler,
	}

	router.GET(options.BaseURL+"/climate-datas", wrapper.GetClimateData)
	router.GET(options.BaseURL+"/houses", wrapper.GetHouses)
	router.POST(options.BaseURL+"/houses/:house-id/devices", wrapper.CreateDevice)
	router.GET(options.BaseURL+"/m304s", wrapper.GetM304s)
	router.GET(options.BaseURL+"/m304s/:m304-id/devices", wrapper.GetDevices)
	router.POST(options.BaseURL+"/time-schedule", wrapper.CreateAndBuildTimeSchedule)
	router.PUT(options.BaseURL+"/time-schedule", wrapper.UpdateAndBuildTimeSchedule)
	router.GET(options.BaseURL+"/time-schedule/:m304-id", wrapper.GetTimeSchedule)
	router.DELETE(options.BaseURL+"/workflow-with-ui/:workflow-id", wrapper.DeleteWorkflowWithUI)
	router.POST(options.BaseURL+"/workflows-with-ui", wrapper.CreateWorkflowWithUI)
	router.PUT(options.BaseURL+"/workflows-with-ui", wrapper.UpdateWorkflowWithUI)
	router.GET(options.BaseURL+"/workflows-with-ui/:m304-id", wrapper.GetWorkflowsWithUI)
}
