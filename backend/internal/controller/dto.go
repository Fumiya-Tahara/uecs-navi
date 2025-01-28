package controller

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
