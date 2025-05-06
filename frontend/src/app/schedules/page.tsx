"use client";

import { TimeSchedule } from "@/features/time-schedule/time-schedule";
import { TimeScheduleInfoProvider } from "@/hooks/time-schedule-info-context";

export default function SettingTimeSchedule() {
  return (
    <TimeScheduleInfoProvider>
      <TimeSchedule />
    </TimeScheduleInfoProvider>
  );
}
