import { TableRow, TableCell, TextField } from "@mui/material";
import { AdapterDayjs } from "@mui/x-date-pickers/AdapterDayjs";
import dayjs from "dayjs";
import { LocalizationProvider, TimePicker } from "@mui/x-date-pickers";
import { WorkflowSelect } from "./select-workflow";
import { EnvConditionForms } from "./env-condition-forms";
import { WorkflowResponse, TimeScheduleResponse } from "@/types/api";
import { useState } from "react";

interface TimeTableRowProps {
  timeSchedule: TimeScheduleResponse | null;
  workflows: WorkflowResponse[];
  index: number;
  onRowChange: (index: number, updatedData: TimeScheduleResponse) => void;
}

export function TimeTableRow(props: TimeTableRowProps) {
  const { timeSchedule, workflows, index, onRowChange } = props;
  const [, setSelectedWorkflows] = useState<WorkflowResponse[]>(
    timeSchedule?.workflows || [{ id: 0, name: "" }]
  );

  const handleStartTimeChange = (value: string) => {
    if (value !== undefined) {
      if (timeSchedule) {
        timeSchedule.start_time = value;

        onRowChange(index, timeSchedule);
      }
    }
  };

  const handleEndTimeChange = (value: string) => {
    if (value !== undefined) {
      if (timeSchedule) {
        timeSchedule.end_time = value;

        onRowChange(index, timeSchedule);
      }
    }
  };

  const handleSelectChange = (index: number, updatedData: WorkflowResponse) => {
    if (timeSchedule) {
      const newWorkflows = [...timeSchedule.workflows];
      newWorkflows[index] = updatedData;
      setSelectedWorkflows(newWorkflows);
    }
  };

  return (
    <TableRow
      sx={{
        "&:hover": {
          backgroundColor: "rgb(245, 245, 245)",
        },
      }}
    >
      <TableCell>
        <LocalizationProvider dateAdapter={AdapterDayjs}>
          <TimePicker
            label="開始時間"
            ampm={false}
            value={
              timeSchedule?.start_time
                ? dayjs(timeSchedule.start_time, "HH:mm")
                : null
            }
            onChange={(value) => {
              const startTime = value?.format("HH:mm");
              if (!startTime) {
                return;
              }
              handleStartTimeChange(startTime);
            }}
            slots={{
              textField: (props) => (
                <TextField
                  {...props}
                  size="small"
                  InputLabelProps={{ shrink: true }}
                />
              ),
            }}
            sx={{ marginRight: "8px" }}
          />
        </LocalizationProvider>
      </TableCell>
      <TableCell>
        <LocalizationProvider dateAdapter={AdapterDayjs}>
          <TimePicker
            label="終了時間"
            ampm={false}
            value={
              timeSchedule?.end_time
                ? dayjs(timeSchedule.end_time, "HH:mm")
                : null
            }
            onChange={(value) => {
              const endTime = value?.format("HH:mm");
              if (!endTime) {
                return;
              }
              handleEndTimeChange(endTime);
            }}
            slots={{
              textField: (props) => (
                <TextField
                  {...props}
                  size="small"
                  InputLabelProps={{ shrink: true }}
                />
              ),
            }}
            sx={{ marginRight: "8px" }}
          />
        </LocalizationProvider>
      </TableCell>
      <TableCell>
        <WorkflowSelect
          initialWorkflow={null}
          index={0}
          workflows={workflows}
          onSelectChange={handleSelectChange}
        />
      </TableCell>
      <TableCell>
        <EnvConditionForms />
      </TableCell>
    </TableRow>
  );
}
