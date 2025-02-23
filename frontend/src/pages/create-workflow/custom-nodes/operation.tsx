import {
  Box,
  Divider,
  Typography,
  FormControl,
  InputLabel,
  Select,
  MenuItem,
  SelectChangeEvent,
} from "@mui/material";
import DevicesIcon from "@mui/icons-material/Devices";
import { Handle, Position, Node, NodeProps } from "@xyflow/react";
import { useCallback, useState } from "react";
import { useNodeInfo } from "@/pages/create-workflow/context/node-info-context";
import { OperationNodeData } from "@/types/workflow";

type OperationNodePropsType = Node<OperationNodeData>;

type OperationNodeProps = NodeProps<OperationNodePropsType>;

export const OperationNode = ({ id, data }: OperationNodeProps) => {
  const { devicesList, addNode, updateNode } = data;
  const [selectedDevice, setSelectedDevice] = useState<string>(
    data.deviceID ? String(data.deviceID) : ""
  );
  const [nodeInfo, setNodeInfo] = useNodeInfo();
  const handleSelectedDeviceChange = useCallback(
    (event: SelectChangeEvent) => {
      const selectedDeviceID: number = parseInt(event.target.value, 10);
      nodeInfo.deviceID = selectedDeviceID;

      setSelectedDevice(event.target.value);
      setNodeInfo(nodeInfo);
      updateNode(id, { ...data, deviceID: selectedDeviceID });
    },
    [id, updateNode, data, setSelectedDevice]
  );

  return (
    <Box
      sx={{
        border: "1px solid #000",
        borderRadius: "10px",
        backgroundColor: "#FFF",
        width: "350px",
      }}
    >
      <Handle
        position={Position.Left}
        type="target"
        style={{ width: 12, height: 12 }}
      />
      <Box>
        <Box
          sx={{
            display: "flex",
            alignItems: "center",
            gap: 1,
            borderRadius: "10px 10px 0 0",
            color: "#FFF",
            backgroundColor: "#66BB6A",
            padding: "4px 8px 4px 8px",
          }}
        >
          <DevicesIcon />
          <Typography variant="h6">Operation</Typography>
        </Box>
        <Divider />
        <Box sx={{ padding: 2, display: "flex", justifyContent: "center" }}>
          <Box sx={{ flex: 2 }}>
            <FormControl fullWidth>
              <InputLabel
                id={`select-device-node-select-label-${id}`}
                size="small"
              >
                デバイス
              </InputLabel>
              <Select
                value={selectedDevice}
                size="small"
                labelId={`select-device-node-select-label-${id}`}
                id={`select-device-node-select-${id}`}
                onChange={handleSelectedDeviceChange}
                label="デバイス"
                inputProps={{ className: "nodrag nopan nowheel" }}
              >
                {devicesList.map((data) => (
                  <MenuItem key={data.id} value={data.id}>
                    {data.name}
                  </MenuItem>
                ))}
              </Select>
            </FormControl>
          </Box>
        </Box>
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
