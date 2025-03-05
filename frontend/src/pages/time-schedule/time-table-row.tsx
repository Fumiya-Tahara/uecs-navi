import { TableRow, TableCell, TextField } from "@mui/material";
import { AdapterDayjs } from "@mui/x-date-pickers/AdapterDayjs";
import dayjs from "dayjs";
import { LocalizationProvider, TimePicker } from "@mui/x-date-pickers";
import { WorkflowSelect } from "./select-workflow";
import { EnvConditionForms } from "./env-condition-forms";
import { Workflow, TimeScheduleRow } from "@/types/api";
import { useEffect, useState } from "react";

interface TimeTableRowProps {
  timeScheduleRow: TimeScheduleRow | null;
  workflows: Workflow[];
  index: number;
  onRowChange: (index: number, updatedData: TimeScheduleRow) => void;
}

export function TimeTableRow(props: TimeTableRowProps) {
  const { timeScheduleRow, workflows, index, onRowChange } = props;
  const [, setSelectedWorkflowID] = useState<number>(
    timeScheduleRow?.selected_workflow_id || 0
  );
  const [, forceRender] = useState(0);

  useEffect(() => {
    forceRender((prev) => prev + 1); // state を変更して強制リレンダリング
  }, [timeScheduleRow]);

  const handleStartTimeChange = (value: string) => {
    if (value !== undefined) {
      if (timeScheduleRow) {
        timeScheduleRow.start_time = value;

        onRowChange(index, timeScheduleRow);
      }
    }
  };

  const handleEndTimeChange = (value: string) => {
    if (value !== undefined) {
      if (timeScheduleRow) {
        timeScheduleRow.end_time = value;

        onRowChange(index, timeScheduleRow);
      }
    }
  };

  const handleSelectChange = (updatedData: Workflow) => {
    if (timeScheduleRow) {
      const selectedWorkflowID: number = updatedData.id;
      setSelectedWorkflowID(selectedWorkflowID);
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
              timeScheduleRow?.start_time
                ? dayjs(timeScheduleRow.start_time, "HH:mm").isValid()
                  ? dayjs(timeScheduleRow.start_time, "HH:mm")
                  : null
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
              timeScheduleRow?.end_time
                ? dayjs(timeScheduleRow.end_time, "HH:mm").isValid()
                  ? dayjs(timeScheduleRow.end_time, "HH:mm")
                  : null
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
