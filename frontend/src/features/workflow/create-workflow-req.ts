import {
  Workflow,
  EdgeRequest,
  NodeRequest,
  WorkflowUIRequest,
  WorkflowWithUIRequest,
  Relays,
  DeviceResponse,
} from "@/types/api";
import { Edge, Node } from "@xyflow/react";

export function createWorkflowWithUIRequest(
  workflowID: number,
  m304ID: number,
  devicesList: DeviceResponse[],
  nodes: Node[],
  edges: Edge[]
): WorkflowWithUIRequest | null {
  const workflow: Workflow | null = createWorkflow(workflowID, m304ID, nodes);
  const workflowUI: WorkflowUIRequest | null = createWorkflowUIRequest(
    nodes,
    edges
  );
  const relays: Relays | null = createRelays(devicesList, nodes);

  if (workflow && workflowUI && relays) {
    const workflowWithUIReq: WorkflowWithUIRequest = {
      workflow: workflow,
      workflow_ui: workflowUI,
      relays: relays,
    };

    return workflowWithUIReq;
  }

  return null;
}

function createWorkflowUIRequest(
  nodes: Node[],
  edges: Edge[]
): WorkflowUIRequest | null {
  const nodesReq: NodeRequest[] = [];
  const edgesReq: EdgeRequest[] = [];

  nodes.forEach((node) => {
    const nodeReq: NodeRequest = {
      workflow_node_id: node.id,
      node_type: String(node.type),
      data: node.data,
      position_x: node.position.x,
      position_y: node.position.y,
    };

    nodesReq.push(nodeReq);
  });

  edges.forEach((edge) => {
    const edgeReq: EdgeRequest = {
      source_node_id: edge.source,
      target_node_id: edge.target,
    };

    edgesReq.push(edgeReq);
  });

  const workflowUIReq: WorkflowUIRequest = {
    nodes: nodesReq,
    edges: edgesReq,
  };

  return workflowUIReq;
}

function createWorkflow(
  workflowID: number,
  m304ID: number,
  nodes: Node[]
): Workflow | null {
  let workflowName: string = "";

  nodes.forEach((node) => {
    if (node.type == "workflow_name") {
      workflowName = String(node.data.deviceID);
    }
  });

  const workflow: Workflow = {
    id: workflowID,
    m304_id: m304ID,
    name: workflowName,
  };

  return workflow;
}

function createRelays(
  devicesList: DeviceResponse[],
  nodes: Node[]
): Relays | null {
  const relays: Relays = {
    relay_1: false,
    relay_2: false,
    relay_3: false,
    relay_4: false,
    relay_5: false,
    relay_6: false,
    relay_7: false,
    relay_8: false,
  };

  nodes.forEach((node) => {
    if (node.type == "operation") {
      const deviceID: number = Number(node.data.deviceID);
      const selectedDevice: DeviceResponse | undefined = devicesList.find(
        (device) => device.id == deviceID
      );

      if (selectedDevice) {
        switch (selectedDevice.rly) {
          case 1:
            relays.relay_1 = true;
            break;
          case 2:
            relays.relay_2 = true;
            break;
          case 3:
            relays.relay_3 = true;
            break;
          case 4:
            relays.relay_4 = true;
            break;
          case 5:
            relays.relay_5 = true;
            break;
          case 6:
            relays.relay_6 = true;
            break;
          case 7:
            relays.relay_7 = true;
            break;
          case 8:
            relays.relay_8 = true;
            break;
          default:
            break;
        }
      }
    }
  });

  return relays;
}
