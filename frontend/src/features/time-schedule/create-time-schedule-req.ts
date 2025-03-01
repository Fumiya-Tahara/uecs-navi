import { Condition, TimeScheduleRequest, TimeScheduleRow } from "@/types/api";

export function createTimeScheduleReq(
  m304ID: number,
  timeScheduleRows: TimeScheduleRow[]
): TimeScheduleRequest | null {
  const timeScheduleReq: TimeScheduleRequest = {
    m304_id: m304ID,
    time_schedule: timeScheduleRows,
  };

  return timeScheduleReq;
}

export function createTimeScheduleRow(
  startTime: string,
  endTime: string,
  workflowID: number,
  condition: Condition
): TimeScheduleRow | null {
  const timeScheduleRow: TimeScheduleRow = {
    start_time: startTime,
    end_time: endTime,
    selected_workflow_id: workflowID,
    condition: condition,
  };

  return timeScheduleRow;
}

export function createCondition(
  climateDataID: number,
  compOpeID: number,
  setPoint: number
): Condition | null {
  const condition: Condition = {
    selected_climate_data_id: climateDataID,
    selected_comparison_operator_id: compOpeID,
    set_point: setPoint,
  };

  return condition;
}
