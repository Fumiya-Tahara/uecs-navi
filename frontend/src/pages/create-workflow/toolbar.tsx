import {
  AppBar,
  Box,
  FormControl,
  InputLabel,
  Select,
  MenuItem,
  Toolbar,
} from "@mui/material";
import Grid from "@mui/material/Grid2";
import { grey } from "@mui/material/colors";
import { useSelectedData } from "./context/selected-data-context";
import { SelectChangeEvent } from "@mui/material";
import { useWorkflowInfo } from "./context/workflow-info-context";
import { useEffect, useState } from "react";
import { DeviceResponse, WorkflowWithUIResponse } from "@/types/api";
import { useM304IDs } from "@/hooks/m304ids-context";
import { useNodesAndEdges } from "@/lib/nodes-and-edges-store";
import { ConfirmButton } from "@/components/confirm-button";
import { saveWorkflow } from "@/features/workflow/save-workflow";
import { deleteWorkflow } from "@/features/workflow/delete-workflow";

export function SelectToolbar() {
  const nodes = useNodesAndEdges((state) => state.nodes);
  const edges = useNodesAndEdges((state) => state.edges);

  const [selectedData] = useSelectedData();
  const [workflowInfo] = useWorkflowInfo();

  const m304ID: number | null = selectedData.selectedM304ID;
  const workflowID: number | undefined =
    selectedData.selectedWorkflow?.workflow.id;

  let deviceList: DeviceResponse[] | undefined = undefined;
  if (m304ID) {
    deviceList = workflowInfo.m304DeviceMap.get(m304ID);
  }

  const handleSaveWorkflow = () => {
    if (!selectedData.selectedWorkflow) {
      console.error("ワークフローが選択されていません。");

      return;
    }
    saveWorkflow(m304ID, workflowID, deviceList, nodes, edges);
  };

  const handleDeleteWorkflow = () => {
    if (workflowID) {
      deleteWorkflow(workflowID);
    }
  };

  return (
    <Box sx={{ flexGrow: 1 }}>
      <AppBar
        position="static"
        sx={{
          backgroundColor: "#FFFFFF",
          color: grey[800],
        }}
      >
        <Toolbar>
          <Grid container spacing={2} sx={{ width: "100%" }}>
            <Grid size={8} sx={{ display: "flex", gap: 1 }}>
              <SelectM304 />
              <SelectWorkflow />
            </Grid>
            <Grid size={4} sx={{ display: "flex", gap: 1 }}>
              <ConfirmButton
                buttonLabel="保存する"
                modalTitle="ワークフローを保存する"
                onConfirm={handleSaveWorkflow}
              />
              <ConfirmButton
                buttonLabel="削除する"
                modalTitle="ワークフローを削除する"
                onConfirm={handleDeleteWorkflow}
                buttonColor="error"
              />
            </Grid>
          </Grid>
        </Toolbar>
      </AppBar>
    </Box>
  );
}

function SelectM304() {
  const [m304IDs] = useM304IDs();
  const [selectedData, setSelectedData] = useSelectedData();

  const handleSelectedM304Change = (event: SelectChangeEvent) => {
    const m304ID = parseInt(event.target.value);

    setSelectedData((prev) => ({
      ...prev,
      selectedM304ID: m304ID,
    }));
  };

  return (
    <FormControl fullWidth size="small">
      <InputLabel id="demo-simple-select-label">M304</InputLabel>
      <Select
        labelId="demo-simple-select-label"
        id="demo-simple-select"
        label="M304"
        value={
          selectedData.selectedM304ID == null
            ? ""
            : String(selectedData.selectedM304ID)
        }
        onChange={handleSelectedM304Change}
      >
        {m304IDs.m304IDs?.map((m304ID, index) => (
          <MenuItem key={index} value={String(m304ID)}>
            M304 ID-{m304ID}
          </MenuItem>
        ))}
      </Select>
    </FormControl>
  );
}

function SelectWorkflow() {
  const [workflowInfo] = useWorkflowInfo();
  const [selectedData, setSelectedData] = useSelectedData();
  const [selectedWorkflowID, setSelectedWorkflowID] = useState<string | null>(
    null
  );
  const [workflows, setWorkflows] = useState<WorkflowWithUIResponse[]>([]);

  useEffect(() => {
    if (
      workflowInfo.m304WorkflowMap != null &&
      selectedData.selectedM304ID != null
    ) {
      const workflowsOption: WorkflowWithUIResponse[] | undefined =
        workflowInfo.m304WorkflowMap.get(selectedData.selectedM304ID);
      if (workflowsOption) {
        setWorkflows(workflowsOption);
      }
    }
  }, [workflowInfo, selectedData]);

  const handleSelectedWorkflowChange = (event: SelectChangeEvent) => {
    const workflowID = parseInt(event.target.value);
    if (workflowID === 0) {
      setSelectedData((prev) => ({
        ...prev,
        selectedWorkflow: {
          workflow: {
            id: 0,
            m304_id: 0,
            name: "",
          },
          workflow_ui: {
            nodes: [],
            edges: [],
          },
        },
      }));

      setSelectedWorkflowID("0");
    } else {
      const selectedWorkflow = workflows.find(
        (workflow) => workflow.workflow.id === workflowID
      );

      if (selectedWorkflow) {
        setSelectedData((prev) => ({
          ...prev,
          selectedWorkflow,
        }));

        setSelectedWorkflowID(String(selectedWorkflow.workflow.id));
      }
    }
  };

  return (
    <FormControl fullWidth size="small">
      <InputLabel id="demo-simple-select-label">ワークフロー</InputLabel>
      <Select
        labelId="demo-simple-select-label"
        id="demo-simple-select"
        label="ワークフロー"
        value={selectedWorkflowID != null ? selectedWorkflowID : ""}
        onChange={handleSelectedWorkflowChange}
      >
        <MenuItem value={String(0)} sx={{ color: "green" }}>
          {"＋新規作成"}
        </MenuItem>
        {workflows.map((workflow, index) => (
          <MenuItem key={index} value={String(workflow.workflow.id)}>
            {workflow.workflow.name}
          </MenuItem>
        ))}
      </Select>
    </FormControl>
  );
}
