import { Box } from "@mui/material";
import { Navigation } from "@/layouts/navigation";
import { WorkflowInfoProvider } from "@/pages/create-workflow/context/workflow-info-context";
import { useWorkflowInfo } from "@/pages/create-workflow/context/workflow-info-context";
import { useEffect } from "react";
import { DeviceResponse, WorkflowWithUIResponse } from "@/types/api";
import {
  getDevices,
  getWorkflowsWithUI,
} from "@/features/api/mocks/workflow-api";
import { getM304s } from "@/features/api/mocks/utility-api";
import { WorkflowEditor } from "./workflow-editor";
import { ReactFlowProvider } from "@xyflow/react";
import { NodeInfoProvider } from "@/pages/create-workflow/context/node-info-context";
import { DnDProvider } from "@/pages/create-workflow/context/dnd-context";
import { SelectedDataProvider } from "./context/selected-data-context";
import { SelectToolbar } from "./toolbar";

function Workflow() {
  const [, setWorkflowInfo] = useWorkflowInfo();

  useEffect(() => {
    const fetchWorkflowData = async () => {
      const m304Res: number[] = await getM304s();
      const m304DeviceMap = new Map<number, DeviceResponse[]>();
      const m304WorkflowMap = new Map<number, WorkflowWithUIResponse[]>();

      for (const m304ID of m304Res) {
        const devicesRes: DeviceResponse[] | undefined = await getDevices(
          m304ID
        );
        const workflowWithUIRes: WorkflowWithUIResponse[] | undefined =
          await getWorkflowsWithUI(m304ID);

        if (devicesRes) {
          m304DeviceMap.set(m304ID, devicesRes);
        }
        if (workflowWithUIRes) {
          m304WorkflowMap.set(m304ID, workflowWithUIRes);
        }
      }

      setWorkflowInfo({
        m304DeviceMap: m304DeviceMap,
        m304WorkflowMap: m304WorkflowMap,
      });
    };

    fetchWorkflowData();
  }, []);

  return (
    <Navigation>
      <Box sx={{ padding: "16px" }}>
        <SelectToolbar />
        <WorkflowEditor m304ID={1} />
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
            <SelectedDataProvider>
              <Workflow />
            </SelectedDataProvider>
          </DnDProvider>
        </NodeInfoProvider>
      </ReactFlowProvider>
    </WorkflowInfoProvider>
  );
}
