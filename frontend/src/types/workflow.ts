import { DeviceResponse } from "./api";

export type CustomNodeData = WorkflowNameNodeData | OperationNodeData;

export type AddNodeFunction = (parentNodeId: string) => void;
export type UpdateNodeFunction = (
  id: string,
  updatedData: CustomNodeData
) => void;

export interface WorkflowNameNodeData {
  [key: string]: unknown;
  addNode: AddNodeFunction;
  updateNode: UpdateNodeFunction;
}

export interface OperationNodeData {
  [key: string]: unknown;
  devicesList: DeviceResponse[];
  addNode: AddNodeFunction;
  updateNode: UpdateNodeFunction;
  device_id: number;
}
