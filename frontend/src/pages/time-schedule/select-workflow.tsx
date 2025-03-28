import { Workflow } from "@/types/api";
import { FormControl, InputLabel, Select, MenuItem } from "@mui/material";
import { useEffect, useState } from "react";
import { SelectChangeEvent } from "@mui/material";

interface WorkflowSelectProps {
  initialWorkflow: Workflow | null;
  workflows: Workflow[];
  index: number;
  onSelectChange: (updatedData: Workflow) => void;
}

export const WorkflowSelect = (props: WorkflowSelectProps) => {
  const { initialWorkflow, workflows, onSelectChange } = props;

  const [options] = useState<Workflow[]>(workflows);
  const [selectedWorkflow, setSelectedWorkflow] = useState<Workflow | null>(
    initialWorkflow
  );

  useEffect(() => {
    if (!initialWorkflow) {
      return;
    }

    setSelectedWorkflow(initialWorkflow);
  }, []);

  const handleWorkflowChange = (event: SelectChangeEvent) => {
    const workflowID = parseInt(event.target.value);
    const workflowRec = options.find((data) => data.id === workflowID);

    if (workflowRec) {
      setSelectedWorkflow(workflowRec);
      onSelectChange(workflowRec);
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
        value={
          selectedWorkflow && selectedWorkflow.id !== 0
            ? String(selectedWorkflow.id)
            : ""
        }
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
