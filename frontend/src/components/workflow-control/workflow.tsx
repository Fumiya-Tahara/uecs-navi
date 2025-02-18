import { Box } from "@mui/material";
import { Navigation } from "@/layouts/navigation";
import { WorkflowInfoProvider } from "@/hooks/workflow-info-context";
import { useWorkflowInfo } from "@/hooks/workflow-info-context";
import { useEffect } from "react";
import {
  ClimateDataResponse,
  DeviceResponse,
  OperationResponse,
} from "@/types/api";
import {
  getClimateDatas,
  getDevices,
  getOperations,
} from "@/features/api/mocks/workflow-api";
import { WorkflowEditor } from "./workflow-editor";
import { ReactFlowProvider } from "@xyflow/react";
import { NodeInfoProvider } from "@/hooks/node-info-context";
import { DnDProvider } from "@/hooks/dnd-context";
import { SelectToolbar } from "./toolbar";

function Workflow() {
  const [, setWorkflowInfo] = useWorkflowInfo();

  useEffect(() => {
    const fetchWorkflowData = async () => {
      const devicesRes: DeviceResponse[] = await getDevices();
      const climateDataRes: ClimateDataResponse[] = await getClimateDatas();
      const operationsRes: OperationResponse[] = await getOperations();

      setWorkflowInfo({
        devices: devicesRes,
        climate_data: climateDataRes,
        operations: operationsRes,
      });
    };

    fetchWorkflowData();
  }, []);

  return (
    <Navigation>
      <Box sx={{ padding: "16px" }}>
        <SelectToolbar />
        <WorkflowEditor />
      </Box>
    </Navigation>
  );
}

export default function WorkflowWrapper() {
  return (
    <WorkflowInfoProvider>
      <ReactFlowProvider>
        <NodeInfoProvider>
          <DnDProvider>
            <Workflow />
          </DnDProvider>
        </NodeInfoProvider>
      </ReactFlowProvider>
    </WorkflowInfoProvider>
  );
}
