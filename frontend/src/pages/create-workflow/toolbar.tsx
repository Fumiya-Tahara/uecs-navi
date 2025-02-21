import {
  AppBar,
  Box,
  Button,
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
import { WorkflowWithUIResponse } from "@/types/api";
import { useM304IDs } from "@/hooks/m304ids-context";

export function SelectToolbar() {
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
            <Grid size={4}>
              <SaveWorkflowButton />
            </Grid>
          </Grid>
        </Toolbar>
      </AppBar>
    </Box>
  );
}

function SaveWorkflowButton() {
  return (
    <Button variant="contained" size="small">
      保存する
    </Button>
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
        {m304IDs.m304IDs.map((m304ID, index) => (
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
    const selectedWorkflow = workflows.find(
      (workflow) => workflow.workflow.id === workflowID
    );

    if (selectedWorkflow) {
      setSelectedData((prev) => ({
        ...prev,
        selectedWorkflow,
      }));
    }
  };

  return (
    <FormControl fullWidth size="small">
      <InputLabel id="demo-simple-select-label">ワークフロー</InputLabel>
      <Select
        labelId="demo-simple-select-label"
        id="demo-simple-select"
        label="ワークフロー"
        value={
          selectedData.selectedWorkflow == null
            ? ""
            : String(selectedData.selectedWorkflow.workflow.id)
        }
        onChange={handleSelectedWorkflowChange}
      >
        {workflows.map((workflow, index) => (
          <MenuItem key={index} value={String(workflow.workflow.id)}>
            {workflow.workflow.name}
          </MenuItem>
        ))}
      </Select>
    </FormControl>
  );
}
