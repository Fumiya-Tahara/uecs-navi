import { Box } from "@mui/material";
import {
  Node,
  Edge,
  Connection,
  useNodesState,
  useEdgesState,
  addEdge,
  ReactFlow,
  Background,
  BackgroundVariant,
  useReactFlow,
} from "@xyflow/react";
import "@xyflow/react/dist/style.css";
import { useMemo, useEffect, useCallback, DragEvent } from "react";
import { Sidebar } from "./sidebar";
import { useDnD } from "@/pages/create-workflow/context/dnd-context";
import { useWorkflowInfo } from "@/pages/create-workflow/context/workflow-info-context";
import {
  WorkflowNameNode,
  WorkflowNameNodeData,
} from "./custom-nodes/workflow-name";
import { OperationNode, OperationNodeData } from "./custom-nodes/operation";
import { useSelectedData } from "./context/selected-data-context";
import { WorkflowWithUIResponse } from "@/types/api";

type CustomNodeData = WorkflowNameNodeData | OperationNodeData;

export type AddNodeFunction = (parentNodeId: string) => void;
export type UpdateNodeFunction = (
  id: string,
  updatedData: CustomNodeData
) => void;

const nodeIdMap: Map<string, number> = new Map();
const getId = (type: string) => {
  const currentId = nodeIdMap.get(type) || 1;
  nodeIdMap.set(type, currentId + 1);
  return `${type}_${currentId}`;
};

interface WorkflowEditorProps {
  m304ID: number;
}

