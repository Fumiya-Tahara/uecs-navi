import { Box, Divider, Typography, TextField } from "@mui/material";
import TimelineIcon from "@mui/icons-material/Timeline";
import { Handle, Position, Node, NodeProps } from "@xyflow/react";
import { AddNodeFunction, UpdateNodeFunction } from "../workflow-editor";

export interface WorkflowNameNodeData {
  [key: string]: unknown;
  addNode: AddNodeFunction;
  updateNode: UpdateNodeFunction;
}

type WorkflowNameNodePropsType = Node<WorkflowNameNodeData>;

type WorkflowNameNodeProps = NodeProps<WorkflowNameNodePropsType>;

export const WorkflowNameNode = ({ id, data }: WorkflowNameNodeProps) => {
  const { addNode } = data;
  return (
    <Box
      sx={{
        border: "1px solid #000",
        borderRadius: "10px",
        backgroundColor: "#FFF",
        width: "300px",
      }}
    >
      <Box
        sx={{
          display: "flex",
          alignItems: "center",
          gap: 1,
          borderRadius: "10px 10px 0 0",
          color: "#FFF",
          backgroundColor: "#42A5F5",
          padding: "4px 8px 4px 8px",
        }}
      >
        <TimelineIcon />
        <Typography variant="h6">Workflow Name</Typography>
      </Box>
      <Divider />
      <Box sx={{ padding: 2, display: "flex", justifyContent: "center" }}>
        <TextField
          id={`workflow-name-node-textfield-${id}`}
          label="ワークフロー名"
          size="small"
        />
      </Box>
      <Handle
        position={Position.Right}
        type="source"
        style={{ width: 12, height: 12 }}
        onClick={(event) => {
          event.stopPropagation();
          addNode(id);
        }}
      />
    </Box>
  );
};
