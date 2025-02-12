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
	ClimateDataId *int     `json:"climate_data_id,omitempty"`
	DeviceName    *string  `json:"device_name,omitempty"`
	Duration      *int     `json:"duration"`
	SetPoint      *float32 `json:"set_point,omitempty"`
	Unit          *string  `json:"unit,omitempty"`
}

// DeviceResponse defines model for DeviceResponse.
type DeviceResponse struct {
	ClimateData *string  `json:"climate_data,omitempty"`
	DeviceName  *string  `json:"device_name,omitempty"`
	Duration    *int     `json:"duration,omitempty"`
	Id          *int     `json:"id,omitempty"`
	SetPoint    *float64 `json:"set_point,omitempty"`
	Unit        *string  `json:"unit,omitempty"`
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

// HouseRequest defines model for HouseRequest.
type HouseRequest = string

// HousesResponse defines model for HousesResponse.
type HousesResponse struct {
	Id   *int    `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

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

// Relays defines model for Relays.
type Relays struct {
	Relay1 *bool `json:"relay_1,omitempty"`
	Relay2 *bool `json:"relay_2,omitempty"`
	Relay3 *bool `json:"relay_3,omitempty"`
	Relay4 *bool `json:"relay_4,omitempty"`
	Relay5 *bool `json:"relay_5,omitempty"`
	Relay6 *bool `json:"relay_6,omitempty"`
	Relay7 *bool `json:"relay_7,omitempty"`
	Relay8 *bool `json:"relay_8,omitempty"`
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
	Relays     *Relays            `json:"relays,omitempty"`
	Workflow   *Workflow          `json:"workflow,omitempty"`
	WorkflowUI *WorkflowUIRequest `json:"workflowUI,omitempty"`
}

// WorkflowWithUIResponse defines model for WorkflowWithUIResponse.
type WorkflowWithUIResponse struct {
	Workflow   *Workflow           `json:"workflow,omitempty"`
	WorkflowUI *WorkflowUIResponse `json:"workflowUI,omitempty"`
}

// HouseId defines model for house-id.
type HouseId = int

// GetTimeSchedulesJSONBody defines parameters for GetTimeSchedules.
type GetTimeSchedulesJSONBody = int

// CreateHouseJSONRequestBody defines body for CreateHouse for application/json ContentType.
type CreateHouseJSONRequestBody = HouseRequest

// CreateDeviceJSONRequestBody defines body for CreateDevice for application/json ContentType.
type CreateDeviceJSONRequestBody = DeviceRequest

// GetTimeSchedulesJSONRequestBody defines body for GetTimeSchedules for application/json ContentType.
type GetTimeSchedulesJSONRequestBody = GetTimeSchedulesJSONBody

// CreateAndBuildTimeScheduleJSONRequestBody defines body for CreateAndBuildTimeSchedule for application/json ContentType.
type CreateAndBuildTimeScheduleJSONRequestBody = TimeScheduleRequest

// GetWorkflowsWithUIJSONRequestBody defines body for GetWorkflowsWithUI for application/json ContentType.
type GetWorkflowsWithUIJSONRequestBody = M304IDRequest

// CreateWorkflowWithUIJSONRequestBody defines body for CreateWorkflowWithUI for application/json ContentType.
type CreateWorkflowWithUIJSONRequestBody = WorkflowWithUIRequest

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// Get climate data list
	// (GET /climate-datas)
	GetClimateData(c *gin.Context)
	// Get houses list
	// (GET /houses)
	GetHouses(c *gin.Context)
	// Create a house
	// (POST /houses)
	CreateHouse(c *gin.Context)
	// Get devices list
	// (GET /houses/{house-id})
	GetDevice(c *gin.Context, houseId HouseId)
	// Create device
	// (POST /houses/{house-id}/devices)
	CreateDevice(c *gin.Context, houseId HouseId)
	// Get time schedule
	// (GET /time-schedule)
	GetTimeSchedules(c *gin.Context)
	// Create and build time schedule
	// (POST /time-schedule)
	CreateAndBuildTimeSchedule(c *gin.Context)
	// Get workflows with the state of React Flow.
	// (GET /workflows-with-ui)
	GetWorkflowsWithUI(c *gin.Context)
	// Create a workflow with UI
	// (POST /workflows-with-ui)
	CreateWorkflowWithUI(c *gin.Context)
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

// CreateHouse operation middleware
func (siw *ServerInterfaceWrapper) CreateHouse(c *gin.Context) {

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.CreateHouse(c)
}

// GetDevice operation middleware
func (siw *ServerInterfaceWrapper) GetDevice(c *gin.Context) {

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

	siw.Handler.GetDevice(c, houseId)
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

// GetTimeSchedules operation middleware
func (siw *ServerInterfaceWrapper) GetTimeSchedules(c *gin.Context) {

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.GetTimeSchedules(c)
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

// GetWorkflowsWithUI operation middleware
func (siw *ServerInterfaceWrapper) GetWorkflowsWithUI(c *gin.Context) {

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.GetWorkflowsWithUI(c)
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
	router.POST(options.BaseURL+"/houses", wrapper.CreateHouse)
	router.GET(options.BaseURL+"/houses/:house-id", wrapper.GetDevice)
	router.POST(options.BaseURL+"/houses/:house-id/devices", wrapper.CreateDevice)
	router.GET(options.BaseURL+"/time-schedule", wrapper.GetTimeSchedules)
	router.POST(options.BaseURL+"/time-schedule", wrapper.CreateAndBuildTimeSchedule)
	router.GET(options.BaseURL+"/workflows-with-ui", wrapper.GetWorkflowsWithUI)
	router.POST(options.BaseURL+"/workflows-with-ui", wrapper.CreateWorkflowWithUI)
}