export function WorkflowEditor(props: WorkflowEditorProps) {
  const { m304ID } = props;
  const [selectedData] = useSelectedData();
  const [nodes, setNodes, onNodesChange] = useNodesState<Node>([]);
  const [edges, setEdges, onEdgesChange] = useEdgesState<Edge>([]);
  const [type] = useDnD();
  const [workflowInfo] = useWorkflowInfo();

  const nodeTypes = useMemo(
    () => ({
      workflow_name: WorkflowNameNode,
      operation: OperationNode,
    }),
    []
  );

  useEffect(() => {
    if (selectedData.selectedWorkflow) {
      const workflow: WorkflowWithUIResponse = selectedData.selectedWorkflow;

      workflow.workflow_ui.nodes.forEach((node) => {
        const currentId = nodeIdMap.get(node.node_type) || 1;
        nodeIdMap.set(node.node_type, currentId + 1);
      });

      const nodes = workflow.workflow_ui.nodes
        .map((node): Node | undefined => {
          switch (node.node_type) {
            case "workflow_name":
              return {
                id: node.workflow_node_id,
                type: node.node_type,
                position: { x: node.position_x, y: node.position_y },
                data: {
                  ...(node.data as Record<string, unknown>),
                  // 必要なデータ
                },
              };
            case "operation":
              if (selectedData.selectedM304ID) {
                return {
                  id: node.workflow_node_id,
                  type: node.node_type,
                  position: { x: node.position_x, y: node.position_y },
                  data: {
                    ...(node.data as Record<string, unknown>),
                    // 必要なデータ
                    devicesList: workflowInfo.m304DeviceMap.get(
                      selectedData.selectedM304ID
                    ),
                  },
                };
              }

              return undefined;
          }
          // if (node.node_type === "select_device") {
          //   return {
          //     id: node.workflow_node_id,
          //     type: node.node_type,
          //     position: { x: node.position_x, y: node.position_y },
          //     data: {
          //       ...(node.data as Record<string, unknown>),
          //       devicesList: workflowInfo.devices,
          //       updateNode: updateNodeData,
          //     },
          //   };
          // } else if (node.node_type === "condition") {
          //   return {
          //     id: node.workflow_node_id,
          //     type: node.node_type,
          //     position: { x: node.position_x, y: node.position_y },
          //     data: {
          //       ...(node.data as Record<string, unknown>),
          //       climateDataList: workflowInfo.climate_data,
          //       updateNode: updateNodeData,
          //     },
          //   };
          // } else if (node.node_type === "device_operation") {
          //   return {
          //     id: node.workflow_node_id,
          //     type: node.node_type,
          //     position: { x: node.position_x, y: node.position_y },
          //     data: {
          //       ...(node.data as Record<string, unknown>),
          //       operationsList: workflowInfo.operations,
          //       updateNode: updateNodeData,
          //     },
          //   };
          // }

          return undefined;
        })
        .filter((node): node is Node => node !== undefined);

      const edges = workflow.workflow_ui.edges.map((edge) => ({
        id: edge.id.toString(),
        source: edge.source_node_id,
        target: edge.target_node_id,
        style: { strokeWidth: 4 },
      }));

      setNodes(nodes);
      setEdges(edges);

      return;
    }

    const initialNode: Node = {
      id: "workflow_name_1",
      type: "workflow_name",
      position: { x: 0, y: 300 },
      data: {
        devicesList: workflowInfo.m304DeviceMap.get(m304ID),
        updateNode: updateNodeData,
      },
    };

    setNodes([initialNode]);
  }, [selectedData.selectedWorkflow]);

  const { screenToFlowPosition } = useReactFlow();

  // イベントハンドラー
  const onConnect = useCallback(
    (params: Connection) => {
      const newEdge = {
        ...params,
        style: { strokeWidth: 3 },
      };
      setEdges((eds) => addEdge(newEdge, eds));
    },
    [setEdges]
  );

  const onDragOver = useCallback((event: DragEvent) => {
    event.preventDefault();
    event.dataTransfer.dropEffect = "move";
  }, []);

  const updateNodeData = useCallback(
    (id: string, updatedData: CustomNodeData) => {
      setNodes((nds) =>
        nds.map((node) =>
          node.id === id
            ? { ...node, data: { ...node.data, ...updatedData } }
            : node
        )
      );
    },
    [setNodes]
  );

  const onDrop = useCallback(
    (event: DragEvent) => {
      event.preventDefault();

      if (!type) {
        return;
      }

      const dataString = event.dataTransfer.getData("application/reactflow");
      const nodeData = dataString ? JSON.parse(dataString) : {};

      const position = screenToFlowPosition({
        x: event.clientX,
        y: event.clientY,
      });
      const newNode = {
        id: getId(type),
        type,
        position,
        data: {
          label: `${type} node`,
          ...nodeData,
          updateNode: updateNodeData,
        },
      };

      setNodes((nds) => nds.concat(newNode));
    },
    [screenToFlowPosition, type]
  );

  // 画面の大きさ調節
  interface Viewport {
    x: number;
    y: number;
    zoom: number;
  }
  const defaultViewport: Viewport = { x: 50, y: 15, zoom: 0 };

  return (
    <>
      <Box sx={{ width: "100%", height: "80vh", backgroundColor: "#eee" }}>
        <Box sx={{ width: "100%", height: "100%", display: "flex" }}>
          <ReactFlow
            nodes={nodes}
            edges={edges}
            nodesDraggable={true} // ノードのドラッグを無効化
            edgesReconnectable={true} // エッジの更新を無効化
            // panOnDrag={false} // 画面全体のドラッグを無効化
            // zoomOnScroll={false} // マウスホイールでのズームを無効化
            zoomOnPinch={false} // ピンチ操作でのズームを無効化
            zoomOnDoubleClick={false} // ダブルクリックでのズームを無効化
            defaultViewport={defaultViewport} // 初期配置と大きさを設定
            nodeTypes={nodeTypes}
            onNodesChange={onNodesChange}
            onEdgesChange={onEdgesChange}
            onConnect={onConnect}
            onDrop={onDrop}
            onDragOver={onDragOver}
          >
            <Background
              color="#000"
              variant={BackgroundVariant.Dots}
            ></Background>
          </ReactFlow>
          <Sidebar />
        </Box>
      </Box>
    </>
  );
}
