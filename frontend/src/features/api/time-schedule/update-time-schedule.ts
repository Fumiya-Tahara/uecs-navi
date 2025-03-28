import { TimeScheduleRequest } from "@/types/api";
import { apiClient } from "@/lib/api-client";

// 戻り値は更新したレコードのID
export async function updateTimeSchedule(
  data: TimeScheduleRequest
): Promise<number> {
  try {
    const response = await apiClient.put(`/time-schedule`, data);
    return response.data;
  } catch (error) {
    console.error("Error updating time schedule:", error);
    throw error;
  }
}
