import {
  Box,
  Select,
  MenuItem,
  SelectChangeEvent,
  Divider,
  FormControl,
  InputLabel,
  Typography,
} from "@mui/material";
import { Handle, Position, Node, NodeProps } from "@xyflow/react";
import { useEffect, useState } from "react";
import { UpdateNodeFunction } from "../workflow";
import { OperationResponse } from "@/types/api";
import { useNodeInfo } from "@/hooks/node-info-context";

export interface DeviceOperationNodeData {
  [key: string]: unknown;
  operationsList: OperationResponse[];
  updateNode: UpdateNodeFunction;
  operationID: number;
}

type DeviceOperationNodePropsType = Node<DeviceOperationNodeData>;

type DeviceOperationNodeProps = NodeProps<DeviceOperationNodePropsType>;

export function DeviceOperationNode({ id, data }: DeviceOperationNodeProps) {
  const { operationsList, updateNode } = data;
  const [selectedOperation, setSelectedOperation] = useState<string>("");
  const handleSelectedDeviceChange = (event: SelectChangeEvent) => {
    const selectedOperationID: number = parseInt(event.target.value, 10);
    setSelectedOperation(event.target.value);
    updateNode(id, { ...data, operationID: selectedOperationID });
  };
  const [deviceOperations, setDeviceOperations] = useState<OperationResponse[]>(
    []
  );
  const [workflowInfo] = useNodeInfo();

  useEffect(() => {
    const selectBoxOperations: OperationResponse[] = operationsList.filter(
      (data) => data.device_id === workflowInfo?.device_id
    );

    console.log(selectBoxOperations);

    setDeviceOperations(selectBoxOperations);

    // const climateDataRec = climateDataList.find(
    //   (data) => data.id === climateDataID
    // );
  }, [operationsList, workflowInfo.device_id]);

  return (
    <Box
      sx={{
        border: "1px solid #000",
        borderRadius: "10px",
        backgroundColor: "#FFF",
        width: "350px",
      }}
    >
      <Handle position={Position.Left} type="target" />
      <Box sx={{ padding: "8px" }}>
        <Typography variant="h6">Operation</Typography>
        <Divider />
        <Box sx={{ padding: "8px", display: "flex", justifyContent: "center" }}>
          <Box sx={{ flex: 2 }}>
            <FormControl fullWidth>
              <InputLabel id="demo-simple-select-label">操作</InputLabel>
              <Select
                value={selectedOperation}
                size="small"
                onChange={handleSelectedDeviceChange}
                label="操作"
                inputProps={{ className: "nodrag nopan nowheel" }}
                disabled={deviceOperations.length === 0}
              >
                {deviceOperations.map((data) => (
                  <MenuItem key={data.id} value={data.id}>
                    {data.name}
                  </MenuItem>
                ))}
              </Select>
            </FormControl>
          </Box>
        </Box>
      </Box>
    </Box>
  );
}