import { DeviceResponse } from "@/types/api";
import { WorkflowWithUIRequest } from "@/types/api";
import { createWorkflowWithUIRequest } from "./create-workflow-req";
import { Node, Edge } from "@xyflow/react";
import {
  createWorkflowWithUI,
  updateWorkflowWithUI,
} from "../api/mocks/workflow-api";

export function saveWorkflow(
  m304ID: number | null,
  workflowID: number | undefined,
  deviceList: DeviceResponse[] | undefined,
  nodes: Node[],
  edges: Edge[]
) {
  if (!m304ID) {
    return;
  }

  if (!workflowID) {
    workflowID = 0;
  }

  if (!deviceList) {
    console.error("device listがありません。");
    return;
  }

  const workflowReq: WorkflowWithUIRequest | null = createWorkflowWithUIRequest(
    workflowID,
    m304ID,
    deviceList,
    nodes,
    edges
  );

  if (!workflowReq) {
    return;
  }

  console.log(workflowReq);

  if (workflowReq?.workflow.id == 0) {
    createWorkflowWithUI(workflowReq);
  } else {
    updateWorkflowWithUI(workflowReq);
  }
}
