export interface ClimateData {
  id: number;
  name: string;
  unit: string;
}

export interface Relays {
  relay_1: boolean;
  relay_2: boolean;
  relay_3: boolean;
  relay_4: boolean;
  relay_5: boolean;
  relay_6: boolean;
  relay_7: boolean;
  relay_8: boolean;
}

export interface Workflow {
  id: number;
  m304_id: number;
  name: string;
}

export interface Condition {
  selected_climate_data_id: number;
  selected_comparison_operator_id: number;
  set_point: number;
}

export interface TimeScheduleRow {
  start_time: string;
  end_time: string;
  selected_workflow_id: number;
  condition: Condition;
}

export type M304ID = number;

// Request Interfaces

export type HouseRequest = string;

export interface DeviceRequest {
  climate_data_id: number;
  m304_id: number;
  name: string;
  rly: number;
}

export interface NodeRequest {
  workflow_node_id: string;
  node_type: string;
  data: object;
  position_x: number;
  position_y: number;
}

export interface EdgeRequest {
  source_node_id: string;
  target_node_id: string;
}

export interface WorkflowUIRequest {
  nodes: NodeRequest[];
  edges: EdgeRequest[];
}

export interface WorkflowWithUIRequest {
  workflow: Workflow;
  workflow_ui: WorkflowUIRequest;
  relays: Relays;
}

export interface TimeScheduleRequest {
  id: number;
  m304_id: number;
  time_schedule: TimeScheduleRow[];
}

// Response Interfaces

export interface HousesResponse {
  id: number;
  name: string;
}

export interface DeviceResponse {
  id: number;
  climate_data_id: number;
  m304_id: number;
  name: string;
  rly: number;
}

export interface NodeResponse {
  id: number;
  workflow_id: number;
  workflow_node_id: string;
  node_type: string;
  data: object;
  position_x: number;
  position_y: number;
}

export interface EdgeResponse {
  id: number;
  workflow_id: number;
  source_node_id: string;
  target_node_id: string;
}

export interface WorkflowUIResponse {
  nodes: NodeResponse[];
  edges: EdgeResponse[];
}

export interface WorkflowWithUIResponse {
  workflow: Workflow;
  workflow_ui: WorkflowUIResponse;
}

export interface TimeScheduleResponse {
  id: number;
  m304_id: number;
  workflows: Workflow[];
  time_schedules: TimeScheduleRow[];
}
