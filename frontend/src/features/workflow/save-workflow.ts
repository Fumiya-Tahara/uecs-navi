import { DeviceResponse } from "@/types/api";
import { WorkflowWithUIRequest } from "@/types/api";
import { createWorkflowWithUIRequest } from "./create-workflow-req";
import { Node, Edge } from "@xyflow/react";
// import {
//   createWorkflowWithUI,
//   updateWorkflowWithUI,
// } from "../api/mocks/workflow-api";
import { createWorkflowWithUI } from "../api/workflow/create-workflow";
import { updateWorkflowWithUI } from "../api/workflow/update-workflow";

export async function saveWorkflow(
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

  try {
    if (workflowReq?.workflow.id == 0) {
      await createWorkflowWithUI(workflowReq);
    } else {
      await updateWorkflowWithUI(workflowReq);
    }
  } catch (error) {
    console.error("Failed to save workflow", error);
  }
}
