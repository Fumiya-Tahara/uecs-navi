import {
  Button,
  Box,
  TableContainer,
  Table,
  TableHead,
  TableRow,
  TableCell,
  TableBody,
  Paper,
} from "@mui/material";
import { TimeTableRow } from "./time-table-row";
import { TimeScheduleResponse, TimeScheduleRow, Workflow } from "@/types/api";
import { useState } from "react";
import { ConfirmButton } from "@/components/confirm-button";
import { saveTimeSchedule } from "@/features/time-schedule/functions/save-time-schedule";

interface TimeTableProps {
  initialSchedules: TimeScheduleResponse | null;
  workflows: Workflow[];
}

export function TimeTable(props: TimeTableProps) {
  const { initialSchedules, workflows } = props;
  const [timeSchedule, setTimeSchedules] =
    useState<TimeScheduleResponse | null>(initialSchedules);

  const handleRowChange = (index: number, updatedData: TimeScheduleRow) => {
    if (!timeSchedule) {
      return;
    }

    const newTimeSchedule: TimeScheduleResponse = {
      ...timeSchedule,
      time_schedules: timeSchedule.time_schedules.map((schedule, i) =>
        i === index ? updatedData : schedule
      ),
    };
    setTimeSchedules(newTimeSchedule);
  };

  const handleAddRow = () => {
    if (!timeSchedule) {
      return;
    }

    if (timeSchedule.time_schedules.length < 8) {
      const newTimeSchedule: TimeScheduleResponse = {
        ...timeSchedule,
        time_schedules: [
          ...timeSchedule.time_schedules,
          {
            start_time: "",
            end_time: "",
            selected_workflow_id: 0,
            condition: {
              selected_climate_data_id: 0,
              selected_comparison_operator_id: 0,
              set_point: 0,
            },
          },
        ],
      };

      setTimeSchedules(newTimeSchedule);
    }
  };

  const handleRemoveRow = () => {
    if (!timeSchedule) {
      return;
    }

    if (timeSchedule.time_schedules.length > 1) {
      const newTimeSchedule: TimeScheduleResponse = {
        ...timeSchedule,
        time_schedules: timeSchedule.time_schedules.slice(0, -1),
      };
      setTimeSchedules(newTimeSchedule);
    }
  };

  const handleSave = () => {
    saveTimeSchedule(timeSchedule);
  };

  return (
    <Box>
      <TableContainer component={Paper}>
        <Table sx={{ minWidth: 650 }} aria-label="simple table">
          <TableHead>
            <TableRow>
              <TableCell sx={{ width: "20%" }}>開始時間</TableCell>
              <TableCell sx={{ width: "20%" }}>終了時間</TableCell>
              <TableCell sx={{ width: "15%" }}>ワークフロー操作</TableCell>
              <TableCell sx={{ width: "45%" }}>稼働条件</TableCell>
            </TableRow>
          </TableHead>
          <TableBody>
            {timeSchedule?.time_schedules.length == 0 ? (
              <TimeTableRow
                key={0}
                timeScheduleRow={null}
                workflows={workflows}
                index={0}
                onRowChange={handleRowChange}
              />
            ) : (
              timeSchedule?.time_schedules.map((schedule, index) => (
                <TimeTableRow
                  key={index}
                  timeScheduleRow={schedule}
                  workflows={workflows}
                  index={index}
                  onRowChange={handleRowChange}
                />
              ))
            )}
          </TableBody>
        </Table>
      </TableContainer>
      <Box sx={{ mt: 1, display: "flex", justifyContent: "space-between" }}>
        <Box sx={{ display: "flex", gap: 1 }}>
          <Button size="small" variant="contained" onClick={handleAddRow}>
            行を追加
          </Button>
          <Button size="small" variant="outlined" onClick={handleRemoveRow}>
            行を削除
          </Button>
        </Box>
        <Box>
          <ConfirmButton
            buttonLabel="保存する"
            modalTitle="スケジュールを保存する"
            onConfirm={handleSave}
          />
        </Box>
      </Box>
    </Box>
  );
}
