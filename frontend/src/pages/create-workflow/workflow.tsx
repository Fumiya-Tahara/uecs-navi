import { Box } from "@mui/material";
import { Navigation } from "@/layouts/navigation";
import { WorkflowInfoProvider } from "@/pages/create-workflow/context/workflow-info-context";
import { useWorkflowInfo } from "@/pages/create-workflow/context/workflow-info-context";
import { useEffect } from "react";
import { DeviceResponse, WorkflowWithUIResponse } from "@/types/api";
// import {
//   getDevices,
//   getWorkflowsWithUI,
// } from "@/features/api/mocks/workflow-api";
import { getDevices } from "@/features/api/device/get-device";
import { getWorkflowsWithUI } from "@/features/api/workflow/get-workflow-with-ui";
import { WorkflowEditor } from "./workflow-editor";
import { ReactFlowProvider } from "@xyflow/react";
import { DnDProvider } from "@/pages/create-workflow/context/dnd-context";
import { SelectedDataProvider } from "./context/selected-data-context";
import { SelectToolbar } from "./toolbar";
import { useM304IDs } from "@/hooks/m304ids-context";

function Workflow() {
  const [, setWorkflowInfo] = useWorkflowInfo();
  const [m304IDs] = useM304IDs();

  useEffect(() => {
    const fetchWorkflowData = async () => {
      const m304DeviceMap = new Map<number, DeviceResponse[]>();
      const m304WorkflowMap = new Map<number, WorkflowWithUIResponse[]>();

      if (!m304IDs.m304IDs) {
        return;
      }

      for (const m304ID of m304IDs.m304IDs) {
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
        <DnDProvider>
          <SelectedDataProvider>
            <Workflow />
          </SelectedDataProvider>
        </DnDProvider>
      </ReactFlowProvider>
    </WorkflowInfoProvider>
  );
}
