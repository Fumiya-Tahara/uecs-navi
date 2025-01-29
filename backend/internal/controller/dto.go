package controller

// Schemas
// General schemas
type ClimateData struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Unit string `unit:"unit"`
}

type Relays struct {
	Relay1 bool `json:"relay_1"`
	Relay2 bool `json:"relay_2"`
	Relay3 bool `json:"relay_3"`
	Relay4 bool `json:"relay_4"`
	Relay5 bool `json:"relay_5"`
	Relay6 bool `json:"relay_6"`
	Relay7 bool `json:"relay_7"`
	Relay8 bool `json:"relay_8"`
}

type Workflow struct {
	Name     string `json:"name"`
	M304ID   int    `json:"m304_id"`
	DeviceID int    `json:"device_id"`
}

type Condition struct {
	SelectedClimateDataID        int     `json:"selected_climate_data_id"`
	SelectedComparisonOperatorID int     `json:"selected_comparison_operator_id"`
	SetPoint                     float32 `json:"set_point"`
}

type TimeSchedule struct {
	StartTime          string    `json:"start_time"`
	EndTime            string    `json:"end_time"`
	SelectedWorkflowID int       `json:"selected_workflow_id"`
	Condition          Condition `json:"condition"`
}

// Request schemas
type HouseRequest string

// type DeviceRequest struct {
// 	DeviceName    string  `json:"device_name"`
// 	ClimateDataID int     `json:"climate_data_id"`
// 	Unit          string  `json:"unit"`
// 	SetPoint      float32 `json:"set_point"`
// 	Duration      int     `json:"duration"`
// }

type NodeRequest struct {
	WorkflowNodeID string                 `json:"workflow_node_id"`
	Type           string                 `json:"type"`
	Data           map[string]interface{} `json:"data"`
	PositionX      float32                `json:"position_x"`
	PositionY      float32                `json:"position_y"`
}

type EdgeRequest struct {
	SourceNodeID string `json:"source_node_id"`
	TargetNodeID string `json:"target_node_id"`
}

type WorkflowUIRequest struct {
	Nodes []NodeRequest `json:"nodes"`
	Edges []EdgeRequest `json:"edges"`
}

type WorkflowWithUIRequest struct {
	Workflow   Workflow          `json:"workflow"`
	WorkflowUI WorkflowUIRequest `json:"workflow_ui"`
	Relays     Relays            `json:"relays"`
}

type TimeScheduleRequest struct {
	M304ID       int            `json:"m304_id"`
	TimeSchedule []TimeSchedule `json:"time_schedule"`
}

// Response schemas
type HousesResponse struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// type DeviceResponse struct {
// 	ID          int     `json:"id"`
// 	DeviceName  string  `json:"device_name"`
// 	ClimateData string  `json:"climate_data"`
// 	Unit        string  `json:"unit"`
// 	SetPoint    float64 `json:"set_point"`
// 	Duration    int     `json:"duration"`
// }

type WorkflowResponse struct {
	ID     int    `json:"id"`
	M304ID int    `json:"m304_id"`
	Name   string `json:"name"`
}

type NodeResponse struct {
	ID             int                    `json:"id"`
	WorkflowID     int                    `json:"workflow_id"`
	WorkflowNodeID string                 `json:"workflow_node_id"`
	Type           string                 `json:"type"`
	Data           map[string]interface{} `json:"data"`
	PositionX      float32                `json:"position_x"`
	PositionY      float32                `json:"position_y"`
}

type EdgeResponse struct {
	ID           int    `json:"id"`
	WorkflowID   int    `json:"workflow_id"`
	SourceNodeID string `json:"source_node_id"`
	TargetNodeID string `json:"target_node_id"`
}

type WorkflowUIResponse struct {
	Nodes []NodeResponse `json:"nodes"`
	Edges []EdgeResponse `json:"edges"`
}

type WorkflowWithUIResponse struct {
	Workflow   WorkflowResponse   `json:"workflow"`
	WorkflowUI WorkflowUIResponse `json:"workflowUI"`
}

type TimeScheduleResponse struct {
	M304ID        int            `json:"m304_id"`
	Workflows     []Workflow     `json:"Workflows"`
	TimeSchedules []TimeSchedule `json:"time_schedules"`
}

// Constructors
// General schema constructors
func NewClimateData(id int, name, unit string) *ClimateData {
	return &ClimateData{
		ID:   id,
		Name: name,
		Unit: unit,
	}
}

func NewRelays(relay1, relay2, relay3, relay4, relay5, relay6, relay7, relay8 bool) *Relays {
	return &Relays{
		Relay1: relay1,
		Relay2: relay2,
		Relay3: relay3,
		Relay4: relay4,
		Relay5: relay5,
		Relay6: relay6,
		Relay7: relay7,
		Relay8: relay8,
	}
}

