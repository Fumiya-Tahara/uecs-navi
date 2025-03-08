import { TimeScheduleRequest, TimeScheduleResponse } from "@/types/api";

export function saveTimeSchedule(timeSchedule: TimeScheduleResponse | null) {
  if (!timeSchedule) {
    return;
  }

  const timeScheduleReq: TimeScheduleRequest = {
    m304_id: timeSchedule.m304_id,
    time_schedule: timeSchedule.time_schedules,
  };

  console.log(timeScheduleReq);
}
