import { TimeScheduleResponse } from "@/types/api";
import { apiClient } from "@/lib/api-client";

export async function getTimeSchedule(
  m304ID: number
): Promise<TimeScheduleResponse> {
  try {
    const response = await apiClient.get(`/time-schedule/${m304ID}`);
    return response.data;
  } catch (error) {
    console.error("Error getting time schedule:", error);
    throw error;
  }
}
