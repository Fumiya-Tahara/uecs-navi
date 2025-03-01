import { Workflow } from "@/types/api";
import { FormControl, InputLabel, Select, MenuItem } from "@mui/material";
import { useState } from "react";
import { SelectChangeEvent } from "@mui/material";

interface WorkflowSelectProps {
  initialWorkflow: Workflow | null;
  workflows: Workflow[];
  index: number;
  onSelectChange: (updatedData: Workflow) => void;
}

export const WorkflowSelect = (props: WorkflowSelectProps) => {
  const { initialWorkflow, workflows } = props;

  const [options] = useState<Workflow[]>(workflows);
  const [selectedWorkflow, setSelectedWorkflow] = useState<Workflow | null>(
    initialWorkflow
  );

  const handleWorkflowChange = (event: SelectChangeEvent) => {
    const workflowID = parseInt(event.target.value);
    const workflowRec = options.find((data) => data.id === workflowID);

    if (workflowRec) {
      setSelectedWorkflow(workflowRec);
    }
  };

  return (
    <FormControl fullWidth size="small" variant="outlined">
      <InputLabel shrink id="workflow-select-label">
        ワークフロー
      </InputLabel>
      <Select
        labelId="workflow-select-label"
        id="workflow-select"
        value={selectedWorkflow ? String(selectedWorkflow.id) : ""}
        onChange={handleWorkflowChange}
        label="ワークフロー"
        notched
      >
        {options.map((data, index) => (
          <MenuItem key={index} value={data.id}>
            {data.name}
          </MenuItem>
        ))}
      </Select>
    </FormControl>
  );
};
