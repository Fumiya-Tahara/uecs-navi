import { TimeScheduleRequest, TimeScheduleResponse } from "@/types/api";
import { createTimeScheduleReq } from "./create-time-schedule-req";
import {
  createTimeSchedule,
  updateTimeSchedule,
} from "@/features/api/mocks/setting-time-schedule-api";

export function saveTimeSchedule(timeSchedule: TimeScheduleResponse | null) {
  if (!timeSchedule) {
    return;
  }

  const timeScheduleReq: TimeScheduleRequest | null = createTimeScheduleReq(
    timeSchedule.id,
    timeSchedule.m304_id,
    timeSchedule.time_schedules
  );

  if (!timeScheduleReq) {
    return;
  }

  console.log(timeScheduleReq);

  const timeScheduleID: number = timeScheduleReq.id;
  if (timeScheduleID == 0) {
    createTimeSchedule(timeScheduleReq);
  } else {
    updateTimeSchedule(timeScheduleReq);
  }
}