func NewWorkflow(name string, m304ID, deviceID int) *Workflow {
	return &Workflow{
		Name:     name,
		M304ID:   m304ID,
		DeviceID: deviceID,
	}
}

func NewCondition(selectedClimateDataID, selectedComparisonOperatorID int, setPoint float32) *Condition {
	return &Condition{
		SelectedClimateDataID:        selectedClimateDataID,
		SelectedComparisonOperatorID: selectedComparisonOperatorID,
		SetPoint:                     setPoint,
	}
}

func NewTimeSchedule(startTime, endTime string, selectedWorkflowID int, condition Condition) *TimeSchedule {
	return &TimeSchedule{
		StartTime:          startTime,
		EndTime:            endTime,
		SelectedWorkflowID: selectedWorkflowID,
		Condition:          condition,
	}
}

// Request schema constructors
func NewHouseRequest(name string) *HouseRequest {
	request := HouseRequest(name)
	return &request
}

// func NewDeviceRequest(deviceName string, climateDataID int, unit string, setPoint float32, duration int) *DeviceRequest {
// 	return &DeviceRequest{
// 		DeviceName:    deviceName,
// 		ClimateDataID: climateDataID,
// 		Unit:          unit,
// 		SetPoint:      setPoint,
// 		Duration:      duration,
// 	}
// }

func NewNodeRequest(workflowNodeID, nodeType string, data map[string]interface{}, positionX, positionY float32) *NodeRequest {
	return &NodeRequest{
		WorkflowNodeID: workflowNodeID,
		Type:           nodeType,
		Data:           data,
		PositionX:      positionX,
		PositionY:      positionY,
	}
}

func NewEdgeRequest(sourceNodeID, targetNodeID string) *EdgeRequest {
	return &EdgeRequest{
		SourceNodeID: sourceNodeID,
		TargetNodeID: targetNodeID,
	}
}

func NewWorkflowUIRequest(nodes []NodeRequest, edges []EdgeRequest) *WorkflowUIRequest {
	return &WorkflowUIRequest{
		Nodes: nodes,
		Edges: edges,
	}
}

func NewWorkflowWithUIRequest(workflow Workflow, workflowUI WorkflowUIRequest, relays Relays) *WorkflowWithUIRequest {
	return &WorkflowWithUIRequest{
		Workflow:   workflow,
		WorkflowUI: workflowUI,
		Relays:     relays,
	}
}

func NewTimeScheduleRequest(m304ID int, timeSchedules []TimeSchedule) *TimeScheduleRequest {
	return &TimeScheduleRequest{
		M304ID:       m304ID,
		TimeSchedule: timeSchedules,
	}
}

// Response schema constructors
func NewHousesResponse(id int, name string) *HousesResponse {
	return &HousesResponse{
		ID:   id,
		Name: name,
	}
}

// func NewDeviceResponse(id int, deviceName, climateData, unit string, setPoint float64, duration int) *DeviceResponse {
// 	return &DeviceResponse{
// 		ID:          id,
// 		DeviceName:  deviceName,
// 		ClimateData: climateData,
// 		Unit:        unit,
// 		SetPoint:    setPoint,
// 		Duration:    duration,
// 	}
// }

func NewWorkflowResponse(id, m304ID int, name string) *WorkflowResponse {
	return &WorkflowResponse{
		ID:     id,
		M304ID: m304ID,
		Name:   name,
	}
}

func NewNodeResponse(id, workflowID int, workflowNodeID, nodeType string, data map[string]interface{}, positionX, positionY float32) *NodeResponse {
	return &NodeResponse{
		ID:             id,
		WorkflowID:     workflowID,
		WorkflowNodeID: workflowNodeID,
		Type:           nodeType,
		Data:           data,
		PositionX:      positionX,
		PositionY:      positionY,
	}
}

func NewEdgeResponse(id, workflowID int, sourceNodeID, targetNodeID string) *EdgeResponse {
	return &EdgeResponse{
		ID:           id,
		WorkflowID:   workflowID,
		SourceNodeID: sourceNodeID,
		TargetNodeID: targetNodeID,
	}
}

func NewWorkflowUIResponse(nodes []NodeResponse, edges []EdgeResponse) *WorkflowUIResponse {
	return &WorkflowUIResponse{
		Nodes: nodes,
		Edges: edges,
	}
}

func NewWorkflowWithUIResponse(workflow WorkflowResponse, workflowUI WorkflowUIResponse) *WorkflowWithUIResponse {
	return &WorkflowWithUIResponse{
		Workflow:   workflow,
		WorkflowUI: workflowUI,
	}
}

func NewTimeScheduleResponse(m304ID int, workflows []Workflow, timeSchedules []TimeSchedule) *TimeScheduleResponse {
	return &TimeScheduleResponse{
		M304ID:        m304ID,
		Workflows:     workflows,
		TimeSchedules: timeSchedules,
	}
}
