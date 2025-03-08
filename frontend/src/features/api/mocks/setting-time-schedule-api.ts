import { TimeScheduleRequest, TimeScheduleResponse } from "@/types/api";

export function getTimeSchedules(m304ID: number): TimeScheduleResponse | null {
  const schedule1: TimeScheduleResponse = {
    id: 1,
    m304_id: 1,
    workflows: [
      {
        id: 1,
        m304_id: 1,
        name: "高めの加温",
      },
      {
        id: 2,
        m304_id: 1,
        name: "加湿",
      },
    ],
    time_schedules: [
      {
        start_time: "03:00",
        end_time: "05:00",
        selected_workflow_id: 1,
        condition: {
          selected_climate_data_id: 1,
          selected_comparison_operator_id: 1,
          set_point: 1,
        },
      },
      {
        start_time: "05:30",
        end_time: "07:30",
        selected_workflow_id: 1,
        condition: {
          selected_climate_data_id: 1,
          selected_comparison_operator_id: 1,
          set_point: 1,
        },
      },
      {
        start_time: "08:00",
        end_time: "10:00",
        selected_workflow_id: 1,
        condition: {
          selected_climate_data_id: 1,
          selected_comparison_operator_id: 1,
          set_point: 1,
        },
      },
      {
        start_time: "10:30",
        end_time: "12:30",
        selected_workflow_id: 1,
        condition: {
          selected_climate_data_id: 1,
          selected_comparison_operator_id: 1,
          set_point: 1,
        },
      },
      {
        start_time: "13:00",
        end_time: "15:00",
        selected_workflow_id: 1,
        condition: {
          selected_climate_data_id: 1,
          selected_comparison_operator_id: 1,
          set_point: 1,
        },
      },
    ],
  };

  const schedule2: TimeScheduleResponse = {
    id: 2,
    m304_id: 2,
    workflows: [
      {
        id: 3,
        m304_id: 2,
        name: "高めの加温",
      },
      {
        id: 4,
        m304_id: 2,
        name: "加湿",
      },
    ],
    time_schedules: [
      {
        start_time: "03:00",
        end_time: "05:00",
        selected_workflow_id: 3,
        condition: {
          selected_climate_data_id: 1,
          selected_comparison_operator_id: 1,
          set_point: 1,
        },
      },
      {
        start_time: "05:30",
        end_time: "07:30",
        selected_workflow_id: 3,
        condition: {
          selected_climate_data_id: 1,
          selected_comparison_operator_id: 1,
          set_point: 1,
        },
      },
      {
        start_time: "08:00",
        end_time: "10:00",
        selected_workflow_id: 3,
        condition: {
          selected_climate_data_id: 1,
          selected_comparison_operator_id: 1,
          set_point: 1,
        },
      },
      {
        start_time: "10:30",
        end_time: "12:30",
        selected_workflow_id: 3,
        condition: {
          selected_climate_data_id: 1,
          selected_comparison_operator_id: 1,
          set_point: 1,
        },
      },
      {
        start_time: "13:00",
        end_time: "15:00",
        selected_workflow_id: 3,
        condition: {
          selected_climate_data_id: 1,
          selected_comparison_operator_id: 1,
          set_point: 1,
        },
      },
    ],
  };

  const schedule3: TimeScheduleResponse = {
    id: 3,
    m304_id: 3,
    workflows: [
      {
        id: 1,
        m304_id: 5,
        name: "高めの加温",
      },
      {
        id: 2,
        m304_id: 6,
        name: "加湿",
      },
    ],
    time_schedules: [
      {
        start_time: "03:00",
        end_time: "05:00",
        selected_workflow_id: 5,
        condition: {
          selected_climate_data_id: 1,
          selected_comparison_operator_id: 1,
          set_point: 1,
        },
      },
    ],
  };

  const timeScheduleMap: Map<number, TimeScheduleResponse> = new Map();

  timeScheduleMap.set(1, schedule1);
  timeScheduleMap.set(2, schedule2);
  timeScheduleMap.set(3, schedule3);

  const result: TimeScheduleResponse | null =
    timeScheduleMap.get(m304ID) || null;

  return result;
}

export function createTimeSchedule(timeSchedule: TimeScheduleRequest) {
  console.log("Created: ", timeSchedule);
}

export function updateTimeSchedule(timeSchedule: TimeScheduleRequest) {
  console.log("Updated: ", timeSchedule);
}
