import { TimeScheduleRequest, TimeScheduleResponse } from "@/types/api";
import { createTimeScheduleReq } from "./create-time-schedule-req";
// import {
//   createTimeSchedule,
//   updateTimeSchedule,
// } from "../api/mocks/setting-time-schedule-api";
import { createTimeSchedule } from "../api/time-schedule/create-time-schedule";
import { updateTimeSchedule } from "../api/time-schedule/update-time-schedule";

export async function saveTimeSchedule(
  timeSchedule: TimeScheduleResponse | null
) {
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
  try {
    if (timeScheduleID === 0) {
      await createTimeSchedule(timeScheduleReq);
    } else {
      await updateTimeSchedule(timeScheduleReq);
    }
  } catch (error) {
    console.error("Failed to save time schedule:", error);
  }
}
